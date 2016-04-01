package gosrp6

import "math/big"

func isSafePrime(i *big.Int) bool {
	q := big.NewInt(32)
	q = q.Sub(i, big.NewInt(1))
	q.Div(q, big.NewInt(2))
	return i.ProbablyPrime(40) && q.ProbablyPrime(40)
}
