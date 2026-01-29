// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_service_token

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustAccessServiceTokenResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the service token.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the service token.",
				Required:    true,
			},
			"previous_client_secret_expires_at": schema.StringAttribute{
				Description: "The expiration of the previous `client_secret`. This can be modified at any point after a rotation. For example, you may extend it further into the future if you need more time to update services with the new secret; or move it into the past to immediately invalidate the previous token in case of compromise.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.Expressions{
						path.MatchRoot("client_secret_version"),
					}...),
				},
			},
			"client_secret_version": schema.Float64Attribute{
				Description: "A version number identifying the current `client_secret` associated with the service token. Incrementing it triggers a rotation; the previous secret will still be accepted until the time indicated by `previous_client_secret_expires_at`.",
				Computed:    true,
				Optional:    true,
				Default: float64default.StaticFloat64(1),
				// Note: AlsoRequires validator removed because client_secret_version has a default value of 1,
				// so it's always set. The previous_client_secret_expires_at is only needed when rotating
				// (incrementing client_secret_version), not on initial creation.
			},
			"duration": schema.StringAttribute{
				Description: "The duration for how long the service token will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The default is 1 year in hours (8760h).",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("8760h"),
			},
			"client_id": schema.StringAttribute{
				Description:   "The Client ID for the service token. Access will check for this value in the `CF-Access-Client-ID` request header.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"client_secret": schema.StringAttribute{
				Description:   "The Client Secret for the service token. Access will check for this value in the `CF-Access-Client-Secret` request header.",
				Computed:      true,
				Sensitive:     true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"expires_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustAccessServiceTokenResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustAccessServiceTokenResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		&clientSecretVersionValidator{},
	}
}

// clientSecretVersionValidator validates that previous_client_secret_expires_at is set
// when client_secret_version is greater than 1 (i.e., during rotation).
type clientSecretVersionValidator struct{}

func (v *clientSecretVersionValidator) Description(ctx context.Context) string {
	return "validates that previous_client_secret_expires_at is set when client_secret_version > 1"
}

func (v *clientSecretVersionValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v *clientSecretVersionValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var clientSecretVersion types.Float64
	var previousClientSecretExpiresAt timetypes.RFC3339

	diags := req.Config.GetAttribute(ctx, path.Root("client_secret_version"), &clientSecretVersion)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.Config.GetAttribute(ctx, path.Root("previous_client_secret_expires_at"), &previousClientSecretExpiresAt)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Only require previous_client_secret_expires_at when client_secret_version > 1 (rotation)
	// If clientSecretVersion is null or unknown, skip validation (will get the default of 1)
	if !clientSecretVersion.IsNull() && !clientSecretVersion.IsUnknown() && clientSecretVersion.ValueFloat64() > 1 && previousClientSecretExpiresAt.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("previous_client_secret_expires_at"),
			"Missing Required Attribute",
			fmt.Sprintf("previous_client_secret_expires_at must be specified when client_secret_version is greater than 1 (current value: %v)", clientSecretVersion.ValueFloat64()),
		)
	}
}
