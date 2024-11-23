package parser

import (
	"advanced-search/token"
	"fmt"
)

// AST node types
type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

// Binary expression (e.g. type=pe AND tag=upx)
type BinaryExpression struct {
	Left     Expression
	Operator *token.Token
	Right    Expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Operator.Literal }

// Comparison expression (e.g. type=pe)
type ComparisonExpression struct {
	Left     string
	Operator *token.Token
	Right    string
}

func (ce *ComparisonExpression) expressionNode()      {}
func (ce *ComparisonExpression) TokenLiteral() string { return ce.Operator.Literal }

// Parser structure
type Parser struct {
	tokens  []*token.Token
	current int
}

func New(tokens []*token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

func (p *Parser) ParseExpression() (Expression, error) {
	expr, err := p.parseComparison()
	if err != nil {
		return nil, err
	}

	// Look for AND/OR operators
	for p.current < len(p.tokens) {
		if p.tokens[p.current].Type != token.AND &&
			p.tokens[p.current].Type != token.OR {
			break
		}

		operator := p.tokens[p.current]
		p.current++

		right, err := p.parseComparison()
		if err != nil {
			return nil, err
		}

		expr = &BinaryExpression{
			Left:     expr,
			Operator: operator,
			Right:    right,
		}
	}

	return expr, nil
}

func (p *Parser) parseComparison() (Expression, error) {
	if p.current >= len(p.tokens) {
		return nil, fmt.Errorf("unexpected end of input")
	}

	left := p.tokens[p.current]
	p.current++

	if p.current >= len(p.tokens) {
		return nil, fmt.Errorf("expected operator after %s", left.Literal)
	}

	operator := p.tokens[p.current]
	p.current++

	if p.current >= len(p.tokens) {
		return nil, fmt.Errorf("expected value after operator")
	}

	right := p.tokens[p.current]
	p.current++

	return &ComparisonExpression{
		Left:     left.Literal,
		Operator: operator,
		Right:    right.Literal,
	}, nil
}
