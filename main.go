package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Color = "\033[1;33m%s\033[0m"

func main() {
	fmt.Printf(Color, "To use the calculator enter an expression of the form: x+y,z*k...\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter expression: ")

	for {
		exps, _ := reader.ReadString('\n')
		exps = strings.ReplaceAll(exps, "\n", "")
		exps = strings.ReplaceAll(exps, " ", "")

		if len(exps) < 3 {
			continue
		}

		mathExps := parse(exps)

		for _, e := range mathExps {
			fmt.Printf("%s = %v", e, e.Solve())
		}
	}

}

func parse(expStrings string) []*MathExpression {
	var result []*MathExpression
	exps := strings.Split(expStrings, ",")
	for _, e := range exps {
		exp, err := NewMathExpressionFromString(e)
		if err != nil {
			fmt.Printf("Invalid expression: %s. Error: %s\n", e, err)
			continue
		}

		result = append(result, exp)
	}
	return result
}
