package leetcode

import (
	"reflect"
	"testing"
)

// 1269.停在原地的方案数
// https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps
func numWays(steps int, arrLen int) int {
	if steps == 1 {
		return 1
	}
	sum := make([]int, arrLen)
	tmpSum := make([]int, arrLen)
	sum[0] = 1
	for i := 0; i < steps; i++ {
		copy(tmpSum, sum)
		for j := 0; j < arrLen; j++ {
			if j > 0 {
				if sum[j-1] == 0 {
					continue
				}
				tmpSum[j] += sum[j-1]
				tmpSum[j] %= mod
			}
			if j < arrLen-1 {
				if sum[j+1] == 0 {
					continue
				}
				tmpSum[j] += sum[j+1]
				tmpSum[j] %= mod
			}
		}
		copy(sum, tmpSum)
	}
	return sum[0]
}

const mod = 1000000007

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		steps  int
		arrLen int
		want   int
	}{
		{
			steps:  3,
			arrLen: 2,
			want:   4,
		},
		{
			steps:  2,
			arrLen: 4,
			want:   2,
		},
		{
			steps:  4,
			arrLen: 2,
			want:   8,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numWays(tC.steps, tC.arrLen)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
