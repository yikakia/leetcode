# 461.汉明距离
[https://leetcode-cn.com/problems/hamming-distance](https://leetcode-cn.com/problems/hamming-distance) 
## 原题
两个整数之间的<a href="https://baike.baidu.com/item/%E6%B1%89%E6%98%8E%E8%B7%9D%E7%A6%BB">汉明距离</a>指的是这两个数字对应二进制位不同的位置的数目。

给出两个整数 `x` 和 `y` ，计算它们之间的汉明距离。

 **注意：** <br />
0 ≤ `x` , `y` < 2^31.

 **示例:** 

```

输入: x = 1, y = 4

输出: 2

解释:
1   (0 0 0 1)
4   (0 1 0 0)
       &uarr;   &uarr;

上面的箭头指出了对应二进制位不同的位置。

```
 
**标签**
`位运算` 


## 
```go
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
```
>执行用时: 0 ms
内存消耗: 2 MB

统计异或后的1的个数即可。