package dice

// D generates random value (like dice roll) within range [1,upper]
// both values are inclusive
func (d Dice) D(upper int) int {
	return d.Rng(upper) + 1
}

// D generates random value (like dice roll) within range [1,upper]
// both values are inclusive
func D(upper int) int {
	return global.D(upper)
}

func (d Dice) D4() int  { return d.D(4) }
func (d Dice) D6() int  { return d.D(6) }
func (d Dice) D8() int  { return d.D(8) }
func (d Dice) D10() int { return d.D(10) }
func (d Dice) D12() int { return d.D(12) }
func (d Dice) D20() int { return d.D(20) }

func D4() int  { return global.D4() }
func D6() int  { return global.D6() }
func D8() int  { return global.D8() }
func D10() int { return global.D10() }
func D12() int { return global.D12() }
func D20() int { return global.D20() }
