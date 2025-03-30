package queue

import "fmt"

type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{elements: []int{}}
}

func (q *Queue) Push(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return -1, fmt.Errorf("Queue vacío")
	}

	value := q.elements[0]
	q.elements = q.elements[1:]

	return value, nil
}

func (q Queue) peek() (int, error) {
	if len(q.elements) == 0 {
		return -1, fmt.Errorf("Queue vacío")
	}
	return q.elements[0], nil
}

func (q Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) ShowQueue() {
	fmt.Println("Queue actual:", q.elements)
}
