# 406. 根据身高重建队列
[https://leetcode-cn.com/problems/queue-reconstruction-by-height/](https://leetcode-cn.com/problems/queue-reconstruction-by-height/) 
## 
```go
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	res := [][]int{}
	for _, person := range people {
		index := person[1]
		res = append(res[:index], append([][]int{person}, res[index:]...)...)
	}
	return res
}
```
>执行用时: 36 ms
内存消耗: 7.7 MB

参照的官方解法，先按照值的大小为第一排序，值的顺序为第二排序来进行排序。然后不断地遍历，把结果插入数组中。

这个地方的可以优化为在原数组中进行操作，这样可以节省掉结果数组的空间和不断地开辟空间的消耗。

不过还是要看能不能对原数组进行修改了。

## 在原数组中进行修改
```go
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	res := [][]int{}
	for i, person := range people {
		copy(people[person[1]+1:i+1], people[person[1]:i])
		people[person[1]] = person
	}
	return res
}
```

简单地说就是在循环到下标为 i 的时候，要做的就是插入，这个时候插入的地方必然是在这个位置之前（或是当前位置）的。

也就是说 copy 这一步是实现了挪移要插入的位置到这个元素当前的位置的所有元素的后一个位置。同时因为在 for range 中获取了这个元素的值，所以这个时候直接在应该插入的位置上赋值就可以了。