package leetcode

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// 179.最大数
// https://leetcode-cn.com/problems/largest-number
func largestNumber(nums []int) string {
	strNums := make([]string, 0, len(nums))
	for _, num := range nums {
		strNums = append(strNums, strconv.Itoa(num))
	}
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})
	if strNums[0] == "0" {
		return "0"
	}
	return strings.Join(strNums, "")
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want string
	}{
		{
			nums: []int{3, 3, 3},
			want: "333",
		},
		{
			nums: []int{0, 0, 0},
			want: "0",
		},
		{
			nums: []int{10, 2},
			want: "210",
		},
		{
			nums: []int{3, 30, 34, 5, 9},
			want: "9534330",
		},
		{
			nums: []int{1},
			want: "1",
		},
		{
			nums: []int{10},
			want: "10",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := largestNumber(tC.nums)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
