package fibonacci

func Calc(n *int) int {
	if *n <= 0 {
		return 0
	}
	if *n > 20 {
		*n = 20
	}
	x, y := 0, 1
	for i := *n; i > 0; i-- {
		x, y = y, x + y
	}
	return y
}