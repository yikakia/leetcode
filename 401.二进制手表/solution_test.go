package leetcode

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

// 401.二进制手表
// https://leetcode-cn.com/problems/binary-watch

var hour map[int]string
var min map[int]string

func initHM() {
	if hour == nil {
		hour = map[int]string{}
		for i := 0; i < 24; i++ {
			hour[i] = strconv.Itoa(i)
		}
	}
	if min == nil {
		min = map[int]string{}
		for i := 0; i < 10; i++ {
			min[i] = "0" + strconv.Itoa(i)
		}
		for i := 10; i < 60; i++ {
			min[i] = strconv.Itoa(i)
		}
	}
}

func genClock(turnedOn int) map[int]bool {
	if turnedOn <= 0 {
		return map[int]bool{0: true}
	}
	genn := genClock(turnedOn - 1)
	n := map[int]bool{}
	for v, _ := range genn {
		for i := 0; i < 10; i++ {
			if b := 1 << i; b&v == 0 &&
				(b|v)&63 < 60 &&
				(b|v)>>6 < 12 {
				n[b|v] = true
			}
		}
	}
	return n
}

func readBinaryWatch(turnedOn int) (ret []string) {
	initHM()
	ret = []string{}
	clocks := genClock(turnedOn)
	for v, _ := range clocks {
		h := v >> 6
		if h > 12 {
			continue
		}
		m := v & 63
		if m > 60 {
			continue
		}
		ret = append(ret, hour[h]+":"+min[m])
	}
	sort.Strings(ret)
	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		turnedOn int
		want     []string
		desc     string
	}{
		{
			turnedOn: 1,
			want:     []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			turnedOn: 9,
			want:     []string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := readBinaryWatch(tC.turnedOn)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
