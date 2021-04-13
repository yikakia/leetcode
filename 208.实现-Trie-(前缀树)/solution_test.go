package leetcode

// 208.实现Trie(前缀树)
// https://leetcode-cn.com/problems/implement-trie-prefix-tree
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

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// func TestSolution(t *testing.T) {
// 	testCases := []struct {
// 		desc string
// 		want
// 	}{
// 		{
//             want: ,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			get :=
// 			if !reflect.DeepEqual(tC.want,get){
// 				t.Errorf("input: %+v get: %v\n",tC,get)
// 			}
// 		})
// 	}
// }
