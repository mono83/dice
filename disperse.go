package dice

// Disperse performs dispersion of given value for amount of factor
// For example given value=100 and factor=10 will produce
// random values in range [90, 110] (both inclusive)
func (d Dice) Disperse(value, factor int) int {
	variance := value * factor / 100
	return value - variance + d.Rng(variance*2+1)
}

// Disperse performs dispersion of given value for amount of factor
// For example given value=100 and factor=10 will produce
// random values in range [90, 110] (both inclusive)
func Disperse(value, factor int) int {
	return global.Disperse(value, factor)
}
