package main

import (
	"fmt"
	"strings"
)

// (1+2)+3
type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left  Expression
	right Expression
}

func Print(e Expression, sb *strings.Builder) {
	de, ok := e.(*DoubleExpression)
	if ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else {
		ae, ok := e.(*AdditionExpression)
		if ok {
			sb.WriteString("(")
			Print(ae.left, sb)
			sb.WriteString("+")
			Print(ae.right, sb)
			sb.WriteString(")")
		}
	}

	// breaks OCP
	// will work incorrectly on missing case
}

func main() {
	// 1+(2+3)
	e := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}
