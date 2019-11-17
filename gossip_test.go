package gossip_test

import (
	"testing"

	gossip "github.com/bgokden/gossip-to-gossip"
)

func TestGossipNegotiation(t *testing.T) {
	services := []string{"localhost:9090"}
	g1 := gossip.NewGossip(9000, 9001, services, []string{"localhost:9000"})
	g2 := gossip.NewGossip(8000, 8001, services, []string{"localhost:8000"})

	g1.StartNegotiationServer()
	g2.StartNegotiationServer()

	// assert negotiation successful
}
