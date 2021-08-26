package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 881.救生艇
// https://leetcode-cn.com/problems/boats-to-save-people
func numRescueBoats(people []int, limit int) (res int) {
	sort.Ints(people)
	l, r := 0, len(people)-1
	for l <= r {
		if people[l]+people[r] > limit {
			r--
			res++
			continue
		}
		l++
		r--
		res++
	}

	return
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		people []int
		limit  int
		want   int
		desc   string
	}{
		{
			people: []int{1, 2},
			limit:  3,
			want:   1,
		},
		{
			people: []int{3, 2, 2, 1},
			limit:  3,
			want:   3,
		},
		{
			people: []int{3, 5, 3, 4},
			limit:  5,
			want:   4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numRescueBoats(tC.people, tC.limit)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
