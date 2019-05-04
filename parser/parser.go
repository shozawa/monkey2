package parser

import (
	"fmt"

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

const (
	_ = iota
	SUM
	PRODUCT
)

var precedences = map[token.TokenType]int{
	token.PLUS:     SUM,
	token.ASTERISK: PRODUCT,
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
	exp := p.parseExpression(0)
	for p.curToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return &ast.ExpressionStatement{Expression: exp}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	var left ast.Expression
	switch p.curToken.Type {
	case token.INT:
		left = p.parseInteger()
	default:
		msg := fmt.Sprintf("unexpected token %q", p.curToken.Type)
		panic(msg)
	}

	for !p.peekTokenIs(token.SEMICOLON) && p.peekPrecedence() > precedence {
		p.nextToken()
		left = p.parseInfixExpression(left)
	}

	return left
}

func (p *Parser) parseInteger() *ast.Integer {
	return &ast.Integer{Token: p.curToken}
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	infix := &ast.InfixExpression{
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken() // consume operator
	infix.Right = p.parseExpression(precedence)
	return infix
}

func (p *Parser) curPrecedence() int {
	return precedences[p.curToken.Type]
}

func (p *Parser) peekPrecedence() int {
	return precedences[p.peekToken.Type]
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(tokeType token.TokenType) bool {
	return p.peekToken.Type == tokeType
}
