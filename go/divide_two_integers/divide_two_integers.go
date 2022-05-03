package divide_two_integers

// 10 / 3 = 3.333 {3, 6, 9} (seq) (3)
// 10 / 3 = 3.333 {6, 12} (exp) (4 - 1 = 3)
// 41 / 5 = 8.2 {5, 10, 15, 20, 25, 30, 35, 40} (seq) (8)
// 41 / 5 = 8.2 {10, 20, 30, 40, 50} (exp) (10 - 2 = 8)
// 4 / 2
func divide(dividend int, divisor int) int {

	dividendActual := dividend
	divisorActual := divisor

	// use absolute value for calculations
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	// Simple cases to deal with before brute force
	switch {
	// broken test case
	case dividendActual == -2147483648 && divisorActual == -1:
		return 2147483647
	// if divisor is 1 and divisor <= dividend, answer is dividen
	case divisor == 1 && divisor <= dividend:
		return evalIfNegative(dividend, dividendActual, divisorActual)
	// if divisor is 2, answer is dividend/2 (via left bitshift)
	case divisor == 2:
		return evalIfNegative((dividend >> 1), dividendActual, divisorActual)
	}

	i := 0
	n := 0
	for {
		// count more quickly in exponential of 2
		n += divisor << 1

		// once we're greater than the dividend, break
		if n > dividend {
			break
		}

		i += 2
	}

	// if the quotient - divisor is less than the dividend then add one
	// to the quotient to deal with numbers in between our doubling count
	if (n - divisor) <= dividend {
		return evalIfNegative((i + 1), dividendActual, divisorActual)
	}

	return evalIfNegative(i, dividendActual, divisorActual)
}

func evalIfNegative(x int, dividend int, divisor int) int {
	if (divisor < 0 && !(dividend < 0)) || (dividend < 0 && !(divisor < 0)) {
		x = -x
	}

	return x
}
