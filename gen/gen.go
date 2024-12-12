package gen

import (
	"advanced-search/parser"
	"advanced-search/token"
	"fmt"
)

func GenerateSQL(expr parser.Expression) (string, error) {
	switch e := expr.(type) {
	case *parser.BinaryExpression:
		return generateBinarySQL(e)
	case *parser.ComparisonExpression:
		return generateComparisonSQL(e)
	default:
		return "", fmt.Errorf("unsupported expression type: %T", expr)
	}
}

func generateBinarySQL(expr *parser.BinaryExpression) (string, error) {
	left, err := GenerateSQL(expr.Left)
	if err != nil {
		return "", err
	}
	right, err := GenerateSQL(expr.Right)
	if err != nil {
		return "", err
	}

	switch expr.Operator.Type {
	case token.AND:
		return fmt.Sprintf("(%s AND %s)", left, right), nil
	case token.OR:
		return fmt.Sprintf("(%s OR %s)", left, right), nil

	default:
		return "", fmt.Errorf("unsupported operator type: %T", expr.Operator.Type)
	}
}

func generateComparisonSQL(expr *parser.ComparisonExpression) (string, error) {
	return fmt.Sprintf("%s %s '%s'", expr.Left, expr.Operator.Literal, expr.Right), nil
}
