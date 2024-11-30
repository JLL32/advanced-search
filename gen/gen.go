package gen

import (
	"advanced-search/parser"
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
	return fmt.Sprintf("(%s %s %s)", left, expr.Operator.Literal, right), nil
}

func generateComparisonSQL(expr *parser.ComparisonExpression) (string, error) {
	return fmt.Sprintf("%s %s '%s'", expr.Left, expr.Operator.Literal, expr.Right), nil
}
