package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
// func BaseToDec(value string, base int) int {
// 	index := "0123456789ABCDEF"
// 	res := 0.0
// 	for i, v := range value {
// 		f := strings.IndexRune(index, v)
// 		res += float64(f) * math.Pow(float64(base), float64(len(value)-i-1))
// 	}
// 	return int(res)
// }

func BaseToDec(value string, base int) int {
	index := "0123456789ABCDEF"
	res := 0
	multiplyer := 1
	for i := len(value); i > 0; i-- {
		v := value[i-1]
		for j, r := range index {
			if v == byte(r) {
				res += j * multiplyer
				multiplyer *= base
				break
			}
		}
	}
	return res
}
