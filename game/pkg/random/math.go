package random

/*
benchmark
https://stackoverflow.com/questions/64108933/how-to-use-math-pow-with-integers-in-golang
*/
func maxBracket(precision int32) int64 {
	if precision == 0 {
		return 1
	}
	if precision == 1 {
		return 10
	}
	val := int64(10)
	for i := int32(1); i < precision; i++ {
		val *= 10
	}
	return val + 1
}
