# 148. 排序链表
[https://leetcode-cn.com/problems/sort-list/](https://leetcode-cn.com/problems/sort-list/) 

## 迭代的归并排序
```go
func lenOfNode(node *ListNode) int {
	n := 0
	for node != nil {
		n++
		node = node.Next
	}
	return n
}

// mergeList 返回按升序排序后合并的链表
func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next, l1 = l1, l1.Next
		} else {
			head.Next, l2 = l2, l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	} else if l2 != nil {
		head.Next = l2
	}
	return res.Next
}
func sortList(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	n := lenOfNode(head)
	// 迭代幅值
	for subListLen := 1; subListLen <= n; subListLen *= 2 {
		// preNode 是合并区间的前一个节点
		// l1 是左边链的起点，l2 是右边链的起点
		preNode := res
		l1, l2 := preNode.Next, res.Next
		// 根据幅值合并一次链表
		for i := 0; i < n; i += 2 * subListLen {
			// 切第一段
			for j := 1; l2 != nil && j < subListLen; j++ {
				l2 = l2.Next
			}
			tmp1 := l2
			//切第二段
			if l2 != nil {
				l2 = l2.Next
			}
			for j := 1; l2 != nil && j < subListLen; j++ {
				l2 = l2.Next
			}
			tmp2 := l2
			tmp3 := l2
			if tmp3 != nil {
				tmp3 = tmp3.Next
			}
			if tmp1 != nil {
				l2 = tmp1.Next
				tmp1.Next = nil
			}
			if tmp2 != nil {
				tmp2.Next = nil
			}
			// 切分成了 [l1,tmp1] [l2,tmp2] [tmp3,]
			// 合并后 l1 和 l2 不一定在原来的位置上
			tmp := mergeList(l1, l2)

			if tmp != nil {
				preNode.Next = tmp
			}
			preNode, l1 = l2, l2
			// 因为 l2 可能为空，所以为了避免报错就根据判断决定
			// preNode 要不要更新到 l2.Next
			if preNode != nil {
				for preNode.Next != nil {
					preNode = preNode.Next
				}
			}
			// 接上末尾
			tmp = res.Next
			for tmp.Next != nil {
				tmp = tmp.Next
			}
			tmp.Next = tmp3

			// 更新 l1 l2 到新的位置
			l1 = tmp3
			l2 = tmp3
		}
	}
	return res.Next
}
```
这个版本超时了，但原理和官方题解的应该差不多，都是拆分链表然后再进行归并操作。这个地方最困扰人的就是拆分链表然后重组链表了。要记录当前指针的下一个位置，然后再把它置空，再把当前指针指向目标位置。

搞定了还没完，还要把当前这个区间的末尾和下一区间的开头相连接。

虽然流程感觉挺简单的，但是第一次上手写各种问题都遇到了。其实主要还是对与指针和链表的用法不太熟悉，以及合并之后位置会改变这一点认识并不清楚。


## 官方题解
```go
// mergeList 返回按升序排序后合并的链表
func mergeList(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	head := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next, l1 = l1, l1.Next
		} else {
			head.Next, l2 = l2, l2.Next
		}
		head = head.Next
	}
	if l1 != nil {
		head.Next = l1
	} else if l2 != nil {
		head.Next = l2
	}
	return res.Next
}
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{Next: head}
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	// subListLen 是迭代幅值
	for subListLen := 1; subListLen <= n; subListLen *= 2 {
        // preNode 记录链表归并后应该在什么节点后面 即 preNode.Next = mergeList()
        // curNode 记录当前到了哪个地方
		preNode, curNode := res, res.Next
		for curNode != nil {
            // 进入循环时的第一个节点就是第一个列表的节点
			list1 := curNode
			for i := 1; curNode.Next != nil && i < subListLen; i++ {
				curNode = curNode.Next
            }
            // 循环 subListLen 次之后的节点为第一个列表的最后一个
            // 而第二个列表的开头是第一个列表的最后一个的下一个
            list2 := curNode.Next
            // 切断链接，现在的 curNode 是第一个列表的最后一个
            curNode.Next = nil
            // 到第二个列表的开头
			curNode = list2
			for i := 1; curNode != nil && curNode.Next != nil && i < subListLen; i++ {
				curNode = curNode.Next
            }
            // next 是下一组的开头，应当指向curNode.Next 
            // 如果 curNode 是 nil 则不用管
			next := curNode
			if curNode != nil {
                next = curNode.Next
                // curNode 现在指向第二个区间的最后一个节点
                // 因此应该断链，将 curNode.Next 置空
				curNode.Next = nil
            }
            // 归并这一组，接在前面的区间的后面
			preNode.Next = mergeList(list1, list2)
            // 找到该区间的最后位置，就是下一个区间的第一个
            // 节点之前的位置
			for preNode.Next != nil {
				preNode = preNode.Next
            }
            // curNode 移动到下一个组的开头
			curNode = next
		}
	}
	return res.Next
}

```
>执行用时: 36 ms
内存消耗: 7.2 MB


和官方题解差不多，具体的都写在代码里了。
