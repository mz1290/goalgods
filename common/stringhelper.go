package common

func StringIsAlpha(str string) bool {
	if str[0] >= 'a' && str[0] <= 'z' {
		return true
	}

	if str[0] >= 'A' && str[0] <= 'Z' {
		return true
	}

	return false
}

func StringIsDigit(str string) bool {
	if str[0] >= '0' && str[0] <= '9' {
		return true
	}

	return false
}

func StringIsNumber(str string) bool {
	for _, ch := range str {
		if !StringIsDigit(string(ch)) {
			return false
		}
	}

	return true
}
