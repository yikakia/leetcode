# 401.二进制手表
[https://leetcode-cn.com/problems/binary-watch](https://leetcode-cn.com/problems/binary-watch) 
## 原题
二进制手表顶部有 4 个 LED 代表 **小时（0-11）** ，底部的 6 个 LED 代表 **分钟（0-59）** 。每个 LED 代表一个 0 或 1，最低位在右侧。
- 例如，下面的二进制手表读取 `"3:25"` 。
<img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2021/03/29/binary_clock_samui_moon.jpg" style="height: 300px; width" />

<small> *（图源：<a href="https://commons.m.wikimedia.org/wiki/File:Binary_clock_samui_moon.jpg">WikiMedia - Binary clock samui moon.jpg</a> ，许可协议：<a href="https://creativecommons.org/licenses/by-sa/3.0/deed.en">Attribution-ShareAlike 3.0 Unported (CC BY-SA 3.0)</a> ）* </small>

给你一个整数 `turnedOn` ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 **按任意顺序** 返回答案。

小时不会以零开头：
- 例如， `"01:00"` 是无效的时间，正确的写法应该是 `"1:00"` 。
分钟必须由两位数组成，可能会以零开头：
- 例如， `"10:2"` 是无效的时间，正确的写法应该是 `"10:02"` 。
 

 **示例 1：** 

```

输入：turnedOn = 1
输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]

```
 **示例 2：** 

```

输入：turnedOn = 9
输出：[]

```
 

 **提示：** 
-  `0 <= turnedOn <= 10` 
 
**标签**
`位运算` `回溯算法` 


## 搜索
```go
var hour map[int]string
var min map[int]string

func initHM() {
	if hour == nil {
		hour = map[int]string{}
		for i := 0; i < 24; i++ {
			hour[i] = strconv.Itoa(i)
		}
	}
	if min == nil {
		min = map[int]string{}
		for i := 0; i < 10; i++ {
			min[i] = "0" + strconv.Itoa(i)
		}
		for i := 10; i < 60; i++ {
			min[i] = strconv.Itoa(i)
		}
	}
}

func genClock(turnedOn int) map[int]bool {
	if turnedOn <= 0 {
		return map[int]bool{0: true}
	}
	genn := genClock(turnedOn - 1)
	n := map[int]bool{}
	for v, _ := range genn {
		for i := 0; i < 10; i++ {
			if b := 1 << i; b&v == 0 &&
				(b|v)&63 < 60 &&
				(b|v)>>6 < 12 {
				n[b|v] = true
			}
		}
	}
	return n
}

func readBinaryWatch(turnedOn int) (ret []string) {
	initHM()
	ret = []string{}
	clocks := genClock(turnedOn)
	for v, _ := range clocks {
		h := v >> 6
		if h > 12 {
			continue
		}
		m := v & 63
		if m > 60 {
			continue
		}
		ret = append(ret, hour[h]+":"+min[m])
	}
	sort.Strings(ret)
	return
}
```
>执行用时: 0 ms
内存消耗: 2.2 MB

