package leetcode

// 1603.设计停车系统
// https://leetcode-cn.com/problems/design-parking-system

type ParkingSystem struct {
	nums []int
	used []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		nums: []int{big, medium, small},
		used: []int{0, 0, 0},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	carType--
	if this.used[carType] < this.nums[carType] {
		this.used[carType]++
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
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
