func WordNumberCase(number int, ifOne string, ifTwo string, ifFive string, addNumber bool) string {
	result := ""
	num := number
	if number < 0 {
		num = -number
	}
	m := num % 10
	switch m {
	case 1:
		result = ifOne
		break
	case 2, 3, 4:
		result = ifTwo
		break
	default:
		result = ifFive
		break
	}
	if m >= 1 && m <= 4 {
		m = num % 100
		if m >= 11 && m <= 14 {
			result = ifFive
		}
	}
	if addNumber {
		result = strconv.Itoa(number) + result
	}
	return result
}