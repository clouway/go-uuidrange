// Package uuidrange provides functions for splitting uuid's on ranges. Mostly
// this functions will be used when you want to dispatch job execution on multiple
// workers without job overlapping.
package uuidrange

import (
	"fmt"
	"math/big"
	"strings"
)

// maxuuid value used to determine size of a single part by dividing
// this value to the number of parts.
var maxuuid = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(128), nil)

// Range value represents a single UUID range as string.
type Range struct {
	From string
	To   string
}

type Ranges []Range

// New creates a a slice of Range's of uuid's by splitting them on even parts
// from 00000000-0000-0000-0000-000000000000 to ffffffff-ffff-ffff-ffff-ffffffffffff.
// For example, a split on 2 will produce:
// 00000000-0000-0000-0000-000000000000 7fffffff-ffff-ffff-ffff-ffffffffffff
// 80000000-0000-0000-0000-000000000000 ffffffff-ffff-ffff-ffff-ffffffffffff
func New(n int) Ranges {
	if n < 0 || n > 1024 {
		panic("invalid range")
	}
	parts := big.NewInt(int64(n))
	chunk := big.NewInt(0).Div(maxuuid, parts)

	ranges := make(Ranges, n)
	for i := 0; i < n; i++ {
		low := big.NewInt(0).Mul(chunk, big.NewInt(int64(i)))
		lastInRange := big.NewInt(0).Mul(chunk, big.NewInt(int64(i+1)))
		high := big.NewInt(0).Sub(lastInRange, big.NewInt(1))
		ranges[i] = Range{From: toUUID(low), To: toUUID(high)}
	}
	return ranges
}

func toUUID(n *big.Int) string {
	memo := fmt.Sprintf("%x", n)
	memo = rightPad2Len(memo, "0", 32)
	var sb strings.Builder
	for i := range memo {
		if i == 8 || i == 12 || i == 16 || i == 20 {
			sb.WriteString("-")
		}
		sb.WriteRune(rune(memo[i]))
	}
	return sb.String()
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
