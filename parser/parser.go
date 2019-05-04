package parser

import (
	"github.com/shozawa/monkey2/ast"
	// "github.com/shozawa/monkey2/token"
	"github.com/shozawa/monkey2/lexer"
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	return p
}

type Parser struct {
	l *lexer.Lexer
}

func (p *Parser) Parse() *ast.Program {
	return nil
}

func (p *Parser) nextToken() ast.Node {
	return nil
}
