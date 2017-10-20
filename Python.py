def WordNumberCase(number, ifOne, ifTwo, ifFive, addNumber = False):
	result = ''
	num = number
	if number < 0:
		num = -number
	m = num % 10
	if m == 1:
		result = ifOne
	elif 2 <= m <= 4:
		result = ifTwo
	else:
		result = ifFive
	if 1 <= m <= 4:
		if 11 <= num % 100 <= 14:
			result = ifFive
	if addNumber:
		result = str(number) + result;
	return result