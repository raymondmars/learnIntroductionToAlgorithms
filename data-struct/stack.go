package main

//Stack implements by golang
type Stack struct {
	Items []interface{}
}

func (q *Stack) Push(item interface{}) {
	q.Items = append(q.Items, item)

}
func (q *Stack) Pop() interface{} {
	if len(q.Items) > 0 {
		item := q.Items[len(q.Items)-1]
		q.Items = q.Items[0 : len(q.Items)-1]
		return item
	} else {
		return nil
	}

}
func (q *Stack) Top() interface{} {
	if len(q.Items) > 0 {
		return q.Items[len(q.Items)-1]
	} else {
		return nil
	}
}
func (q *Stack) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Stack) Len() int {
	return len(q.Items)
}
