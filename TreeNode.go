package datastructures

//binary tree node
type TreeNode[t any] struct {
	Left, Right *TreeNode[t]
	Val         t
}
