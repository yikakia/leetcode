package leetcode

import (
	"reflect"
	"testing"
)

// 1109.航班预订统计
// https://leetcode-cn.com/problems/corporate-flight-bookings
func corpFlightBookings(bookings [][]int, n int) []int {
	answer := make([]int, n+2)
	for _, booking := range bookings {
		from, to, orders := booking[0], booking[1], booking[2]
		answer[from] += orders
		answer[to+1] -= orders
	}
	for i := 1; i <= n; i++ {
		answer[i] += answer[i-1]
	}
	return answer[1 : len(answer)-1]
}
func TestSolution(t *testing.T) {
	testCases := []struct {
		bookings [][]int
		n        int
		want     []int
		desc     string
	}{
		{
			bookings: [][]int{
				{1, 2, 10},
				{2, 3, 20},
				{2, 5, 25},
			},
			n:    5,
			want: []int{10, 55, 45, 25, 25},
		},
		{
			bookings: [][]int{
				{1, 2, 10},
				{2, 2, 15},
			},
			n:    2,
			want: []int{10, 25},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := corpFlightBookings(tC.bookings, tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
