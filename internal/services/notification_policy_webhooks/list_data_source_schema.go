// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &NotificationPolicyWebhooksListDataSource{}

func (d *NotificationPolicyWebhooksListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account id",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of a webhook",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Timestamp of when the webhook destination was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_failure": schema.StringAttribute{
							Description: "Timestamp of the last time an attempt to dispatch a notification to this webhook failed.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"last_success": schema.StringAttribute{
							Description: "Timestamp of the last time Cloudflare was able to successfully dispatch a notification using this webhook.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "The name of the webhook destination. This will be included in the request body when you receive a webhook notification.",
							Computed:    true,
							Optional:    true,
						},
						"secret": schema.StringAttribute{
							Description: "Optional secret that will be passed in the `cf-webhook-auth` header when dispatching generic webhook notifications or formatted for supported destinations. Secrets are not returned in any API response body.",
							Computed:    true,
							Optional:    true,
						},
						"type": schema.StringAttribute{
							Description: "Type of webhook endpoint.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("slack", "generic", "gchat"),
							},
						},
						"url": schema.StringAttribute{
							Description: "The POST endpoint to call when dispatching a notification.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (d *NotificationPolicyWebhooksListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
