package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

// 331.验证二叉树的前序序列化
// https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree
func isValidSerialization(preorder string) bool {
	split := strings.Split(preorder, ",")
	freeNodes := 1
	for _, s := range split[:] {
		if freeNodes == 0 {
			return false
		}
		if s != "#" {
			freeNodes++
		} else {
			freeNodes--
		}
	}
	return freeNodes == 0
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		want     bool
		preorder string
	}{
		{
			preorder: "#",
			want:     true,
		},
		{
			preorder: "#,7,6,9,#,#,#",
			want:     false,
		},
		{
			preorder: "9,3,4,#,#,1,#,#,2,#,6,#,#",
			want:     true,
		},
		{
			preorder: "1,#",
			want:     false,
		},
		{
			preorder: "9,#,#,1",
			want:     false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := isValidSerialization(tC.preorder)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
