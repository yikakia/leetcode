# 227.基本计算器 II
[https://leetcode-cn.com/problems/basic-calculator-ii](https://leetcode-cn.com/problems/basic-calculator-ii) 
## 原题
给你一个字符串表达式 `s` ，请你实现一个基本计算器来计算并返回它的值。

整数除法仅保留整数部分。
 

 **示例 1：** 

```

输入：s = "3+2*2"
输出：7

```
 **示例 2：** 

```

输入：s = " 3/2 "
输出：1

```
 **示例 3：** 

```

输入：s = " 3+5 / 2 "
输出：5

```
 

 **提示：** 
-  `1 <= s.length <= 3 * 10^5` 
-  `s` 由整数和算符 `('+', '-', '*', '/')` 组成，中间由一些空格隔开
-  `s` 表示一个 **有效表达式** 
- 表达式中的所有整数都是非负整数，且在范围 `[0, 2^31 - 1]` 内
- 题目数据保证答案是一个 **32-bit 整数** 
 
**标签**
`栈` `字符串` 


## 
```go
const (
	plus  = '+'
	minus = '-'
	mul   = '*'
	div   = '/'
	lbr   = '('
	rbr   = ')'
)

var priority = map[byte]int{
	plus:  0,
	minus: 0,
	mul:   1,
	div:   1,
	lbr:   -1,
	rbr:   3}

func calc(nums []int, ops []byte) ([]int, []byte) {
	if len(nums) < 2 {
		return nums, ops
	}
	if len(ops) == 0 {
		return nums, ops
	}
	a, b := nums[len(nums)-2], nums[len(nums)-1]
	nums = nums[:len(nums)-2]
	op := ops[len(ops)-1]
	ops = ops[:len(ops)-1]
	res := 0
	switch op {
	case plus:
		res = a + b
	case minus:
		res = a - b
	case mul:
		res = a * b
	case div:
		res = a / b
	}
	nums = append(nums, res)
	return nums, ops
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "(-", "(0-")

	ops := []byte{}
	nums := []int{0}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch == lbr:
			ops = append(ops, ch)
		case ch == rbr:
			for len(ops) > 0 {
				if ops[len(ops)-1] != lbr {
					nums, ops = calc(nums, ops)
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		case unicode.IsDigit(rune(ch)):
			var j = i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			t, _ := strconv.Atoi(s[i:j])
			nums = append(nums, t)
			i = j - 1
		default:
			for len(ops) > 0 && ops[len(ops)-1] != lbr {
				op := ops[len(ops)-1]
				if priority[op] >= priority[ch] {
					nums, ops = calc(nums, ops)
				} else {
					break
				}
			}
			ops = append(ops, ch)
		}
	}
	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		nums, ops = calc(nums, ops)
	}
	return nums[len(nums)-1]
}
```
>执行用时: 8 ms
内存消耗: 2.7 MB

用优先级来做这道题，和上一道题[224.基本计算器](../224.基本计算器/README.md) 相比就是添加了乘除法而已。然后把右括号的优先级增大了。因为右括号应该优先计算出来。

### 完全依靠优先级运算的版本
```go

const (
	plus  = '+'
	minus = '-'
	mul   = '*'
	div   = '/'
	lbr   = '('
	rbr   = ')'
)

var priority = map[byte]int{
	plus:  0,
	minus: 0,
	mul:   1,
	div:   1,
	lbr:   3,
	rbr:   4}

func calc(nums []int, ops []byte) ([]int, []byte) {
	if len(nums) < 2 {
		return nums, ops
	}
	if len(ops) == 0 {
		return nums, ops
	}
	a, b := nums[len(nums)-2], nums[len(nums)-1]
	nums = nums[:len(nums)-2]
	op := ops[len(ops)-1]
	ops = ops[:len(ops)-1]
	res := 0
	switch op {
	case plus:
		res = a + b
	case minus:
		res = a - b
	case mul:
		res = a * b
	case div:
		res = a / b
	case lbr:
		panic("should not")
	case rbr:
		calced := false
		for len(ops) > 0 {
			if ops[len(ops)-1] != lbr {
				nums = append(nums, a, b)
				nums, ops = calc(nums, ops)
				a, b = nums[len(nums)-2], nums[len(nums)-1]
				calced = true
			} else {
				if !calced {
					nums = append(nums, a, b)
				}
				ops = ops[:len(ops)-1]
				return nums, ops
			}
		}
	}
	nums = append(nums, res)
	return nums, ops
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "(-", "(0-")

	ops := []byte{}
	nums := []int{0}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case unicode.IsDigit(rune(ch)):
			var j = i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			t, _ := strconv.Atoi(s[i:j])
			nums = append(nums, t)
			i = j - 1
		default:
			for len(ops) > 0 && ops[len(ops)-1] != lbr {
				op := ops[len(ops)-1]
				if priority[op] >= priority[ch] {
					nums, ops = calc(nums, ops)
				} else {
					break
				}
			}
			ops = append(ops, ch)
		}
	}
	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		nums, ops = calc(nums, ops)
	}
	return nums[len(nums)-1]
}
```
>执行用时: 8 ms
内存消耗: 2.7 MB

这就是完全依靠优先级的版本了。和之前不同的是这个时候左括号的优先级可以提高，然后在最后由右括号进行消除。这个写法主要是写着别扭，还是把括号的逻辑拆分出来比较好。