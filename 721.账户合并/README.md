# 721.账户合并
[https://leetcode-cn.com/problems/accounts-merge/](https://leetcode-cn.com/problems/accounts-merge/) 
## 原题
给定一个列表 `accounts`，每个元素 `accounts[i]` 是一个字符串列表，其中第一个元素 `accounts[i][0]` 是 *名称 (name)*，其余元素是 *emails *表示该账户的邮箱地址。

现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。

合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是按顺序排列的邮箱地址。账户本身可以以任意顺序返回。

 

**示例 1：** 

```

输入：
accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
输出：
[["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
解释：
第一个和第三个 John 是同一个人，因为他们有共同的邮箱地址 "johnsmith@mail.com"。 
第二个 John 和 Mary 是不同的人，因为他们的邮箱地址没有被其他帐户使用。
可以以任何顺序返回这些列表，例如答案 [['Mary'，'mary@mail.com']，['John'，'johnnybravo@mail.com']，
['John'，'john00@mail.com'，'john_newyork@mail.com'，'johnsmith@mail.com']] 也是正确的。

```
 

<b>提示：</b>
- `accounts`的长度将在`[1，1000]`的范围内。
- `accounts[i]`的长度将在`[1，10]`的范围内。
- `accounts[i][j]`的长度将在`[1，30]`的范围内。
 
**标签**
`深度优先搜索` `并查集` 


## 
```go
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
```
>执行用时: 52 ms
内存消耗: 8.4 MB

相对来说还是比较简单的。主要是要自己建立并查集的关系，于是就要先储存这个 `email` 是否之前出现过，如果出现过就说明这个账户和另外一个账户相关，于是添加一个关系。我这里是简单地把账户的下标作为出现过的线索，当前账户下的`email`都指向自己的账户的下标，如果这个出现过，那么添加当前的账户的下标和之前已经出现过的账户的下标。

合并的时候就很简单了，就是一般的并查集。

最后有点坑的地方在于合并结果，于是就用另外一个 `map` 来储存下标到合并后的结果的映射关系。为了最后排序的方便，我就用了一个结构体来储存。最后就是先排序结构体中的 `emails`，然后再把结果按照 `name ,emails `的顺序放入结果数组中就可以了。最开始就是没看清楚是要排序所以交错了一次。很不爽。