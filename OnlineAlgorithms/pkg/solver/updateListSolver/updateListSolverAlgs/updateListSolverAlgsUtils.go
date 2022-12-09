package updatelistsolveralgs

func createList(size int) []int {
	list := make([]int, size)

	for i := range list {
		list[i] = i
	}

	return list
}
