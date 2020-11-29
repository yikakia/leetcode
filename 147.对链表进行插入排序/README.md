# 147. 对链表进行插入排序
[https://leetcode-cn.com/problems/insertion-sort-list/](https://leetcode-cn.com/problems/insertion-sort-list/) 
##  插入排序
```go
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := &ListNode{Next: head}

	n := 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		n++
	}
	for i := 0; i < n; i++ {
		cur, curFather := res.Next, res
		for j := 0; j < i; j++ {
			cur, curFather = cur.Next, curFather.Next
		}
		tmp, tmpFather, min, minFather := cur, curFather, cur, curFather
		for tmp != nil {
			if tmp.Val < min.Val {
				minFather = tmpFather
				min = tmp
			}
			tmp, tmpFather = tmp.Next, tmpFather.Next
		}
		if cur.Next == min {
			cur.Next, min.Next, curFather.Next = min.Next, cur, min
			continue
		}
		min.Next, cur.Next = cur.Next, min.Next
		curFather.Next, minFather.Next = minFather.Next, curFather.Next
	}

	return res.Next
}
```
>执行用时: 96 ms
内存消耗: 3.3 MB

其实和题解的思路差不多，不过没想到利用前面已经是排好了序的性质去寻找合适的插入位置。所以时间效率很低。链表还是不熟练啊。

我的思路是遍历 n 次，然后为每一位寻找该插入的元素，然后把该插入的元素和当前元素进行交换。

而题解的思路是从左到右遍历的时候如果不满足偏序关系，则从头开始给它找到一个合适的满足偏序关系的位置，然后交换。
