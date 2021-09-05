package leetcode

// 470.用Rand7()实现Rand10()
// https://leetcode-cn.com/problems/implement-rand10-using-rand7
func rand10() int {
	for {
		row := rand7()
		col := rand7()
		idx := (row-1)*7 + col
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}

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
