package queue

import "fmt"

/*
[1, 2, 3, 4, 5]
1 -> 2 -> 3 -> 4 -> 5
*/
type valType int

type Queue struct {
	array []int64
}

func Create() Queue {
	array := make([]int64, 0)
	newQueue := Queue{array: array}
	return newQueue
}

func (q *Queue) Enqueue(value int64) {
	q.array = append(q.array, value)
}

func (q *Queue) Dequeue() {
	q.array = remove(q.array, 0)
}

func (q *Queue) String() string {
	str := fmt.Sprintf("%v", q.array)
	return str
}

func remove(slice []int64, i int64) []int64 {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
