// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_risk_scoring_integration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustRiskScoringIntegrationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The id of the integration, a UUIDv4.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"integration_type": schema.StringAttribute{
				Description: "available values: \"Okta\"",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Okta"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"tenant_url": schema.StringAttribute{
				Description: "The base url of the tenant, e.g. \"https://tenant.okta.com\"",
				Required:    true,
			},
			"active": schema.BoolAttribute{
				Description: "Whether this integration is enabled. If disabled, no risk changes will be exported to the third-party.",
				Optional:    true,
			},
			"reference_id": schema.StringAttribute{
				Description: "A reference id that can be supplied by the client. Currently this should be set to the Access-Okta IDP ID (a UUIDv4).\nhttps://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider",
				Optional:    true,
			},
			"account_tag": schema.StringAttribute{
				Description: "The Cloudflare account tag.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the integration was created in RFC3339 format.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"well_known_url": schema.StringAttribute{
				Description: "The URL for the Shared Signals Framework configuration, e.g. \"/.well-known/sse-configuration/{integration_uuid}/\". https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1",
				Computed:    true,
			},
		},
	}
}

func (r *ZeroTrustRiskScoringIntegrationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustRiskScoringIntegrationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
