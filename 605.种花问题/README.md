# 605. 种花问题
[https://leetcode-cn.com/problems/can-place-flowers/](https://leetcode-cn.com/problems/can-place-flowers/) 
## 原题
假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 **n ** 。能否在不打破种植规则的情况下种入 **n ** 朵花？能则返回True，不能则返回False。
**示例 1:** 
```
输入: flowerbed = [1,0,0,0,1], n = 1
输出: True
```
**示例 2:** 
```
输入: flowerbed = [1,0,0,0,1], n = 2
输出: False
```
**注意:** 
- 数组内已种好的花不会违反种植规则。
- 输入的数组长度范围为 [1, 20000]。
- **n**  是非负整数，且不会超过输入数组的大小。


## 检查是否违反规则
```go
func beforeCheck(flowers []int, index int) bool {
	if index == 0 {
		return true
	}
	return index-1 >= 0 && flowers[index-1] != 1
}
func afterCheck(flowers []int, index int) bool {
	if index == len(flowers)-1 {
		return true
	}
	return index+1 < len(flowers) && flowers[index+1] != 1
}
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := range flowerbed {
		if flowerbed[i] == 0 {
			if beforeCheck(flowerbed, i) && afterCheck(flowerbed, i) {
				count++
				flowerbed[i] = 1
			}
		}
	}
	return count >= n
}

func test() {
	type TestType struct {
		flowerbed []int
		n         int
		want      bool
	}
	ts := []TestType{
		TestType{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		TestType{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
		TestType{
			flowerbed: []int{1},
			n:         0,
			want:      true,
		},
		TestType{
			flowerbed: []int{0},
			n:         1,
			want:      true,
		},
		TestType{
			flowerbed: []int{0},
			n:         3,
			want:      false,
		},
	}
	for _, t := range ts {
		get := canPlaceFlowers(t.flowerbed, t.n)
		if get != t.want {
			// 填充输出格式
			fmt.Printf("%+v get:%v\n", t, get)
		}
	}
	fmt.Println("Finished Test!")
}
```
>执行用时: 20 ms
内存消耗: 6 MB

模拟检查，看是否能插入。能够插入就记录下来，并且把它插入。最后能够插入的数量`count`大于需要插入的数量`n`的时候就返回`true`