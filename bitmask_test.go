package bitmask_test

import (
	"testing"

	"github.com/kerelape/bitmask"
	"github.com/stretchr/testify/assert"
)

type Mask bitmask.Mask

const (
	EmptyMask Mask               = 0x0000_0000_0000_0000_0000_0000_0000_0000
	EmptyFlag bitmask.Flag[Mask] = 0x0000_0000_0000_0000_0000_0000_0000_0000
)

const (
	Flag1 bitmask.Flag[Mask] = 1 << iota
	Flag2
)

func TestNew(t *testing.T) {
	t.Run("Without initial flags", func(t *testing.T) {
		assert.Equal(
			t,
			EmptyMask,
			bitmask.New[Mask](),
			"New without initial flags must return an empty mask",
		)
	})
	t.Run("With empty initial flag", func(t *testing.T) {
		assert.Equal(
			t,
			(Mask)(EmptyFlag),
			bitmask.New[Mask](EmptyFlag),
			"New with an empty initial flag should return an empty mask",
		)
	})
	t.Run("With one initial flag", func(t *testing.T) {
		assert.Equal(
			t,
			(Mask)(Flag1),
			bitmask.New[Mask](Flag1),
			"New with one initial flag must return a mask matching that flag",
		)
	})
	t.Run("With several initial flags", func(t *testing.T) {
		assert.Equal(
			t,
			bitmask.Set(
				bitmask.Set(
					bitmask.New[Mask](),
					Flag1,
				),
				Flag2,
			),
			bitmask.New[Mask](Flag1, Flag2),
			"New with several initial flags must returns a mask containing all the flags",
		)
	})
}

func TestSet(t *testing.T) {
	t.Run("With empty flag", func(t *testing.T) {
		assert.Equal(
			t,
			(Mask)(EmptyFlag),
			bitmask.Set(
				bitmask.New[Mask](),
				EmptyFlag,
			),
			"Set with empty flag must return the original mask",
		)
	})
	t.Run("With the existing flag", func(t *testing.T) {
		assert.Equal(
			t,
			(Mask)(Flag1),
			bitmask.Set(
				bitmask.New[Mask](Flag1),
				Flag1,
			),
			"Set with an existing flag must return the original mask",
		)
	})
	t.Run("New flag", func(t *testing.T) {
		assert.Equal(
			t,
			(Mask)(Flag1),
			bitmask.Set(
				bitmask.New[Mask](),
				Flag1,
			),
			"Set with a new flag on an empty mask must return a mask matching the flag",
		)
		assert.Equal(
			t,
			bitmask.New[Mask](Flag1, Flag2),
			bitmask.Set(
				bitmask.New[Mask](Flag1),
				Flag2,
			),
			"Set with a new flag on a non-empty mask must return the original mask including the new flag",
		)
	})
}

func TestClear(t *testing.T) {
	t.Run("With empty flag", func(t *testing.T) {
		assert.Equal(
			t,
			bitmask.New[Mask](Flag1, Flag2),
			bitmask.Clear(
				bitmask.New[Mask](Flag1, Flag2),
				EmptyFlag,
			),
			"Clear with empty flag must not modify the mask",
		)
	})
	t.Run("With unset flag", func(t *testing.T) {
		assert.Equal(
			t,
			bitmask.New[Mask](Flag2),
			bitmask.Clear(
				bitmask.New[Mask](Flag2),
				Flag1,
			),
			"Clear with an unset flag must not modify the mask",
		)
	})
	t.Run("With empty mask", func(t *testing.T) {
		assert.Equal(
			t,
			EmptyMask,
			bitmask.Clear(
				EmptyMask,
				Flag1,
			),
			"Clear with an empty mask must not modify it",
		)
	})
	t.Run("With real flag", func(t *testing.T) {
		assert.Equal(
			t,
			EmptyMask,
			bitmask.Clear(
				bitmask.New[Mask](Flag1),
				Flag1,
			),
			"Clear with a flag on a mask with that flag must clear the flag",
		)
	})
}
