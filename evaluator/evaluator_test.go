package evaluator

import (
	"testing"

	"github.com/shozawa/monkey2/lexer"
	"github.com/shozawa/monkey2/object"
	"github.com/shozawa/monkey2/parser"
)

func TestEval(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	evaluated := Eval(program)
	integer, ok := evaluated.(*object.Integer)
	if !ok {
		t.Errorf("evaluated is not object.Integer. got=%T", evaluated)
	}
	if integer.Value != 5 {
		t.Errorf("integer.Value is not 5. got=%d", integer.Value)
	}
}
