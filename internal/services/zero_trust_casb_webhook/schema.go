// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_webhook

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustCasbWebhookResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zero Trust Read",
				"Zero Trust Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the specific webhook configuration.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"authentication_type": schema.StringAttribute{
				Description: "Type of authentication used for the webhook.\nAvailable values: \"Basic Auth\", \"None\", \"Bearer Auth\", \"Static Headers\", \"HMAC-Signing\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"Basic Auth",
						"None",
						"Bearer Auth",
						"Static Headers",
						"HMAC-Signing",
					),
				},
			},
			"destination_url": schema.StringAttribute{
				Description: "Target URL for the webhook configuration. Where resulting data will be sent.",
				Required:    true,
			},
			"label": schema.StringAttribute{
				Description: "Account-specified display label for the webhook configuration.",
				Required:    true,
			},
			"signing_secret": schema.StringAttribute{
				Description: `Secret key used for HMAC signing when authentication_type is "HMAC-Signing".`,
				Optional:    true,
				Sensitive:   true,
			},
			"headers": schema.ListNestedAttribute{
				Description: "List of custom headers to include in webhook requests.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "Header key name.",
							Required:    true,
						},
						"value": schema.StringAttribute{
							Description: "Header value. Required on Create and Evaluate. On Update, omit or set to null to keep existing value.",
							Optional:    true,
							Sensitive:   true,
						},
					},
				},
			},
			"status": schema.StringAttribute{
				Description: "Status of the webhook configuration. Defaults to enabled when omitted.\nAvailable values: \"enabled\", \"disabled\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
				},
				Default: stringdefault.StaticString("enabled"),
			},
			"created_at": schema.StringAttribute{
				Description:   "Timestamp when the webhook configuration was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp when the webhook configuration was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"version": schema.Int64Attribute{
				Description: "Version number of the configuration.",
				Computed:    true,
			},
		},
	}
}

func (r *ZeroTrustCasbWebhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustCasbWebhookResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
