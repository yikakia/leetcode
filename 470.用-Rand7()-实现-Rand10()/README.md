# 470.用 Rand7() 实现 Rand10()
[https://leetcode-cn.com/problems/implement-rand10-using-rand7](https://leetcode-cn.com/problems/implement-rand10-using-rand7) 
## 原题
已有方法 `rand7` 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 `rand10` 生成 1 到 10 范围内的均匀随机整数。

不要使用系统的 `Math.random()` 方法。
 

 **示例 1:** 

```

输入: 1
输出: [7]

```
 **示例 2:** 

```

输入: 2
输出: [8,4]

```
 **示例 3:** 

```

输入: 3
输出: [8,1,10]

```
 

 **提示:** 
-  `rand7` 已定义。
- 传入参数: `n` 表示 `rand10` 的调用次数。
 

 **进阶:** 
-  `rand7()` 调用次数的 <a href="https://en.wikipedia.org/wiki/Expected_value" target="_blank">期望值</a> 是多少 ?
- 你能否尽量少调用 `rand7()` ?
 
**标签**
`数学` `拒绝采样` `概率与统计` `随机化` 


## 拒绝采样
```go
func rand10() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}

```
>执行用时：12 ms
内存消耗：5.5 MB

原理就是先生成部分的随机数，