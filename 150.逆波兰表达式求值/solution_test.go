package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

// 150.逆波兰表达式求值
// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation

func evalRPN(tokens []string) int {
	nums := []int{}
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err != nil {
			a, b := nums[len(nums)-2], nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			switch token {
			case "+":
				a += b
			case "-":
				a -= b
			case "*":
				a *= b
			case "/":
				a /= b
			}
			nums[len(nums)-1] = a
		} else {
			nums = append(nums, num)
		}
	}

	return nums[0]
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := evalRPN(tC.tokens)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
