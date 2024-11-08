package lexer

import (
	"advanced-search/token"
	"testing"
)

var tests = []struct {
	input    string
	expected []token.Token
}{
	{
		input: `size >= 1000kb type = pe fs < 2020-12-31 name = "example" positives != 0`,
		expected: []token.Token{
			{token.SIZE, "size", token.TYPE_INT},
			{token.GE, ">=", ""},
			{token.INT, "1000", token.TYPE_INT},
			{token.UNIT, "kb", ""},
			{token.TYPE, "type", token.TYPE_ENUM},
			{token.ASSIGN, "=", ""},
			{token.STRING, "pe", token.TYPE_STRING},
			{token.FS, "fs", token.TYPE_DATE},
			{token.LT, "<", ""},
			{token.DATE, "2020-12-31", token.TYPE_DATE},
			{token.NAME, "name", token.TYPE_STRING},
			{token.ASSIGN, "=", ""},
			{token.STRING, "example", token.TYPE_STRING},
			{token.POSITIVES, "positives", token.TYPE_INT},
			{token.NOT_EQ, "!=", ""},
			{token.INT, "0", token.TYPE_INT},
			{token.EOF, "", ""},
		},
	},
	{
		input: `type = elf and fs >= 2021-01-01T00:00:00Z or positives < 5`,
		expected: []token.Token{
			{token.TYPE, "type", token.TYPE_ENUM},
			{token.ASSIGN, "=", ""},
			{token.STRING, "elf", token.TYPE_STRING},
			{token.AND, "and", ""},
			{token.FS, "fs", token.TYPE_DATE},
			{token.GE, ">=", ""},
			{token.DATE, "2021-01-01T00:00:00Z", token.TYPE_DATE},
			{token.OR, "or", ""},
			{token.POSITIVES, "positives", token.TYPE_INT},
			{token.LT, "<", ""},
			{token.INT, "5", token.TYPE_INT},
			{token.EOF, "", ""},
		},
	},
	{
		input: `extension = dll or (type = macho and positives > 10)`,
		expected: []token.Token{
			{token.EXTENSION, "extension", token.TYPE_ENUM},
			{token.ASSIGN, "=", ""},
			{token.STRING, "dll", token.TYPE_STRING},
			{token.OR, "or", ""},
			{token.LPAREN, "(", ""},
			{token.TYPE, "type", token.TYPE_ENUM},
			{token.ASSIGN, "=", ""},
			{token.STRING, "macho", token.TYPE_STRING},
			{token.AND, "and", ""},
			{token.POSITIVES, "positives", token.TYPE_INT},
			{token.GT, ">", ""},
			{token.INT, "10", token.TYPE_INT},
			{token.RPAREN, ")", ""},
			{token.EOF, "", ""},
		},
	},
}

func TestNextToken(t *testing.T) {
	for i, tt := range tests {

		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			l := New(tt.input)
			t.Parallel()

			for j, expectedToken := range tt.expected {
				tok := l.NextToken()

				if tok.Type != expectedToken.Type {
					t.Fatalf("tests[%d][%d] - tokentype wrong. expected=%q, got=%q",
						i, j, expectedToken.Type, tok.Type)
				}

				if tok.Literal != expectedToken.Literal {
					t.Fatalf("tests[%d][%d] - literal wrong. expected=%q, got=%q",
						i, j, expectedToken.Literal, tok.Literal)
				}

				if tok.ValueType != expectedToken.ValueType {
					t.Fatalf("tests[%d][%d] - valuetype wrong. expected=%q, got=%q",
						i, j, expectedToken.ValueType, tok.ValueType)
				}
			}

		})
	}
}
