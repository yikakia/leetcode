# 208.实现 Trie (前缀树)
[https://leetcode-cn.com/problems/implement-trie-prefix-tree](https://leetcode-cn.com/problems/implement-trie-prefix-tree) 
## 原题
 **<a href="https://baike.baidu.com/item/字典树/9825209?fr=aladdin" target="_blank">Trie</a>** （发音类似 "try"）或者说 **前缀树** 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：
-  `Trie()` 初始化前缀树对象。
-  `void insert(String word)` 向前缀树中插入字符串 `word` 。
-  `boolean search(String word)` 如果字符串 `word` 在前缀树中，返回 `true` （即，在检索之前已经插入）；否则，返回 `false` 。
-  `boolean startsWith(String prefix)` 如果之前已经插入的字符串  `word` 的前缀之一为 `prefix` ，返回 `true` ；否则，返回 `false` 。
 

 **示例：** 

```

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True

```
 

 **提示：** 
-  `1 <= word.length, prefix.length <= 2000` 
-  `word` 和 `prefix` 仅由小写英文字母组成
-  `insert` 、 `search` 和 `startsWith` 调用次数 **总计** 不超过 `3 * 10^4` 次
 
**标签**
`设计` `字典树` 


## 
```go
type Trie struct {
	next  [26]*Trie
	exist bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := range word {
		index := word[i] - 'a'
		if cur.next[index] == nil {
			newTrie := Constructor()
			cur.next[index] = &newTrie
		}
		cur = cur.next[index]
	}
	cur.exist = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := range word {
		index := word[i] - 'a'
		if cur.next[index] == nil {
			return false
		}
		cur = cur.next[index]
	}
	return cur.exist
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := range prefix {
		index := prefix[i] - 'a'
		if cur.next[index] == nil {
			return false
		}
		cur = cur.next[index]
	}
	return true
}

```
>执行用时: 60 ms
内存消耗: 17.7 MB

之前没有写过字典树，这还是第一次写，没想到竟然直接就 `AC` 了，看来对于这种定义比较清晰的数据结构，我的基础还不错啊。

这里简单写一下思路，字典树主要就是为了 `Search` 和 `StartWith` 两个操作，`Search` 是返回之前存过这个单词，而 `StartWith` 则是返回之前存过的单词的前缀是否出现过。

而题目又给了一个范围是在小写英文范围内，因此就可以开一个 `[26] *Trie` 的数组来储存对应的这个字母出现过没有。当不为 `nil` 的时候就说明存在，则 `StartWith` 就返回 `true` 。而 `Search` 则是要判断这条链路的这个字母是否存在过以此为结尾的单词（比如 A->P->P->L->E 这个查找的顺序就是一个单词）。因此就需要另外开一个元素来存以此为结尾的单词是否存在。

了解了以上两点就能很容易地实现了。当然也可以用 `map[rune]*Trie` 来存，只要不是非常密集地存放，用 `map` 还是比数组省空间。