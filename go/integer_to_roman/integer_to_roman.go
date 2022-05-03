package integer_to_roman

func intToRoman(num int) string {
	var r []byte
	m := 1
	for num > 0 {
		d := num % 10                  // get the digit (the remainder of num / 10)
		r = append(convert(d*m), r...) // push the converted place value
		m *= 10                        // multiplier is x10 (1, 10, 100...)
		num /= 10                      // number is divided by ten
	}

	return string(r)
}

func convert(val int) []byte {
	var r []byte
	switch {
	case val >= 1000:
		r = append(r, repeat("M", val/1000)...)
	case val == 900:
		r = append(r, "CM"...)
	case val >= 500:
		r = append(r, "D"+repeat("C", (val%500)/100)...)
	case val == 400:
		r = append(r, "CD"...)
	case val >= 100:
		r = append(r, repeat("C", val/100)...)
	case val == 90:
		r = append(r, "XC"...)
	case val >= 50:
		r = append(r, "L"+repeat("X", (val%50)/10)...)
	case val == 40:
		r = append(r, "XL"...)
	case val >= 10:
		r = append(r, repeat("X", val/10)...)
	case val == 9:
		r = append(r, "IX"...)
	case val >= 5:
		r = append(r, "V"+repeat("I", val%5)...)
	case val == 4:
		r = append(r, "IV"...)
	case val >= 1:
		r = append(r, repeat("I", val)...)
	}

	return r
}

func repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}

	var b []byte
	for i := 0; i < count; i++ {
		b = append(b, s...)
	}

	return string(b)
}
