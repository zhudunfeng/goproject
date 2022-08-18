package utils

func Cal(num1 float64, num2 float64, operator byte) float64 {
	var result float64

	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		result = num1 / num2
	}
	return result
}
