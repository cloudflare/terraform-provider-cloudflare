package zero_trust_risk_score_integration

import (
	"context"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r *RiskScoreIntegrationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
			The [Risk Score Integration](https://developers.cloudflare.com/cloudflare-one/insights/risk-score/#send-risk-score-to-okta) resource allows you to transmit changes in User Risk Score to a specified vendor such as Okta.
	`),

		Attributes: map[string]schema.Attribute{
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			consts.IDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.IDSchemaDescription,
				Computed:            true,
			},
			"integration_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The type of integration, e.g. 'Okta'. Full list of allowed values can be found here: https://developers.cloudflare.com/api/operations/dlp-zt-risk-score-integration-create#request-body",
			},
			"tenant_url": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The base url of the tenant, e.g. 'https://tenant.okta.com'. Must be your Okta Tenant URL and not your custom domain.",
			},
			"reference_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "A reference id that can be supplied by the client. Currently this should be set to the Access-Okta IDP ID (a UUIDv4). If omitted, a random UUIDv4 is used. https://developers.cloudflare.com/api/operations/access-identity-providers-get-an-access-identity-provider",
			},
			"active": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether this integration is enabled. If disabled, no risk changes will be exported to the third-party.",
			},
			"well_known_url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The URL for the Shared Signals Framework configuration, e.g. '/.well-known/sse-configuration/{integration_uuid}/'. https://openid.net/specs/openid-sse-framework-1_0.html#rfc.section.6.2.1",
			},
		},
	}
}
