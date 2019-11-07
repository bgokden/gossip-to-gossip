package gossip

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
	grpcPeer "google.golang.org/grpc/peer"

	pb "github.com/bgokden/gossip-to-gossip/protos"
)

// Gossip struct keeps info of a peer
type Gossip struct {
	BroadcastAddresses []string
	StaticServices     []string
	Services           *cache.Cache // []*string
	Peers              *cache.Cache // []*pb.Peer
}

// NewGossip return a new instance of gossip
func NewGossip(services []string, broadcastAddresses []string) *Gossip {
	return &Gossip{
		BroadcastAddresses: broadcastAddresses,
		StaticServices:     services,
		Services:           cache.New(60*time.Minute, 10*time.Minute),
		Peers:              cache.New(60*time.Minute, 10*time.Minute),
	}
}

// Check method is regular check network method
func (g *Gossip) Check() error {
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
	g.Peers.Add(peerStruct.Address, peerStruct, cache.DefaultExpiration)
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
