package main

import (
	"encoding/json"
	"testing"
	
	"github.com/tidwall/gjson"
)

// State transformation tests for zero_trust_access_group
func TestZeroTrustAccessGroupStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms_v4_include_with_email_arrays",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [{
					"email": ["user1@example.com", "user2@example.com"]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"email": {"email": "user1@example.com"}},
					{"email": {"email": "user2@example.com"}}
				]
			}`,
		},
		{
			Name: "transforms_v4_include_with_email_domain_arrays",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [{
					"email_domain": ["example.com", "test.com"]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"email_domain": {"domain": "example.com"}},
					{"email_domain": {"domain": "test.com"}}
				]
			}`,
		},
		{
			Name: "transforms_v4_include_with_ip_arrays",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [{
					"ip": ["192.0.2.1/32", "192.0.2.2/32"]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"ip": {"ip": "192.0.2.1/32"}},
					{"ip": {"ip": "192.0.2.2/32"}}
				]
			}`,
		},
		{
			Name: "transforms_v4_boolean_fields_to_empty_objects",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [{
					"everyone": true,
					"certificate": true,
					"any_valid_service_token": true
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"everyone": {}},
					{"certificate": {}},
					{"any_valid_service_token": {}}
				]
			}`,
		},
		{
			Name: "handles_exclude_rules",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"exclude": [{
					"email": ["blocked@example.com"],
					"ip": ["192.168.1.1/32"]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"exclude": [
					{"email": {"email": "blocked@example.com"}},
					{"ip": {"ip": "192.168.1.1/32"}}
				]
			}`,
		},
		{
			Name: "handles_require_rules",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"require": [{
					"email_domain": ["company.com"],
					"certificate": true
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"require": [
					{"email_domain": {"domain": "company.com"}},
					{"certificate": {}}
				]
			}`,
		},
		{
			Name: "handles_v5_format_unchanged",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"email": {"email": "user@example.com"}},
					{"everyone": {}}
				]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"email": {"email": "user@example.com"}},
					{"everyone": {}}
				]
			}`,
		},
		{
			Name: "handles_mixed_rule_types",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [{
					"email": ["user1@example.com", "user2@example.com"],
					"ip": ["10.0.0.1/32"],
					"everyone": true
				}],
				"exclude": [{
					"geo": ["CN", "RU"]
				}],
				"require": [{
					"email_domain": ["company.com"]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"email": {"email": "user1@example.com"}},
					{"email": {"email": "user2@example.com"}},
					{"ip": {"ip": "10.0.0.1/32"}},
					{"everyone": {}}
				],
				"exclude": [
					{"geo": {"country_code": "CN"}},
					{"geo": {"country_code": "RU"}}
				],
				"require": [
					{"email_domain": {"domain": "company.com"}}
				]
			}`,
		},
		{
			Name: "handles_empty_rules",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [],
				"exclude": [],
				"require": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [],
				"exclude": [],
				"require": []
			}`,
		},
		{
			Name: "handles_additional_attributes",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [{
					"group": ["group1", "group2"],
					"service_token": ["token1"],
					"email_list": ["list1"],
					"ip_list": ["iplist1"],
					"login_method": ["method1"],
					"device_posture": ["posture1"]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"group": {"id": "group1"}},
					{"group": {"id": "group2"}},
					{"service_token": {"token_id": "token1"}},
					{"email_list": {"id": "list1"}},
					{"ip_list": {"id": "iplist1"}},
					{"login_method": {"id": "method1"}},
					{"device_posture": {"integration_uid": "posture1"}}
				]
			}`,
		},
		{
			Name: "handles_geo_attribute",
			Input: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [{
					"geo": ["US", "CA", "GB"]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-group",
				"account_id": "acc-123",
				"include": [
					{"geo": {"country_code": "US"}},
					{"geo": {"country_code": "CA"}},
					{"geo": {"country_code": "GB"}}
				]
			}`,
		},
	}

	// Run the tests using the full state transformation (since we operate on JSON)
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			// Wrap the input JSON in a minimal state structure
			stateJSON := `{
				"resources": [{
					"type": "cloudflare_zero_trust_access_group",
					"instances": [{
						"attributes": ` + tc.Input + `
					}]
				}]
			}`

			// Transform using the state transformation function
			result := transformZeroTrustAccessGroupStateJSON(stateJSON, "resources.0.instances.0")

			// Extract the transformed attributes
			transformedAttrs := gjson.Get(result, "resources.0.instances.0.attributes")

			// Compare the JSON structures
			var expectedMap, actualMap interface{}
			json.Unmarshal([]byte(tc.Expected), &expectedMap)
			json.Unmarshal([]byte(transformedAttrs.Raw), &actualMap)

			expectedJSON, _ := json.Marshal(expectedMap)
			actualJSON, _ := json.Marshal(actualMap)

			if string(expectedJSON) != string(actualJSON) {
				// For better error output, pretty print both
				expectedPretty, _ := json.MarshalIndent(expectedMap, "", "  ")
				actualPretty, _ := json.MarshalIndent(actualMap, "", "  ")
				t.Errorf("Transformation failed\nExpected:\n%s\n\nGot:\n%s", expectedPretty, actualPretty)
			}
		})
	}
}