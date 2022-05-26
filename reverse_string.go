package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	runes := []rune(input)
	output = ""
	for i := len(runes) - 1; i >= 0; i-- {
		if string(runes[i]) == "\r" && string(runes[i+1]) == "\n" {
			output = output[0:len([]rune(output))-1] + "\r\n"
			continue
		}
		output = output + string(runes[i])
	}
	return output
}
