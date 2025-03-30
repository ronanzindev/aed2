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

func MoveZeros(arr []int) []int {
	j := 0
	size := len(arr)
	for i := 0; i < size; i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		if arr[j] != 0 {
			j++
		}
	}
	return arr
}
