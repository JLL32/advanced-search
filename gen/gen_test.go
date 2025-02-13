package gen

import (
	"testing"

	"github.com/saferwall/advanced-search/lexer"
	"github.com/saferwall/advanced-search/parser"
	"github.com/saferwall/advanced-search/token"
)

func TestGenerateSQL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"type=pe", "type = 'pe'"},
		{"size>1000", "size > '1000'"},
		{"name!=test.exe", "name != 'test.exe'"},
		{"fs<=2023-01-01", "fs <= '2023-01-01'"},
		{"type=pe and tag=upx OR size>1000mb", "((type = 'pe' AND tag = 'upx') OR size > '1000000000')"},
		{"type=pe or tag=upx size>1000 kb", "(type = 'pe' OR (tag = 'upx' AND size > '1000000'))"},
		{"(type=pe or tag=upx) size>1000 kb", "((type = 'pe' OR tag = 'upx') AND size > '1000000')"},
		{"tag=upx", "tag = 'upx'"},
		{"size<500", "size < '500'"},
		{"name=test.exe", "name = 'test.exe'"},
		{"fs>=2022-12-31", "fs >= '2022-12-31'"},
		{"type=pe and (tag=upx or size>1000mb)", "(type = 'pe' AND (tag = 'upx' OR size > '1000000000'))"},
		{"(type=pe and tag=upx) or size>1000kb", "((type = 'pe' AND tag = 'upx') OR size > '1000000')"},
		{"type=pe and tag=upx and size>1000kb", "(type = 'pe' AND (tag = 'upx' AND size > '1000000'))"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			l := lexer.New(tt.input)
			var tokens []*token.Token
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				tokCopy := tok
				tokens = append(tokens, &tokCopy)
			}

			p := parser.New(tokens)
			expr, err := p.Parse()
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

			sql, err := GenerateSQL(expr)
			if err != nil {
				t.Fatalf("generate SQL error: %v", err)
			}

			if sql != tt.expected {
				t.Errorf("wrong SQL: got %q, want %q", sql, tt.expected)
			}
		})
	}
}
