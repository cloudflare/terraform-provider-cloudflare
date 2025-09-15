package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTransformStateJSON(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms_load_balancer_pool_with_empty_arrays",
			Input: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"identity_schema_version": 0,
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"load_shedding": [],
							"origin_steering": [],
							"origins": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"origins": []
						}
					}]
				}]
			}`,
		},
		{
			Name: "keeps_non_empty_arrays",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"load_shedding": [{"default_percent": 10}],
							"origin_steering": [{"policy": "random"}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"load_shedding": {"default_percent": 10},
							"origin_steering": {"policy": "random"}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_non_load_balancer_pool_resources",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zone",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"some_array": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zone",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"some_array": []
						}
					}]
				}]
			}`,
		},
		{
			Name: "removes_empty_header_arrays_in_origins",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"origins": [
								{
									"name": "origin-1",
									"address": "192.0.2.1",
									"header": []
								},
								{
									"name": "origin-2",
									"address": "192.0.2.2",
									"header": []
								}
							]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"origins": [
								{
									"name": "origin-1",
									"address": "192.0.2.1"
								},
								{
									"name": "origin-2",
									"address": "192.0.2.2"
								}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_header_formats_in_state",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"origins": [{
								"name": "origin-1",
								"address": "192.0.2.1",
								"header": [{
									"header": "Host",
									"values": ["example.com"]
								}]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"origins": [{
								"name": "origin-1",
								"address": "192.0.2.1",
								"header": {
									"host": ["example.com"]
								}
							}]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_load_balancer_transformations",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"fallback_pool_id": "pool-123",
							"default_pool_ids": ["pool-1", "pool-2"],
							"adaptive_routing": [],
							"country_pools": [],
							"rules": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"fallback_pool": "pool-123",
							"default_pools": ["pool-1", "pool-2"],
							"country_pools": {},
							"rules": []
						}
					}]
				}]
			}`,
		},
		{
			Name: "complete_transformation_with_multiple_resources",
			Input: `{
				"version": 4,
				"resources": [
					{
						"type": "cloudflare_load_balancer_pool",
						"name": "pool1",
						"instances": [{
							"identity_schema_version": 0,
							"attributes": {
								"id": "pool-id",
								"name": "test-pool",
								"load_shedding": [{"default_percent": 10}],
								"origin_steering": [],
								"origins": [{
									"name": "origin-1",
									"address": "192.0.2.1",
									"header": [{
										"header": "Host",
										"values": ["example.com"]
									}]
								}]
							}
						}]
					},
					{
						"type": "cloudflare_load_balancer",
						"name": "lb1",
						"instances": [{
							"attributes": {
								"id": "lb-id",
								"name": "test-lb",
								"fallback_pool_id": "pool-123",
								"default_pool_ids": ["pool-1", "pool-2"],
								"adaptive_routing": [],
								"country_pools": [],
								"pop_pools": [],
								"rules": []
							}
						}]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"resources": [
					{
						"type": "cloudflare_load_balancer_pool",
						"name": "pool1",
						"instances": [{
							"attributes": {
								"id": "pool-id",
								"name": "test-pool",
								"load_shedding": {"default_percent": 10},
								"origins": [{
									"name": "origin-1",
									"address": "192.0.2.1",
									"header": {
										"host": ["example.com"]
									}
								}]
							}
						}]
					},
					{
						"type": "cloudflare_load_balancer",
						"name": "lb1",
						"instances": [{
							"attributes": {
								"id": "lb-id",
								"name": "test-lb",
								"fallback_pool": "pool-123",
								"default_pools": ["pool-1", "pool-2"],
								"country_pools": {},
								"pop_pools": {},
								"rules": []
							}
						}]
					}
				]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}

func TestTransformManagedTransformsState(t *testing.T) {
	// Test state with managed_transforms that's missing response headers
	testState := `{
		"version": 4,
		"resources": [{
			"type": "cloudflare_managed_transforms",
			"instances": [{
				"attributes": {
					"id": "test-zone-id",
					"zone_id": "test-zone-id",
					"managed_request_headers": [
						{"id": "add_true_client_ip_headers", "enabled": true}
					]
				}
			}]
		}]
	}`

	// Transform the state
	result, err := transformStateJSON([]byte(testState))
	if err != nil {
		t.Fatalf("Failed to transform state: %v", err)
	}

	// Parse the result
	var resultMap map[string]interface{}
	if err := json.Unmarshal(result, &resultMap); err != nil {
		t.Fatalf("Failed to parse result: %v", err)
	}

	// Check if managed_response_headers was added
	resources := resultMap["resources"].([]interface{})
	resource := resources[0].(map[string]interface{})
	instances := resource["instances"].([]interface{})
	instance := instances[0].(map[string]interface{})
	attributes := instance["attributes"].(map[string]interface{})

	if _, exists := attributes["managed_response_headers"]; !exists {
		t.Error("managed_response_headers was not added to the state")
	}

	// Verify it's an empty array
	responseHeaders := attributes["managed_response_headers"].([]interface{})
	if len(responseHeaders) != 0 {
		t.Errorf("Expected empty managed_response_headers, got %v", responseHeaders)
	}
}

func TestTransformSnippetStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
	}{
		{
			name: "v4 indexed format with single file",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet",
					"name": "test",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"zone_id": "abc123",
							"name": "test_snippet",
							"main_module": "main.js",
							"files.#": "1",
							"files.0.name": "main.js",
							"files.0.content": "export default { async fetch(request) { return fetch(request); } };",
							"id": "some-id"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id":      "abc123",
				"snippet_name": "test_snippet",
				"metadata": map[string]interface{}{
					"main_module": "main.js",
				},
				"files": []interface{}{
					map[string]interface{}{
						"name":    "main.js",
						"content": "export default { async fetch(request) { return fetch(request); } };",
					},
				},
			},
		},
		{
			name: "v4 indexed format with multiple files",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet",
					"name": "test",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"zone_id": "abc123",
							"name": "multi_file_snippet",
							"main_module": "main.js",
							"files.#": "2",
							"files.0.name": "main.js",
							"files.0.content": "import {helper} from './helper.js';",
							"files.1.name": "helper.js",
							"files.1.content": "export function helper() {}",
							"id": "snippet-id"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id":      "abc123",
				"snippet_name": "multi_file_snippet",
				"metadata": map[string]interface{}{
					"main_module": "main.js",
				},
				"files": []interface{}{
					map[string]interface{}{
						"name":    "main.js",
						"content": "import {helper} from './helper.js';",
					},
					map[string]interface{}{
						"name":    "helper.js",
						"content": "export function helper() {}",
					},
				},
			},
		},
		{
			name: "v4 indexed format with empty files",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet",
					"name": "test",
					"instances": [{
						"schema_version": 1,
						"attributes": {
							"zone_id": "abc123",
							"name": "empty_snippet",
							"main_module": "main.js",
							"files.#": "0",
							"id": "empty-id"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id":      "abc123",
				"snippet_name": "empty_snippet",
				"metadata": map[string]interface{}{
					"main_module": "main.js",
				},
				"files": []interface{}{},
			},
		},
		{
			name: "v5 format passthrough",
			input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_snippet",
					"name": "test",
					"instances": [{
						"schema_version": 0,
						"attributes": {
							"zone_id": "abc123",
							"snippet_name": "v5_snippet",
							"metadata": {
								"main_module": "main.js"
							},
							"files": [
								{
									"name": "main.js",
									"content": "export default {};"
								}
							],
							"created_on": "2024-01-01T00:00:00Z",
							"modified_on": "2024-01-02T00:00:00Z"
						}
					}]
				}]
			}`,
			expected: map[string]interface{}{
				"zone_id":      "abc123",
				"snippet_name": "v5_snippet",
				"metadata": map[string]interface{}{
					"main_module": "main.js",
				},
				"files": []interface{}{
					map[string]interface{}{
						"name":    "main.js",
						"content": "export default {};",
					},
				},
				"created_on":  "2024-01-01T00:00:00Z",
				"modified_on": "2024-01-02T00:00:00Z",
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

			// Verify removed fields
			removedFields := []string{"id", "name", "main_module"}
			for _, field := range removedFields {
				if field == "name" || field == "main_module" {
					// These should be removed only if snippet_name and metadata exist
					if _, hasSnippetName := attributes["snippet_name"]; hasSnippetName {
						if _, exists := attributes[field]; exists {
							t.Errorf("Field %q should have been removed", field)
						}
					}
				} else if field == "id" {
					// id should always be removed
					if _, exists := attributes[field]; exists {
						t.Errorf("Field %q should have been removed", field)
					}
				}
			}

			// Verify indexed format is cleaned up
			for i := 0; i < 10; i++ {
				fileNameKey := fmt.Sprintf("files.%d.name", i)
				fileContentKey := fmt.Sprintf("files.%d.content", i)
				if _, exists := attributes[fileNameKey]; exists {
					t.Errorf("Indexed file attribute %q should have been removed", fileNameKey)
				}
				if _, exists := attributes[fileContentKey]; exists {
					t.Errorf("Indexed file attribute %q should have been removed", fileContentKey)
				}
			}
			if _, exists := attributes["files.#"]; exists {
				t.Error("files.# should have been removed")
			}
			if _, exists := attributes["files.%"]; exists {
				t.Error("files.% should have been removed")
			}
		})
	}
}
func TestCustomPagesStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms type to identifier for account-level custom pages",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"type": "500_errors",
									"state": "customized",
									"url": "https://example.com/500.html",
									"id": "custom-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "500_errors",
									"state": "customized",
									"url": "https://example.com/500.html",
									"id": "custom-page-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "transforms type to identifier for zone-level custom pages",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"zone_id": "abcdef123456",
									"type": "ip_block",
									"state": "default",
									"url": "",
									"id": "custom-page-zone-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"zone_id": "abcdef123456",
									"identifier": "ip_block",
									"state": "default",
									"url": "",
									"id": "custom-page-zone-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "handles multiple custom pages in one state file",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "account_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"type": "basic_challenge",
									"state": "customized",
									"url": "https://example.com/challenge.html",
									"id": "account-page-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "zone_page",
						"instances": [
							{
								"attributes": {
									"zone_id": "fedcba987654321",
									"type": "managed_challenge",
									"state": "default",
									"url": "",
									"id": "zone-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "account_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "basic_challenge",
									"state": "customized",
									"url": "https://example.com/challenge.html",
									"id": "account-page-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "zone_page",
						"instances": [
							{
								"attributes": {
									"zone_id": "fedcba987654321",
									"identifier": "managed_challenge",
									"state": "default",
									"url": "",
									"id": "zone-page-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "leaves already migrated custom pages unchanged",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "already_migrated",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "waf_challenge",
									"state": "customized",
									"url": "https://example.com/waf.html",
									"id": "migrated-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "already_migrated",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "waf_challenge",
									"state": "customized",
									"url": "https://example.com/waf.html",
									"id": "migrated-page-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "handles mixed resource types with custom pages",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zone",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"account": {
										"id": "123456789abcdef0"
									},
									"name": "example.com",
									"id": "zone-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "error_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"type": "1000_errors",
									"state": "customized",
									"url": "https://example.com/1000.html",
									"id": "error-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zone",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"account": {
										"id": "123456789abcdef0"
									},
									"name": "example.com",
									"id": "zone-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "error_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "1000_errors",
									"state": "customized",
									"url": "https://example.com/1000.html",
									"id": "error-page-id"
								}
							}
						]
					}
				]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}

func TestTransformZeroTrustAccessApplicationStateJSON(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms_cors_headers_from_array_to_object",
			Input: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"identity_schema_version": 0,
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"type": "self_hosted",
							"cors_headers": [{
								"allowed_methods": ["GET", "POST", "OPTIONS"],
								"allowed_origins": ["https://example.com"],
								"allow_credentials": true,
								"max_age": 600
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [{
					"type": "cloudflare_zero_trust_access_application", 
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"type": "self_hosted",
							"cors_headers": {
								"allowed_methods": ["GET", "POST", "OPTIONS"],
								"allowed_origins": ["https://example.com"],
								"allow_credentials": true,
								"max_age": 600
							}
						},
						"identity_schema_version": 0
					}]
				}]
			}`,
		},
		{
			Name: "handles_empty_cors_headers_array",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"cors_headers": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123", 
							"name": "Test App",
							"cors_headers": null
						}
					}]
				}]
			}`,
		},
		{
			Name: "preserves_cors_headers_when_already_object",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App", 
							"cors_headers": {
								"allowed_methods": ["GET", "POST"],
								"allowed_origins": ["https://test.com"]
							}
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app", 
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"cors_headers": {
								"allowed_methods": ["GET", "POST"],
								"allowed_origins": ["https://test.com"]
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_landing_page_design_from_array_to_object",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"landing_page_design": [{
								"title": "Welcome",
								"message": "Please sign in",
								"image_url": "https://example.com/logo.png"
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"landing_page_design": {
								"title": "Welcome",
								"message": "Please sign in",
								"image_url": "https://example.com/logo.png"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_empty_landing_page_design_array",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"landing_page_design": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"landing_page_design": null
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_saas_app_from_array_to_object",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test SAAS App",
							"type": "saas",
							"saas_app": [{
								"consumer_service_url": "https://example.com/sso/saml/consume",
								"sp_entity_id": "example.com",
								"name_id_format": "email",
								"auth_type": "saml"
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test SAAS App",
							"type": "saas",
							"saas_app": {
								"consumer_service_url": "https://example.com/sso/saml/consume",
								"sp_entity_id": "example.com",
								"name_id_format": "email",
								"auth_type": "saml"
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_empty_saas_app_array",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"type": "saas",
							"saas_app": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"type": "saas",
							"saas_app": null
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_scim_config_from_array_to_object",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"type": "saas",
							"scim_config": [{
								"enabled": true,
								"remote_uri": "https://example.com/scim/v2",
								"deactivate_on_delete": true
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"type": "saas",
							"scim_config": {
								"enabled": true,
								"remote_uri": "https://example.com/scim/v2",
								"deactivate_on_delete": true
							}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_empty_scim_config_array",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"scim_config": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"scim_config": null
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_multiple_attributes_together",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"cors_headers": [{
								"allowed_methods": ["GET", "POST"],
								"allow_credentials": false
							}],
							"landing_page_design": [{
								"title": "Welcome",
								"message": "Please sign in"
							}],
							"saas_app": [{
								"consumer_service_url": "https://example.com/callback",
								"sp_entity_id": "example.com"
							}],
							"scim_config": [{
								"enabled": true,
								"remote_uri": "https://example.com/scim"
							}],
							"policies": ["policy-123"],
							"allowed_idps": ["idp-1", "idp-2"],
							"custom_pages": ["page-1"]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zero_trust_access_application",
					"name": "app",
					"instances": [{
						"attributes": {
							"id": "app-id-123",
							"name": "Test App",
							"cors_headers": {
								"allowed_methods": ["GET", "POST"],
								"allow_credentials": false
							},
							"landing_page_design": {
								"title": "Welcome",
								"message": "Please sign in"
							},
							"saas_app": {
								"consumer_service_url": "https://example.com/callback",
								"sp_entity_id": "example.com"
							},
							"scim_config": {
								"enabled": true,
								"remote_uri": "https://example.com/scim"
							},
							"policies": [{"id": "policy-123"}],
							"allowed_idps": ["idp-1", "idp-2"],
							"custom_pages": ["page-1"]
						}
					}]
				}]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}
