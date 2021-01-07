package main

import (
	"fmt"
	"reflect"
)

// 830.较大分组的位置
// https://leetcode-cn.com/problems/positions-of-large-groups/

func largeGroupPositions(s string) [][]int {
	res := [][]int{}
	s += "0"
	curRune := '0'
	count := 1
	for i, ch := range s {
		if ch != curRune {
			if count >= 3 {
				res = append(res, []int{i - count, i - 1})
			}
			curRune = ch
			count = 1
			continue
		} else {
			count++
		}
	}

	return res
}
func test() {
	type TestType struct {
		s    string
		want [][]int
	}
	ts := []TestType{
		TestType{
			s: "aaa",
			want: [][]int{
				[]int{0, 2}},
		},
		TestType{
			s: "abbxxxxzzy",
			want: [][]int{
				[]int{3, 6},
			},
		},
		TestType{
			s:    "abc",
			want: [][]int{},
		},
		TestType{
			s: "abcdddeeeeaabbbcd",
			want: [][]int{
				[]int{3, 5},
				[]int{6, 9},
				[]int{12, 14},
			},
		},
		TestType{
			s:    "aba",
			want: [][]int{},
		},
		TestType{
			s: "aaaabbbb",
			want: [][]int{
				[]int{0, 3},
				[]int{4, 7},
			},
		},
	}
	for _, t := range ts {
		get := largeGroupPositions(t.s)
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
