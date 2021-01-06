# 399. 除法求值
[https://leetcode-cn.com/problems/evaluate-division/](https://leetcode-cn.com/problems/evaluate-division/) 
## 原题
给你一个变量对数组 `equations` 和一个实数值数组 `values` 作为已知条件，其中 `equations[i] = [A<sub>i</sub>, B<sub>i</sub>]` 和 `values[i]` 共同表示等式 `A<sub>i</sub> / B<sub>i</sub> = values[i]` 。每个 `A<sub>i</sub>` 或 `B<sub>i</sub>` 是一个表示单个变量的字符串。
另有一些以数组 `queries` 表示的问题，其中 `queries[j] = [C<sub>j</sub>, D<sub>j</sub>]` 表示第 `j` 个问题，请你根据已知条件找出 `C<sub>j</sub> / D<sub>j</sub> = ?` 的结果作为答案。
返回 **所有问题的答案**  。如果存在某个无法确定的答案，则用 `-1.0` 替代这个答案。
 
**注意：** 输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
 
**示例 1：** 
```
输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
解释：
条件：a / b = 2.0, b / c = 3.0
问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
```
**示例 2：** 
```
输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
输出：[3.75000,0.40000,5.00000,0.20000]
```
**示例 3：** 
```
输入：equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
输出：[0.50000,2.00000,-1.00000,-1.00000]
```
 
**提示：** 
- `1 <= equations.length <= 20`
- `equations[i].length == 2`
- `1 <= A<sub>i</sub>.length, B<sub>i</sub>.length <= 5`
- `values.length == equations.length`
- `0.0 < values[i] <= 20.0`
- `1 <= queries.length <= 20`
- `queries[i].length == 2`
- `1 <= C<sub>j</sub>.length, D<sub>j</sub>.length <= 5`
- `A<sub>i</sub>, B<sub>i</sub>, C<sub>j</sub>, D<sub>j</sub>` 由小写英文字母与数字组成


## Floyd 算法
```go
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 初始化变量
	id := make(map[string]int)
	for _, e := range equations {
		if _, ok := id[e[0]]; !ok {
			id[e[0]] = len(id)
		}
		if _, ok := id[e[1]]; !ok {
			id[e[1]] = len(id)
		}
	}

	graph := make([][]float64, len(id))
	for i := range graph {
		graph[i] = make([]float64, len(id))
	}
	for i, e := range equations {
		v, w := id[e[0]], id[e[1]]
		graph[v][w] = values[i]
		graph[w][v] = 1 / values[i]
	}

	// 计算每个点之间的距离
	// 即每个变量相对于其他变量的比值
	for k := range graph {
		for i := range graph {
			for j := range graph {
				// 题目里的条件 0.0 < values[i] <= 20.0
				// 因此只要大于 0 就是这条边存在。
				if graph[i][k] > 0 && graph[k][j] > 0 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		v, ok1 := id[q[0]]
		w, ok2 := id[q[1]]
		if !(ok1 && ok2) || graph[v][w] == 0 {
			res[i] = -1
			continue
		}
		res[i] = graph[v][w]
	}

	return res
}
```
>执行用时: 0 ms
内存消耗: 2.1 MB

把变量看成一个点，变量之间的比例关系看成边，就可以转化成求图里面的任意两点间的距离了。

当`query`中有一个不存在时，即这条边不存在，按照题意该返回 `-1.0` 。还有就是当`query`两个都存在，但是边不存在的时候，也该返回 `-1.0`。

## 并查集
```go
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
            // 因为初始倍率为 1 所以多次查找也不会增加
            // w[x] 的值，因为 w[fa[x]] == 1（当fa[x]不变时）。
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}

    // 根据已知条件初始化
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
```
>执行用时: 0 ms
内存消耗: 2.1 MB

从查询上看，可以利用并查集的思想。因为如果已知 $\frac{a}{b}$和 $\frac{b}{c}$ 那么要求$\frac{a}{c}$的时候 只用求 $\frac{a}{b} \cdot \frac{b}{c}$即可得到$\frac{a}{c}$。


于是我们先对每一个变量编号，然后把除数指向被除数所指向的数（初始化的时候都指向自己），最终得到的一定是会指向一个或多个数的数组（如果初始化后再对每个数进行 find 操作）。

上面的操作就同时完成了查和并的两个方法。而处理 `query` 的时候就是直接用除数`query[0]`对根节点的倍率和被除数`query[1]`对根节点的倍率相除即可。