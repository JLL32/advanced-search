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

func (p *Parser) eatToken() *token.Token {
	if p.current >= len(p.tokens) {
		return &token.Token{Type: token.EOF, Literal: ""}
	}

	tok := p.tokens[p.current]
	p.current++
	return tok
}

func (p *Parser) peekToken() *token.Token {
	if p.current >= len(p.tokens) {
		return &token.Token{Type: token.EOF, Literal: ""}
	}

	return p.tokens[p.current]
}

func (p *Parser) match(tokenType token.TokenType) bool {
	if p.current >= len(p.tokens) {
		return false
	}

	return p.tokens[p.current].Type == tokenType
}

func (p *Parser) Parse() (Expression, error) {
	return p.ParseExpression()
}

func (p *Parser) ParseExpression() (Expression, error) {
	expr, err := p.parseComparison()
	if err != nil {
		return nil, err
	}

	// Look for AND/OR operators
	for p.current < len(p.tokens) {
		if !p.match(token.AND) &&
			!p.match(token.OR) {
			break
		}

		operator := p.eatToken()

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

	left := p.eatToken()

	if p.current >= len(p.tokens) {
		return nil, fmt.Errorf("expected operator after %s", left.Literal)
	}

	operator := p.eatToken()

	if p.current >= len(p.tokens) {
		return nil, fmt.Errorf("expected value after operator")
	}

	right := p.eatToken()

	return &ComparisonExpression{
		Left:     left.Literal,
		Operator: operator,
		Right:    right.Literal,
	}, nil
}
