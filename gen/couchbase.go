package gen

import (
	"advanced-search/parser"
	"advanced-search/token"
	"fmt"
)

func GenerateCouchbaseFTS(expr parser.Expression) (map[string]interface{}, error) {
	switch e := expr.(type) {
	case *parser.BinaryExpression:
		return generateBinaryCouchbase(e)
	case *parser.ComparisonExpression:
		return generateComparisonCouchbase(e)
	default:
		return nil, fmt.Errorf("unsupported expression type: %T", expr)
	}
}

func generateBinaryCouchbase(expr *parser.BinaryExpression) (map[string]interface{}, error) {
	left, err := GenerateCouchbaseFTS(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := GenerateCouchbaseFTS(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Type {
	case token.AND:
		return map[string]interface{}{
			"conjuncts": []interface{}{left, right},
		}, nil
	case token.OR:
		return map[string]interface{}{
			"disjuncts": []interface{}{left, right},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported operator type: %T", expr.Operator.Type)
	}
}

func generateComparisonCouchbase(expr *parser.ComparisonExpression) (map[string]interface{}, error) {
	switch expr.Operator.Type {
	case token.ASSIGN:
		return map[string]interface{}{
			"match": expr.Right,
			"field": expr.Left,
		}, nil
	case token.NOT_EQ:
		return map[string]interface{}{
			"must_not": map[string]interface{}{
				"match": expr.Right,
				"field": expr.Left,
			},
		}, nil
	case token.GT, token.GE, token.LT, token.LE:
		return generateRangeQuery(expr)
	default:
		return nil, fmt.Errorf("unsupported comparison operator: %s", expr.Operator.Type)
	}
}

func generateRangeQuery(expr *parser.ComparisonExpression) (map[string]interface{}, error) {
	rangeQuery := map[string]interface{}{
		"field": expr.Left,
	}

	switch expr.Operator.Type {
	case token.GT:
		rangeQuery["min"] = expr.Right
		rangeQuery["inclusive_min"] = false
	case token.GE:
		rangeQuery["min"] = expr.Right
		rangeQuery["inclusive_min"] = true
	case token.LT:
		rangeQuery["max"] = expr.Right
		rangeQuery["inclusive_max"] = false
	case token.LE:
		rangeQuery["max"] = expr.Right
		rangeQuery["inclusive_max"] = true
	}

	return map[string]interface{}{
		"range": rangeQuery,
	}, nil
}
