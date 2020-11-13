# 剑指 Offer 12. 矩阵中的路径
[https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/](https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/)

## `DFS`
```go
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, x, y int, steps int) bool {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return false
	}
	if board[x][y] != word[steps] {
		return false
	}
	if steps == len(word)-1 {
		return true
	}
	tmp := board[x][y]
	board[x][y] = ' '
	if dfs(board, word, x+1, y, steps+1) ||
		dfs(board, word, x-1, y, steps+1) ||
		dfs(board, word, x, y+1, steps+1) ||
		dfs(board, word, x, y-1, steps+1) {
		return true
	}

	board[x][y] = tmp
	return false
}
```

很明显的`DFS`。搜出一条路出来。`DFS`解题主要是找清楚边界条件是什么。这个题目里面的是是否数组越界，要找的字符和要找的位置的字符是否相等，以及是否搜索完了。

要注意的是这个地方不能走回头路，所以要有一个记录这个地方是否已经走过了的方法。这里因为不会出现空字符（题意没说清楚，但应该不会有空格存在），所以把走过的位置置空，便于后面查找。如果有空格的存在的话，就得另外建一个 `seen`数组来记录这个位置是否访问过。

在没有找到的话就把置空的位置恢复，或者把`seen`数组对应的位置修改为`unseen`。因为 Golang 的切片在没有对指向的对象进行修改时是不会有变化的（没有指向新的内存区域），就是说表现出了`址传递`的特性。这个要注意，因为会影响到回溯。

`Golang`的数组相关可以看这篇文章👉[Go中slice作为参数传递的一些“坑”](https://juejin.im/post/6844903570706284551)