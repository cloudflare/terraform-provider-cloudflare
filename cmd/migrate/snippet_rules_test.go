package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMigrateCloudflareSnippetRules(t *testing.T) {
	tests := []TestCase{
		{
			Name: "basic snippet_rules migration",
			Config: `
resource "cloudflare_snippet_rules" "test" {
  zone_id = "abc123"
  rules {
    expression   = "true"
    snippet_name = "test_snippet_0"
  }
  rules {
    enabled      = true
    expression   = "http.request.uri.path contains \"/api\""
    description  = "API route handler"
    snippet_name = "test_snippet_1"
  }
}`,
			Expected: []string{`resource "cloudflare_snippet_rules" "test" {
  zone_id = "abc123"
  rules = [
    {
      expression   = "true"
      snippet_name = "test_snippet_0"
    },
    {
      enabled      = true
      expression   = "http.request.uri.path contains \"/api\""
      description  = "API route handler"
      snippet_name = "test_snippet_1"
    }
  ]
}
`},
		},
		{
			Name: "snippet_rules with all attributes",
			Config: `
resource "cloudflare_snippet_rules" "test" {
  zone_id = "abc123"
  rules {
    enabled      = false
    expression   = "http.host eq \"example.com\""
    snippet_name = "snippet_a"
    description  = "First rule"
  }
  rules {
    enabled      = true
    expression   = "http.request.method eq \"POST\""
    snippet_name = "snippet_b"
    description  = "Second rule"
  }
  rules {
    expression   = "ip.src in {10.0.0.0/8}"
    snippet_name = "snippet_c"
  }
}`,
			Expected: []string{`resource "cloudflare_snippet_rules" "test" {
  zone_id = "abc123"
  rules = [
    {
      enabled      = false
      expression   = "http.host eq \"example.com\""
      description  = "First rule"
      snippet_name = "snippet_a"
    },
    {
      enabled      = true
      expression   = "http.request.method eq \"POST\""
      description  = "Second rule"
      snippet_name = "snippet_b"
    },
    {
      expression   = "ip.src in {10.0.0.0/8}"
      snippet_name = "snippet_c"
    }
  ]
}
`},
		},
		{
			Name: "snippet_rules with cloudflare_snippet references",
			Config: `
resource "cloudflare_snippet" "redirect_snippet" {
  zone_id      = "abc123"
  name         = "redirect"
  main_module  = "main.js"
  files {
    name    = "main.js"
    content = "export default {};"
  }
}

resource "cloudflare_snippet" "api_snippet" {
  zone_id      = "abc123"
  name         = "api_handler"
  main_module  = "main.js"
  files {
    name    = "main.js"
    content = "export default {};"
  }
}

resource "cloudflare_snippet_rules" "test" {
  zone_id = "abc123"
  rules {
    expression   = "http.host eq \"example.com\""
    snippet_name = cloudflare_snippet.redirect_snippet.name
  }
  rules {
    enabled      = true
    expression   = "http.request.uri.path contains \"/api\""
    snippet_name = cloudflare_snippet.api_snippet.name
    description  = "API handler"
  }
}`,
			Expected: []string{
				`resource "cloudflare_snippet" "redirect_snippet" {
  zone_id      = "abc123"
  snippet_name = "redirect"
  metadata = {
    main_module = "main.js"
  }
  files = [
    {
      name    = "main.js"
      content = "export default {};"
    }
  ]
}`,
				`resource "cloudflare_snippet" "api_snippet" {
  zone_id      = "abc123"
  snippet_name = "api_handler"
  metadata = {
    main_module = "main.js"
  }
  files = [
    {
      name    = "main.js"
      content = "export default {};"
    }
  ]
}`,
				`resource "cloudflare_snippet_rules" "test" {
  zone_id = "abc123"
  rules = [
    {
      expression   = "http.host eq \"example.com\""
      snippet_name = cloudflare_snippet.redirect_snippet.snippet_name
    },
    {
      enabled      = true
      expression   = "http.request.uri.path contains \"/api\""
      description  = "API handler"
      snippet_name = cloudflare_snippet.api_snippet.snippet_name
    }
  ]
}`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

func TestTransformSnippetRulesStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
	}{
		{
			name: "v4 indexed format with basic rules",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet_rules",
					"name": "test",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"zone_id": "abc123",
							"rules.#": "2",
							"rules.0.expression": "true",
							"rules.0.snippet_name": "test_snippet_0",
							"rules.1.enabled": true,
							"rules.1.expression": "http.request.uri.path contains \"/api\"",
							"rules.1.snippet_name": "test_snippet_1",
							"rules.1.description": "API route handler"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id": "abc123",
				"rules": []interface{}{
					map[string]interface{}{
						"enabled":      true, // v4 default when not specified
						"expression":   "true",
						"snippet_name": "test_snippet_0",
						"description":  "", // v5 default when not specified
					},
					map[string]interface{}{
						"enabled":      true,
						"expression":   "http.request.uri.path contains \"/api\"",
						"snippet_name": "test_snippet_1",
						"description":  "API route handler",
					},
				},
			},
		},
		{
			name: "v4 format with explicit enabled=false",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet_rules",
					"name": "test",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"zone_id": "abc123",
							"rules.#": "1",
							"rules.0.enabled": false,
							"rules.0.expression": "http.host eq \"example.com\"",
							"rules.0.snippet_name": "snippet_a",
							"rules.0.description": "Test rule"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id": "abc123",
				"rules": []interface{}{
					map[string]interface{}{
						"enabled":      false,
						"expression":   "http.host eq \"example.com\"",
						"snippet_name": "snippet_a",
						"description":  "Test rule",
					},
				},
			},
		},
		{
			name: "v4 format with empty rules",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet_rules",
					"name": "test",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"zone_id": "abc123",
							"rules.#": "0"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id": "abc123",
				"rules":   []interface{}{},
			},
		},
		{
			name: "v5 format passthrough",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet_rules",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"zone_id": "abc123",
							"rules": [
								{
									"id": "rule-1",
									"enabled": false,
									"expression": "true",
									"snippet_name": "snippet_1",
									"description": "First rule",
									"last_updated": "2024-01-01T00:00:00Z"
								}
							]
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id": "abc123",
				"rules": []interface{}{
					map[string]interface{}{
						"id":           "rule-1",
						"enabled":      false,
						"expression":   "true",
						"snippet_name": "snippet_1",
						"description":  "First rule",
						"last_updated": "2024-01-01T00:00:00Z",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse input JSON
			var inputMap map[string]interface{}
			if err := json.Unmarshal([]byte(tt.input), &inputMap); err != nil {
				t.Fatalf("Failed to parse input JSON: %v", err)
			}

			// Transform the state
			result, err := transformStateJSON([]byte(tt.input))
			if err != nil {
				t.Fatalf("transformStateJSON failed: %v", err)
			}

			// Parse result
			var resultMap map[string]interface{}
			if err := json.Unmarshal(result, &resultMap); err != nil {
				t.Fatalf("Failed to parse result JSON: %v", err)
			}

			// Extract attributes from result
			resources := resultMap["resources"].([]interface{})
			resource := resources[0].(map[string]interface{})
			instances := resource["instances"].([]interface{})
			instance := instances[0].(map[string]interface{})
			attributes := instance["attributes"].(map[string]interface{})

			// Verify schema_version is set to 0
			if schemaVersion := instance["schema_version"]; schemaVersion != float64(0) {
				t.Errorf("Expected schema_version to be 0, got %v", schemaVersion)
			}

			// Verify the transformation
			for key, expectedValue := range tt.expected {
				actualValue, exists := attributes[key]
				if !exists {
					t.Errorf("Missing expected attribute %q", key)
					continue
				}

				// Compare JSON representation for complex types
				expectedJSON, _ := json.Marshal(expectedValue)
				actualJSON, _ := json.Marshal(actualValue)
				if string(expectedJSON) != string(actualJSON) {
					t.Errorf("Attribute %q mismatch:\nExpected: %s\nActual: %s",
						key, string(expectedJSON), string(actualJSON))
				}
			}

			// Verify indexed format is cleaned up for v4 inputs
			if tt.name != "v5 format passthrough" {
				for i := 0; i < 10; i++ {
					indexedKeys := []string{
						fmt.Sprintf("rules.%d.enabled", i),
						fmt.Sprintf("rules.%d.expression", i),
						fmt.Sprintf("rules.%d.snippet_name", i),
						fmt.Sprintf("rules.%d.description", i),
						fmt.Sprintf("rules.%d", i),
					}
					for _, key := range indexedKeys {
						if _, exists := attributes[key]; exists {
							t.Errorf("Indexed attribute %q should have been removed", key)
						}
					}
				}
				if _, exists := attributes["rules.#"]; exists {
					t.Error("rules.# should have been removed")
				}
				if _, exists := attributes["rules.%"]; exists {
					t.Error("rules.% should have been removed")
				}
			}
		})
	}
}