package main

//Stack implements by golang
type Stack struct {
	Items []interface{}
}

func (q *Stack) Push(item interface{}) {
	q.Items = append(q.Items, item)

}
func (q *Stack) Pop() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	item := q.Items[len(q.Items)-1]
	q.Items = q.Items[0 : len(q.Items)-1]
	return item

}
func (q *Stack) Top() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	return q.Items[len(q.Items)-1]
}
func (q *Stack) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Stack) Len() int {
	return len(q.Items)
}
