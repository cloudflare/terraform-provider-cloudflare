package main

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

// Tests for transformQueryStringIncludeInTokens - improving coverage from 8.6%
func TestTransformQueryStringIncludeInTokens(t *testing.T) {
	tests := []struct {
		name     string
		input    hclwrite.Tokens
		expected string
	}{
		{
			name: "transform query_string include list to object",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("query_string")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("include")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("param1")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenComma, Bytes: []byte(",")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("param2")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
			},
			expected: "query_string={include={list=[\"param1\",\"param2\"]}}",
		},
		{
			name: "no query_string - no change",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("other_field")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("value")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
			},
			expected: "other_field=\"value\"",
		},
		{
			name: "query_string without include",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("query_string")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("exclude")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
			},
			expected: "query_string={exclude=[]}",
		},
		{
			name: "nested query_string with multiple depths",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("cache_key")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("query_string")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("include")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("q")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
			},
			expected: "cache_key={query_string={include={list=[\"q\"]}}}",
		},
		{
			name: "empty include list",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("query_string")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("include")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
			},
			expected: "query_string={include=}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformQueryStringIncludeInTokens(tt.input)
			resultStr := strings.ReplaceAll(string(result.Bytes()), " ", "")
			resultStr = strings.ReplaceAll(resultStr, "\n", "")
			expectedStr := strings.ReplaceAll(tt.expected, " ", "")
			assert.Equal(t, expectedStr, resultStr)
		})
	}
}

// Tests for transformHeadersInTokens - improving coverage from 12%
func TestTransformHeadersInTokens(t *testing.T) {
	tests := []struct {
		name     string
		input    hclwrite.Tokens
		contains []string
	}{
		{
			name: "transform headers list to map",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("headers")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("name")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("X-Custom-Header")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenComma, Bytes: []byte(",")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("operation")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("set")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenComma, Bytes: []byte(",")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("value")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("test")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
			},
			contains: []string{"headers", "{", "X-Custom-Header", "operation", "set", "value", "test"},
		},
		{
			name: "no headers - no change",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("other")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("value")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
			},
			contains: []string{"other", "value"},
		},
		{
			name: "headers not followed by equals and bracket",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("headers")},
				{Type: hclsyntax.TokenDot, Bytes: []byte(".")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("something")},
			},
			contains: []string{"headers", ".", "something"},
		},
		{
			name: "multiple headers in list",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("headers")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("name")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("Header1")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenComma, Bytes: []byte(",")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("name")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("Header2")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
			},
			contains: []string{"headers", "Header1", "Header2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformHeadersInTokens(tt.input)
			resultStr := string(result.Bytes())
			for _, expected := range tt.contains {
				assert.Contains(t, resultStr, expected, "Should contain %s", expected)
			}
		})
	}
}

// Tests for fixNestedListToObject - improving coverage from 27%
func TestFixNestedListToObject(t *testing.T) {
	tests := []struct {
		name     string
		input    hclwrite.Tokens
		expected string
	}{
		{
			name: "fix cache_key list to object",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("cache_key")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("custom_key")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
			},
			expected: "cache_key={custom_key={}}",
		},
		{
			name: "fix query_string list to object",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("query_string")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("include")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("all")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
			},
			expected: "query_string={include=\"all\"}",
		},
		{
			name: "no change for non-object fields",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("rules")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
			},
			expected: "rules=[{}]",
		},
		{
			name: "fix nested from_value",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("from_value")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("target_url")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("value")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenQuotedLit, Bytes: []byte("https://example.com")},
				{Type: hclsyntax.TokenCQuote, Bytes: []byte(`"`)},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
			},
			expected: "from_value={target_url={value=\"https://example.com\"}}",
		},
		{
			name: "fix edge_ttl list to object",
			input: hclwrite.Tokens{
				{Type: hclsyntax.TokenIdent, Bytes: []byte("edge_ttl")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenOBrack, Bytes: []byte("[")},
				{Type: hclsyntax.TokenOBrace, Bytes: []byte("{")},
				{Type: hclsyntax.TokenIdent, Bytes: []byte("default")},
				{Type: hclsyntax.TokenEqual, Bytes: []byte("=")},
				{Type: hclsyntax.TokenNumberLit, Bytes: []byte("3600")},
				{Type: hclsyntax.TokenCBrace, Bytes: []byte("}")},
				{Type: hclsyntax.TokenCBrack, Bytes: []byte("]")},
			},
			expected: "edge_ttl={default=3600}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fixNestedListToObject(tt.input)
			resultStr := strings.ReplaceAll(string(result.Bytes()), " ", "")
			resultStr = strings.ReplaceAll(resultStr, "\n", "")
			expectedStr := strings.ReplaceAll(tt.expected, " ", "")
			assert.Equal(t, expectedStr, resultStr)
		})
	}
}

