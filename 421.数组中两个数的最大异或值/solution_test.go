package leetcode

import (
	"reflect"
	"testing"
)

// 421.数组中两个数的最大异或值
// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array
func findMaximumXOR(nums []int) int {
	t := trie{}
	for _, num := range nums {
		t.Add(num)
	}
	res := 0
	for _, num := range nums {
		if t := t.MaxXor(num); t > res {
			res = t
		}
	}
	return res
}

type trie struct {
	zero, one *trie
}

const highBin = 30

func (t *trie) Add(num int) {
	cur := t
	for i := highBin; i >= 0; i-- {
		if bit := num >> i & 1; bit == 0 {
			if cur.zero == nil {
				cur.zero = &trie{}
			}
			cur = cur.zero
		} else {
			if cur.one == nil {
				cur.one = &trie{}
			}
			cur = cur.one
		}
	}
}

func (t *trie) MaxXor(num int) (res int) {
	cur := t
	for i := highBin; i >= 0; i-- {
		if bit := num >> i & 1; bit == 0 {
			if cur.one != nil {
				cur = cur.one
				res = res*2 + 1
			} else {
				cur = cur.zero
				res *= 2
			}
		} else {
			if cur.zero != nil {
				cur = cur.zero
				res = res*2 + 1
			} else {
				cur = cur.one
				res *= 2
			}
		}
	}
	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			nums: []int{3, 10, 5, 25, 2, 8},
			want: 28,
		},
		{
			nums: []int{0},
			want: 0,
		},
		{
			nums: []int{2, 4},
			want: 6,
		},
		{
			nums: []int{8, 10, 2},
			want: 10,
		},
		{
			nums: []int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70},
			want: 127,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findMaximumXOR(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
