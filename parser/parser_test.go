package parser

import (
	"testing"
	// "github.com/shozawa/monkey2/token"
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
}
