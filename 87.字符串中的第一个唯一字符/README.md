# 87. 字符串中的第一个唯一字符
[https://leetcode-cn.com/problems/first-unique-character-in-a-string/](https://leetcode-cn.com/problems/first-unique-character-in-a-string/) 
## 原题
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 
**示例：** 
```
s = "leetcode"
返回 0
s = "loveleetcode"
返回 2
```
 
**提示：** 你可以假定该字符串只包含小写字母。


## 先统计后排序
```go
func firstUniqChar(s string) int {
	countMap := make(map[rune]int)
	for index, char := range s {
		if _, ok := countMap[char]; ok {
			countMap[char] = -1
		} else {
			countMap[char] = index
		}
	}
	res := []int{}
	for _, v := range countMap {
		if v != -1 {
			res = append(res, v)
		}
	}
	sort.Ints(res)
	if len(res) > 0 {
		return res[0]
	}
	return -1
}
```
>执行用时: 44 ms
内存消耗: 5.4 MB



## 两次遍历字符串
```go
func firstUniqChar(s string) int {
	arr := [26]int{}

	for i, v := range s {
		arr[v-'a'] = i
	}

	for i, v := range s {
		if arr[v-'a'] == i {
			return i
		}
		arr[v-'a'] = -1
	}

	return -1
}
```
>执行用时: 8 ms
内存消耗: 5.2 MB

思路就是开辟一个数组，遍历字符串**s**记录每个字符最后出现的位置。

然后再次遍历字符串**s**查找当前位置是最后一次出现的位置的字符（从头到尾遍历的话，该字符最后一次出现的位置和当前位置相等的话那么就是第一次出现的位置，不等就置为 **-1**）。