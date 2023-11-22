package crypto

import "crypto/ed25519"

const (
	privKeyLen = 64
	pubKeyLen  = 32
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

// This is because the private key of the ed25519 crypto package is an array of bytes, hence we create a Bytes function that takes a input which is the pointer the private
// key of which returns a byte array of the private key
func (p *PrivateKey) Bytes() []byte {
	return p.key
}

// This the function that uses the private key to sign a message and return the signed message bytes array
func (p *PrivateKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}

func (p *PrivateKey) Public() *PublicKey {
	b := make([]byte, pubKeyLen)
	copy(b, p.key[32:])

	return &PublicKey{
		key:b,
	}
}

type PublicKey struct {
	key ed25519.PublicKey
}
