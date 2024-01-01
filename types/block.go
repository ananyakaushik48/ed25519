package types

import (
	"crypto/sha256"

	"github.com/ananyakaushik48/ed25519/proto"
	pb "google.golang.org/protobuf/proto"
)

// This function returns a SHA256 of the header.
func HashBlock(block *proto.Block) []byte {
	b, err := pb.Marshal(block)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	// As sha256 returns an Array we have to convert it into a slice
	return hash[:]
}
