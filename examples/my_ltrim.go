package examples

// To Do: This function is used to remove spaces to the left of text
func MyLtrim(text string) string {
	var i int

	for ; i < len(text); i++ {
		if text[i] != 32 {
			break
		}
	}

	return text[i:]
}
