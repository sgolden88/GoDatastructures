package datastructures

//pretty simple if you cant use it git gud
type LinkedListNode[t any] struct {
	Next *LinkedListNode[t]
	Val  t
}
