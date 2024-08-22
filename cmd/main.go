package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lekan-pvp/calc/internal/arabiccalc"
	"github.com/lekan-pvp/calc/internal/parsing"
	"github.com/lekan-pvp/calc/internal/romancalc"
)

func main() {
	for {
		s, err := GetString("> ")
		if err != nil {
			panic("cannot read string from commandline")
		}

		exp, err := parsing.Parse(s)
		if err != nil {
			panic(err)
		}

		if e, ok := exp.(*arabiccalc.ArabNums); ok {
			res, err := e.Calculate()
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
		}

		if e, ok := exp.(*romancalc.RomanNumbers); ok {
			res, err := e.Calculate()
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
		}

	}
}

func GetString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
