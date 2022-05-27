package datastructures

//AdjacencyMatrix implemented with a SparseMatrix, more efficient if majority of elements will be sparse
type SparseAdjacencyMatrix struct {
	vertices SparseMatrix[float32]
	size     int
}

func NewSparseAdjacencyMatrix(size int) SparseAdjacencyMatrix {
	return (SparseAdjacencyMatrix{NewSparseMatrix[float32](size, size), size})
}
func (am *SparseAdjacencyMatrix) GetSize() int {
	return am.size
}
func (am *SparseAdjacencyMatrix) AddEdge(from, to int, val float32) {
	if from > am.GetSize() || to > am.GetSize() || from < 0 || to < 0 {
		panic("Node is out of bounds")
	}
	am.vertices.Set(from, to, val)
}
func (am *SparseAdjacencyMatrix) RemoveEdge(from, to int) {
	if from > am.GetSize() || to > am.GetSize() || from < 0 || to < 0 {
		panic("Node is out of bounds")
	}
	am.vertices.Set(from, to, 0)
}
func (am *SparseAdjacencyMatrix) GetEdge(from, to int) float32 {
	return am.vertices.Get(from, to)
}

//returns list of vertices that specified vertex has an outgoing edge for
func (am *SparseAdjacencyMatrix) OutEdges(vertex int) []int {
	var ret []int
	for x := 0; x < am.size; x++ {
		if am.vertices.Get(vertex, x) > 0 {
			ret = append(ret, x)
		}
	}
	return ret
}

//returns a list of vertices that have edges towards the specified vertex
func (am *SparseAdjacencyMatrix) InEdges(vertex int) []int {
	var ret []int
	for x := 0; x < am.size; x++ {
		if am.vertices.Get(x, vertex) > 0 {
			ret = append(ret, x)
		}
	}
	return ret
}
