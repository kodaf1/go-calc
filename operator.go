package main

const (
	sum = iota
	diff
	mul
	div
)

type Operator int

func (a Operator) String() string {
	switch a {
	case sum:
		return "+"
	case diff:
		return "-"
	case mul:
		return "*"
	case div:
		return "/"
	}
	return "(incorrect operator)"
}
