package main

import (
	"fmt"
	"reflect"
)

// 1018.可被5整除的二进制前缀
// https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/
func prefixesDivBy5(A []int) []bool {
	if len(A) == 0 {
		return nil
	}
	res := make([]bool, len(A))
	pre := 0
	for i := range A {
		pre *= 2
		pre += A[i]
		pre %= 10
		if pre == 0 || pre == 5 {
			res[i] = true
		}
	}
	return res
}
func test() {
	type TestType struct {
		A    []int
		want []bool
	}
	ts := []TestType{
		TestType{},
	}
	for _, t := range ts {
		get := prefixesDivBy5(t.A)
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
