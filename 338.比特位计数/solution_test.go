package leetcode

import (
	"reflect"
	"testing"
)

// 338.比特位计数
// https://leetcode-cn.com/problems/counting-bits
func countBits(num int) (res []int) {
	res = make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		num  int
		want []int
	}{
		{
			num:  2,
			want: []int{0, 1, 1},
		},
		{
			num:  5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := countBits(tC.num)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
