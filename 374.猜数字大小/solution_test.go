package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 374.猜数字大小
// https://leetcode-cn.com/problems/guess-number-higher-or-lower
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

var target int

func guess(num int) int {
	if num == target {
		return 0
	}
	if num > target {
		return -1
	}
	return 1
}

func guessNumber(n int) int {
	return sort.Search(n, func(i int) bool { return guess(i) <= 0 })
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		pick int
		want int
		desc string
	}{
		{
			n:    10,
			pick: 6,
			want: 6,
		},
		{
			n:    1,
			pick: 1,
			want: 1,
		},
		{
			n:    2,
			pick: 1,
			want: 1,
		},
		{
			n:    2,
			pick: 2,
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			target = tC.pick
			get := guessNumber(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
