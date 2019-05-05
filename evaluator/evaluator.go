package evaluator

import (
	"fmt"
	"strconv"

	"github.com/shozawa/monkey2/ast"
	"github.com/shozawa/monkey2/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		var result object.Object
		for _, stmt := range node.Statements {
			result = Eval(stmt)
		}
		return result
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.InfixExpression:
		return evalInfixExpression(node)
	case ast.Expression:
		return evalExpression(node)
	default:
	}
	return &object.Integer{}
}

// FIXME: Expressionの評価を切り出さずに、ASTの種類ごとに評価関数を定義したほうが良さそう
func evalExpression(exp ast.Expression) object.Object {
	switch exp := exp.(type) {
	case *ast.Integer:
		value, err := strconv.Atoi(exp.TokenLiteral())
		if err != nil {
			// TODO: 適切なエラーハンドリング
			msg := fmt.Sprintf("exp.TokenLiteral() is invalid. got=%v", exp.TokenLiteral())
			panic(msg)
		}
		return &object.Integer{Value: int64(value)}
	case *ast.InfixExpression:
		// TODO: 中間演算子の評価
		return nil
	default:
		msg := fmt.Sprintf("unexpected expression %T", exp)
		panic(msg)
	}
}

func evalInfixExpression(exp *ast.InfixExpression) object.Object {
	left, ok := Eval(exp.Left).(*object.Integer)
	right, ok := Eval(exp.Right).(*object.Integer)
	if !ok {
		// TODO: 整数同士の加算以外はエラーを起こす
		return nil
	}
	switch exp.Operator {
	case "+":
		return &object.Integer{Value: left.Value + right.Value}
	case "*":
		return &object.Integer{Value: left.Value * right.Value}
	default:
		// TODO: "未対応の演算子" エラーを起こす
		return nil
	}
}
