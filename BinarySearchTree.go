package datastructures

type BinarySearchTree[t comparable] struct {
	root       TreeNode[t]
	size       int
	comparator func(a, b t) bool
}

//TODO Finish removal function, test
func MakeBinarySearchTree[t comparable](val t, comparator func(a, b t) bool) BinarySearchTree[t] {
	return (BinarySearchTree[t]{TreeNode[t]{nil, nil, val}, 0, comparator})
}

func (bst *BinarySearchTree[t]) Add(v t) {
	parent := &bst.root
	node := TreeNode[t]{nil, nil, v}
	for {
		if bst.comparator(parent.Val, v) {
			//go to the left
			if parent.Left == nil {
				parent.Left = &node
				break
			} else {
				parent = parent.Left
			}
		} else {
			//go to the right
			if parent.Right == nil {
				parent.Right = &node
				break
			} else {
				parent = parent.Right
			}
		}
	}
	bst.size++
}
func (bst *BinarySearchTree[t]) Contains(val t) bool {
	if bst.Find(val) != nil {
		return true
	}
	return false
}
func (bst *BinarySearchTree[t]) Find(val t) *TreeNode[t] {
	ret := &bst.root
	for ret != nil {

		if ret.Val == val {
			return ret
		} else if bst.comparator(val, ret.Val) {
			//go to the right
			ret = ret.Right

		} else {
			ret = ret.Left
		}
	}
	return nil
}
func (bst *BinarySearchTree[t]) FindMax() t {
	parent := bst.root
	for parent.Right != nil {
		parent = *parent.Right
	}
	return (parent.Val)
}
func (bst *BinarySearchTree[t]) FindMin() t {
	parent := bst.root
	for parent.Left != nil {
		parent = *parent.Left
	}
	return (parent.Val)
}
func (bst *BinarySearchTree[t]) GetSize() int {
	return (bst.size)
}

//i REFUSE to use recursion until golang has tail call optimization
func (bst *BinarySearchTree[t]) Remove(v t) {
	toRemove := &bst.root
	var prev *TreeNode[t]
	//find value and save previous so i can KILL IT
	for toRemove != nil && toRemove.Val != v {
		prev = toRemove
		if bst.comparator(toRemove.Val, v) {
			toRemove = toRemove.Left
		} else {
			toRemove = toRemove.Right
		}
	}
	if toRemove == nil {
		return
	}
	if toRemove.Left == nil || toRemove.Right == nil {
		var direction **TreeNode[t]
		if prev == nil {
			direction = &toRemove
		} else if toRemove == prev.Left {
			direction = &prev.Left
		} else {
			direction = &prev.Right
		}
		if toRemove.Left == nil {
			*direction = toRemove.Right
		} else {
			*direction = toRemove.Left
		}
	} else {
		//TODO Finish this code
		toSwap := toRemove.Left
		var prev2 *TreeNode[t]
		for toSwap.Right != nil {
			prev2 = toSwap
			toSwap = toSwap.Right
		}
		toRemove.Val = toSwap.Val
		if toSwap.Left != nil {
			*toSwap = *toSwap.Left
		} else {
			prev2.Right = nil
		}
	}
	bst.size--
}

//ill use recursion because i dont care about this function i used for testing
func PostOrder[t any](a *TreeNode[t]) {
	if a == nil {
		return
	}
	PostOrder(a.Left)
	PostOrder(a.Right)
	print(a.Val, " ")
}
