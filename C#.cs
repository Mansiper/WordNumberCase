public static string WordNumberCase(int number, string ifOne, string ifTwo, string ifFive, bool addNumber = false)
{
		string result;
		switch (number % 10)
		{
				case 1:
						result = ifOne;
						break;
				case 2:
				case 3:
				case 4:
						result = ifTwo;
						break;
				default:
						result = ifFive;
						break;
		}
		var m = number % 100;
		if (m >= 11 && m <= 14)
				result = ifFive;
		if (addNumber)
				result = $"{number}{result}";
		return result;
}