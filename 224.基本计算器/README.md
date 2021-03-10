# 224.基本计算器
[https://leetcode-cn.com/problems/basic-calculator](https://leetcode-cn.com/problems/basic-calculator) 
## 原题
实现一个基本的计算器来计算一个简单的字符串表达式 `s` 的值。

 

 **示例 1：** 

```

输入：s = "1 + 1"
输出：2

```
 **示例 2：** 

```

输入：s = " 2-1 + 2 "
输出：3

```
 **示例 3：** 

```

输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23

```
 

 **提示：** 
-  `1 <= s.length <= 3 * 10^5` 
-  `s` 由数字、 `'+'` 、 `'-'` 、 `'('` 、 `')'` 、和 `' '` 组成
-  `s` 表示一个有效的表达式
 
**标签**
`栈` `数学` 


## 栈+优先级
```go
const (
	plus  = '+'
	minus = '-'
	lbr   = '('
	rbr   = ')'
)

var priority = map[byte]int{
	plus:  0,
	minus: 0,
	lbr:   1,
	rbr:   0}

func doOp(a, b int, op byte) int {
	switch op {
	case plus:
		return a + b
	case minus:
		return a - b
	default:
		return 0
	}
}
func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "(-", "(0-")

	ops := []byte{}
	nums := []int{0}

	calc := func() {
		if len(nums) < 2 {
			return
		}
		if len(ops) == 0 {
			return
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
		}
		nums = append(nums, res)
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch == lbr:
			ops = append(ops, ch)
		case ch == rbr:
			for len(ops) > 0 {
				if ops[len(ops)-1] != lbr {
					calc()
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		case ch >= '0' && ch <= '9':
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
					calc()
				} else {
					break
				}
			}
			ops = append(ops, ch)
		}
	}
	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		calc()
	}
	return nums[len(nums)-1]
}
```
>执行用时: 8 ms
内存消耗: 2.9 MB

用了类似于编译原理的方法，构建不同的优先级来实现运算的优先级。不过这里把左括号和右括号特殊了。其实可以通过优先级的方式合并起来的。