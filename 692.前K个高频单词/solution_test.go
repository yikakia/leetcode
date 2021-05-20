package leetcode

import (
	"container/heap"
	"reflect"
	"testing"
)

// 692.前K个高频单词
// https://leetcode-cn.com/problems/top-k-frequent-words

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

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		words []string
		k     int
		want  []string
	}{
		{
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     2,
			want:  []string{"i", "love"},
		},
		{
			words: []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:     4,
			want:  []string{"i", "love", "coding", "leetcode"},
		},
		{
			words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:     4,
			want:  []string{"the", "is", "sunny", "day"},
		},
		{
			words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:     1,
			want:  []string{"the"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := topKFrequent(tC.words, tC.k)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
