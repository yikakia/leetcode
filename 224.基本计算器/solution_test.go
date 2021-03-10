package leetcode

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// 224.基本计算器
// https://leetcode-cn.com/problems/basic-calculator

const (
	plus  = '+'
	minus = '-'
	lbr   = '('
	rbr   = ')'
)

var priority = map[byte]int{
	plus:  0,
	minus: 0,
	lbr:   1,
	rbr:   0}

func doOp(a, b int, op byte) int {
	switch op {
	case plus:
		return a + b
	case minus:
		return a - b
	default:
		return 0
	}
}
func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "(-", "(0-")

	ops := []byte{}
	nums := []int{0}

	calc := func() {
		if len(nums) < 2 {
			return
		}
		if len(ops) == 0 {
			return
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
		}
		nums = append(nums, res)
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch == lbr:
			ops = append(ops, ch)
		case ch == rbr:
			for len(ops) > 0 {
				if ops[len(ops)-1] != lbr {
					calc()
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		case ch >= '0' && ch <= '9':
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
					calc()
				} else {
					break
				}
			}
			ops = append(ops, ch)
		}
	}
	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		calc()
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
			s:    "(7)-(0)+(4)",
			want: 11,
		},
		{
			s:    "23",
			want: 23,
		},
		{
			s:    "+21",
			want: 21,
		},
		{
			s:    " 2-1 + 2 ",
			want: 3,
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
