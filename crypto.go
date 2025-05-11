package dice

import (
	"crypto/rand"
	"math/big"
)

// NewCryptoRand creates dice using crypto/rand package
// It will be 10x times slower that math/rand but provide
// less predictable values.
func NewCryptoRand() Dice {
	return NewFunc(func(i int) int {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}

		return int(n.Int64())
	})
}
