// Package binarysearch provides binary search functionality.
// Example use, assuming sortedFloats is a []float64 sorted smallest first:
// i := binarysearch.TooLowOrHit(0, len(sortedFloats), -1, func(index int) binarysearch.Evaluation {
//     x := sortedFloats[index]
//     if x > 5.5 { return binarysearch.TooHigh }
//     if x < 5.5 { return binarysearch.TooLow }
//     return binarysearch.Hit
// })
// if i == -1 { fmt.Println("all floats are > 5.5") }
// else if sortedFloats[i] < 5.5 { fmt.Println("no float is 5.5 and the largest index of a float < 5.5 is", i) }
// else { fmt.Println("the lowest index of a float equal to 5.5 is", i) }
package binarysearch

// Evaluation indicates whether an index is too low given the target, too high, or is a hit.
type Evaluation int

const (
	// Evaluation constants.
	TooLow  = Evaluation(iota) // if there's a hit, it's at a higher index
	TooHigh = Evaluation(iota) // if there's a hit, it's at a lower index
	Hit     = Evaluation(iota) // this is a hit
)

// Evaluator is a caller supplied function used by the binary search.
// Its result controls which half of the indexes search space will be explored next.
type Evaluator func(index int) Evaluation

// TooLowOrHit executes a binary search using indexes in range [first, first+size-1).
// If some of the indexes are exact matches (ie. the Evaluator returns Hit), it returns the lowest of those indexes.
// Otherwise, it returns the highest of the indexes for which the Evaluator returns TooLow, if there's at least one of these.
// Otherwise, it returns its missIndex argument.
// If size is negative or 0, it returns missIndex.
func TooLowOrHit(first, size, missIndex int, evaluator Evaluator) int {
	if size <= 1 {
		if size <= 0 || evaluator(first) == TooHigh {
			return missIndex
		}
		return first
	}
	beyond := first + size
	a, b := first, beyond
	for {
		i := (a + b) / 2
		switch evaluator(i) {
		case TooLow:
			if b-i <= 1 {
				if b < beyond && evaluator(b) == Hit {
					return b
				}
				return i
			}
			a = i
		case TooHigh:
			if i-a <= 1 {
				if a <= first && evaluator(first) == TooHigh {
					return missIndex
				}
				return a
			}
			b = i
		case Hit:
			if i-a <= 1 {
				if a >= i || evaluator(a) == Hit {
					return a
				}
				return i
			}
			b = i
		}
	}
}
