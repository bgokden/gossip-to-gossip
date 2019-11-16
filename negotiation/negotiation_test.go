package negotiation_test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/bgokden/gossip-to-gossip/negotiation"
	"github.com/stretchr/testify/assert"
)

func TestNegotiation(t *testing.T) {

	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)

	assert.Nil(t, err)

	prefix := negotiation.RandomBytes(50)

	pk1 := negotiation.NewPredefinedKey(key, prefix)

	pk2 := negotiation.NewPredefinedKey(key, prefix)

	ns1 := pk1.CreateNegotiationSession()

	encyrpted, err := ns1.Get()

	assert.Nil(t, err)

	decrpyted, err := pk2.Decrypt(encyrpted)

	assert.Nil(t, err)
	assert.NotNil(t, decrpyted)

	encyrptedEcho, err := pk2.EncryptWithPrefix(decrpyted)

	assert.Nil(t, err)
	assert.NotNil(t, encyrptedEcho)

	validation, err := ns1.Validate(encyrptedEcho)
	assert.Nil(t, err)
	assert.True(t, validation)
}

func TestNegotiationAndCertTransfer(t *testing.T) {

	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)

	assert.Nil(t, err)

	prefix := negotiation.RandomBytes(50)

	pk1 := negotiation.NewPredefinedKey(key, prefix)

	pk2 := negotiation.NewPredefinedKey(key, prefix)

	ns1 := pk1.CreateNegotiationSession()

	encyrpted, err := ns1.Get()

	assert.Nil(t, err)

	decrpyted, err := pk2.Decrypt(encyrpted)

	assert.Nil(t, err)
	assert.NotNil(t, decrpyted)

	encyrptedEcho, err := pk2.EncryptWithPrefix(decrpyted)

	assert.Nil(t, err)
	assert.NotNil(t, encyrptedEcho)

	validation, err := ns1.Validate(encyrptedEcho)
	assert.Nil(t, err)
	assert.True(t, validation)
}
