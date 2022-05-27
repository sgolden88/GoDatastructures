package datastructures

type Ordered interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 |
		~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~uintptr | ~string
}

//A queue that remains ordered, addition and removal takes O(Log(n)) time
type PriorityQueue[t any] struct {
	items      []t
	comparator func(a, b t) bool
}

//comparator returns true if A is higher priority, false if A is lower priority
func MakePriorityQueue[t any](comparator func(a, b t) bool) PriorityQueue[t] {
	return (PriorityQueue[t]{make([]t, 1), comparator})
}
func HighOrder[T Ordered](a, b T) bool {
	return a > b
}
func LowOrder[T Ordered](a, b T) bool {
	return a < b
}
func swap[t any](slice []t, a, b int) {
	slice[b], slice[a] = slice[a], slice[b]
}

func (heap *PriorityQueue[t]) reheapUp() {
	for x := len(heap.items) - 1; x > 1; x = x / 2 {
		if heap.comparator(heap.items[x], heap.items[x/2]) {
			swap(heap.items, x, x/2)
		} else {
			return
		}
	}

}
func (heap *PriorityQueue[t]) reheapDown() {
	for x := 1; (x * 2) < len(heap.items); {
		var y int
		//swap with higher priority of 2 children
		if heap.comparator(heap.items[x*2], heap.items[(x*2)+1]) {
			y = x * 2
		} else {
			y = (x * 2) + 1
		}
		if heap.comparator(heap.items[y], heap.items[x]) {
			swap(heap.items, y, x)
			x = y
		} else {
			return
		}

	}
}

func (heap *PriorityQueue[t]) Add(val t) {
	heap.items = append(heap.items, val)
	heap.reheapUp()

}
func (heap *PriorityQueue[t]) Remove() t {
	var ret t
	if len(heap.items) <= 1 {
		return ret
	}
	ret = heap.items[1]
	heap.items[1] = heap.items[len(heap.items)-1]
	heap.items = heap.items[:len(heap.items)-1]
	heap.reheapDown()
	return (ret)
}
func (heap *PriorityQueue[t]) Peek() t {
	return (heap.items[1])
}
func (heap *PriorityQueue[t]) ToString() {
	for x := 1; x < len(heap.items); x++ {
		print(heap.items[x], " ")
	}
	println()
}
func (heap *PriorityQueue[t]) Size() int {
	return (len(heap.items) - 1)
}
