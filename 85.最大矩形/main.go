package main

import "fmt"

// 85.最大矩形
// https://leetcode-cn.com/problems/maximal-rectangle/

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i, row := range matrix {
		dp[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + 1
			}
		}
	}
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			width := dp[i][j]
			area := width
			for k := i - 1; k >= 0; k-- {
				width = min(width, dp[k][j])
				area = max(area, (i-k+1)*width)
			}
			ans = max(ans, area)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test() {
	type TestType struct {
		matrix [][]byte
		want   int
	}
	ts := []TestType{
		TestType{},
	}
	for _, t := range ts {
		get := maximalRectangle(t.matrix)
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
