package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// builds an evaluator that returns Hit if the index is between low and high (inclusive), TooLow if the index is < low, TooHigh otherwise.
func newEvaluator(low, high int, t *testing.T, first, size int) Evaluator {
	return func(index int) Evaluation {
		t.Helper()
		assert.GreaterOrEqual(t, index, first)
		assert.Less(t, index, first+size)
		index -= first
		switch {
		case index < low:
			return TooLow
		case index <= high:
			return Hit
		}
		return TooHigh
	}
}

// builds an evaluator that returns always the same thing
func newDummyEvaluator(d Evaluation, t *testing.T, first, size int) Evaluator {
	return func(index int) Evaluation {
		t.Helper()
		assert.GreaterOrEqual(t, index, first)
		assert.Less(t, index, first+size)
		return d
	}
}

// builds an evaluator that must never be called
func newNeverCallEvaluatorf(t *testing.T, format string, a ...interface{}) Evaluator {
	return func(int) Evaluation {
		t.Helper()
		t.Errorf(format, a...)
		return Hit
	}
}

func TestHigherOrHitNaively(t *testing.T) {
	for _, tc := range []struct {
		target   int
		ints     []int
		expected int
	}{
		{5, []int{}, -1},
		{5, []int{5}, 0},
		{5, []int{5, 9}, 0},
		{5, []int{9}, -1},
		{5, []int{9, 10}, -1},
		{5, []int{4, 5, 10}, 1},
		{5, []int{2, 3, 5, 6, 7}, 2},
		{5, []int{2, 3, 4, 5, 5, 6, 7}, 3},
		{5, []int{2, 3, 5, 5, 5, 6, 7}, 2},
		{5, []int{2, 3, 4, 5, 5, 5, 6, 7}, 3},
		{5, []int{2, 3, 5, 5, 5, 6, 7}, 2},
	} {
		result := HigherOrHit(0, len(tc.ints), -1, func(i int) Evaluation {
			if tc.ints[i] > tc.target {
				return TooHigh
			}
			if tc.ints[i] < tc.target {
				return TooLow
			}
			return Hit
		})
		assert.Equalf(t, tc.expected, result, "%#v", tc)
	}
}

func TestHigherOrHitExtensively(t *testing.T) {
	// size <= 0
	for first := -3; first <= 3; first++ {
		for size := 0; size >= -3; size-- {
			assert.Equal(t, -12, HigherOrHit(first, size, -12, newNeverCallEvaluatorf(t, "the evaluator should never be called if the size is 0 (first = %d)", first)))
		}
	}

	// for all indexes, the evaluator returns the same Evaluation
	for size := 1; size < 8; size++ {
		for first := -3; first <= 3; first++ {
			assert.Equal(t, -12, HigherOrHit(first, size, -12, newDummyEvaluator(TooHigh, t, first, size)))
			assert.Equal(t, first+size-1, HigherOrHit(first, size, -12, newDummyEvaluator(TooLow, t, first, size)))
			assert.Equal(t, first, HigherOrHit(first, size, -12, newDummyEvaluator(Hit, t, first, size)))
		}
	}

	// some indexes are too low, some too high, no exact match
	for tooLow := 1; tooLow < 8; tooLow++ {
		for tooHigh := 1; tooHigh < 8; tooHigh++ {
			size := tooLow + tooHigh
			for first := -3; first <= 3; first++ {
				e := newEvaluator(tooLow+1, tooLow, t, first, size)
				assert.Equal(t, first+tooLow, HigherOrHit(first, size, -12, e))
			}
		}
	}

	// there are exact matches, and 0 to N1 too low indexes, and 0 to N2 too high indexes
	for tooLow := 0; tooLow < 8; tooLow++ {
		for matches := 1; matches < 8; matches++ {
			for tooHigh := 0; tooHigh < 8; tooHigh++ {
				size := tooLow + matches + tooHigh
				for first := -3; first <= 3; first++ {
					e := newEvaluator(tooLow, tooLow+matches, t, first, size)
					assert.Equal(t, first+tooLow, HigherOrHit(first, size, -12, e))
				}
			}
		}
	}
}
