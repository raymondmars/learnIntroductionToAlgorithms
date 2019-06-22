package main

//Queue implements by golang
type Queue struct {
	Items []interface{}
}

func (q *Queue) Push(item interface{}) {
	q.Items = append(q.Items, item)

}
func (q *Queue) Pop() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	item := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return item

}
func (q *Queue) Top() interface{} {
	if len(q.Items) == 0 {
		return nil
	}
	return q.Items[0]
}
func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) Len() int {
	return len(q.Items)
}
