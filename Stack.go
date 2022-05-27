package datastructures

//First in last out
type Stack[t any] struct {
	stack []t
}

func NewStack[t any]() Stack[t] {
	return (Stack[t]{})
}
func (s *Stack[t]) Push(val t) {
	s.stack = append(s.stack, val)
}
func (s *Stack[t]) Pop() t {
	var ret t
	if len(s.stack) == 0 {
		return ret
	}
	ret = s.stack[len(s.stack)-1]
	s.stack = s.stack[:(len(s.stack) - 1)]
	return (ret)
}
func (s *Stack[t]) Peek() t {
	return (s.stack[len(s.stack)-1])
}
func (s *Stack[t]) Size() int {
	return (len(s.stack))
}
