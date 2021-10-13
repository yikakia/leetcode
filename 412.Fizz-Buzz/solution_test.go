package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

// 412.FizzBuzz
// https://leetcode-cn.com/problems/fizz-buzz
func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := range ret {
		num := i + 1
		if (num)%3 != 0 && (num)%5 != 0 {
			ret[i] = fmt.Sprint(i + 1)
		}
		if (num)%3 == 0 {
			ret[i] = "Fizz"
		}
		if (num)%5 == 0 {
			ret[i] += "Buzz"
		}
	}
	return ret
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		n    int
		want []string
		desc string
	}{
		{
			n: 15,
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := fizzBuzz(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
