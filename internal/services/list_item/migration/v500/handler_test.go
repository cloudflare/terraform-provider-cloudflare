package v500_test

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item"
)

// TestUpgradeFromV1Ambiguous_V5HostnameState simulates the exact scenario from
// GitHub issue #7073: a cloudflare_list_item with hostname created natively in
// v5 (schema_version=1) must survive the upgrade to schema_version=500 without error.
//
// In v5, hostname is a SingleNestedAttribute stored as a JSON object:
//
//	{"url_hostname": "example.com", "exclude_exact_hostname": null}
//
// The v4.52.5 format would be an array: [{"url_hostname": "example.com"}]
//
// UpgradeFromV1Ambiguous must detect the object format and perform a no-op
// version bump instead of trying to parse it as a v4.52.5 array.
func TestUpgradeFromV1Ambiguous_V5HostnameState(t *testing.T) {
	t.Parallel()

	// This is the exact state shape stored by v5.0-v5.18 for a hostname list item.
	v5State := map[string]interface{}{
		"account_id": "abc123",
		"list_id":    "list456",
		"id":         "item789",
		"ip":         nil,
		"asn":        nil,
		"comment":    "test hostname",
		"hostname": map[string]interface{}{
			"url_hostname":           "example.com",
			"exclude_exact_hostname": nil,
		},
		"redirect":     nil,
		"operation_id": nil,
		"modified_on":  nil,
		"created_on":   nil,
	}

	result := runAmbiguousUpgradeTest(t, v5State)

	hostname := getNestedObject(t, result, "hostname")
	assertStringField(t, hostname, "url_hostname", "example.com")
}

// TestUpgradeFromV1Ambiguous_V5RedirectState tests that a v5 native redirect list item
// (stored as a JSON object) survives the ambiguous upgrade.
func TestUpgradeFromV1Ambiguous_V5RedirectState(t *testing.T) {
	t.Parallel()

	v5State := map[string]interface{}{
		"account_id": "abc123",
		"list_id":    "list456",
		"id":         "item789",
		"ip":         nil,
		"asn":        nil,
		"comment":    "test redirect",
		"hostname":   nil,
		"redirect": map[string]interface{}{
			"source_url":            "https://example.com/old",
			"target_url":            "https://example.com/new",
			"status_code":           float64(301),
			"include_subdomains":    true,
			"subpath_matching":      false,
			"preserve_query_string": true,
			"preserve_path_suffix":  false,
		},
		"operation_id": nil,
		"modified_on":  nil,
		"created_on":   nil,
	}

	result := runAmbiguousUpgradeTest(t, v5State)

	redirect := getNestedObject(t, result, "redirect")
	assertStringField(t, redirect, "source_url", "https://example.com/old")
	assertStringField(t, redirect, "target_url", "https://example.com/new")
}

// TestUpgradeFromV1Ambiguous_V5IPOnlyState tests that an IP-only list item (no hostname
// or redirect) also survives the ambiguous upgrade (no-op path).
func TestUpgradeFromV1Ambiguous_V5IPOnlyState(t *testing.T) {
	t.Parallel()

	v5State := map[string]interface{}{
		"account_id":   "abc123",
		"list_id":      "list456",
		"id":           "item789",
		"ip":           "192.0.2.1",
		"asn":          nil,
		"comment":      "test ip",
		"hostname":     nil,
		"redirect":     nil,
		"operation_id": nil,
		"modified_on":  nil,
		"created_on":   nil,
	}

	result := runAmbiguousUpgradeTest(t, v5State)

	assertStringField(t, result, "ip", "192.0.2.1")
	assertStringField(t, result, "comment", "test ip")
}

// TestUpgradeFromV1Ambiguous_V4HostnameState tests that a v4.52.5 hostname list item
// (stored as a JSON array) is correctly detected and transformed to v5 object format.
func TestUpgradeFromV1Ambiguous_V4HostnameState(t *testing.T) {
	t.Parallel()

	// v4.52.5 stores hostname as an array (ListNestedBlock)
	v4State := map[string]interface{}{
		"account_id": "abc123",
		"list_id":    "list456",
		"id":         "item789",
		"ip":         nil,
		"asn":        nil,
		"comment":    "test hostname v4",
		"hostname": []interface{}{
			map[string]interface{}{
				"url_hostname": "example.com",
			},
		},
		"redirect": []interface{}{},
	}

	result := runAmbiguousUpgradeTest(t, v4State)

	// After upgrade, hostname should be an object (not an array)
	hostname := getNestedObject(t, result, "hostname")
	assertStringField(t, hostname, "url_hostname", "example.com")
}

