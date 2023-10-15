package bitmask

// Flag is a bit flag associated with a mask.
type Flag[M ~Mask] uint64

// SetIn sets this flag in the mask.
func (f Flag[M]) SetIn(mask *M) {
	*mask = With(*mask, f)
}

// ClearIn clears this flag in the mask.
func (f Flag[M]) ClearIn(mask *M) {
	*mask = Without(*mask, f)
}

// SwapIn swaps this flag in the mask.
func (f Flag[M]) SwapIn(mask *M) {
	*mask = Swap(*mask, f)
}
