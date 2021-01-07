package main

import "fmt"

// 605.种花问题
// https://leetcode-cn.com/problems/can-place-flowers/

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
func main() {
	test()
}
