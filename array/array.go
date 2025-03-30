package array

func ReverseArray(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)
	start := 0
	end := size - 1
	for start < end {
		newArr[start] = arr[end]
		newArr[end] = arr[start]
		start++
		end--
	}
	return newArr
}
