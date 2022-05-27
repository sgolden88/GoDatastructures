package datastructures

//First in first out
type Queue[t any] struct {
	queue []t
}

func NewQueue[t any]() Queue[t] {
	return (Queue[t]{})
}
func (q *Queue[t]) Add(val t) {
	q.queue = append(q.queue, val)
}
func (q *Queue[t]) Remove() t {
	var ret t
	if len(q.queue) == 0 {
		return ret
	}
	ret = q.queue[0]
	q.queue = q.queue[1:len(q.queue)]
	return (ret)
}
func (q *Queue[t]) Peek() t {
	return (q.queue[0])
}
func (q *Queue[t]) Size() int {
	return (len(q.queue))
}
