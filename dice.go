package dice

import "math/rand"

// Dice is a random number generator with some handy functions.
type Dice struct {
	Rng func(int) int
}

// NewFunc constructs new Dice using given random generator func.
func NewFunc(rng func(int) int) Dice { return Dice{Rng: rng} }

// New constructs new Dice using math.rand
func New() Dice { return NewFunc(rand.Intn) }

var global = New()

// SetGlobal registers Dice to be used for calls on non-instanced dices.
func SetGlobal(d Dice) { global = d }

// SetGlobalFunc registers Dice with given random generator
func SetGlobalFunc(rng func(int) int) { SetGlobal(NewFunc(rng)) }