// Tests for convertArraysToObjects - improving coverage from 43.3%
func TestConvertArraysToObjects(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "convert single-element array to object",
			input: map[string]interface{}{
				"action": "rewrite",
				"action_parameters": []interface{}{
					map[string]interface{}{
						"uri": []interface{}{
							map[string]interface{}{
								"path": []interface{}{
									map[string]interface{}{
										"value": "/new-path",
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"action": "rewrite",
				"action_parameters": map[string]interface{}{
					"uri": map[string]interface{}{
						"path": map[string]interface{}{
							"value": "/new-path",
						},
					},
				},
			},
		},
		{
			name: "remove disable_railgun from action_parameters",
			input: map[string]interface{}{
				"action": "set_cache_settings",
				"action_parameters": map[string]interface{}{
					"cache": true,
					"disable_railgun": true,
					"edge_ttl": []interface{}{
						map[string]interface{}{
							"default": 3600,
						},
					},
				},
			},
			expected: map[string]interface{}{
				"action": "set_cache_settings",
				"action_parameters": map[string]interface{}{
					"cache": true,
					"edge_ttl": map[string]interface{}{
						"default": 3600,
					},
				},
			},
		},
		{
			name: "handle nested structures",
			input: map[string]interface{}{
				"action_parameters": []interface{}{
					map[string]interface{}{
						"cache_key": []interface{}{
							map[string]interface{}{
								"custom_key": []interface{}{
									map[string]interface{}{
										"query_string": []interface{}{
											map[string]interface{}{
												"include": []string{"param1", "param2"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"action_parameters": map[string]interface{}{
					"cache_key": map[string]interface{}{
						"custom_key": map[string]interface{}{
							"query_string": map[string]interface{}{
								"include": []string{"param1", "param2"},
							},
						},
					},
				},
			},
		},
		{
			name: "keep multi-element arrays as arrays",
			input: map[string]interface{}{
				"rules": []interface{}{
					map[string]interface{}{"id": "rule1"},
					map[string]interface{}{"id": "rule2"},
				},
			},
			expected: map[string]interface{}{
				"rules": []interface{}{
					map[string]interface{}{"id": "rule1"},
					map[string]interface{}{"id": "rule2"},
				},
			},
		},
		{
			name: "handle from_value with target_url",
			input: map[string]interface{}{
				"action_parameters": []interface{}{
					map[string]interface{}{
						"from_value": []interface{}{
							map[string]interface{}{
								"target_url": []interface{}{
									map[string]interface{}{
										"value": "https://example.com",
									},
								},
								"preserve_query_string": true,
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"action_parameters": map[string]interface{}{
					"from_value": map[string]interface{}{
						"target_url": map[string]interface{}{
							"value": "https://example.com",
						},
						"preserve_query_string": true,
					},
				},
			},
		},
		{
			name: "handle logging and ratelimit",
			input: map[string]interface{}{
				"logging": []interface{}{
					map[string]interface{}{
						"enabled": true,
					},
				},
				"ratelimit": []interface{}{
					map[string]interface{}{
						"period": 60,
						"requests_per_period": 100,
					},
				},
			},
			expected: map[string]interface{}{
				"logging": map[string]interface{}{
					"enabled": true,
				},
				"ratelimit": map[string]interface{}{
					"period": 60,
					"requests_per_period": 100,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertArraysToObjects(tt.input)
			
			// Compare JSON representations for easier debugging
			expectedJSON, _ := json.Marshal(tt.expected)
			resultJSON, _ := json.Marshal(result)
			
			assert.JSONEq(t, string(expectedJSON), string(resultJSON))
		})
	}
}