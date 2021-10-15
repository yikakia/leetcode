package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// 38.外观数列
// https://leetcode-cn.com/problems/count-and-say
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	preRet := countAndSay(n - 1)
	ret := ""
	curCh := preRet[0]
	count := 0
	for i := range preRet {
		ch := preRet[i]
		if ch != curCh {
			ret += fmt.Sprintf("%d%d", count, curCh-'0')
			count = 0
			curCh = ch
		}
		count++
	}

	if count != 0 {
		ret += fmt.Sprintf("%d%d", count, curCh-'0')
	}

	return ret
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		want string
		desc string
	}{
		{
			n:    1,
			want: "1",
		},
		{
			n:    4,
			want: "1211",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := countAndSay(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
