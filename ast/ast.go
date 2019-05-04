package ast

import (
	"fmt"

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
	String() string
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

func (i *Integer) ExpressionNode()      {}
func (i *Integer) TokenLiteral() string { return i.Token.Literal }
func (i *Integer) String() string       { return i.TokenLiteral() }

type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (i *InfixExpression) ExpressionNode() {}
func (i *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Left.String(), i.Operator, i.Right.String())
}
