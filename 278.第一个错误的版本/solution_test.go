package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 278.第一个错误的版本
// https://leetcode-cn.com/problems/first-bad-version
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
var firstVersion int

func isBadVersion(version int) bool {
	return version >= firstVersion
}

func firstBadVersion(n int) int {
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i)
	})
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		n       int
		version int
		want    int
		desc    string
	}{
		{
			n:       5,
			version: 4,
			want:    4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			firstVersion = tC.version
			get := firstBadVersion(tC.n)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
