package negotiation

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha1"
	"log"
	"math/rand"
	"time"
	"unsafe"
)

type NegotiationSession struct {
	pk      *PredefinedKey
	message []byte
}

type PredefinedKey struct {
	seed    []byte
	prefix  []byte
	private *rsa.PrivateKey
}

func NewPredefinedKey(private *rsa.PrivateKey, seed []byte, prefix []byte) *PredefinedKey {
	return &PredefinedKey{
		seed:    seed,
		prefix:  prefix,
		private: private,
	}
}

func (pk *PredefinedKey) EncryptWithPrefix(message []byte) ([]byte, error) {
	message = append(pk.prefix, message...)
	return pk.Encrypt(message)
}

func (pk *PredefinedKey) Encrypt(message []byte) ([]byte, error) {
	sha1 := sha1.New()
	randomSource := bytes.NewReader(pk.seed)
	encrypted, err := rsa.EncryptOAEP(sha1, randomSource, &pk.private.PublicKey, message, nil)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil, err
	}
	return encrypted, nil
}

func (pk *PredefinedKey) Decrypt(encrypted []byte) ([]byte, error) {
	sha1 := sha1.New()
	decrypted, err := rsa.DecryptOAEP(sha1, nil, pk.private, encrypted, nil)
	if err != nil {
		log.Printf("error: %s\n", err)
		return nil, err
	}
	return decrypted, nil
}

func (pk *PredefinedKey) CreateNegotiationSession() *NegotiationSession {
	message := RandomString(50)
	return &NegotiationSession{
		pk:      pk,
		message: []byte(message),
	}
}

func (ns *NegotiationSession) Get() ([]byte, error) {
	return ns.pk.Encrypt(ns.message)
}

func (ns *NegotiationSession) Validate(message []byte) (bool, error) {
	decrpyted, err := ns.pk.Decrypt(message)
	if err != nil {
		return false, err
	}
	valid := append(ns.pk.prefix, ns.message...)
	result := bytes.Compare(decrpyted, valid)
	if result == 0 {
		return true, nil
	}
	return false, nil
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
