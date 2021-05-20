# 692.前K个高频单词
[https://leetcode-cn.com/problems/top-k-frequent-words](https://leetcode-cn.com/problems/top-k-frequent-words) 
## 原题
给一非空的单词列表，返回前 *k* 个出现次数最多的单词。

返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。

 **示例 1：** 

```

输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]
解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
    注意，按字母顺序 "i" 在 "love" 之前。

```
 

 **示例 2：** 

```

输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
输出: ["the", "is", "sunny", "day"]
解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
    出现次数依次为 4, 3, 2 和 1 次。

```
 

 **注意：** 
- 假定 *k* 总为有效值， 1 ≤ *k* ≤ 集合元素数。
- 输入的单词均由小写字母组成。
 

 **扩展练习：** 
- 尝试以 *O* ( *n* log *k* ) 时间复杂度和 *O* ( *n* ) 空间复杂度解决。
 
**标签**
`堆` `字典树` `哈希表` 


## 
```go
type WordFrq struct {
	word string
	frq  int
}

type WordsFrq struct {
	wf []WordFrq
}

func (w WordsFrq) Len() int {
	return len(w.wf)
}

func (w WordsFrq) Less(i, j int) bool {
	return !(w.wf[i].frq > w.wf[j].frq ||
		w.wf[i].frq == w.wf[j].frq && w.wf[i].word < w.wf[j].word)
}

func (w *WordsFrq) Swap(i, j int) {
	w.wf[i], w.wf[j] = w.wf[j], w.wf[i]
}

func (w *WordsFrq) Pop() interface{} {
	old := w.wf[len(w.wf)-1]
	w.wf = w.wf[:len(w.wf)-1]
	return old
}

func (w *WordsFrq) Push(x interface{}) {
	w.wf = append(w.wf, x.(WordFrq))
}

var _ heap.Interface = (*WordsFrq)(nil)

func reverse(strs []string) {
	n := len(strs)
	for i := 0; i < n/2; i++ {
		strs[i], strs[n-i-1] = strs[n-i-1], strs[i]
	}
}

func topKFrequent(words []string, k int) (res []string) {
	frq := map[string]int{}
	for _, word := range words {
		frq[word]++
	}
	hp := heap.Interface(&WordsFrq{})
	for word, fq := range frq {
		heap.Push(hp, WordFrq{word: word, frq: fq})
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	for hp.Len() > 0 {
		wf := heap.Pop(hp).(WordFrq)
		res = append(res, wf.word)
	}
	reverse(res)
	return
}
```
>执行用时: 4 ms
内存消耗: 4.5 MB

原理很简单，先统计出词频，然后用堆排序求`Top-K`就行。说是中等难度，但可以调库的话就是`Easy`难度了。
