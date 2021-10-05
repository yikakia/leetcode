package leetcode

// 284.窥探迭代器
// https://leetcode-cn.com/problems/peeking-iterator
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements
	return false
}
func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	iter    *Iterator
	hasPeek bool
	peekNum int
}

func Constructor(iter *Iterator) *PeekingIterator {
	ret := &PeekingIterator{
		iter:    iter,
		hasPeek: iter.hasNext(),
		peekNum: iter.next(),
	}

	return ret
}

func (pt *PeekingIterator) hasNext() bool {
	return pt.hasPeek
}

func (pt *PeekingIterator) next() int {
	tmp := pt.peek()
	pt.hasPeek = pt.iter.hasNext()
	if pt.hasPeek {
		pt.peekNum = pt.iter.next()
	}
	return tmp
}

func (pt *PeekingIterator) peek() int {
	return pt.peekNum
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
