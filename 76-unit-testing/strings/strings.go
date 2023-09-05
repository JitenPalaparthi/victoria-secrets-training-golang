package strings

func Reverse(str string) string {

	rstr := ""
	for _, v := range str {
		rstr = string(v) + rstr
	}

	return rstr
}

func Length(str string) int {
	return len(str)
}
