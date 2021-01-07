package main

import (
	"fmt"
	"reflect"
)

// 509.斐波那契数
// https://leetcode-cn.com/problems/fibonacci-number/

func fib(n int) int {
	if n < 2 {
		return n
	}
	f1, f2 := 0, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

func test() {
	type TestType struct {
		n    int
		want int
	}
	ts := []TestType{
		TestType{},
	}
	for _, t := range ts {
		get := fib(t.n)
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
