package reverse_string

func ReverseString(input string) (output string) {
	if input == "" {
		return
	}
	arr := []rune(input)
	length := len(arr)
	for i := 0; i < length/2; i++ {
		arr[i], arr[length-i-1] = arr[length-i-1], arr[i]
	}
	output = string(arr)
	return
}
