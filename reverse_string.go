package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	result := make([]rune, len(runes))
	for i := 0; i < len(runes); i++ {
		result[i] = runes[len(runes)-i-1]
	}
	output = string(result)
	return
}
