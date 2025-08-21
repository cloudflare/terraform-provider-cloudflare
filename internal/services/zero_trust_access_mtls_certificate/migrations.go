// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_certificate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessMTLSCertificateResource)(nil)

// zeroTrustAccessMTLSCertificateResourceSchemaV0 defines the v0 schema (v4 provider format)
var zeroTrustAccessMTLSCertificateResourceSchemaV0 = schema.Schema{
	Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"account_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"zone_id": schema.StringAttribute{
			Optional: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"certificate": schema.StringAttribute{
			Optional: true, // Was optional in v4, required in v5
		},
		"associated_hostnames": schema.SetAttribute{
			ElementType: types.StringType,
			Optional:    true,
		},
		"fingerprint": schema.StringAttribute{
			Computed: true,
		},
		// expires_on is not in v4, it's new in v5
	},
}

func (r *ZeroTrustAccessMTLSCertificateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema:   &zeroTrustAccessMTLSCertificateResourceSchemaV0,
			StateUpgrader: upgradeZeroTrustAccessMTLSCertificateStateV0toV1,
		},
	}
}

// upgradeZeroTrustAccessMTLSCertificateStateV0toV1 migrates from v4 provider state format to v5
func upgradeZeroTrustAccessMTLSCertificateStateV0toV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Debug(ctx, "Starting state migration from v0 to v1 for zero_trust_access_mtls_certificate")

	var priorStateData struct {
		ID                  types.String   `tfsdk:"id"`
		AccountID           types.String   `tfsdk:"account_id"`
		ZoneID              types.String   `tfsdk:"zone_id"`
		Name                types.String   `tfsdk:"name"`
		Certificate         types.String   `tfsdk:"certificate"`
		AssociatedHostnames types.Set      `tfsdk:"associated_hostnames"`
		Fingerprint         types.String   `tfsdk:"fingerprint"`
	}

	resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to extract prior state: %v", resp.Diagnostics.Errors()))
		return
	}

	// Create new state structure
	var newState ZeroTrustAccessMTLSCertificateModel

	// Migrate basic attributes - these are the same between v4 and v5
	newState.ID = priorStateData.ID
	newState.AccountID = priorStateData.AccountID
	newState.ZoneID = priorStateData.ZoneID
	newState.Name = priorStateData.Name
	newState.Fingerprint = priorStateData.Fingerprint

	// Handle certificate field migration (optional in v4 → required in v5)
	if priorStateData.Certificate.IsNull() || priorStateData.Certificate.ValueString() == "" {
		// Certificate is required in v5 but was optional in v4
		// If it's missing, we need to fail the migration with a helpful error
		resp.Diagnostics.AddError(
			"Certificate Required for Migration",
			"The certificate field was optional in the v4 provider but is required in v5. "+
				"Please ensure your certificate configuration includes a valid certificate value before migrating. "+
				"If you're using an existing certificate from your infrastructure, you'll need to re-import it with the certificate content.",
		)
		return
	}
	newState.Certificate = priorStateData.Certificate

	// Handle associated_hostnames - framework should handle Set → SetAttribute conversion automatically
	if !priorStateData.AssociatedHostnames.IsNull() && !priorStateData.AssociatedHostnames.IsUnknown() {
		// Convert types.Set to *[]types.String for the new model
		var hostnames []types.String
		resp.Diagnostics.Append(priorStateData.AssociatedHostnames.ElementsAs(ctx, &hostnames, false)...)
		if resp.Diagnostics.HasError() {
			tflog.Error(ctx, fmt.Sprintf("Failed to convert associated_hostnames: %v", resp.Diagnostics.Errors()))
			return
		}
		
		if len(hostnames) > 0 {
			newState.AssociatedHostnames = &hostnames
		}
	}

	// expires_on is a new computed field in v5 - it will be populated on the next refresh
	// We don't need to set it during migration as it's computed

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, fmt.Sprintf("Failed to set new state: %v", resp.Diagnostics.Errors()))
		return
	}

	tflog.Debug(ctx, "Successfully completed state migration from v0 to v1 for zero_trust_access_mtls_certificate")
}
