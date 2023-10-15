package bitmask_test

import (
	"testing"

	"github.com/kerelape/bitmask"
	"github.com/stretchr/testify/assert"
)

type TestMask bitmask.Mask

const (
	TestMaskFlagNone bitmask.Flag[TestMask] = 1<<iota - 1
	TestMaskFlag1
)

func TestFlag_SetIn(t *testing.T) {
	mask := bitmask.Init[TestMask]()
	TestMaskFlag1.SetIn(&mask)
	assert.True(t, bitmask.IsExactly(mask, TestMaskFlag1))
}

func TestFlag_ClearIn(t *testing.T) {
	mask := bitmask.Init[TestMask](TestMaskFlag1)
	TestMaskFlag1.ClearIn(&mask)
	assert.False(t, bitmask.Has(mask, TestMaskFlag1))
	assert.True(t, bitmask.IsExactly(mask, TestMaskFlagNone))
}

func TestFlag_SwapIn(t *testing.T) {
	mask := bitmask.Init[TestMask](TestMaskFlag1)
	assert.True(t, bitmask.Has(mask, TestMaskFlag1))
	TestMaskFlag1.SwapIn(&mask)
	assert.False(t, bitmask.Has(mask, TestMaskFlag1))
	TestMaskFlag1.SwapIn(&mask)
	assert.True(t, bitmask.Has(mask, TestMaskFlag1))
}
