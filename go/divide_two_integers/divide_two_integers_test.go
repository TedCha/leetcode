package divide_two_integers

import (
	"testing"
)

func Test_Divide_Two_Integers_Case1(t *testing.T) {
	r := divide(10, 3)
	if r != 3 {
		t.FailNow()
	}
}

func Test_Divide_Two_Integers_Case2(t *testing.T) {
	r := divide(7, -3)
	if r != -2 {
		t.FailNow()
	}
}

func Test_Divide_Two_Integers_Case3(t *testing.T) {
	r := divide(41, 5)
	if r != 8 {
		t.FailNow()
	}
}

func Test_Divide_Two_Integers_Case4(t *testing.T) {
	r := divide(2, 2)
	if r != 1 {
		t.FailNow()
	}
}

func Test_Divide_Two_Integers_Case5(t *testing.T) {
	r := divide(1, 2)
	if r != 0 {
		t.FailNow()
	}
}

func Test_Divide_Two_Integers_Case6(t *testing.T) {
	r := divide(-2147483648, -1)
	if r != 2147483647 {
		t.FailNow()
	}
}

func Test_Divide_Two_Integers_Case7(t *testing.T) {
	r := divide(10, -10)
	if r != -1 {
		t.FailNow()
	}
}

func Test_Divide_Two_Integers_Case8(t *testing.T) {
	r := divide(-2147483648, 4)
	if r != -536870912 {
		t.FailNow()
	}
}
