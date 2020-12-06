# 5620. 连接连续二进制数字
[https://leetcode-cn.com/problems/concatenation-of-consecutive-binary-numbers/](https://leetcode-cn.com/problems/concatenation-of-consecutive-binary-numbers/) 
## 位运算
```go
func concatenatedBinary(n int) int {
	const mod = 1000000007
	res := 0
	lenth := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			lenth++
		}
		res = int((int64(res)<<lenth + int64(i)) % mod)
	}
	return res
}
```
>执行用时: 32 ms
内存消耗: 2.1 MB

假设已经处理了 `1~i-1`的数，他们的结果是$res_{1 \sim i-1}$。我们此时要处理的数为`i`，那么要做的就是把$res_{1 \sim i-1}$转换成二进制，记作$resb_{1 \sim i-1}$ ，把它移动 `lenb(i)`位，然后把移动后的结果加上`i`，再对`1e9+7`取模。

这里的`lenb(x int) int`是作为函数，返回输入的数`x`的二进制位有多少位。

当然这个地方可以通过函数来求，不过还有个办法就是在遍历的时候进行求取。

我们让 `lenth` 记录当前的值的二进制位长度。当它需要更新的时候是什么情况呢？就是前一个数`i-1`为`1111`，而当前的数`i`为`10000`。发现了么？当前的数与其之前的数相与的结果正好为`0`，这个时候就是需要更新的了。因为我们只是对比相邻的两个数，所以这个条件就足够了。不会出现`1010`和`0101`相与也为零这种情况。

当然还有种方法，我们可以观察到，当需要更新的时候，`i`是2的幂次方，因此`i`能够被大于它的2的幂次整除的时候，就需要更新了。因为数据的范围最大为`10^5`，所以为了稳妥取个大一点的，`131072 = 2^17`就是大于所以小于`10^5`的2的幂次的2的幂次。因此可以改写为
```go
func concatenatedBinary(n int) int {
	const mod = 1000000007
	res := 0
	lenth := 0
	for i := 1; i <= n; i++ {
		if 131072 % i == 0 {
			lenth++
		}
		res = int((int64(res)<<lenth + int64(i)) % mod)
	}
	return res
}
```