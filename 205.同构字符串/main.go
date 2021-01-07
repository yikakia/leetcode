package main

import "fmt"

// 205.同构字符串
// https://leetcode-cn.com/problems/isomorphic-strings/

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
func main() {
	test()
}
