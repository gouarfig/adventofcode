package main

type Operator string

const (
	Noop   Operator = ""
	And    Operator = "AND"
	Or     Operator = "OR"
	LShift Operator = "LSHIFT"
	RShift Operator = "RSHIFT"
	Not    Operator = "NOT"
)

type OpType int

const (
	OneInput OpType = iota
	TwoInput
)

var (
	Operators = []struct {
		Operator Operator
		Code     string
		Type     OpType
	}{
		{And, string(And), TwoInput},
		{Or, string(Or), TwoInput},
		{LShift, string(LShift), TwoInput},
		{RShift, string(RShift), TwoInput},
		{Not, string(Not), OneInput},
	}
)

type Rule struct {
	Operator Operator
	Sources  []string
}
