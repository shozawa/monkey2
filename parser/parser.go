package parser

import (
	"github.com/shozawa/monkey2/ast"
	"github.com/shozawa/monkey2/lexer"
	"github.com/shozawa/monkey2/token"
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseExpressionStatement()
		program.Statements = append(program.Statements, stmt)
		p.nextToken()
	}
	return program
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	var exp ast.Expression
	switch p.curToken.Type {
	case token.INT:
		exp = p.parseInteger()
	}
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return &ast.ExpressionStatement{Expression: exp}
}

func (p *Parser) parseInteger() *ast.Integer {
	return &ast.Integer{Token: p.curToken}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
