# bitmask

bitmask provides a type-safe and verbose way to conveniently
operate on bit masks.

## How to use

First, to use this library you need to add it to your go module:
```shell
$ go get github.com/kerelape/bitmask
```

### Mask

Functions in this module can work with any `uint64` derivative,
but for convenience there is a predefined type alias called
`Mask`, and it is recommended to use it as the base type for
your own masks.

To start working with a new mask, first you should declare it:
```go
import (
    "github.com/kerelape/bitmask"
)

// MyMask is a new fancy mask type.
type MyMask bitmask.Mask

```

Then declare the flags, that can be set in the mask:
```go
const (
    MyMaskFlag1 bitmask.Flag[MyMask] = 1 << iota
    MyMaskFlag2
    MyMaskFlag3
    ...
    MyMaskFlag64
)
```

Then, to create and instance of the mask, you can utilize
the `New[M](...Flag[M]) M` function:

```go
mask := bitmask.New[MyMask]()

// or with optional initial flags
mask := bitmask.New[MyMask](MyMaskFlag1, MyMaskFlag2)
```

The function simply initializes the mask of type `M` (`MyMask` in
this case), and can be provided with an arbitary amount of initial
flags to be set to the mask.

Then you can work with the mask:
```go
hasFlag1 := bitmask.Has(mask, MyMaskFlag1) // true if the flag is set in the mask
mask = bitmask.Set(mask, MyMaskFlag2) // sets the flag to the mask and returns it
mask = bitmask.Clear(mask, MyMaskFlag1) // unsets the flag in the mask and returns it
```

### Flag

Declare flags:

```go
const (
    Flag1 bitmask.Flag[MyMask] = 1 << iota
    Flag2
    Flag3
    ...
    FlagN
)
```

And use them:

```go
mask := bitmask.New[MyMask](Flag1, Flag2)
mask = bitmask.Set(mask, Flag3)
mask = bitmask.Clear(mask, Flag2)
```
