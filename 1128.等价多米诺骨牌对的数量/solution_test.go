package leetcode

import (
	"reflect"
	"testing"
)

// 1128.等价多米诺骨牌对的数量
// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/
func numEquivDominoPairs(dominoes [][]int) int {
	var m [10][10]int
	res := 0
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		res += m[d[0]][d[1]]
		m[d[0]][d[1]]++
	}
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		want     int
		dominoes [][]int
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := numEquivDominoPairs(tC.dominoes)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
