package datastructures

type SparseMatrix[t any] struct {
	r        int
	c        int
	elements map[int]t
}

//create a sparse SparseMatrix
func NewSparseMatrix[t any](r, c int) SparseMatrix[t] {
	return (SparseMatrix[t]{r, c, make(map[int]t)})
}

func (mat SparseMatrix[t]) Get(r, c int) t {
	if r >= mat.r || c >= mat.c || c < 0 || r < 0 {
		panic("Out of bounds")
	}
	ret, _ := mat.elements[(r*mat.c)+c]
	return ret
}
func (mat *SparseMatrix[t]) Set(r, c int, e t) {
	if r > mat.r || c > mat.c || r < 0 || c < 0 {
		panic("Out of bounds")
	}
	mat.elements[(r*mat.c)+c] = e
}
func (mat *SparseMatrix[t]) Remove(r, c int) {
	delete(mat.elements, (r*mat.c)+c)
}
func (mat *SparseMatrix[t]) Rows() (r int) {
	return mat.r
}
func (mat *SparseMatrix[t]) Columns() (c int) {
	return (mat.c)
}
