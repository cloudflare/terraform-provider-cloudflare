package main

import (
	"encoding/json"
	"testing"
)

func TestTransformLoadBalancerMonitorStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
	}{
		{
			name: "single header transformation",
			input: `{
				"attributes": {
					"header": [
						{
							"header": "Host",
							"values": ["api.example.com"]
						}
					]
				}
			}`,
			expected: map[string]interface{}{
				"Host": []string{"api.example.com"},
			},
		},
		{
			name: "multiple headers transformation",
			input: `{
				"attributes": {
					"header": [
						{
							"header": "Host",
							"values": ["api.example.com"]
						},
						{
							"header": "X-App-ID",
							"values": ["abc123", "def456"]
						}
					]
				}
			}`,
			expected: map[string]interface{}{
				"Host":     []string{"api.example.com"},
				"X-App-ID": []string{"abc123", "def456"},
			},
		},
		{
			name: "empty header array",
			input: `{
				"attributes": {
					"header": []
				}
			}`,
			expected: nil, // Should remove the attribute
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Transform the JSON
			result := transformLoadBalancerMonitorStateJSON(tt.input, "")

			t.Logf("Transformed JSON: %s", result)

			// Parse the result
			var resultData map[string]interface{}
			if err := json.Unmarshal([]byte(result), &resultData); err != nil {
				t.Fatalf("Failed to parse result JSON: %v", err)
			}

			// Get the header from attributes
			attrs, ok := resultData["attributes"].(map[string]interface{})
			if !ok {
				t.Fatal("Failed to get attributes from result")
			}

			header, exists := attrs["header"]
			if tt.expected == nil {
				if exists {
					t.Errorf("Expected header to be removed, but it exists: %v", header)
				}
				return
			}

			if !exists {
				t.Errorf("Expected header to exist, but it doesn't")
				return
			}

			// Compare the header
			headerMap, ok := header.(map[string]interface{})
			if !ok {
				t.Errorf("Header is not a map: %T", header)
				return
			}

			// Check all expected headers
			for key, expectedValues := range tt.expected {
				actualValues, ok := headerMap[key].([]interface{})
				if !ok {
					t.Errorf("Header %s not found or not an array", key)
					continue
				}

				expectedSlice := expectedValues.([]string)
				if len(actualValues) != len(expectedSlice) {
					t.Errorf("Header %s: expected %d values, got %d", key, len(expectedSlice), len(actualValues))
					continue
				}

				for i, v := range actualValues {
					if v.(string) != expectedSlice[i] {
						t.Errorf("Header %s[%d]: expected %s, got %s", key, i, expectedSlice[i], v.(string))
					}
				}
			}
		})
	}
}

func TestLoadBalancerMonitorFullStateTransformation(t *testing.T) {
	// Full state file test
	oldState := `{
		"version": 4,
		"terraform_version": "1.5.0",
		"resources": [
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "api-healthcheck",
				"instances": [
					{
						"attributes": {
							"id": "monitor123",
							"account_id": "account123",
							"expected_body": "healthy",
							"expected_codes": "200",
							"method": "GET",
							"timeout": 5,
							"path": "/health",
							"interval": 60,
							"retries": 2,
							"port": 443,
							"description": "API health check",
							"header": [
								{
									"header": "Host",
									"values": ["api.example.com"]
								},
								{
									"header": "X-Custom-Header",
									"values": ["value1", "value2"]
								}
							]
						}
					}
				]
			}
		]
	}`

	// Transform the state
	transformedBytes, err := transformStateJSON([]byte(oldState))
	if err != nil {
		t.Fatalf("Failed to transform state: %v", err)
	}

	// Parse the transformed state
	var transformedState map[string]interface{}
	if err := json.Unmarshal(transformedBytes, &transformedState); err != nil {
		t.Fatalf("Failed to parse transformed state: %v", err)
	}

	// Navigate to the header
	resources := transformedState["resources"].([]interface{})
	resource := resources[0].(map[string]interface{})
	instances := resource["instances"].([]interface{})
	instance := instances[0].(map[string]interface{})
	attributes := instance["attributes"].(map[string]interface{})
	header := attributes["header"].(map[string]interface{})

	// Verify the transformation
	if header["Host"] == nil {
		t.Error("Host header not found in transformed state")
	}
	hostValues := header["Host"].([]interface{})
	if len(hostValues) != 1 || hostValues[0].(string) != "api.example.com" {
		t.Errorf("Host header incorrect: %v", hostValues)
	}

	if header["X-Custom-Header"] == nil {
		t.Error("X-Custom-Header not found in transformed state")
	}
	customValues := header["X-Custom-Header"].([]interface{})
	if len(customValues) != 2 || customValues[0].(string) != "value1" || customValues[1].(string) != "value2" {
		t.Errorf("X-Custom-Header incorrect: %v", customValues)
	}

	t.Logf("Transformed state header: %v", header)
}

