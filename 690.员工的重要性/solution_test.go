package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

// 690.员工的重要性
// https://leetcode-cn.com/problems/employee-importance
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) (res int) {

	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Id < employees[j].Id
	})

	if index := sort.Search(
		len(employees),
		func(i int) bool {
			return employees[i].Id >= id
		},
	); index >= len(employees) || employees[index].Id != id {
		return
	}

	queue := []int{id}

	for len(queue) != 0 {
		curId := queue[0]
		queue = queue[1:]
		index := sort.Search(
			len(employees),
			func(i int) bool {
				return employees[i].Id >= curId
			},
		)
		if index >= len(employees) || employees[index].Id != curId {
			continue
		}
		res += employees[index].Importance
		queue = append(queue, employees[index].Subordinates...)
	}

	return
}

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc      string
		employees []*Employee
		id        int
		want      int
	}{
		{
			employees: []*Employee{
				{
					1, 5, []int{2, 3},
				},
				{
					2, 3, []int{},
				},
				{
					3, 3, []int{},
				},
			},
			id:   1,
			want: 11,
		},
		{
			employees: []*Employee{
				{
					1, 15, []int{2},
				},
				{
					2, 10, []int{3},
				},
				{
					3, 5, []int{},
				},
			},
			id:   1,
			want: 30,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := getImportance(tC.employees, tC.id)
			if !reflect.DeepEqual(tC.want, get) {
				t.Errorf("input: %+v get: %v\n", tC, get)
			}
		})
	}
}
