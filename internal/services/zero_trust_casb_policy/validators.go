package zero_trust_casb_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type integrationScopeValidator struct{}

func (integrationScopeValidator) Description(context.Context) string {
	return "integration_ids must contain at least one value when applies_to_all_integrations is false"
}

func (v integrationScopeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (integrationScopeValidator) ValidateResource(
	ctx context.Context,
	req resource.ValidateConfigRequest,
	resp *resource.ValidateConfigResponse,
) {
	var appliesToAll types.Bool
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("applies_to_all_integrations"), &appliesToAll)...)

	var integrationIDs customfield.List[types.String]
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("integration_ids"), &integrationIDs)...)
	if resp.Diagnostics.HasError() || !integrationScopeInvalid(appliesToAll, integrationIDs) {
		return
	}

	resp.Diagnostics.AddAttributeError(
		path.Root("integration_ids"),
		"Missing integration IDs",
		"integration_ids must contain at least one value when applies_to_all_integrations is false.",
	)
}

func integrationScopeInvalid(appliesToAll types.Bool, integrationIDs customfield.List[types.String]) bool {
	if appliesToAll.IsNull() || appliesToAll.IsUnknown() || integrationIDs.IsUnknown() {
		return false
	}

	return !appliesToAll.ValueBool() && (integrationIDs.IsNull() || len(integrationIDs.Elements()) == 0)
}
