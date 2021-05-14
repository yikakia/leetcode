package leetcode

import (
	"reflect"
	"testing"
)

// 12.整数转罗马数字
// https://leetcode-cn.com/problems/integer-to-roman
func intToRoman(num int) string {
	res := ""
	for num != 0 {
		switch {
		case num >= 1000:
			num -= 1000
			res += "M"
		case num >= 900:
			num -= 900
			res += "CM"
		case num >= 500:
			num -= 500
			res += "D"
		case num >= 400:
			num -= 400
			res += "CD"
		case num >= 100:
			num -= 100
			res += "C"
		case num >= 90:
			num -= 90
			res += "XC"
		case num >= 50:
			num -= 50
			res += "L"
		case num >= 40:
			num -= 40
			res += "XL"
		case num >= 10:
			num -= 10
			res += "X"
		case num >= 9:
			num -= 9
			res += "IX"
		case num >= 5:
			num -= 5
			res += "V"
		case num >= 4:
			num -= 4
			res += "IV"
		case num >= 1:
			num -= 1
			res += "I"
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		num  int
		want string
	}{
		{
			num:  3,
			want: "III",
		},
		{
			num:  4,
			want: "IV",
		},
		{
			num:  9,
			want: "IX",
		},
		{
			num:  58,
			want: "LVIII",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := intToRoman(tC.num)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
