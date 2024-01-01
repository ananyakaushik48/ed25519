package util

import (
	"crypto/rand"
	"io"
	mrand "math/rand"
	"time"

	"github.com/ananyakaushik48/ed25519/proto"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(rand.Reader, hash)
	return hash
}

func RandomBlock() *proto.Block {
	header := &proto.Header{
		Version:   1,
		Height:    int32(mrand.Intn(1000)),
		PrevHash:  RandomHash(),
		RootHash:  RandomHash(),
		Timestamp: time.Now().UnixNano(),
	}

	return &proto.Block{
		Header: header,
	}
}
