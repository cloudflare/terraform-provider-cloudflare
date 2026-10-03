package zero_trust_casb_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// attributeGetter is satisfied by both tfsdk.Plan and tfsdk.Config, allowing the
// scope check to read whichever source is authoritative for the operation.
type attributeGetter interface {
	GetAttribute(ctx context.Context, path path.Path, target interface{}) diag.Diagnostics
}

// validateIntegrationScope rejects policies that are scoped to specific
// integrations but do not resolve to at least one integration ID.
//
// integration_ids is optional and computed: an existing policy may omit it from
// configuration and still carry IDs in state, which a plan modifier preserves.
// Validating configuration alone would therefore reject those policies, so this
// runs at plan time where both signals are available:
//
//   - on create there is no prior state, so configuration is authoritative;
//   - on update the resolved plan already carries any preserved IDs.
func validateIntegrationScope(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	// Destroy plans have no configuration to validate.
	if req.Plan.Raw.IsNull() {
		return
	}

	// On create there is no prior state, so configuration is authoritative.
	// On update the resolved plan already carries any preserved IDs.
	var source attributeGetter = req.Plan
	if req.State.Raw.IsNull() {
		source = req.Config
	}

	var appliesToAll types.Bool
	resp.Diagnostics.Append(source.GetAttribute(ctx, path.Root("applies_to_all_integrations"), &appliesToAll)...)

	var integrationIDs customfield.List[types.String]
	resp.Diagnostics.Append(source.GetAttribute(ctx, path.Root("integration_ids"), &integrationIDs)...)

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
