package evaluator

import (
	"testing"

	"github.com/shozawa/monkey2/lexer"
	"github.com/shozawa/monkey2/object"
	"github.com/shozawa/monkey2/parser"
)

func TestEval(t *testing.T) {
	evaluated := testEval("5;")
	testIntegerLiteral(t, evaluated, 5)
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	return Eval(program)
}

func testIntegerLiteral(
	t *testing.T,
	evaluated object.Object,
	want int64,
) bool {
	integer, ok := evaluated.(*object.Integer)
	if !ok {
		t.Errorf("evaluated is not object.Integer. got=%T", evaluated)
		return false
	}
	if integer.Value != want {
		t.Errorf("integer.Value is not %d. got=%d", want, integer.Value)
		return false
	}
	return true
}
