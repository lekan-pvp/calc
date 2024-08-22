package romancalc

import (
	"errors"

	"github.com/lekan-pvp/calc/internal/roman"
)

var (
	ErrNegativeValue    = errors.New("the value must be positive")
	ErrUnknownOperation = errors.New("unknown operation")
)

type RomanNumbers struct {
	A  int
	B  int
	Op string
}

func (r *RomanNumbers) Calculate() (string, error) {
	switch r.Op {
	case "+":
		return r.add()
	case "-":
		return r.sub()
	case "*":
		return r.mul()
	case "/":
		return r.div()
	default:
		return "", ErrUnknownOperation
	}
}

func (r RomanNumbers) add() (string, error) {
	result := r.A + r.B

	return roman.FromInt(result)
}

func (r RomanNumbers) sub() (string, error) {
	result := r.A - r.B

	if result < 1 {
		return "", ErrNegativeValue
	}

	return roman.FromInt(result)
}

func (r RomanNumbers) mul() (string, error) {
	result := r.A * r.B

	return roman.FromInt(result)
}

func (r RomanNumbers) div() (string, error) {
	result := r.A / r.B

	if result < 1 {
		return "", ErrNegativeValue
	}

	return roman.FromInt(result)
}
