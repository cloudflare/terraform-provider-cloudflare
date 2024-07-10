// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &NotificationPolicyWebhooksDataSource{}
var _ datasource.DataSourceWithValidateConfig = &NotificationPolicyWebhooksDataSource{}

func (r NotificationPolicyWebhooksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account id",
				Optional:    true,
			},
			"webhook_id": schema.StringAttribute{
				Description: "The unique identifier of a webhook",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of a webhook",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of when the webhook destination was created.",
				Computed:    true,
			},
			"last_failure": schema.StringAttribute{
				Description: "Timestamp of the last time an attempt to dispatch a notification to this webhook failed.",
				Computed:    true,
			},
			"last_success": schema.StringAttribute{
				Description: "Timestamp of the last time Cloudflare was able to successfully dispatch a notification using this webhook.",
				Computed:    true,
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
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "The account id",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *NotificationPolicyWebhooksDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *NotificationPolicyWebhooksDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
