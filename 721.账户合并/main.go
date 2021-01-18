package main

import (
	"fmt"
	"reflect"
	"sort"
)

// 721.账户合并
// https://leetcode-cn.com/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	email2Num := make(map[string]int)
	num2Name := make([]string, len(accounts))
	fa := make([]int, len(accounts))
	pairs := [][]int{}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		fa[fy] = fx
	}

	// 构建关系
	for i := range accounts {
		num2Name[i] = accounts[i][0]
		fa[i] = i
		for _, email := range accounts[i][1:] {
			if v, ok := email2Num[email]; ok {
				pairs = append(pairs, []int{i, v})
			} else {
				email2Num[email] = fa[i]
			}
		}
	}

	// 合并
	for _, pair := range pairs {
		merge(pair[0], pair[1])
	}

	type Account struct {
		name   string
		emails []string
	}
	tmp := map[int]Account{}

	for k := range email2Num {
		v := find(email2Num[k])
		if _, ok := tmp[v]; ok {
			account := tmp[v]
			account.emails = append(account.emails, k)
			tmp[v] = account
		} else {
			tmp[v] = Account{name: num2Name[v], emails: []string{k}}
		}
	}

	res := [][]string{}
	for _, account := range tmp {
		e := account.emails
		sort.Slice(e, func(i, j int) bool {
			return e[i] < e[j]
		})
		t := append([]string{account.name}, e...)
		res = append(res, t)
	}
	return res
}
func test() {
	type TestType struct {
		accounts [][]string
		want     [][]string
	}
	ts := []TestType{
		TestType{
			accounts: [][]string{
				[]string{"John", "johnsmith@mail.com", "john00@mail.com"},
				[]string{"John", "johnnybravo@mail.com"},
				[]string{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				[]string{"Mary", "mary@mail.com"},
			},
			want: [][]string{
				[]string{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				[]string{"John", "johnnybravo@mail.com"},
				[]string{"Mary", "mary@mail.com"},
			},
		},
	}
	for _, t := range ts {
		get := accountsMerge(t.accounts)
		if !reflect.DeepEqual(get, t.want) {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")
}
func main() {
	test()
}
