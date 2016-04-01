package gosrp6

import (
	"math/big"
	"testing"
)

func TestIsSafePrime(t *testing.T) {
	i := big.NewInt(2)
	i = i.Exp(i, big.NewInt(256), big.NewInt(0))
	t.Logf("%v", i)
	i = i.Sub(i, big.NewInt(189))
	t.Logf("%v", i)
	p := isSafePrime(i)
	if !p {
		t.Errorf("%v not a safe prime\n", i)
	}
}
