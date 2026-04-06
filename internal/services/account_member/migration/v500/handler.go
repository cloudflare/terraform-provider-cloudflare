package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v500.
//
// This performs a full transformation from v4 → v5 format:
//   - email_address → email (rename)
//   - role_ids → roles (rename + type conversion)
//   - policies: initialized as null (not in v4)
//   - user: initialized as null (not in v4)
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_member state from v4 SDKv2 provider (schema_version=0)")

	var v4State SourceV4AccountMemberModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := TransformV4toV500(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v500 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 stepping stone (schema_version=1) to v500.
//
// This handles the schema changes between v5.13 and current v5:
//   - policies: ListNestedAttribute → SetNestedAttribute, remove 'id' field
//   - permission_groups: ListNestedAttribute → SetNestedAttribute
//   - resource_groups: ListNestedAttribute → SetNestedAttribute
//   - roles: ListAttribute → SetAttribute
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_member state from version=1 (v5.13) to version=500")

	var v513State SourceV513AccountMemberModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v513State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v500State, diags := TransformV513toV500(ctx, v513State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v500State)...)
	tflog.Info(ctx, "State upgrade from v5.13 to v500 completed successfully")
}
