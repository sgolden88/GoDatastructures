package datastructures

//If elements i and j share an edge, vertices[i][j] will have an edge weight, best used if nodes will be tightly packed and edgeweight is required
//also supports directional edges
type AdjacencyMatrix struct {
	vertices [][]float32
	size     int
}

func NewAdjacencyMatrix(size int) AdjacencyMatrix {
	ret := make([][]float32, size)
	for i := range ret {
		ret[i] = make([]float32, size)
	}
	return (AdjacencyMatrix{ret, size})
}
func (am *AdjacencyMatrix) GetSize() int {
	return am.size
}
func (am *AdjacencyMatrix) AddEdge(from, to int, val float32) {
	if from > am.GetSize() || to > am.GetSize() || from < 0 || to < 0 {
		panic("Node is out of bounds")
	}
	am.vertices[from][to] = val
}
func (am *AdjacencyMatrix) RemoveEdge(from, to int) {
	if from > am.GetSize() || to > am.GetSize() || from < 0 || to < 0 {
		panic("Node is out of bounds")
	}
	am.vertices[from][to] = 0
}
func (am *AdjacencyMatrix) GetEdge(from, to int) float32 {
	return am.vertices[from][to]
}

//returns list of vertices that specified vertex has an outgoing edge for
func (am *AdjacencyMatrix) OutEdges(vertex int) []int {
	var ret []int
	for x := 0; x < am.size; x++ {
		if am.vertices[vertex][x] > 0 {
			ret = append(ret, x)
		}
	}
	return ret
}

//returns a list of vertices that have edges towards the specified vertex
func (am *AdjacencyMatrix) InEdges(vertex int) []int {
	var ret []int
	for x := 0; x < am.size; x++ {
		if am.vertices[x][vertex] > 0 {
			ret = append(ret, x)
		}
	}
	return ret
}
