package ast

import (
	"github.com/shozawa/monkey2/token"
)

type Node interface {
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Program struct {
	Statements []Statement
}

type ExpressionStatement struct {
	Expression Expression
}

func (stmt *ExpressionStatement) StatementNode() {}

type Integer struct {
	Token token.Token
}

func (i *Integer) ExpressionNode() {}

type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (i *InfixExpression) ExpressionNode() {}
