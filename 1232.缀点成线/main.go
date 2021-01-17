package main

import (
	"fmt"
	"reflect"
)

// 1232.缀点成线
// https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	c0, c1 := coordinates[0], coordinates[1]
	x0, y0 := c0[0], c0[1]
	x1, y1 := c1[0], c1[1]
	for _, c := range coordinates[2:] {
		x2, y2 := c[0], c[1]
		if (y1-y0)*(x2-x0) != (y2-y0)*(x1-x0) {
			return false
		}
	}

	return true
}
func test() {
	type TestType struct {
		coordinates [][]int
		want        bool
	}
	ts := []TestType{
		TestType{
			coordinates: [][]int{
				[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{6, 7},
			},
			want: true,
		},
		TestType{
			coordinates: [][]int{
				[]int{1, -8}, []int{2, -3}, []int{1, 2},
			},
			want: false,
		},
	}
	for _, t := range ts {
		get := checkStraightLine(t.coordinates)
		if !reflect.DeepEqual(get, t.want) {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")
}
func main() {
	test()
}
