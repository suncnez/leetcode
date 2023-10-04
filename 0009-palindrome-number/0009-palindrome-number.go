func isPalindrome(x int) bool {
    if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	if x == 1 || x == 2 || x == 3 || x == 4 || x == 5 || x == 6 || x == 7 || x == 8 || x == 9 {
        return true
    }

	xReversed := reverse(x)
	return xReversed == x
}
func reverse(x int) int {
	sign := "true"
	if x >= 0 {
		sign = "true"
	} else {
		sign = "false"
	}

	x = int(math.Abs(float64(x)))
	var reverseDigit int

	for x > 0 {
		lastDigit := x % 10
		reverseDigit = reverseDigit*10 + lastDigit
		x = x / 10
	}
	if sign == "false" {
		reverseDigit = reverseDigit * -1
	}
	return reverseDigit
}

