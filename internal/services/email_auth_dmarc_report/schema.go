// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_auth_dmarc_report

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*EmailAuthDMARCReportResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Email Security DMARC Reports Read",
				"Email Security DMARC Reports Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Description: "Enable or disable DMARC reports for this zone. When enabling, the API ensures the DMARC RUA record exists in DNS.",
				Optional:    true,
				Computed:    true,
			},
			"skip_wizard": schema.BoolAttribute{
				Description: "Skip the DMARC setup wizard.",
				Optional:    true,
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Creation timestamp",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_at": schema.StringAttribute{
				Description: "Last modification timestamp",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"rua_prefix": schema.StringAttribute{
				Description: "Prefix for DMARC RUA addresses (32-char hex string).",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "DMARC configuration status. Omitted by the API when DMARC is correctly configured.\nAvailable values: \"missing-dmarc-report\", \"multiple-dmarc-reports\", \"missing-dmarc-rua\", \"cname-on-dmarc-record\", \"unauthorized-reporting-domain\".",
				Computed:    true,
			},
		},
	}
}

func (r *EmailAuthDMARCReportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailAuthDMARCReportResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
