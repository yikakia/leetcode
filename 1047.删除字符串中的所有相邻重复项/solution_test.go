package leetcode

import (
	"reflect"
	"testing"
)

// 1047.删除字符串中的所有相邻重复项
// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string
func removeDuplicates(S string) string {
	stack := make([]byte, 0, len(S)/2)
	for i := range S {
		var changed bool
		for len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
			changed = true
		}
		if !changed {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		S    string
		want string
	}{
		{
			S:    "abbaca",
			want: "ca",
		},
		{
			S:    "a",
			want: "a",
		},
		{
			S:    "aa",
			want: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := removeDuplicates(tC.S)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
