package dice

// DiceSlice returns random value from given slice.
// Randomness is determined using given dice.
func DiceSlice[T any](d Dice, values []T) (out T) {
	l := len(values)
	if l == 0 {
		return
	} else if l == 1 {
		return values[0]
	}

	return values[d.Rng(l)]
}

// Slice returns random value from given slice.
// Randomess is determined using global dice.
func Slice[T any](values []T) T {
	return DiceSlice(global, values)
}
