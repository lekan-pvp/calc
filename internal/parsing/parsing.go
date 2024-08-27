package parsing

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lekan-pvp/calc/internal/arabiccalc"
	"github.com/lekan-pvp/calc/internal/roman"
	"github.com/lekan-pvp/calc/internal/romancalc"
	"github.com/lekan-pvp/calc/internal/utils/checkinput"
)

var (
	ErrConvertInt         = errors.New("convert int failure")
	ErrNumberOfExprUnit   = errors.New("the number of unit exression must be 3")
	ErrNotEqualNumberType = errors.New("the type of operands is not equal")
	ErrMustBeGraterThen0  = errors.New("operands must be grater then 0")
)

func Parse(s string) (interface{}, error) {
	opList := strings.Split(s, " ")

	if len(opList) != 3 {
		return nil, ErrNumberOfExprUnit
	}

	if checkinput.IsRoman(opList[0]) != checkinput.IsRoman(opList[2]) {
		return nil, ErrNotEqualNumberType
	}

	if checkinput.IsRoman(opList[0]) {
		a, err := roman.FromRoman(opList[0])
		if err != nil {
			return nil, err
		}
		b, err := roman.FromRoman(opList[2])
		if err != nil {
			return nil, err
		}
		return &romancalc.RomanNumbers{
			A:  a,
			B:  b,
			Op: strings.TrimSpace(opList[1]),
		}, nil
	}

	a, err := strconv.Atoi(opList[0])
	if err != nil {
		return nil, ErrConvertInt
	}

	b, err := strconv.Atoi(opList[2])
	if err != nil {
		return nil, ErrConvertInt
	}

	if a < 1 || b < 1 {
		return nil, ErrMustBeGraterThen0
	}

	return &arabiccalc.ArabNums{
		A:  a,
		B:  b,
		Op: strings.TrimSpace(opList[1]),
	}, nil
}
