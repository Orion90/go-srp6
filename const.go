package gosrp6

import (
	"crypto/sha256"
	"io"
	"math/big"
)

type Constant *big.Int

//init creates the N
func init() {
	N.SetString("009d4eaf224f51a59d343f1d0f32562e15b0785963ab99be264478d6a0eef17445f476fa006fe2f0e6bc7a8c50c2781e1737a8805eed1d7ade902965013dfd64732596658c5b3c8b067e2f74d0dc1bd7ccf98f835d72f644db64a12b4b85a9a3dcafd6adf13f678af770757bb7f390eecd470a68eafb186ac2ffdfaf10c26d5593", 16)
	kh := sha256.New()
	io.WriteString(kh, N.String())
	io.WriteString(kh, "2")
	k.SetBytes(kh.Sum(nil))
}

const (
	g = 2
)

var (
	k = new(big.Int)
	N = new(big.Int)
)
