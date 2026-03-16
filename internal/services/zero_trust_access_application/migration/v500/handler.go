package v500

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=0 to current version.
// Covers v5.12–v5.16 which had no schema Version set (defaulted to 0).
// State is already in v5 format — no-op pass-through.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_application state from schema_version=0 (no-op)")
	resp.State.Raw = req.State.Raw
	tflog.Info(ctx, "State upgrade from schema_version=0 completed")
}

// UpgradeFromV1 handles state upgrades from schema_version=1 to current version.
// Covers v5.17–v5.18 which set Version: 1.
// State is already in v5 format — no-op pass-through.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_application state from schema_version=1 (no-op)")
	resp.State.Raw = req.State.Raw
	tflog.Info(ctx, "State upgrade from schema_version=1 completed")
}

// MoveFromAccessApplication handles moves from cloudflare_access_application (v4) to
// cloudflare_zero_trust_access_application (v5). Triggered by the `moved` block that
// tf-migrate generates when renaming cloudflare_access_application resources.
//
// SourceSchema is not set on the StateMover (see migrations.go), so req.SourceState is
// always nil. We unmarshal req.SourceRawState directly with the v4 schema to ensure
// parsing errors are explicit diagnostics rather than silent skips.
func MoveFromAccessApplication(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
	if req.SourceTypeName != "cloudflare_access_application" {
		return
	}

	if !isCloudflareProvider(req.SourceProviderAddress) {
		return
	}

	tflog.Info(ctx, "Moving state from cloudflare_access_application to cloudflare_zero_trust_access_application",
		map[string]interface{}{
			"source_type":           req.SourceTypeName,
			"source_schema_version": req.SourceSchemaVersion,
			"source_provider":       req.SourceProviderAddress,
		})

	if req.SourceRawState == nil {
		resp.Diagnostics.AddError(
			"Missing source state",
			"MoveResourceState received nil SourceRawState for cloudflare_access_application",
		)
		return
	}

	// Unmarshal the raw v4 state using the v4 schema type.
	v4Schema := SourceAccessApplicationSchema()
	v4SchemaType := v4Schema.Type().TerraformType(ctx)
	rawVal, err := req.SourceRawState.UnmarshalWithOpts(v4SchemaType, tfprotov6.UnmarshalOpts{
		ValueFromJSONOpts: tftypes.ValueFromJSONOpts{
			IgnoreUndefinedAttributes: true,
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to parse source state for MoveResourceState",
			fmt.Sprintf("Error reading cloudflare_access_application state: %s", err),
		)
		return
	}

	srcState := tfsdk.State{Raw: rawVal, Schema: v4Schema}

	var v4State SourceAccessApplicationModel
	resp.Diagnostics.Append(srcState.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.TargetState.Set(ctx, v5State)...)

	tflog.Info(ctx, "State move from cloudflare_access_application completed successfully")
}

// isCloudflareProvider checks if the provider address is the Cloudflare provider.
func isCloudflareProvider(addr string) bool {
	return strings.Contains(addr, "cloudflare/cloudflare") ||
		strings.Contains(addr, "registry.terraform.io/cloudflare")
}
