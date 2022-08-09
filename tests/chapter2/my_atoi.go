package chapter2

import (
	"math"

	"learning/unit-tests/examples/chapter1"
)

// To Do: Used to convert a number value of type string to int (same as Atoi).
//        - If there is a value other than a number, then end the process and return it according to the previously processed value.
//        - When the resulting value is below -2^31 use the value -2^31 (-2147483648).
//        - When the resulting value is above 2^31 - 1 use the value 2^31 - 1 (2147483647).
func MyAtoi(value string) int {
	if value = chapter1.MyLtrim(value); len(value) == 0 {
		return 0
	}

	var isNegative bool
	if string(value[0]) == "-" {
		isNegative = true
	}

	if string(value[0]) == "-" || string(value[0]) == "+" {
		value = value[1:]
	}

	var min, max, res = int(math.Pow(-2, 31)), (int(math.Pow(2, 31)) - 1), 0

	for i := 0; i < len(value); i++ {
		if value[i] < 48 || value[i] > 57 || res > max {
			break
		}
		res = res*10 + int(value[i]-'0')
	}

	if isNegative {
		res = -1 * res
	}

	if res < min {
		res = min

	} else if res > max {
		res = max

	}

	return res
}
