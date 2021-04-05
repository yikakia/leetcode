package leetcode

import (
	"reflect"
	"testing"
)

// 781.森林中的兔子
// https://leetcode-cn.com/problems/rabbits-in-forest

func numRabbits(answers []int) int {
	res := 0
	m := map[int]int{}
	for _, num := range answers {
		m[num]++
	}
	for k, v := range m {
		if k+1 >= v {
			res += k + 1
		} else {
			nogroups := v % (k + 1)
			if nogroups > 0 {
				res += v - nogroups + k + 1
			} else if nogroups == 0 {
				res += v
			}
		}
	}
	return res
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		answers []int
		want    int
	}{
		{
			answers: []int{1, 1, 2},
			want:    5,
		},
		{
			answers: []int{10, 10, 10},
			want:    11,
		},
		{
			answers: []int{},
			want:    0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numRabbits(tC.answers)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
