func WordNumberCase(number int, ifOne string, ifTwo string, ifFive string, addNumber bool) string {
	result := ""
	switch number % 10 {
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
	m := number % 100
	if m >= 11 && m <= 14 {
		result = ifFive
	}
	if addNumber {
		result = strconv.Itoa(number) + result
	}
	return result
}