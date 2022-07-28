package updatelistsolver

type UpdateListSolvingAlg interface {
	UpdateList(request int) int
}

func CreateList(size int) []int {
	list := make([]int, size)

	for i := range list {
		list[i] = i
	}

	return list
}
