package leetcode

// 341.扁平化嵌套列表迭代器
// https://leetcode-cn.com/problems/flatten-nested-list-iterator
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool           { return true }
func (this NestedInteger) GetInteger() int           { return 0 }
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }

type NestedIterator struct {
	stacks [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stacks: [][]*NestedInteger{nestedList},
	}
}

func (this *NestedIterator) Next() int {
	last := this.stacks[len(this.stacks)-1]
	val := last[0].GetInteger()
	this.stacks[len(this.stacks)-1] = last[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stacks) > 0 {
		last := this.stacks[len(this.stacks)-1]
		if len(last) == 0 {
			this.stacks = this.stacks[:len(this.stacks)-1]
			continue
		}
		nest := last[0]
		if nest.IsInteger() {
			return true
		}
		this.stacks[len(this.stacks)-1] = last[1:]
		this.stacks = append(this.stacks, nest.GetList())
	}
	return false
}

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
