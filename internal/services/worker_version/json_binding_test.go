package worker_version_test

import (
	"context"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_version"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestJsonBindingMarshal verifies that the json field in bindings is properly
// serialized as raw JSON, not as a double-encoded string.
// This is a regression test for https://github.com/cloudflare/terraform-provider-cloudflare/issues/6699
func TestJsonBindingMarshal(t *testing.T) {
	t.Parallel()

	binding := worker_version.WorkerVersionBindingsModel{
		Name: types.StringValue("JSON"),
		Type: types.StringValue("json"),
		Json: jsontypes.NewNormalizedValue(`{"key1":"value1","key2":"value2"}`),
	}

	data, err := apijson.Marshal(binding)
	if err != nil {
		t.Fatalf("Failed to marshal binding: %v", err)
	}

	jsonStr := string(data)

	// The JSON value should appear as a raw object, NOT as an escaped string
	// Good: "json":{"key1":"value1","key2":"value2"}
	// Bad:  "json":"{\"key1\":\"value1\",\"key2\":\"value2\"}"
	if strings.Contains(jsonStr, `"{\"key1\"`) {
		t.Errorf("JSON binding is double-encoded (bug #6699):\n  got: %s", jsonStr)
	}

	if !strings.Contains(jsonStr, `"json":{"key1":"value1","key2":"value2"}`) {
		t.Errorf("JSON binding should contain raw JSON object:\n  got: %s\n  want to contain: %s", jsonStr, `"json":{"key1":"value1","key2":"value2"}`)
	}
}

// TestJsonBindingUnmarshal verifies that the json field in bindings is properly
// unmarshaled from API responses.
func TestJsonBindingUnmarshal(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	apiResponse := []byte(`{
		"result": {
			"id": "test-version-id",
			"number": 1,
			"source": "api",
			"usage_model": "standard",
			"created_on": "2024-04-04T21:18:22.776392Z",
			"main_module": "index.js",
			"bindings": [
				{
					"name": "JSON_VAR",
					"type": "json",
					"json": {"key1": "value1", "key2": "value2"}
				}
			]
		}
	}`)

	var envelope worker_version.WorkerVersionResultEnvelope
	err := apijson.Unmarshal(apiResponse, &envelope)
	if err != nil {
		t.Fatalf("Failed to unmarshal API response: %v", err)
	}

	bindings, diags := envelope.Result.Bindings.AsStructSliceT(ctx)
	if diags.HasError() {
		t.Fatalf("Failed to get bindings: %v", diags)
	}

	if len(bindings) != 1 {
		t.Fatalf("Expected 1 binding, got %d", len(bindings))
	}

	binding := bindings[0]

	if binding.Name.ValueString() != "JSON_VAR" {
		t.Errorf("Name mismatch: got %q, want %q", binding.Name.ValueString(), "JSON_VAR")
	}

	if binding.Type.ValueString() != "json" {
		t.Errorf("Type mismatch: got %q, want %q", binding.Type.ValueString(), "json")
	}

	// The JSON value should be preserved as the raw JSON object string
	jsonValue := binding.Json.ValueString()
	expectedJson := `{"key1":"value1","key2":"value2"}`

	// Normalize whitespace for comparison
	if strings.ReplaceAll(jsonValue, " ", "") != strings.ReplaceAll(expectedJson, " ", "") {
		t.Errorf("JSON value mismatch:\n  got:  %q\n  want: %q", jsonValue, expectedJson)
	}
}

// TestJsonBindingFullModel tests the complete worker version model with a JSON
// binding, simulating the exact scenario from issue #6699.
func TestJsonBindingFullModel(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Create a model similar to what the user would configure in Terraform
	model := worker_version.WorkerVersionModel{
		AccountID:  types.StringValue("test-account-id"),
		WorkerID:   types.StringValue("test-worker-id"),
		MainModule: types.StringValue("index.js"),
	}

	// Create bindings with a JSON type - this simulates:
	// bindings = [{
	//   name = "JSON"
	//   type = "json"
	//   json = jsonencode({ key1 = "value1", key2 = "value2" })
	// }]
	bindings := []worker_version.WorkerVersionBindingsModel{
		{
			Name: types.StringValue("JSON"),
			Type: types.StringValue("json"),
			// jsonencode({ key1 = "value1", key2 = "value2" }) produces this string
			Json: jsontypes.NewNormalizedValue(`{"key1":"value1","key2":"value2"}`),
		},
	}

	// Set the bindings on the model
	bindingsList, diags := customfield.NewObjectList(ctx, bindings)
	if diags.HasError() {
		t.Fatalf("Failed to create bindings list: %v", diags)
	}
	model.Bindings = bindingsList

	// Marshal the model to JSON (this is what gets sent to the API)
	data, err := model.MarshalJSON()
	if err != nil {
		t.Fatalf("Failed to marshal model: %v", err)
	}

	jsonStr := string(data)

	// Verify the JSON is NOT double-encoded
	// Bug #6699: The value was being sent as "{\"key1\":\"value1\",\"key2\":\"value2\"}"
	// instead of {"key1":"value1","key2":"value2"}
	if strings.Contains(jsonStr, `\\"key1\\"`) || strings.Contains(jsonStr, `"{\"key1\"`) {
		t.Errorf("JSON binding is double-encoded (bug #6699):\n  API request body: %s", jsonStr)
	}

	// Verify the JSON object is present as raw JSON
	if !strings.Contains(jsonStr, `"json":{"key1":"value1","key2":"value2"}`) {
		t.Errorf("JSON binding should contain raw JSON object in API request.\n  got: %s\n  want to contain: %s",
			jsonStr, `"json":{"key1":"value1","key2":"value2"}`)
	}

	t.Logf("API request body (correct format): %s", jsonStr)
}

// TestJsonBindingNotDoubleEncoded is a negative test that explicitly verifies
// the old buggy behavior (double-encoding) is NOT happening.
func TestJsonBindingNotDoubleEncoded(t *testing.T) {
	t.Parallel()

	// This test demonstrates what the bug looked like:
	// When json was types.String, the value {"key1":"value1"} would be
	// serialized as "{\"key1\":\"value1\"}" (a JSON-encoded string containing JSON)

	binding := worker_version.WorkerVersionBindingsModel{
		Name: types.StringValue("CONFIG"),
		Type: types.StringValue("json"),
		Json: jsontypes.NewNormalizedValue(`{"nested":{"deep":"value"},"array":[1,2,3]}`),
	}

	data, err := apijson.Marshal(binding)
	if err != nil {
		t.Fatalf("Failed to marshal binding: %v", err)
	}

	jsonStr := string(data)

	// These patterns would appear if double-encoding was happening
	doubleEncodedPatterns := []string{
		`"{\"`,      // Opening quote before escaped quote
		`\"}"`,      // Escaped quote before closing quote
		`\\\"`,      // Double-escaped quotes
		`\\"nested`, // Escaped quote before key
	}

	for _, pattern := range doubleEncodedPatterns {
		if strings.Contains(jsonStr, pattern) {
			t.Errorf("Found double-encoding pattern %q in output.\nThis indicates bug #6699 is NOT fixed.\nOutput: %s", pattern, jsonStr)
		}
	}

	// Verify correct patterns ARE present (raw JSON structure)
	correctPatterns := []string{
		`"json":{`,        // json field starts an object
		`"nested":{`,      // nested object
		`"array":[1,2,3]`, // array with numbers
	}

	for _, pattern := range correctPatterns {
		if !strings.Contains(jsonStr, pattern) {
			t.Errorf("Missing expected pattern %q in output.\nOutput: %s", pattern, jsonStr)
		}
	}

	t.Logf("Correctly formatted output: %s", jsonStr)
}

// TestJsonBindingBackwardsCompatibility verifies that the type change from
// types.String to jsontypes.Normalized is backwards compatible with existing
// Terraform state. This test simulates reading old state values.
func TestJsonBindingBackwardsCompatibility(t *testing.T) {
	t.Parallel()

	// Simulate an API response that contains a JSON binding
	// This is what would be stored in Cloudflare and returned by the API
	apiResponseWithJsonBinding := []byte(`{
		"result": {
			"id": "test-version-id",
			"number": 1,
			"source": "api",
			"usage_model": "standard",
			"created_on": "2024-04-04T21:18:22.776392Z",
			"main_module": "index.js",
			"bindings": [
				{
					"name": "MY_JSON",
					"type": "json",
					"json": {"config": {"timeout": 30, "enabled": true}}
				}
			]
		}
	}`)

	// Unmarshal should work correctly with the new jsontypes.Normalized type
	var envelope worker_version.WorkerVersionResultEnvelope
	err := apijson.Unmarshal(apiResponseWithJsonBinding, &envelope)
	if err != nil {
		t.Fatalf("Failed to unmarshal API response: %v", err)
	}

	// Verify we can access the binding
	ctx := context.Background()
	bindings, diags := envelope.Result.Bindings.AsStructSliceT(ctx)
	if diags.HasError() {
		t.Fatalf("Failed to get bindings: %v", diags)
	}

	if len(bindings) != 1 {
		t.Fatalf("Expected 1 binding, got %d", len(bindings))
	}

	binding := bindings[0]

	// Verify the json value is correctly parsed
	jsonValue := binding.Json.ValueString()
	if jsonValue == "" {
		t.Error("JSON value should not be empty")
	}

	// The value should contain the expected keys (order may vary)
	if !strings.Contains(jsonValue, `"timeout"`) || !strings.Contains(jsonValue, `"enabled"`) {
		t.Errorf("JSON value missing expected keys: %s", jsonValue)
	}

	t.Logf("Successfully read JSON binding value: %s", jsonValue)

	// Now test that we can marshal it back correctly
	data, err := apijson.Marshal(binding)
	if err != nil {
		t.Fatalf("Failed to marshal binding: %v", err)
	}

	// Verify no double-encoding in the output
	output := string(data)
	if strings.Contains(output, `\\"`) {
		t.Errorf("Output contains double-encoding: %s", output)
	}

	t.Logf("Successfully marshaled binding: %s", output)
}
