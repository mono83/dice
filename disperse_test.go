package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisperse(t *testing.T) {
	value := 100
	factor := 30

	// Checking lower bound
	assert.Equal(
		t,
		value-factor,
		NewFunc(func(int) int { return 0 }).Disperse(value, factor),
	)

	// Checking upper bound
	assert.Equal(
		t,
		value+factor,
		NewFunc(func(int) int { return factor * 2 }).Disperse(value, factor),
	)

	// Checking mid
	assert.Equal(
		t,
		value,
		NewFunc(func(int) int { return factor }).Disperse(value, factor),
	)
}
