// Package bitmask provides type-safe functionality
// to operate on bit masks in more convenient way, than
// bitwise operators on simple integer types, making
// the client code more readable, and, thus, more
// maintanable.
package bitmask

// Mask is a shorthand for uint64.
type Mask = uint64

// Flag is a bit flag associated with a mask.
type Flag[M ~Mask] uint64

// New returns a new mask with flags applied to it.
func New[M ~Mask](flags ...Flag[M]) M {
	var mask M
	for _, flag := range flags {
		mask = Set(mask, flag)
	}
	return mask
}

// Merge returns a with b bits set.
//
// It's equivalent to a | b.
func Merge[M ~Mask](a, b M) M {
	return a | b
}

// Screen returns a without the bits, that are not
// set in b.
//
// It's equivalent to a & b.
func Screen[M ~Mask](a, b M) M {
	return a & b
}

// Invert returns mask with all the bits inverted.
//
// It's equivalent to ^mask.
func Invert[M ~Mask](mask M) M {
	return ^mask
}

// Subtract returns a without b bits.
//
// It's equivalent to a &^ b.
func Subtract[M ~Mask](a, b M) M {
	return Screen(a, Invert(b))
}

// Toggle returns a XORed by b.
func Toggle[M ~Mask](a, b M) M {
	return a ^ b
}

// Set returns the mask with the flag set.
func Set[M ~Mask](mask M, flag Flag[M]) M {
	return Merge(mask, (M)(flag))
}

// Clear returns the mask without the flag.
func Clear[M ~Mask](mask M, flag Flag[M]) M {
	return Subtract(mask, (M)(flag))
}

// Swap returns the mask with the flag swapped.
func Swap[M ~Mask](mask M, flag Flag[M]) M {
	return Toggle(mask, (M)(flag))
}

// Has reports whether the flag is set in mask.
func Has[M ~Mask](mask M, flag Flag[M]) bool {
	return IsExactly(Screen(mask, (M)(flag)), flag)
}

// IsExactly reports whether the mask exactly matches the flag.
func IsExactly[M ~Mask](mask M, flag Flag[M]) bool {
	return mask == (M)(flag)
}
