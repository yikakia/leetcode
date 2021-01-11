package main

import (
	"fmt"
	"reflect"
	"sort"
)

// 1202.交换字符串中的元素
// https://leetcode-cn.com/problems/smallest-string-with-swaps/
func smallestStringWithSwaps(s string, pairs [][]int) string {
	if len(pairs) == 0 {
		return s
	}
	n := len(s)

	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(from, to int) {
		fFrom, fTo := find(from), find(to)
		if fFrom == fTo {
			return
		}
		if rank[fTo] < rank[fFrom] {
			fFrom, fTo = fTo, from
		}
		rank[fTo] += rank[from]
		fa[fFrom] = fTo
	}

	for _, pair := range pairs {
		merge(pair[0], pair[1])
	}

	set := make(map[int][]byte)
	for i := range s {
		f := find(i)
		set[f] = append(set[f], s[i])
	}

	for _, bytes := range set {
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
	}

	res := make([]byte, n)
	for i := range res {
		f := find(i)
		res[i] = set[f][0]
		set[f] = set[f][1:]
	}
	return string(res)
}
func test() {
	type TestType struct {
		s     string
		pairs [][]int
		want  string
	}
	ts := []TestType{
		TestType{},
	}
	for _, t := range ts {
		get := smallestStringWithSwaps(t.s, t.pairs)
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
