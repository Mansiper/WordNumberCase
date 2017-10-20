def WordNumberCase(number, ifOne, ifTwo, ifFive, addNumber = False):
	result = ''
	m = number % 10
	if m == 1:
		result = ifOne
	elif 2 <= m <= 4:
		result = ifTwo
	else:
		result = ifFive
	m = number % 100
	if 11 <= m <= 14:
		result = ifFive
	if addNumber:
		result = str(number) + result;
	return result