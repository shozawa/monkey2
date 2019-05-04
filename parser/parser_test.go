package parser

import (
	"testing"

	"github.com/shozawa/monkey2/ast"
	"github.com/shozawa/monkey2/lexer"
)

func TestParseIntegerLiteral(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 1 {
		t.Errorf("len(program.Statements) is not 1. got=%d", got)
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	integer, ok := stmt.Expression.(*ast.Integer)
	if !ok {
		t.Errorf("stmt.Expression is not ast.Integer. got=%T", stmt.Expression)
	}
	if got := integer.TokenLiteral(); got != "5" {
		t.Errorf("integer.TokenLiteral is not 5. got=%q", got)
	}
}
