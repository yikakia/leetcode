package leetcode

import (
	"reflect"
	"testing"
)

// 674.最长连续递增序列
// https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/
func findLengthOfLCIS(nums []int) int {

}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want 
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := 
			if !reflect.DeepEqual(tC.want,get){
				t.Errorf("input: %+v get: %v\n",tC,get)
			}
		})
	}
}