func TestComprehensiveLoadBalancerMonitorStateMigration(t *testing.T) {
	// Test with all the monitors mentioned in the issue
	oldState := `{
		"version": 4,
		"terraform_version": "1.5.0",
		"resources": [
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "api-healthcheck",
				"instances": [{
					"attributes": {
						"id": "monitor1",
						"account_id": "account123",
						"expected_body": "healthy",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/health",
						"interval": 60,
						"retries": 2,
						"port": 443,
						"description": "API health check",
						"header": [{"header": "Host", "values": ["api.example.com"]}]
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "chatgpt-snc",
				"instances": [{
					"attributes": {
						"id": "monitor2",
						"account_id": "account123",
						"expected_body": "ok",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/status",
						"interval": 30,
						"retries": 3,
						"port": 443,
						"description": "ChatGPT SNC monitor",
						"header": [
							{"header": "Host", "values": ["chatgpt-snc.example.com"]},
							{"header": "X-API-Key", "values": ["secret123"]}
						]
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "chatgpt-origins-monitor-favicon",
				"instances": [{
					"attributes": {
						"id": "monitor3",
						"account_id": "account123",
						"expected_body": "",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/favicon.ico",
						"interval": 60,
						"retries": 2,
						"port": 443,
						"description": "Favicon monitor",
						"header": []
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "kenny-test",
				"instances": [{
					"attributes": {
						"id": "monitor4",
						"account_id": "account123",
						"expected_body": "test",
						"expected_codes": "2xx",
						"method": "POST",
						"timeout": 10,
						"path": "/test",
						"interval": 120,
						"retries": 1,
						"port": 8080,
						"description": "Kenny's test monitor",
						"header": [
							{"header": "Content-Type", "values": ["application/json"]},
							{"header": "Authorization", "values": ["Bearer token123"]}
						]
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "api-staging-monitor",
				"instances": [{
					"attributes": {
						"id": "monitor5",
						"account_id": "account123",
						"expected_body": "staging",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/health",
						"interval": 60,
						"retries": 2,
						"port": 443,
						"description": "API staging monitor",
						"header": [{"header": "Host", "values": ["staging.api.example.com"]}]
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "staging-identity-edge-testing-monitor",
				"instances": [{
					"attributes": {
						"id": "monitor6",
						"account_id": "account123",
						"expected_body": "edge",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/edge",
						"interval": 60,
						"retries": 2,
						"port": 443,
						"description": "Edge testing monitor",
						"header": [{"header": "Host", "values": ["edge.staging.example.com"]}]
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "api-healthcheck-monitor",
				"instances": [{
					"attributes": {
						"id": "monitor7",
						"account_id": "account123",
						"expected_body": "alive",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/healthcheck",
						"interval": 60,
						"retries": 2,
						"port": 443,
						"description": "API healthcheck monitor",
						"header": [{"header": "Host", "values": ["api.example.com"]}]
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "sora-healthcheck-monitor",
				"instances": [{
					"attributes": {
						"id": "monitor8",
						"account_id": "account123",
						"expected_body": "sora",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/health",
						"interval": 60,
						"retries": 2,
						"port": 443,
						"description": "Sora healthcheck monitor",
						"header": [{"header": "Host", "values": ["sora.example.com"]}]
					}
				}]
			},
			{
				"type": "cloudflare_load_balancer_monitor",
				"name": "platform-healthcheck",
				"instances": [{
					"attributes": {
						"id": "monitor9",
						"account_id": "account123",
						"expected_body": "platform",
						"expected_codes": "200",
						"method": "GET",
						"timeout": 5,
						"path": "/health",
						"interval": 60,
						"retries": 2,
						"port": 443,
						"description": "Platform healthcheck",
						"header": [{"header": "Host", "values": ["platform.example.com"]}]
					}
				}]
			}
		]
	}`

	// Transform the state
	transformedBytes, err := transformStateJSON([]byte(oldState))
	if err != nil {
		t.Fatalf("Failed to transform state: %v", err)
	}

	// Parse the transformed state
	var transformedState map[string]interface{}
	if err := json.Unmarshal(transformedBytes, &transformedState); err != nil {
		t.Fatalf("Failed to parse transformed state: %v", err)
	}

	// Verify all monitors were transformed correctly
	resources := transformedState["resources"].([]interface{})

	// Test cases for each monitor
	testCases := []struct {
		index           int
		name            string
		expectedHeaders map[string][]string
	}{
		{0, "api-healthcheck", map[string][]string{"Host": {"api.example.com"}}},
		{1, "chatgpt-snc", map[string][]string{
			"Host":      {"chatgpt-snc.example.com"},
			"X-API-Key": {"secret123"},
		}},
		{2, "chatgpt-origins-monitor-favicon", nil}, // Empty array should be removed
		{3, "kenny-test", map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer token123"},
		}},
		{4, "api-staging-monitor", map[string][]string{"Host": {"staging.api.example.com"}}},
		{5, "staging-identity-edge-testing-monitor", map[string][]string{"Host": {"edge.staging.example.com"}}},
		{6, "api-healthcheck-monitor", map[string][]string{"Host": {"api.example.com"}}},
		{7, "sora-healthcheck-monitor", map[string][]string{"Host": {"sora.example.com"}}},
		{8, "platform-healthcheck", map[string][]string{"Host": {"platform.example.com"}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource := resources[tc.index].(map[string]interface{})
			instances := resource["instances"].([]interface{})
			instance := instances[0].(map[string]interface{})
			attributes := instance["attributes"].(map[string]interface{})

			// Check the name matches
			if resource["name"] != tc.name {
				t.Errorf("Expected name %s, got %s", tc.name, resource["name"])
			}

			// Check the header transformation
			header, exists := attributes["header"]
			if tc.expectedHeaders == nil {
				if exists {
					t.Errorf("Expected header to be removed for %s, but it exists: %v", tc.name, header)
				}
				return
			}

			if !exists {
				t.Errorf("Expected header to exist for %s, but it doesn't", tc.name)
				return
			}

			headerMap, ok := header.(map[string]interface{})
			if !ok {
				t.Errorf("Header for %s is not a map: %T", tc.name, header)
				return
			}

			// Verify each expected header
			for key, expectedValues := range tc.expectedHeaders {
				actualValues, ok := headerMap[key].([]interface{})
				if !ok {
					t.Errorf("%s: Header %s not found or not an array", tc.name, key)
					continue
				}

				if len(actualValues) != len(expectedValues) {
					t.Errorf("%s: Header %s: expected %d values, got %d", tc.name, key, len(expectedValues), len(actualValues))
					continue
				}

				for i, v := range actualValues {
					if v.(string) != expectedValues[i] {
						t.Errorf("%s: Header %s[%d]: expected %s, got %s", tc.name, key, i, expectedValues[i], v.(string))
					}
				}
			}

			t.Logf("%s: Header successfully transformed to %v", tc.name, headerMap)
		})
	}
}
