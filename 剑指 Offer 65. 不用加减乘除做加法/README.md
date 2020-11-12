# 剑指 Offer 65. 不用加减乘除做加法
[https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/](https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/)
## 位运算
```go
func add(a int, b int) int {
    sum:= a^b
    carry := (a&b)<<1
    for carry != 0{
        sum,carry= sum^carry,(sum&carry)<<1
    }
    return sum
}
```
>执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了68.08%的用户

简单地说就是用位运算模拟加法，具体可以参考这篇文章[位运算实现加、减、乘、除运算](https://www.jianshu.com/p/7bba031b11e7)