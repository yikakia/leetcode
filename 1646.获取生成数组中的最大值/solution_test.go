package leetcode

import (
	"reflect"
	"testing"
)

// 1646.获取生成数组中的最大值
// https://leetcode-cn.com/problems/get-maximum-in-generated-array
func getMaximumGenerated(n int) int {
	switch n {
	case 0:
		return 0
	}
	nums := make([]int, n+1)
	ans := 0
	nums[1] = 1
	for i := 0; i <= n; i++ {
		t := 2 * i
		if t >= 2 && t <= n {
			nums[t] = nums[i]
			ans = max(ans, nums[t])
		}
		if t+1 >= 2 && t+1 <= n {
			nums[t+1] = nums[i] + nums[i+1]
			ans = max(ans, nums[t+1])
		}
	}
	return max(ans, nums[n])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		want int
		desc string
	}{
		{
			n:    0,
			want: 0,
		},
		{
			n:    1,
			want: 1,
		},
		{
			n:    7,
			want: 3,
		},
		{
			n:    2,
			want: 1,
		},
		{
			n:    3,
			want: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := getMaximumGenerated(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
