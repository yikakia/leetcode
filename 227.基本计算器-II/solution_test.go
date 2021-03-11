package leetcode

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// 227.基本计算器II
// https://leetcode-cn.com/problems/basic-calculator-ii

const (
	plus  = '+'
	minus = '-'
	mul   = '*'
	div   = '/'
	lbr   = '('
	rbr   = ')'
)

var priority = map[byte]int{
	plus:  0,
	minus: 0,
	mul:   1,
	div:   1,
	lbr:   3,
	rbr:   4}

func calc(nums []int, ops []byte) ([]int, []byte) {
	if len(nums) < 2 {
		return nums, ops
	}
	if len(ops) == 0 {
		return nums, ops
	}
	a, b := nums[len(nums)-2], nums[len(nums)-1]
	nums = nums[:len(nums)-2]
	op := ops[len(ops)-1]
	ops = ops[:len(ops)-1]
	res := 0
	switch op {
	case plus:
		res = a + b
	case minus:
		res = a - b
	case mul:
		res = a * b
	case div:
		res = a / b
	case lbr:
		panic("should not")
	case rbr:
		calced := false
		for len(ops) > 0 {
			if ops[len(ops)-1] != lbr {
				nums = append(nums, a, b)
				nums, ops = calc(nums, ops)
				a, b = nums[len(nums)-2], nums[len(nums)-1]
				calced = true
			} else {
				if !calced {
					nums = append(nums, a, b)
				}
				ops = ops[:len(ops)-1]
				return nums, ops
			}
		}
	}
	nums = append(nums, res)
	return nums, ops
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "(-", "(0-")

	ops := []byte{}
	nums := []int{0}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case unicode.IsDigit(rune(ch)):
			var j = i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			t, _ := strconv.Atoi(s[i:j])
			nums = append(nums, t)
			i = j - 1
		default:
			for len(ops) > 0 && ops[len(ops)-1] != lbr {
				op := ops[len(ops)-1]
				if priority[op] >= priority[ch] {
					nums, ops = calc(nums, ops)
				} else {
					break
				}
			}
			ops = append(ops, ch)
		}
	}
	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		nums, ops = calc(nums, ops)
	}
	return nums[len(nums)-1]
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		s    string
		want int
	}{
		{
			s:    "(((10)+20)*(30-35))",
			want: -150,
		},
		{
			s:    "(1+2)",
			want: 3,
		},
		{
			s:    "(10)+(2*(-1))",
			want: 8,
		},
		{
			s:    "3+2*2",
			want: 7,
		},
		{
			s:    " 3/2 ",
			want: 1,
		},
		{
			s:    " 3+5 / 2 ",
			want: 5,
		},
		{
			s:    "-2+3/5",
			want: -2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := calculate(tC.s)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
