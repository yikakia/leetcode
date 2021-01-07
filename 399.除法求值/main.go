package main

import (
	"fmt"
	"reflect"
)

// 399.除法求值
// https://leetcode-cn.com/problems/evaluate-division/

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	// fa[x] 表示编号为 x 的变量指向第几个变量。
	fa := make([]int, len(id))
	// w 初值为 1 便于相乘 ，如果权值是相加就改为初值为 0
	w := make([]float64, len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}

	// 要递归调用就要声明再使用
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}

	for i, eq := range equations {
		from, to, val := id[eq[0]], id[eq[1]], values[i]
		fFrom, fTo := find(from), find(to)
		w[fFrom] = val * w[to] / w[from]
		fa[fFrom] = fTo
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE && find(start) == find(end) {
			ans[i] = w[start] / w[end]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func test() {
	type TestType struct {
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}
	ts := []TestType{
		TestType{
			equations: [][]string{
				[]string{"a", "b"},
				[]string{"a", "c"},
			},
			values: []float64{2, 3},
			queries: [][]string{
				[]string{"b", "c"},
				[]string{"c", "a"},
			},
			want: []float64{1.5, 0.3333333333333333},
		},
		TestType{
			equations: [][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
			},
			values: []float64{
				2.0,
				3.0},
			queries: [][]string{
				[]string{"a", "c"},
				[]string{"b", "a"},
				[]string{"a", "e"},
				[]string{"a", "a"},
				[]string{"x", "x"},
			},
			want: []float64{
				6.0,
				0.5,
				-1.0,
				1,
				-1.0,
			},
		},
		TestType{
			equations: [][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
				[]string{"bc", "cd"},
			},
			values: []float64{
				1.5, 2.5, 5.0},
			queries: [][]string{
				[]string{"a", "c"},
				[]string{"c", "b"},
				[]string{"bc", "cd"},
				[]string{"cd", "bc"},
			},
			want: []float64{
				3.75000, 0.40000, 5.00000, 0.20000},
		},
	}
	for _, t := range ts {
		get := calcEquation(t.equations, t.values, t.queries)
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
