# 204. 计数质数
[https://leetcode-cn.com/problems/count-primes/](https://leetcode-cn.com/problems/count-primes/) 
## 筛法
```go
func countPrimes(n int) (res int) {
	if n < 2 {
		return
	}
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := range isPrime {
		if isPrime[i] {
			res++
			for j := i * i; j < len(isPrime); j += i {
				isPrime[j] = false
			}
		}
	}
	return
}
```
>执行用时: 12 ms
内存消耗: 4.9 MB

简单地说就是新建一个数组让每一个数初始状态为`true`表示它是素数。然后就是从小到大遍历，如果它是素数，那么就结果加一，并且把${[x,n)}$中$x$的倍数标记为`false`。

当然了，0和1不是素数，这个先自己手动设置为`false`就可以了。

## 线性筛
```go
func countPrimes(n int) (res int) {
	if n < 2 {
		return
	}
	primes := make([]int, 0)
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := range isPrime {
		if isPrime[i] {
			res++
			primes = append(primes, i)
		}
		for _, prime := range primes {
			if i*prime >= n {
				break
			}
			isPrime[i*prime] = false
			if i%prime == 0 {
				break
			}
		}
	}
	return
}
```
>执行用时: 40 ms
内存消耗: 7.1 MB

这里是另外储存了素数数组。在标记`false`的时候是对`isPrime[i*prime]`进行标记。退出条件是当下标溢出或者`i%prime == 0`的时候。这样就可以只对每个数标记一次，避免重复标记。这个的本质是找最小公倍数，每次只更新最小的部分，因为更大的数会被以后标记到。

因此它的时间复杂度应该是$n$，但是执行用时却更多了，~~我觉得应该是动态开辟primes的时候的问题。~~ 直接开辟**n**的空间结果还是一样。我觉得可能是数据不够大吧。
