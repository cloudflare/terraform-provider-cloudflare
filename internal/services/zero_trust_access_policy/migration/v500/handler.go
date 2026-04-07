package v500

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// detectAccessPolicyV4State returns true if the raw state looks like v4 format.
// Detection: v4 stores connection_rules as a JSON array []; v4 also has application_id
// or precedence fields that v5 removes.
func detectAccessPolicyV4State(req resource.UpgradeStateRequest) bool {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false
	}
	var raw map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false
	}
	// v4 had application_id and precedence fields that v5 removes
	_, hasPrecedence := raw["precedence"]
	_, hasAppID := raw["application_id"]
	if hasPrecedence || hasAppID {
		return true
	}
	// v4 stores connection_rules as array [], v5 as null or object {}
	if cr, ok := raw["connection_rules"]; ok && cr != nil {
		_, isArray := cr.([]interface{})
		return isArray
	}
	return false
}

// upgradeAccessPolicyFromV4 transforms v4 state to v5 using the source schema and Transform function.
func upgradeAccessPolicyFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	v4Schema := SourceAccessPolicySchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError("Failed to unmarshal v4 access policy state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s\n\nThis resource may need to be imported manually.", err))
		return
	}

	state := tfsdk.State{Raw: rawValue, Schema: v4Schema}
	var v4State SourceAccessPolicyModel
	resp.Diagnostics.Append(state.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "v4 access policy state transformation completed via UpgradeFromSchemaV0")
}

// UpgradeFromSchemaV0 handles state upgrades from schema_version=0 to current version.
// Both v4 cloudflare_access_policy (when resource name was already zero_trust_access_policy)
// and early v5 cloudflare_zero_trust_access_policy have schema_version=0.
//
// This handler detects the format and either transforms v4 state or passes v5 state through.
func UpgradeFromSchemaV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_policy state from schema_version=0")

	if detectAccessPolicyV4State(req) {
		tflog.Info(ctx, "Detected v4 format in schema_version=0 state, transforming")
		upgradeAccessPolicyFromV4(ctx, req, resp)
		return
	}

	// Early v5 state (v5.12-v5.15) — pass through unchanged
	tflog.Info(ctx, "Detected v5 format in schema_version=0 state, no-op upgrade")

	// When PriorSchema is nil (as set in migrations.go for v4 compatibility),
	// req.State is nil and only req.RawState is available. We need to unmarshal
	// the raw state using the target schema type.
	v5Type := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, err := req.RawState.Unmarshal(v5Type)
	if err != nil {
		resp.Diagnostics.AddError("Failed to unmarshal v5 state during upgrade",
			fmt.Sprintf("Could not parse state as v5 format: %s", err))
		return
	}
	resp.State.Raw = rawValue
}

// UpgradeFromSchemaV1 handles state upgrades from schema_version=1 to current version.
// This is a no-op upgrade since the schema is compatible - just copy state through.
func UpgradeFromSchemaV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_policy state from schema_version=1 (no-op)")

	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State upgrade from schema_version=1 completed")
}

// MoveFromAccessPolicy handles moves from cloudflare_access_policy (v4) to cloudflare_zero_trust_access_policy (v5).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_policy.example
//	    to   = cloudflare_zero_trust_access_policy.example
//	}
//
// This is the ONLY supported path for v4 → v5 migration. `terraform state mv` is NOT supported
// because both v4 and early v5 have schema_version=0, and the upgrader expects v5 format.
func MoveFromAccessPolicy(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if req.SourceTypeName != "cloudflare_access_policy" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Starting state move from cloudflare_access_policy to cloudflare_zero_trust_access_policy",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	var v4State SourceAccessPolicyModel
	resp.Diagnostics.Append(req.SourceState.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.TargetState.Set(ctx, v5State)...)

	tflog.Info(ctx, "State move from cloudflare_access_policy to cloudflare_zero_trust_access_policy completed successfully")
}

// isCloudflareProvider checks if the provider address is the Cloudflare provider.
func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
