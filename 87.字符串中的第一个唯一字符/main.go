package main

import (
	"fmt"
)

// 87.字符串中的第一个唯一字符
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	arr := [26]int{}

	for i, v := range s {
		arr[v-'a'] = i
	}

	for i, v := range s {
		if arr[v-'a'] == i {
			return i
		}
		arr[v-'a'] = -1
	}

	return -1
}

func test() {
	type TestType struct {
		s    string
		want int
	}
	ts := []TestType{
		TestType{
			s:    "leetcode",
			want: 0,
		},
		TestType{
			s:    "loveleetcode",
			want: 2,
		},
	}
	for _, t := range ts {
		get := firstUniqChar(t.s)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")

}
func main() {
	test()
}
