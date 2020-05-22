package remove_same_neighbor

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func RemoveSameNeighbor(data []string) []string {
	for i := len(data) - 1; i > 0; i-- {
		if data[i] == data[i-1] {
			data = remove(data, i)
		}
	}
	return data
}
