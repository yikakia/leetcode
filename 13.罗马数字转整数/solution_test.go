package leetcode

import (
	"reflect"
	"testing"
)

// 13.罗马数字转整数
// https://leetcode-cn.com/problems/roman-to-integer

var m = map[string]int{
	"M": 1000, "CM": 900, "D": 500, "CD": 400,
	"C": 100, "XC": 90, "L": 50, "XL": 40,
	"X": 10, "IX": 9, "V": 5, "IV": 4,
	"I": 1,
}

func romanToInt(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		if len(s)-i > 1 {
			if v, ok := m[s[i:i+2]]; ok {
				res += v
				i++
				continue
			}
		}
		res += m[s[i:i+1]]
	}

	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		want int
	}{
		{
			s:    "III",
			want: 3,
		},
		{
			s:    "IV",
			want: 4,
		},
		{
			s:    "IX",
			want: 9,
		},
		{
			s:    "LVIII",
			want: 58,
		},
		{
			s:    "MCMXCIV",
			want: 1994,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := romanToInt(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
