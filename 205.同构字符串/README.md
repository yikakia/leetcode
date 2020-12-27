# 205. 同构字符串
[https://leetcode-cn.com/problems/isomorphic-strings/](https://leetcode-cn.com/problems/isomorphic-strings/) 
## 原题
给定两个字符串 **s** 和 **t** ，判断它们是否是同构的。
如果 **s** 中的字符可以被替换得到 **t** ，那么这两个字符串是同构的。
所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
**示例 1:** 
```
输入: s = "egg", t = "add"
输出: true
```
**示例 2:** 
```
输入: s = "foo", t = "bar"
输出: false
```
**示例 3:** 
```
输入: s = "paper", t = "title"
输出: true
```
**说明:** <br>
你可以假设 **s** 和 **t**具有相同的长度。


## 两个哈希表储存两个映射
```go
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2tMap := make(map[byte]byte)
	t2sMap := make(map[byte]byte)
	lenth := len(s)
	for i := 0; i < lenth; i++ {
		if v, ok := s2tMap[s[i]]; ok {
			if v == t[i] {
				if v, ok = t2sMap[t[i]]; ok {
					if v != s[i] {
						return false
					}
				} else {
					return false
				}
				continue
			} else {
				return false
			}
		} else {
			if _, ok = t2sMap[t[i]]; ok {
				return false
			}
			s2tMap[s[i]] = t[i]
			t2sMap[t[i]] = s[i]
		}

	}

	return true
}

func test() {
	type TestType struct {
		s    string
		t    string
		want bool
	}
	ts := []TestType{
		TestType{
			s: "egg", t: "add", want: true,
		},
		TestType{
			s: "foo", t: "bar", want: false,
		},
		TestType{
			s: "paper", t: "title", want: true,
		},
		TestType{
			s: "abcd", t: "eeee", want: false,
		},
	}
	for _, t := range ts {
		get := isIsomorphic(t.s, t.t)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")

}

```
>执行用时: 4 ms
内存消耗: 2.5 MB

和之前的一道题比较类似，也是储存双向的映射关系，然后判断是不是一一对应的映射就行了。