package main

import (
	"fmt"
	"reflect"
)

// 1720.解码异或后的数组
// https://leetcode-cn.com/problems/decode-xored-array/
func decode(encoded []int, first int) []int {
	n := len(encoded)
	if n == 0 {
		return []int{}
	}
	res := make([]int, n+1)
	res[0] = first
	for i := range encoded {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}
func test() {
	type TestType struct {
		encoded []int
		first   int
		want    []int
	}
	ts := []TestType{}
	for _, t := range ts {
		get := decode(t.encoded, t.first)
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
