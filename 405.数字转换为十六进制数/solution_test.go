package leetcode

import (
	"reflect"
	"strings"
	"testing"
)

// 405.数字转换为十六进制数
// https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal
func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := &strings.Builder{}
	for i := 7; i >= 0; i-- {
		val := num >> (4 * i) & 0xf
		if val > 0 || sb.Len() > 0 {
			sb.WriteByte(getByte(val))
		}
	}
	return sb.String()
}

func getByte(a int) (ret byte) {
	if a < 10 {
		ret = '0' + byte(a)
	} else {
		ret = 'a' + byte(a-10)
	}
	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		num  int
		want string
		desc string
	}{
		{
			num:  26,
			want: "1a",
		},
		{
			num:  -1,
			want: "ffffffff",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := toHex(tC.num)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
