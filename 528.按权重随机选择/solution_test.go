package leetcode

import (
	"math/rand"
	"sort"
	"time"
)

// 528.按权重随机选择
// https://leetcode-cn.com/problems/random-pick-with-weight
type Solution struct {
	preSum []int
	total  int
}

func Constructor(w []int) Solution {
	preSum := make([]int, len(w))
	preSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		preSum[i] = preSum[i-1] + w[i]
	}
	rand.Seed(time.Now().Unix())
	total := preSum[len(preSum)-1]
	return Solution{
		preSum: preSum,
		total:  total,
	}
}

func (this *Solution) PickIndex() int {
	curFind := rand.Intn(this.total) + 1
	return sort.SearchInts(this.preSum, curFind)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// func TestSolution(t *testing.T) {
// 	testCases := []struct {

// 		want
// 		desc string
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
