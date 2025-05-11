Dice
====

Simple yet efficient library for randomness with some syntax sugar.

Provides `dice.Dice` abstraction, able to provide random dice rolls for standard dices like
`D4`, `D6`, `D8`, `D10`, `D12`,`D20` and arbitrary `D(int)`.

New `dice.Dice` instance can be created using any random source (even deterministic for
unit test). All API from `dice.Dice` also available on package level - using global `Dice` 
instance that uses `math/rand`

## Basic usage

```go
import "github.com/mono83/dice"

// Create new dice with math/rand source
d := dice.New()
// Obtain random value in range [1,6] inclusive
x := d.D6()

// WIthout dice creation
y := dice.D6()
```

## Advanced usage

```go
import "github.com/mono83/dice"

// Create dice with crypto/rand source
cd := dice.NewCryptoRand()

// Create dice with deterministic source
dd := dice.NewFunc(func(int) int { return 0 })
```
