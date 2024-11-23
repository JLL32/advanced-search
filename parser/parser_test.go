package parser

import (
    "advanced-search/lexer"
    "advanced-search/token"
    "testing"
)

func TestSimpleComparison(t *testing.T) {
    tests := []struct {
        input    string
        wantLeft string
        wantOp   token.TokenType
        wantRight string
    }{
        {"type=pe", "type", token.ASSIGN, "pe"},
        {"size>1000", "size", token.GT, "1000"},
        {"name!=test.exe", "name", token.NOT_EQ, "test.exe"},
        {"fs<=2023-01-01", "fs", token.LE, "2023-01-01"},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            l := lexer.New(tt.input)
            var tokens []*token.Token
            for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
                tokCopy := tok
                tokens = append(tokens, &tokCopy)
            }

            p := New(tokens)
            expr, err := p.ParseExpression()
            if err != nil {
                t.Fatalf("parse error: %v", err)
            }

            compExpr, ok := expr.(*ComparisonExpression)
            if !ok {
                t.Fatalf("expected ComparisonExpression, got %T", expr)
            }

            if compExpr.Left != tt.wantLeft {
                t.Errorf("wrong left value: got %q, want %q", compExpr.Left, tt.wantLeft)
            }
            if compExpr.Operator.Type != tt.wantOp {
                t.Errorf("wrong operator: got %q, want %q", compExpr.Operator.Type, tt.wantOp)
            }
            if compExpr.Right != tt.wantRight {
                t.Errorf("wrong right value: got %q, want %q", compExpr.Right, tt.wantRight)
            }
        })
    }
}

func TestBinaryExpression(t *testing.T) {
    tests := []struct {
        input string
        want  string
    }{
        {
            input: "type=pe AND tag=upx",
            want:  "(type=pe AND tag=upx)",
        },
        {
            input: "size>1000 OR name=test.exe",
            want:  "(size>1000 OR name=test.exe)",
        },
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            l := lexer.New(tt.input)
            var tokens []*token.Token
            for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
                tokCopy := tok
                tokens = append(tokens, &tokCopy)
            }

            p := New(tokens)
            expr, err := p.ParseExpression()
            if err != nil {
                t.Fatalf("parse error: %v", err)
            }

            binExpr, ok := expr.(*BinaryExpression)
            if !ok {
                t.Fatalf("expected BinaryExpression, got %T", expr)
            }

            // Simple validation that we have both sides and an operator
            if binExpr.Left == nil {
                t.Error("left expression is nil")
            }
            if binExpr.Right == nil {
                t.Error("right expression is nil")
            }
            if binExpr.Operator == nil {
                t.Error("operator is nil")
            }
        })
    }
}

func TestParseErrors(t *testing.T) {
    tests := []struct {
        input       string
        wantErrMsg string
    }{
        {"type", "expected operator after type"},
        {"type=", "expected value after operator"},
        {"", "unexpected end of input"},
    }

    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            l := lexer.New(tt.input)
            var tokens []*token.Token
            for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
                tokCopy := tok
                tokens = append(tokens, &tokCopy)
            }

            p := New(tokens)
            _, err := p.ParseExpression()
            if err == nil {
                t.Fatal("expected error, got nil")
            }
            if err.Error() != tt.wantErrMsg {
                t.Errorf("wrong error message: got %q, want %q", err.Error(), tt.wantErrMsg)
            }
        })
    }
}
