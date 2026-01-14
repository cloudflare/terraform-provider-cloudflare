package worker_version_test

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_version"
)

// TestMainScriptBase64Unmarshal verifies that the main_script_base64 field
// is properly unmarshaled from API responses. This field is returned by the
// Cloudflare API for workers using the older "service worker syntax"
// (e.g., addEventListener('fetch', ...)) instead of ES modules.
func TestMainScriptBase64Unmarshal(t *testing.T) {
	t.Parallel()

	// This is a sample API response for a service worker syntax worker.
	// The main_script_base64 field contains base64-encoded JavaScript code.
	apiResponse := []byte(`{
		"result": {
			"id": "9560de65-bb66-4b75-8c73-d4d143fe09ef",
			"number": 1,
			"source": "unknown",
			"usage_model": "bundled",
			"created_on": "2024-04-04T21:18:22.776392Z",
			"main_script_base64": "YWRkRXZlbnRMaXN0ZW5lcignZmV0Y2gnLCBldmVudCA9PiB7IGV2ZW50LnJlc3BvbmRXaXRoKG5ldyBSZXNwb25zZSgnSGVsbG8nKSkgfSk=",
			"annotations": {
				"workers/message": "Automatic migration",
				"workers/triggered_by": "upload"
			},
			"bindings": [
				{
					"name": "MY_VAR",
					"type": "plain_text",
					"text": "hello"
				}
			]
		}
	}`)

	var envelope worker_version.WorkerVersionResultEnvelope
	err := apijson.Unmarshal(apiResponse, &envelope)
	if err != nil {
		t.Fatalf("Failed to unmarshal API response: %v", err)
	}

	// Verify that main_script_base64 was properly populated
	mainScriptBase64 := envelope.Result.MainScriptBase64.ValueString()
	expectedBase64 := "YWRkRXZlbnRMaXN0ZW5lcignZmV0Y2gnLCBldmVudCA9PiB7IGV2ZW50LnJlc3BvbmRXaXRoKG5ldyBSZXNwb25zZSgnSGVsbG8nKSkgfSk="

	if mainScriptBase64 != expectedBase64 {
		t.Errorf("main_script_base64 mismatch:\n  got:  %q\n  want: %q", mainScriptBase64, expectedBase64)
	}

	// Verify that other fields were also properly populated
	if envelope.Result.ID.ValueString() != "9560de65-bb66-4b75-8c73-d4d143fe09ef" {
		t.Errorf("ID mismatch: got %q", envelope.Result.ID.ValueString())
	}

	if envelope.Result.Number.ValueInt64() != 1 {
		t.Errorf("Number mismatch: got %d", envelope.Result.Number.ValueInt64())
	}

	if envelope.Result.Source.ValueString() != "unknown" {
		t.Errorf("Source mismatch: got %q", envelope.Result.Source.ValueString())
	}

	if envelope.Result.UsageModel.ValueString() != "bundled" {
		t.Errorf("UsageModel mismatch: got %q", envelope.Result.UsageModel.ValueString())
	}
}

// TestMainScriptBase64UnmarshalNull verifies that main_script_base64 is null
// when not present in the API response (e.g., for ES module workers).
func TestMainScriptBase64UnmarshalNull(t *testing.T) {
	t.Parallel()

	// This is a sample API response for an ES module worker.
	// The main_script_base64 field is not present.
	apiResponse := []byte(`{
		"result": {
			"id": "abc123",
			"number": 1,
			"source": "api",
			"usage_model": "standard",
			"created_on": "2024-04-04T21:18:22.776392Z",
			"main_module": "index.js"
		}
	}`)

	var envelope worker_version.WorkerVersionResultEnvelope
	err := apijson.Unmarshal(apiResponse, &envelope)
	if err != nil {
		t.Fatalf("Failed to unmarshal API response: %v", err)
	}

	// Verify that main_script_base64 is null when not present
	if !envelope.Result.MainScriptBase64.IsNull() {
		t.Errorf("main_script_base64 should be null when not in response, got: %q", envelope.Result.MainScriptBase64.ValueString())
	}

	// Verify that main_module was populated instead
	if envelope.Result.MainModule.ValueString() != "index.js" {
		t.Errorf("MainModule mismatch: got %q", envelope.Result.MainModule.ValueString())
	}
}

// TestDataSourceMainScriptBase64Unmarshal verifies that the main_script_base64
// field is properly unmarshaled for the data source model.
func TestDataSourceMainScriptBase64Unmarshal(t *testing.T) {
	t.Parallel()

	apiResponse := []byte(`{
		"result": {
			"id": "test-version-id",
			"number": 1,
			"source": "unknown",
			"usage_model": "bundled",
			"created_on": "2024-04-04T21:18:22.776392Z",
			"main_script_base64": "Y29uc29sZS5sb2coJ2hlbGxvJyk="
		}
	}`)

	var envelope worker_version.WorkerVersionResultDataSourceEnvelope
	err := apijson.Unmarshal(apiResponse, &envelope)
	if err != nil {
		t.Fatalf("Failed to unmarshal API response: %v", err)
	}

	mainScriptBase64 := envelope.Result.MainScriptBase64.ValueString()
	expectedBase64 := "Y29uc29sZS5sb2coJ2hlbGxvJyk="

	if mainScriptBase64 != expectedBase64 {
		t.Errorf("main_script_base64 mismatch:\n  got:  %q\n  want: %q", mainScriptBase64, expectedBase64)
	}
}

