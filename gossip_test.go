package gossip_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	gossip "github.com/bgokden/gossip-to-gossip"
	negotiation "github.com/bgokden/gossip-to-gossip/negotiation"
)

func TestGossipNegotiation(t *testing.T) {

	rs := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(rs)

	port1 := uint32(1025 + randSource.Intn(1000))
	port2 := uint32(1025 + randSource.Intn(1000))

	pk := negotiation.NewRandomPredefinedKey()

	services := []string{fmt.Sprintf("localhost:%d", port1)}
	t.Logf("services: %v\n", services)
	g1 := gossip.NewGossip(9000, port1, services, []string{"localhost:9000"}, pk)
	g2 := gossip.NewGossip(8000, port2, services, []string{"localhost:8000"}, pk)

	t.Log("Starting registration")
	go func() {
		err := g1.StartRegistrationServer()
		if err != nil {
			t.Logf("G1 err: %v\n", err)
		}
	}()

	time.Sleep(500 * time.Millisecond)

	go g1.Run()
	go g2.Run()
	// assert negotiation successful
	time.Sleep(300 * time.Millisecond)
	isRegistredAsInterface, ok := g2.Services.Get(services[0])
	assert.True(t, ok)
	assert.NotNil(t, isRegistredAsInterface)
	isRegistred, ok := isRegistredAsInterface.(bool)
	assert.True(t, ok)
	assert.True(t, isRegistred)
}
