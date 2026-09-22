// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_webhook

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustCasbWebhookDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zero Trust Read",
				"Zero Trust Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"webhook_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"authentication_type": schema.StringAttribute{
				Description: "Type of authentication used for the webhook.\nAvailable values: \"Basic Auth\", \"None\", \"Bearer Auth\", \"Static Headers\", \"HMAC-Signing\".",
				Computed:    true,
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
			"created_at": schema.StringAttribute{
				Description: "Timestamp when the webhook configuration was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"destination_url": schema.StringAttribute{
				Description: "Target URL for the webhook configuration. Where resulting data will be sent.",
				Computed:    true,
			},
			"label": schema.StringAttribute{
				Description: "Account-specified display label for the webhook configuration.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current status of the webhook configuration. If disabled, data cannot be sent through this configuration.\nAvailable values: \"enabled\", \"disabled\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
				},
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
			"headers": schema.ListNestedAttribute{
				Description: "List of header keys configured for this webhook. Values are not included for security reasons.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustCasbWebhookHeadersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "Header key name (lowercase).",
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "Header value. This field is never returned in API responses for security reasons.",
							Computed:    true,
							Sensitive:   true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustCasbWebhookDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustCasbWebhookDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
