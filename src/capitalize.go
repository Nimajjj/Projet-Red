package main

func isAlpha(s string) bool {
	for _, char := range s {
		if !(rune(char) >= 97 && rune(char) <= 122) && !(rune(char) >= 65 && rune(char) <= 90) && !(rune(char) >= 48 && rune(char) <= 57) {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	for _, char := range s {
		if !(rune(char) >= 97 && rune(char) <= 122) {
			return false
		}
	}
	return true
}

func toLower(s string) string {
	res := ""
	for _, char := range s {
		if rune(char) >= 65 && rune(char) <= 90 {
			new_rune := rune(char) + 32
			res += string(new_rune)
		} else {
			res += string(char)
		}
	}
	return res
}

func toUpper(s string) string {
	res := ""
	for _, char := range s {
		if rune(char) >= 97 && rune(char) <= 122 {
			new_rune := rune(char) - 32
			res += string(new_rune)
		} else {
			res += string(char)
		}
	}
	return res
}

func isUpper(s string) bool {
	for _, char := range s {
		if !(rune(char) >= 65 && rune(char) <= 90) {
			return false
		}
	}
	return true
}

func Capitalize(s string) string {
	str := ""
	for index, char := range s {
		new_char := ""
		if index == 0 && isLower(string(char)) {
			new_char = toUpper(string(char))
		} else if isLower(string(char)) && !(isAlpha(string(s[index-1]))) {
			new_char = toUpper(string(char))
		} else {
			if index != 0 && isUpper(string(char)) && isAlpha(string(s[index-1])) {
				new_char = toLower(string(char))
			} else {
				new_char = string(char)
			}
		}
		str += new_char
	}
	return str
}
