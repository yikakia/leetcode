package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

// 443.压缩字符串
// https://leetcode-cn.com/problems/string-compression
func compress(chars []byte) int {
	ns := ""
	for i := 0; i < len(chars); {
		j := i + 1
		for ; j < len(chars); j++ {
			if chars[j] != chars[i] {
				break
			}
		}
		ns += string(chars[i])
		if t := j - i; t > 1 {
			ns += strconv.Itoa(t)
		}
		i = j
	}
	for i := range ns {
		chars[i] = ns[i]
	}
	return len(ns)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		chars []byte
		want  int
		desc  string
	}{
		{
			chars: []byte("aabbccc"),
			want:  6,
		},
		{
			chars: []byte("a"),
			want:  1,
		},
		{
			chars: []byte("abbbbbbbbbbbb"),
			want:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := compress(tC.chars)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
