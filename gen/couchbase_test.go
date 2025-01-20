package gen

import (
	"reflect"
	"testing"

	"github.com/saferwall/advanced-search/lexer"
	"github.com/saferwall/advanced-search/parser"
	"github.com/saferwall/advanced-search/token"
)

func TestGenerateCouchbaseFTS(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]interface{}
	}{
		{
			"type=pe",
			map[string]interface{}{
				"match": "pe",
				"field": "type",
			},
		},
		{
			"size>1000",
			map[string]interface{}{
				"range": map[string]interface{}{
					"field":         "size",
					"min":           "1000",
					"inclusive_min": false,
				},
			},
		},
		{
			"type=pe AND tag=upx",
			map[string]interface{}{
				"conjuncts": []interface{}{
					map[string]interface{}{
						"match": "pe",
						"field": "type",
					},
					map[string]interface{}{
						"match": "upx",
						"field": "tag",
					},
				},
			},
		},
		{
			"type=pe OR size>1000",
			map[string]interface{}{
				"disjuncts": []interface{}{
					map[string]interface{}{
						"match": "pe",
						"field": "type",
					},
					map[string]interface{}{
						"range": map[string]interface{}{
							"field":         "size",
							"min":           "1000",
							"inclusive_min": false,
						},
					},
				},
			},
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

			p := parser.New(tokens)
			expr, err := p.Parse()
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}

			result, err := GenerateCouchbaseFTS(expr)
			if err != nil {
				t.Fatalf("generate error: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("wrong query.\nexpected: %#v\ngot: %#v", tt.expected, result)
			}
		})
	}
}
