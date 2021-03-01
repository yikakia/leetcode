package leetcode

// 303.区域和检索-数组不可变
// https://leetcode-cn.com/problems/range-sum-query-immutable

type NumArray struct {
	pre []int
}

func Constructor(nums []int) NumArray {
	pre := make([]int, len(nums)+1)
	for i := 1; i < len(pre); i++ {
		pre[i] += pre[i-1] + nums[i-1]
	}
	return NumArray{
		pre: pre,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.pre[j+1] - this.pre[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
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
