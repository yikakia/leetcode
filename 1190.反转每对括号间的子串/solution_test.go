package leetcode

import (
	"reflect"
	"testing"
)

// 1190.反转每对括号间的子串
// https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses
func reverseParentheses(s string) string {
	q := []string{""}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			q = append(q, "")
		} else if s[i] == ')' {
			str := reverseStr(q[len(q)-1])
			q = q[:len(q)-1]
			q[len(q)-1] += str
		} else {
			q[len(q)-1] = q[len(q)-1] + string(s[i])
		}
	}

	return q[0]
}

func reverseStr(s string) string {
	bStr := []byte(s)
	for i := 0; i < len(bStr)/2; i++ {
		bStr[i], bStr[len(bStr)-i-1] = bStr[len(bStr)-i-1], bStr[i]
	}
	return string(bStr)
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		want string
	}{
		{
			s:    "",
			want: "",
		},
		{
			s:    "(abcd)",
			want: "dcba",
		},
		{
			s:    "(u(love)i)",
			want: "iloveu",
		},
		{
			s:    "(ed(et(oc))el)",
			want: "leetcode",
		},
		{
			s:    "a(bcdefghijkl(mno)p)q",
			want: "apmnolkjihgfedcbq",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := reverseParentheses(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
