<?php
	function WordNumberCase($number, $ifOne, $ifTwo, $ifFive, $addNumber = false)
	{
		$result = '';
		$num = $number >= 0 ? $number : -$number;
		$m = $num % 10;
		switch ($m)
		{
			case 1:
				$result = $ifOne;
				break;
			case 2:
			case 3:
			case 4:
				$result = $ifTwo;
				break;
			default:
				$result = $ifFive;
				break;
		}
		if ($m >= 1 && $m <= 4) {
			$m = $num % 100;
			if ($m >= 11 && $m <= 14)
				$result = $ifFive;
		}
		if ($addNumber)
			$result = $number.$result;
		return $result;
	}
?>