// TestUpgradeFromV1Ambiguous_V4RedirectState tests that a v4.52.5 redirect list item
// (stored as a JSON array) is correctly detected and transformed to v5 object format.
func TestUpgradeFromV1Ambiguous_V4RedirectState(t *testing.T) {
	t.Parallel()

	v4State := map[string]interface{}{
		"account_id": "abc123",
		"list_id":    "list456",
		"id":         "item789",
		"ip":         nil,
		"asn":        nil,
		"comment":    "test redirect v4",
		"hostname":   []interface{}{},
		"redirect": []interface{}{
			map[string]interface{}{
				"source_url":            "https://example.com/old",
				"target_url":            "https://example.com/new",
				"status_code":           float64(301),
				"include_subdomains":    true,
				"subpath_matching":      false,
				"preserve_query_string": true,
				"preserve_path_suffix":  false,
			},
		},
	}

	result := runAmbiguousUpgradeTest(t, v4State)

	// After upgrade, redirect should be an object (not an array)
	redirect := getNestedObject(t, result, "redirect")
	assertStringField(t, redirect, "source_url", "https://example.com/old")
	assertStringField(t, redirect, "target_url", "https://example.com/new")
}

// runAmbiguousUpgradeTest constructs the UpgradeStateRequest/Response with
// PriorSchema=nil (matching how migrations.go registers version 1), invokes
// UpgradeFromV1Ambiguous, and returns the resulting state as a JSON map.
//
// It verifies no diagnostics errors occurred during the upgrade.
func runAmbiguousUpgradeTest(t *testing.T, stateJSON map[string]interface{}) map[string]interface{} {
	t.Helper()

	ctx := context.Background()

	rawJSON, err := json.Marshal(stateJSON)
	if err != nil {
		t.Fatalf("failed to marshal test state: %v", err)
	}

	// Build the target schema (v5/v500) that resp.State will use.
	targetSchema := list_item.ResourceSchema(ctx)

	resp := &resource.UpgradeStateResponse{
		State: tfsdk.State{
			Schema: targetSchema,
		},
	}

	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{
			JSON: rawJSON,
		},
		// State is nil because PriorSchema is nil in the migrations.go registration
		State: nil,
	}

	v500.UpgradeFromV1Ambiguous(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		for _, d := range resp.Diagnostics.Errors() {
			t.Errorf("diagnostic error: %s: %s", d.Summary(), d.Detail())
		}
		t.FailNow()
	}

	// Extract the resulting tftypes.Value into a Go map for field-level assertions.
	// We cannot use resp.State.Get into a Go struct because the migration
	// handler writes TargetListItemModel (with NestedObject[TargetHostnameModel])
	// while the schema's CustomType expects NestedObject[ListItemHostnameModel].
	// The types are structurally identical so tftypes works fine, but Go generics
	// prevent direct Get deserialization. Instead, extract via tftypes.Value.As().
	var rawMap map[string]tftypes.Value
	if err := resp.State.Raw.As(&rawMap); err != nil {
		t.Fatalf("failed to extract state as map: %v", err)
	}

	result := tftypesMapToInterface(t, rawMap)
	return result
}

// getNestedObject extracts a nested object field from the result map.
func getNestedObject(t *testing.T, result map[string]interface{}, key string) map[string]interface{} {
	t.Helper()
	val, ok := result[key]
	if !ok || val == nil {
		t.Fatalf("expected %q field to be present and non-nil", key)
	}
	obj, ok := val.(map[string]interface{})
	if !ok {
		t.Fatalf("expected %q to be an object, got %T", key, val)
	}
	return obj
}

// assertStringField checks that a string field in a map has the expected value.
func assertStringField(t *testing.T, obj map[string]interface{}, key, expected string) {
	t.Helper()
	val, ok := obj[key]
	if !ok {
		t.Errorf("expected field %q to be present", key)
		return
	}
	str, ok := val.(string)
	if !ok {
		t.Errorf("expected field %q to be a string, got %T", key, val)
		return
	}
	if str != expected {
		t.Errorf("expected %s=%q, got %q", key, expected, str)
	}
}

// tftypesMapToInterface converts a map[string]tftypes.Value into a
// map[string]interface{} for easy test assertions. Handles nested objects
// and primitive types (string, number, bool, nil).
func tftypesMapToInterface(t *testing.T, m map[string]tftypes.Value) map[string]interface{} {
	t.Helper()
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		result[k] = tftypesValueToInterface(t, v)
	}
	return result
}

func tftypesValueToInterface(t *testing.T, v tftypes.Value) interface{} {
	t.Helper()

	if !v.IsKnown() || v.IsNull() {
		return nil
	}

	// Try string
	var s string
	if err := v.As(&s); err == nil {
		return s
	}

	// Try bool
	var b bool
	if err := v.As(&b); err == nil {
		return b
	}

	// Try number (via big.Float -> float64)
	var bf *big.Float
	if err := v.As(&bf); err == nil {
		f, _ := bf.Float64()
		return f
	}

	// Try object (map)
	var m map[string]tftypes.Value
	if err := v.As(&m); err == nil {
		return tftypesMapToInterface(t, m)
	}

	// Try list/set/tuple
	var l []tftypes.Value
	if err := v.As(&l); err == nil {
		result := make([]interface{}, len(l))
		for i, item := range l {
			result[i] = tftypesValueToInterface(t, item)
		}
		return result
	}

	t.Fatalf("unsupported tftypes.Value type: %s", v.Type())
	return nil
}
