package strings

func Reverse(str *string) string {
	rstr := ""
	if str == nil {
		return rstr
	}

	for _, v := range *str {
		rstr = string(v) + rstr
	}

	return rstr
}

func Length(str *string) int {
	if str == nil {
		return -1
	}
	return len(*str)
}
