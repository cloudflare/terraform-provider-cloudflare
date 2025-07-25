package argo_tunnel

import (
	"context"
	"testing"
)

func TestArgoTunnelSchema_MatchesV3(t *testing.T) {
	ctx := context.Background()
	schema := ResourceSchema(ctx)

	// Verify required attributes exist
	requiredAttrs := []string{"account_id", "name", "secret"}
	for _, attr := range requiredAttrs {
		if schema.Attributes[attr] == nil {
			t.Errorf("Required attribute '%s' missing from schema", attr)
		}
	}

	// Verify computed attributes exist
	computedAttrs := []string{"cname", "tunnel_token"}
	for _, attr := range computedAttrs {
		if schema.Attributes[attr] == nil {
			t.Errorf("Computed attribute '%s' missing from schema", attr)
		}
	}

	// Verify no extra attributes that weren't in v3
	expectedAttrs := map[string]bool{
		"id":           true,
		"account_id":   true,
		"name":         true,
		"secret":       true,
		"cname":        true,
		"tunnel_token": true,
	}

	for attr := range schema.Attributes {
		if !expectedAttrs[attr] {
			t.Errorf("Unexpected attribute '%s' in schema - not present in v3", attr)
		}
	}

	// Verify secret is marked as sensitive
	if !schema.Attributes["secret"].IsSensitive() {
		t.Error("Secret attribute should be marked as sensitive")
	}

	// Verify tunnel_token is marked as sensitive
	if !schema.Attributes["tunnel_token"].IsSensitive() {
		t.Error("Tunnel token attribute should be marked as sensitive")
	}
}
