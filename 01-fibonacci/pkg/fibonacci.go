package fibonacci

import (
	"errors"
)

func Calc(n int) (int, error){

	if n > 20 || n < 0 {
		return 0, errors.New("Wrong number")
	}

	if n == 0 {
		return n, nil
	}

	x,y := 0,1

	for ;n > 0; n-- {
		x, y = y, x + y
	}

	return y, nil
}