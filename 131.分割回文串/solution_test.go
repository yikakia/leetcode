package leetcode

import (
	"reflect"
	"testing"
)

// 131.分割回文串
// https://leetcode-cn.com/problems/palindrome-partitioning

func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]int8, n)
	for i := range f {
		f[i] = make([]int8, n)
	}

	// 0 表示尚未搜索，1 表示是回文串，-1 表示不是回文串
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if f[i][j] != 0 {
			return f[i][j]
		}
		f[i][j] = -1
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		}
		return f[i][j]
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) > 0 {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		want [][]string
	}{
		{
			s: "cabac",
			want: [][]string{
				{"c", "a", "b", "a", "c"},
				{"c", "aba", "c"},
				{"cabac"},
			},
		},
		{
			s: "aab",
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := partition(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
