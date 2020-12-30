# 1046. 最后一块石头的重量
[https://leetcode-cn.com/problems/last-stone-weight/](https://leetcode-cn.com/problems/last-stone-weight/) 
## 原题
有一堆石头，每块石头的重量都是正整数。
每一回合，从中选出两块 **最重的**  石头，然后将它们一起粉碎。假设石头的重量分别为 `x` 和 `y`，且 `x <= y`。那么粉碎的可能结果如下：
- 如果 `x == y`，那么两块石头都会被完全粉碎；
- 如果 `x != y`，那么重量为 `x` 的石头将会完全粉碎，而重量为 `y` 的石头新重量为 `y-x`。
最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 `0`。
 
**示例：** 
```
输入：[2,7,4,1,8,1]
输出：1
解释：
先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
```
 
**提示：** 
- `1 <= stones.length <= 30`
- `1 <= stones[i] <= 1000`


## 
```go
type intHeap []int

func (h intHeap) Len() int { return len(h) }
func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func lastStoneWeight(stones []int) int {
	h := &intHeap{}
	for _, stone := range stones {
		heap.Push(h, stone)
	}

	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		if x > y {
			heap.Push(h, x-y)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h.Pop().(int)
}

func test() {
	type TestType struct {
		stones []int
		want   int
	}
	ts := []TestType{
		TestType{
			stones: []int{2, 7, 4, 1, 8, 1}, want: 1,
		},
	}
	for _, t := range ts {
		get := lastStoneWeight(t.stones)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")

}
```
>执行用时: 0 ms
内存消耗: 2 MB


这种要动态处理的都是用堆来处理。
