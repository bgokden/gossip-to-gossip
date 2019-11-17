package gossip

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
	"google.golang.org/grpc"
	grpcPeer "google.golang.org/grpc/peer"

	"github.com/bgokden/gossip-to-gossip/negotiation"
	pb "github.com/bgokden/gossip-to-gossip/protos"
)

// Gossip struct keeps info of a peer
type Gossip struct {
	GossipPort         uint32
	RegistrationPort   uint32
	PredefinedKey      *negotiation.PredefinedKey
	ClientTokens       []string
	ServerTokens       []string
	BroadcastAddresses []string
	Certs              map[string][]byte
	StaticServices     []string
	Services           *cache.Cache // []*string
	Peers              *cache.Cache // []*pb.Peer
}

// SecureConnectionInfo defines information requred for a tls connection
type SecureConnectionInfo struct {
	Address    string
	PrivateKey []byte
	PublicKey  []byte
}

// NewGossip return a new instance of gossip
func NewGossip(gossipPort uint32, registrationPort uint32, services []string, broadcastAddresses []string, predefinedKey *negotiation.PredefinedKey) *Gossip {
	g := &Gossip{
		GossipPort:         gossipPort,
		RegistrationPort:   registrationPort,
		PredefinedKey:      predefinedKey,
		BroadcastAddresses: broadcastAddresses,
		StaticServices:     services,
		Services:           cache.New(60*time.Minute, 10*time.Minute),
		Peers:              cache.New(60*time.Minute, 10*time.Minute),
	}
	for _, service := range services {
		g.Services.Set(service, false, cache.NoExpiration)
	}
	return g
}

func (g *Gossip) StartRegistrationServer() error {
	log.Printf("Starting registration server at %v\n", fmt.Sprintf("0.0.0.0:%d", g.RegistrationPort))
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", g.RegistrationPort))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRegistrationServer(grpcServer, g)
	log.Printf("Serve registration server at %v\n", fmt.Sprintf("0.0.0.0:%d", g.RegistrationPort))
	return grpcServer.Serve(lis)
}

func (g *Gossip) getRegistrationClient(address string) (pb.RegistrationClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("fail to dial: %v\n", err)
		return nil, nil, err
	}
	client := pb.NewRegistrationClient(conn)
	return client, conn, nil
}

func (g *Gossip) RegisterToKnownServices() error {
	for service, item := range g.Services.Items() {
		if isRegistred, ok := item.Object.(bool); ok && !isRegistred {
			log.Printf("registering to %v\n", service)
			client, conn, err := g.getRegistrationClient(service)
			if err != nil {
				log.Printf("connection failed to %v %v\n", service, err)
			}
			defer conn.Close()
			err = g.callRegister(client)
			if err == nil {
				log.Printf("registration successful to %v\n", service)
				g.Services.Set(service, false, cache.NoExpiration)
			} else {
				log.Printf("registration failed to %v %v\n", service, err)
			}
		}
	}
	return nil
}

// Check method is regular check network method
func (g *Gossip) Check() error {
	err := g.RegisterToKnownServices()
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	return nil
}

// Run is the main loop
func (g *Gossip) Run() error {
	pollInterval := 100

	timerCh := time.Tick(time.Duration(pollInterval) * time.Millisecond)

	for range timerCh {
		g.Check()
	}
	return nil
}

func (g *Gossip) callRegister(client pb.RegistrationClient) error {
	stream, err := client.Register(context.Background())
	if err != nil {
		return err
	}
	ns := g.PredefinedKey.CreateNegotiationSession()

	firstMessage, err := ns.Get()
	if err != nil {
		return err
	}
	err = stream.Send(&pb.BytesMessage{
		Value: firstMessage,
	})
	if err != nil {
		return err
	}
	echoMessage, err := stream.Recv()
	if err != nil {
		return err
	}
	ns.Validate(echoMessage.GetValue())
	err = stream.CloseSend()
	if err != nil {
		return err
	}
	return nil
}

// Register fucntion to init a connection to join gossip network
func (g *Gossip) Register(stream pb.Registration_RegisterServer) error {
	firstMessage, err := stream.Recv()
	if err != nil {
		return nil
	}
	decrpytedFirstMessage, err := g.PredefinedKey.Decrypt(firstMessage.GetValue())
	if err != nil {
		return err
	}
	encyrptedEcho, err := g.PredefinedKey.EncryptWithPrefix(decrpytedFirstMessage)
	if err != nil {
		return err
	}
	err = stream.Send(&pb.BytesMessage{
		Value: encyrptedEcho,
	})
	if err != nil {
		return err
	}
	return nil
}

func (g *Gossip) callJoin(client pb.GossipClient) error {
	request := &pb.JoinRequest{}

	log.Printf("Call Join Request %v", *request)

	resp, err := client.Join(context.Background(), request)
	if err != nil {
		log.Printf("(Call Join) There is an error %v", err)
		return err
	}
	log.Printf("(Call Join) Respinse %v\n", resp)
	/*
		if s.address != resp.GetAddress() {
			s.address = resp.GetAddress()
		}
	*/
	return nil
}

func (g *Gossip) Join(ctx context.Context, in *pb.JoinRequest) (*pb.JoinResponse, error) {
	peerInfo, ok := grpcPeer.FromContext(ctx)
	if !ok {
		log.Printf("Peer can not be get from context %v\n", peerInfo)
		return nil, errors.New("Peer can not be get from context")
	}
	address := strings.Split(peerInfo.Addr.String(), ":")[0]
	// + ":" + strconv.FormatInt(int64(in.GetPort()), 10)
	// log.Printf("Peer with Addr: %s called Join", address)
	peerStruct := in.Peer
	g.Peers.Add(peerStruct.ConnectionInfo.Addresses[0], peerStruct, cache.DefaultExpiration)
	return &pb.JoinResponse{Address: address}, nil
}

func (g *Gossip) callExchangeServices(client pb.GossipClient) error {
	outputServiceList := make([]string, 0)

	services := g.Services.Items()

	for serviceName := range services {
		outputServiceList = append(outputServiceList, serviceName)
	}

	request := &pb.ServiceMessage{
		Services: outputServiceList,
	}
	resp, err := client.ExchangeServices(context.Background(), request)
	if err != nil {
		log.Printf("(callExchangeServices) There is an error %v", err)
		return err
	}
	inputServiceList := resp.GetServices()
	for i := 0; i < len(inputServiceList); i++ {
		g.Services.Set(inputServiceList[i], true, cache.DefaultExpiration)
	}
	log.Printf("Services exchanged")
	return nil
}

func (s *Gossip) ExchangeServices(ctx context.Context, in *pb.ServiceMessage) (*pb.ServiceMessage, error) {
	return &pb.ServiceMessage{}, nil
}

func (s *Gossip) ExchangePeers(ctx context.Context, in *pb.PeerMessage) (*pb.PeerMessage, error) {
	return &pb.PeerMessage{}, nil
}
