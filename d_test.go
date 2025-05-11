package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD(t *testing.T) {
	counter := -1 // First value from rng func will be 0
	rng := func(i int) int {
		counter++
		return counter % i
	}

	dice := NewFunc(rng)
	for i := 0; i < 100; i++ {
		r := dice.D(11)
		if !assert.Equal(t, i%11+1, r) {
			return
		}
		if r < 1 || r > 11 {
			assert.Fail(t, "unexpected value")
		}
	}
}
