package v500_test

import (
	"context"
	_ "embed"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone/migration/v500"
)

const v4ZoneState = `{
	"id": "023e105f4ecef8ad9ca31a8372d0c353",
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"zone": "example.com",
	"jump_start": false,
	"paused": false,
	"plan": "free",
	"type": "full"
}`

// v5 wrote zone state at schema version 0 through 5.16.0; 5.17.0 bumped it to 1
// (BUGS-2027).
//
// A 5.16-era state with every attribute of that release, including nulls.
//
//go:embed testdata/v5_16_state_v0.json
var v5_16ZoneStateAtV0 string

// A 5.0-era state, before plan, tenant, tenant_unit, permissions, cname_suffix
// and verification_key existed, so parsing it with the current schema relies on
// every later change being additive.
//
//go:embed testdata/v5_0_state_v0.json
var v5_0ZoneStateAtV0 string

func upgradeV0(t *testing.T, raw string) (*resource.UpgradeStateResponse, map[string]tftypes.Value) {
	t.Helper()
	ctx := context.Background()
	s := zone.ResourceSchema(ctx)
	upgraders := (&zone.ZoneResource{}).UpgradeState(ctx)
	req := resource.UpgradeStateRequest{RawState: &tfprotov6.RawState{JSON: []byte(raw)}}
	resp := &resource.UpgradeStateResponse{State: tfsdk.State{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}}
	upgraders[0].StateUpgrader(ctx, req, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("upgrade failed: %v", resp.Diagnostics)
	}
	var attrs map[string]tftypes.Value
	if err := resp.State.Raw.As(&attrs); err != nil {
		t.Fatalf("reading state: %v", err)
	}
	return resp, attrs
}

func TestIsV5State(t *testing.T) {
	t.Parallel()
	for name, tc := range map[string]struct {
		raw  *tfprotov6.RawState
		want bool
	}{
		"v4":    {&tfprotov6.RawState{JSON: []byte(v4ZoneState)}, false},
		"v5.16": {&tfprotov6.RawState{JSON: []byte(v5_16ZoneStateAtV0)}, true},
		"v5.0":  {&tfprotov6.RawState{JSON: []byte(v5_0ZoneStateAtV0)}, true},
		// Neither version writes both keys. Treat it as v4, so the v4 parse
		// reports the unexpected attribute.
		"account and zone": {&tfprotov6.RawState{JSON: []byte(`{"id":"023e105f4ecef8ad9ca31a8372d0c353","zone":"example.com","account":{"id":"f037e56e89293a057740de681ac9abbe"}}`)}, false},
		// v4 never writes an account key, so a null one is still v5 shape.
		"null account": {&tfprotov6.RawState{JSON: []byte(`{"id":"023e105f4ecef8ad9ca31a8372d0c353","name":"example.com","account":null}`)}, true},
		"nil":          {nil, false},
		"empty":        {&tfprotov6.RawState{}, false},
	} {
		if got := v500.IsV5State(tc.raw); got != tc.want {
			t.Errorf("%s: IsV5State = %v, want %v", name, got, tc.want)
		}
	}
}

func TestUpgradeFromV0_V4State(t *testing.T) {
	t.Parallel()
	_, attrs := upgradeV0(t, v4ZoneState)
	var name string
	if err := attrs["name"].As(&name); err != nil || name != "example.com" {
		t.Fatalf("name = %q, %v; want example.com", name, err)
	}
}

func stringAttr(t *testing.T, v tftypes.Value) string {
	t.Helper()
	var s string
	if err := v.As(&s); err != nil {
		t.Fatalf("reading string: %v", err)
	}
	return s
}

func objectAttr(t *testing.T, v tftypes.Value) map[string]tftypes.Value {
	t.Helper()
	var m map[string]tftypes.Value
	if err := v.As(&m); err != nil {
		t.Fatalf("reading object: %v", err)
	}
	return m
}

func TestUpgradeFromV0_V5_16State(t *testing.T) {
	t.Parallel()
	_, attrs := upgradeV0(t, v5_16ZoneStateAtV0)

	if got := stringAttr(t, attrs["name"]); got != "example.com" {
		t.Errorf("name = %q", got)
	}
	if got := stringAttr(t, objectAttr(t, attrs["account"])["id"]); got != "f037e56e89293a057740de681ac9abbe" {
		t.Errorf("account.id = %q", got)
	}
	if got := stringAttr(t, attrs["activated_on"]); got != "2025-06-12T18:04:31.123456Z" {
		t.Errorf("activated_on = %q", got)
	}
	if got := stringAttr(t, objectAttr(t, attrs["owner"])["type"]); got != "organization" {
		t.Errorf("owner.type = %q", got)
	}
	var pageRuleQuota big.Float
	if err := objectAttr(t, attrs["meta"])["page_rule_quota"].As(&pageRuleQuota); err != nil || pageRuleQuota.Cmp(big.NewFloat(125)) != 0 {
		t.Errorf("meta.page_rule_quota = %v, %v", pageRuleQuota.String(), err)
	}
	if got := stringAttr(t, objectAttr(t, attrs["plan"])["legacy_id"]); got != "enterprise" {
		t.Errorf("plan.legacy_id = %q", got)
	}
	var vanity []tftypes.Value
	if err := attrs["vanity_name_servers"].As(&vanity); err != nil || len(vanity) != 2 {
		t.Errorf("vanity_name_servers = %v, %v", vanity, err)
	}
	for _, key := range []string{"cname_suffix", "verification_key", "original_dnshost"} {
		if !attrs[key].IsNull() {
			t.Errorf("%s should stay null", key)
		}
	}
}

func TestUpgradeFromV0_V5_0State(t *testing.T) {
	t.Parallel()
	_, attrs := upgradeV0(t, v5_0ZoneStateAtV0)

	if got := stringAttr(t, attrs["name"]); got != "example.com" {
		t.Errorf("name = %q", got)
	}
	if got := stringAttr(t, objectAttr(t, attrs["account"])["id"]); got != "f037e56e89293a057740de681ac9abbe" {
		t.Errorf("account.id = %q", got)
	}
	if got := stringAttr(t, objectAttr(t, attrs["owner"])["type"]); got != "user" {
		t.Errorf("owner.type = %q", got)
	}
	// Attributes added after 5.0 come back null, for the next refresh to fill.
	for _, key := range []string{"plan", "tenant", "tenant_unit", "permissions", "cname_suffix", "verification_key"} {
		if !attrs[key].IsNull() {
			t.Errorf("%s should be null for 5.0-era state", key)
		}
	}
}

func TestUpgradeFromV0_MissingRawState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	s := zone.ResourceSchema(ctx)
	upgrader := (&zone.ZoneResource{}).UpgradeState(ctx)[0].StateUpgrader
	for name, raw := range map[string]*tfprotov6.RawState{
		"nil":   nil,
		"empty": {},
	} {
		resp := &resource.UpgradeStateResponse{State: tfsdk.State{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}}
		upgrader(ctx, resource.UpgradeStateRequest{RawState: raw}, resp)
		if !resp.Diagnostics.HasError() {
			t.Errorf("%s: expected an error diagnostic", name)
		}
	}
}
