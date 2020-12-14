# 49. 字母异位词分组
[https://leetcode-cn.com/problems/group-anagrams/](https://leetcode-cn.com/problems/group-anagrams/) 
## 原题
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
**示例:** 
```
输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```
**说明：** 
- 所有输入均为小写字母。
- 不考虑答案输出的顺序。


## 词袋模型
```go
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	bow := make(map[[26]int][]string)
	for _, str := range strs {
		tmp := str2bow(str)
		bow[tmp] = append(bow[tmp], str)
	}
	for _, strs := range bow {
		res = append(res, strs)
	}
	return res
}
func str2bow(str string) [26]int {
	res := [26]int{}
	for _, char := range str {
		res[char-'a']++
	}
	return res
}

```
>执行用时: 24 ms
内存消耗: 8.8 MB

看题目的要求，异位词就是每个字母的出现次数是相同的词。

于是想到了自然语言处理的词袋模型。不过我们是记录这个词里面的每个字母出现的次数。因为 **golang** 可以直接比较数组，于是我们就可以直接用 **map** 来把词袋作为 **key**，把 **单词** 作为 **value** 。最后把结果合并起来就可以了。