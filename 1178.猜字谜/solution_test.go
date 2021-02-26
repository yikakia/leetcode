package leetcode

import (
	"math/bits"
	"reflect"
	"testing"
)

// 1178.猜字谜
// https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle

func convertByte(b byte) (res uint32) {
	return 1 << uint32(b-'a')
}
func convertStr(str string) (res uint32) {
	for i := range str {
		res |= convertByte(str[i])
	}
	return
}
func findNumOfValidWords(words []string, puzzles []string) []int {
	const puzzleLen int = 7
	wordCnt := map[uint32]int{}
	for _, word := range words {
		mask := convertStr(word)
		if bits.OnesCount(uint(mask)) <= puzzleLen {
			wordCnt[mask]++
		}
	}
	res := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		first := convertByte(puzzle[0])
		mask := convertStr(puzzle[1:])
		subset := mask
		for ok := true; ok; ok = subset != mask {
			res[i] += wordCnt[subset|first]
			subset = (subset - 1) & mask
			// uint32(0)-1 等于 uint32(1111111..)
			// 此时相与得到的即 mask 本身。
		}
	}
	return res
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		words   []string
		puzzles []string
		want    []int
	}{
		{
			words:   []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
			puzzles: []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			want:    []int{1, 1, 3, 2, 4, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := findNumOfValidWords(tC.words, tC.puzzles)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
