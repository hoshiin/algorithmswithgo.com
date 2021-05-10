package module01

// 1.467s
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
	// if b == 0 {
	// 	return a
	// }
	// return GCD(b, a%b)
}
