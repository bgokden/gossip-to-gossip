package negotiation

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"io"
)

type NegotiationSession struct {
	pk      *PredefinedKey
	message []byte
}

type PredefinedKey struct {
	prefix  []byte
	private *rsa.PrivateKey
}

func NewPredefinedKey(private *rsa.PrivateKey, prefix []byte) *PredefinedKey {
	return &PredefinedKey{
		prefix:  prefix,
		private: private,
	}
}

func NewRandomPredefinedKey() *PredefinedKey {
	reader := rand.Reader
	bitSize := 2048

	key, _ := rsa.GenerateKey(reader, bitSize)

	prefix := RandomBytes(50)

	return NewPredefinedKey(key, prefix)
}

func (pk *PredefinedKey) EncryptWithPrefix(message []byte) ([]byte, error) {
	message = append(pk.prefix, message...)
	return pk.Encrypt(message)
}

func (pk *PredefinedKey) Encrypt(message []byte) ([]byte, error) {
	sha1 := sha1.New()
	encrypted, err := rsa.EncryptOAEP(sha1, rand.Reader, &pk.private.PublicKey, message, nil)
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

func (pk *PredefinedKey) Decrypt(encrypted []byte) ([]byte, error) {
	sha1 := sha1.New()
	decrypted, err := rsa.DecryptOAEP(sha1, rand.Reader, pk.private, encrypted, nil)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}

func (pk *PredefinedKey) CreateNegotiationSession() *NegotiationSession {
	message := RandomBytes(50)
	return &NegotiationSession{
		pk:      pk,
		message: message,
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

func (ns *NegotiationSession) ValidateWithPrefix(prefix []byte, message []byte) (bool, error) {
	decrpyted, err := ns.pk.Decrypt(message)
	if err != nil {
		return false, err
	}
	valid := append(prefix, ns.message...)
	result := bytes.Compare(decrpyted, valid)
	if result == 0 {
		return true, nil
	}
	return false, nil
}

// RandomBytes generates cryptographically secure pseudorandom numbers byte array
func RandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b) // ignore error
	return b
}

func createHashMD5(key []byte) []byte {
	hasher := md5.New()
	hasher.Write(key)
	return hasher.Sum(nil)
}

func (ns *NegotiationSession) encryptAES(data []byte, passphrase []byte) ([]byte, error) {
	block, err := aes.NewCipher(createHashMD5(passphrase))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func (ns *NegotiationSession) decryptAES(data []byte, passphrase []byte) ([]byte, error) {
	block, err := aes.NewCipher(createHashMD5(passphrase))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
