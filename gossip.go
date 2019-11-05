package gossip

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/bgokden/gossip-to-gossip/protos"
)

// Gossip struct keeps info of a peer
type Gossip struct {
	BootstrapServices []string
	Services          sync.Map // []*string
	Peers             sync.Map // []*pb.Peer
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
	return &pb.JoinResponse{Address: "..."}, nil
}

func (g *Gossip) callExchangeServices(client pb.GossipClient) error {
	outputServiceList := make([]string, 0)
	g.Services.Range(func(key, value interface{}) bool {
		serviceName := key.(string)
		outputServiceList = append(outputServiceList, serviceName)
		return true
	})
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
		g.Services.Store(inputServiceList[i], true)
	}
	log.Printf("Services exhanged")
	return nil
}

func (s *Gossip) ExchangeServices(ctx context.Context, in *pb.ServiceMessage) (*pb.ServiceMessage, error) {
	return &pb.ServiceMessage{}, nil
}

func (s *Gossip) ExchangePeers(ctx context.Context, in *pb.PeerMessage) (*pb.PeerMessage, error) {
	return &pb.PeerMessage{}, nil
}
