// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy_webhooks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r NotificationPolicyWebhooksResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "The account id",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the webhook destination. This will be included in the request body when you receive a webhook notification.",
				Required:    true,
			},
			"url": schema.StringAttribute{
				Description: "The POST endpoint to call when dispatching a notification.",
				Required:    true,
			},
			"secret": schema.StringAttribute{
				Description: "Optional secret that will be passed in the `cf-webhook-auth` header when dispatching generic webhook notifications or formatted for supported destinations. Secrets are not returned in any API response body.",
				Optional:    true,
			},
			"errors": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1000),
							},
						},
						"message": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"messages": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.Int64Attribute{
							Required: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1000),
							},
						},
						"message": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"success": schema.BoolAttribute{
				Description: "Whether the API call was successful",
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
			"type": schema.StringAttribute{
				Description: "Type of webhook endpoint.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("slack", "generic", "gchat"),
				},
			},
		},
	}
}
