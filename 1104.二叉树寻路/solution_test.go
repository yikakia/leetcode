package leetcode

import (
	"math/bits"
	"reflect"
	"testing"
)

// 1104.二叉树寻路
// https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree
func pathInZigZagTree(label int) []int {
	ret := []int{label}
	if label == 1 {
		return ret
	}

	level := bits.Len(uint(label)) - 1
	upper_max := 1<<level - 1
	upper_min := 1 << (level - 1)

	var father int
	if l, r := label/2-upper_min, upper_max-label/2; l > r {
		father = upper_max - l
	} else {
		father = upper_min + r
	}

	return append(pathInZigZagTree(father), ret...)
}
func getFirstOneIndex(num int) (index int) {
	for i := 20; i >= 0; i -= 1 {
		if num&(1<<i) != 0 {
			return i
		}
	}
	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		label int
		want  []int
		desc  string
	}{
		{
			label: 1,
			want:  []int{1},
		},
		{
			label: 3,
			want:  []int{1, 3},
		},
		{
			label: 14,
			want:  []int{1, 3, 4, 14},
		},
		{
			label: 26,
			want:  []int{1, 2, 6, 10, 26},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := pathInZigZagTree(tC.label)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
