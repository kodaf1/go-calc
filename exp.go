package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MathExpression struct {
	X, Y int64
	Operator
}

func NewMathExpressionFromString(s string) (*MathExpression, error) {
	validOperators := map[string]int{
		"+": sum,
		"-": diff,
		"*": mul,
		"/": div,
	}

	var (
		u bool
	)

	for i, v := range validOperators {
		if strings.Contains(s, i) && u {
			return nil, errors.New("multiple operators in one expression")
		}

		if !strings.Contains(s, i) {
			continue
		}

		o := Operator(v)
		x, y, err := parseInts(s, i)
		if err != nil {
			return nil, err
		}

		return &MathExpression{x, y, o}, nil
	}
	return nil, errors.New("operator was not found")
}

func parseInts(str, sep string) (int64, int64, error) {
	arr := strings.Split(str, sep)

	x, err := strconv.ParseInt(arr[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func (m *MathExpression) Solve() float64 {
	switch m.Operator {
	case sum:
		return float64(m.X + m.Y)
	case diff:
		return float64(m.X - m.Y)
	case mul:
		return float64(m.X * m.Y)
	case div:
		return float64(m.X) / float64(m.Y)
	}

	return 0
}

func (m *MathExpression) String() string {
	return fmt.Sprintf("%d %s %d", m.X, m.Operator.String(), m.Y)
}
