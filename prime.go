package gosrp6

import (
	"log"
	"math/big"
)

func isSafePrime(q *big.Int) bool {
	log.Printf("%v", q)
	part := q.Sub(q, big.NewInt(1))
	log.Printf("%v", q)
	part = part.Div(part, big.NewInt(2))
	return q.ProbablyPrime(4) && part.ProbablyPrime(4)
}
