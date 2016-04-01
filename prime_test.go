package gosrp6

import (
	"math/big"
	"testing"
)

func TestIsSafePrime(t *testing.T) {
	primes := []int64{5, 7, 11, 23, 47, 59, 83, 107, 167, 179, 227, 263, 347, 359, 383, 467, 479, 503, 563, 587, 719, 839, 863, 887, 983, 1019, 1187, 1283, 1307, 1319, 1367, 1439, 1487, 1523, 1619, 1823, 1907, 2027, 2039, 2063, 2099, 2207, 2447, 2459, 2579, 2819, 2879, 2903}

	for _, q := range primes {
		i := big.NewInt(q)
		p := isSafePrime(i)
		if !p {
			t.Errorf("%v not a safe prime\n", i)
		}
	}
}

func TestIsSGPrime(t *testing.T) {
	primes := []int64{2, 3, 5, 11, 23, 29, 41, 53, 83, 89, 113, 131, 173, 179, 191, 233, 239, 251, 281, 293, 359, 419, 431, 443, 491, 509, 593, 641, 653, 659, 683, 719, 743, 761, 809, 911, 953, 1013, 1019, 1031, 1049, 1103, 1223, 1229, 1289, 1409, 1439, 1451, 1481, 1499, 1511, 1559}
	for _, q := range primes {
		i := big.NewInt(q)
		p := isSGPrime(i)
		if !p {
			t.Errorf("%v not a Sophie Germain prime\n", i)
		}
	}
}
