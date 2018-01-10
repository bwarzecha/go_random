package slice_queue

type IntSliceQueue struct {
	slice []int
}

func New(initialCapacity int) *IntSliceQueue {
	return &IntSliceQueue{slice: make([]int, 0, initialCapacity)}
}

func (q *IntSliceQueue) Put(element int) {
	q.slice = append(q.slice, element)
}

func (q *IntSliceQueue) Peek() int {
	return q.slice[len(q.slice)-1]
}

func (q *IntSliceQueue) Pop() *int {
	if len(q.slice) < 1 {
		return nil
	}
	element := q.slice[0]
	q.slice = q.slice[1:]
	return &element
}
