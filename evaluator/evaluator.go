package evaluator

import (
	"github.com/kambehmw/golang-monkey/ast"
	"github.com/kambehmw/golang-monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statement
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	
	// Expression
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}