package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func UpgradeStateV0toV500(sourceSchema schema.Schema) resource.StateUpgrader {
	return resource.StateUpgrader{
		PriorSchema:   &sourceSchema,
		StateUpgrader: upgradeStateV0toV500,
	}
}

func UpgradeStateV1toV500(targetSchema schema.Schema) resource.StateUpgrader {
	return resource.StateUpgrader{
		PriorSchema:   &targetSchema,
		StateUpgrader: upgradeStateV1toV500,
	}
}

func upgradeStateV0toV500(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading workers_custom_domain state from schema_version=0")

	var source SourceV4WorkersCustomDomainModel
	resp.Diagnostics.Append(req.State.Get(ctx, &source)...)
	if resp.Diagnostics.HasError() {
		return
	}

	target, diags := TransformV4toV500(ctx, source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, target)...)
}

func upgradeStateV1toV500(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading workers_custom_domain state from schema_version=1 (no-op)")
	resp.State.Raw = req.State.Raw
}
