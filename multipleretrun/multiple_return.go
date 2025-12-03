package multiplereturn

import "errors"

func MultipleReturn(numerator int, denomirator int) (int, int, error) {
	if denomirator == 0 {
		return 0, 0, errors.New("can not devide by zero")
	}
	return numerator / denomirator, numerator % denomirator, nil
}
