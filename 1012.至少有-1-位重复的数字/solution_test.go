package leetcode

import (
	"reflect"
	"testing"
)

// 1012.至少有1位重复的数字
// https://leetcode-cn.com/problems/numbers-with-repeated-digits
func numDupDigitsAtMostN(N int) int {
	digits := []int{}
	for tn := N; tn != 0; tn /= 10 {
		digits = append(digits, tn%10)
	}
	n := len(digits)
	total := 0

	for i := 1; i < n; i++ {
		total += 9 * A(9, i-1)
	}

	used := [10]int{}
	for i := n - 1; i >= 0; i-- {
		num := digits[i]
		j := 0
		if i == n-1 {
			j = 1
		}

		for ; j < num; j++ {
			if used[j] != 0 {
				continue
			}
			total += A(10-(n-i), i)
		}
		used[num]++
		if used[num] > 1 {
			break
		}
		if i == 0 {
			total++
		}
	}

	return N - total
}

func A(n, m int) int {
	return fact(n) / fact(n-m)
}

var f map[int]int

func fact(n int) int {
	if f == nil {
		f = map[int]int{}
	}
	if v, ok := f[n]; ok {
		return v
	}
	if n < 2 {
		f[n] = 1
		return 1
	}
	f[n] = n * fact(n-1)
	return f[n]
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		N    int
		want int
	}{
		{
			N:    20,
			want: 1,
		},
		{
			N:    100,
			want: 10,
		},
		{
			N:    1000,
			want: 262,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numDupDigitsAtMostN(tC.N)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
