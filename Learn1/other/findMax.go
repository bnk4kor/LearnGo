package other


func GetMax(arr []int) int {
	var max int = arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}