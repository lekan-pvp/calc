package arabiccalc

import (
	"errors"
	"fmt"
)

var (
	ErrUnknownOperation = errors.New("unknown operation")
)

type ArabNums struct {
	A  int
	B  int
	Op string
}

func (a *ArabNums) Calculate() (string, error) {
	switch a.Op {
	case "+":
		return a.add()
	case "-":
		return a.sub()
	case "*":
		return a.mul()
	case "/":
		return a.div()
	default:
		return "", ErrUnknownOperation
	}
}

func (a *ArabNums) add() (string, error) {
	result := a.A + a.B

	return fmt.Sprintf("%d", result), nil
}

func (a *ArabNums) sub() (string, error) {
	result := a.A - a.B

	return fmt.Sprintf("%d", result), nil
}

func (a *ArabNums) mul() (string, error) {
	result := a.A * a.B

	return fmt.Sprintf("%d", result), nil
}

func (a *ArabNums) div() (string, error) {
	result := a.A / a.B

	return fmt.Sprintf("%d", result), nil
}
