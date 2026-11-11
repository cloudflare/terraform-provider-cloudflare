// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ShareResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Resource Sharing Read",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Share identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Account identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"excluded_recipients": schema.ListNestedAttribute{
				Description: "Accounts to exclude from an organization-targeted share. Only allowed when `recipients` contains no account entries. The share owner cannot be excluded. Duplicate entries are ignored.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"account_id": schema.StringAttribute{
							Description: "The account to exclude from the organization-targeted share.",
							Required:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"recipients": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"organization_id": schema.StringAttribute{
							Description: "Organization identifier.",
							Optional:    true,
						},
						"recipient_account_id": schema.StringAttribute{
							Description: "The account that will receive the share.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"resources": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"meta": schema.StringAttribute{
							Description: "Resource Metadata.",
							Required:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"resource_account_id": schema.StringAttribute{
							Description: "The account that owns the resource. Must match the `account_id` in the path.",
							Required:    true,
						},
						"resource_id": schema.StringAttribute{
							Description: "Identifier assigned by the service that owns the shared resource.",
							Required:    true,
						},
						"resource_type": schema.StringAttribute{
							Description: "Resource Type.\nAvailable values: \"custom-ruleset\", \"gateway-policy\", \"gateway-destination-ip\", \"gateway-block-page-settings\", \"gateway-extended-email-matching\", \"idp-federation-grant\", \"trust-grant\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"custom-ruleset",
									"gateway-policy",
									"gateway-destination-ip",
									"gateway-block-page-settings",
									"gateway-extended-email-matching",
									"idp-federation-grant",
									"trust-grant",
								),
							},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"include_recipient_counts": schema.BoolAttribute{
				Description: "Include recipient counts in the response.",
				Optional:    true,
			},
			"include_resources": schema.BoolAttribute{
				Description: "Include resources in the response.",
				Optional:    true,
			},
			"include_tags_from": schema.StringAttribute{
				Description: "Include share tags in the response. `api` reads tags from the tag service; `mirror` reads a local, eventually consistent replica. If the tag lookup fails, the response omits `tags`.\nAvailable values: \"api\", \"mirror\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("api", "mirror"),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the share.",
				Required:    true,
			},
			"tags": schema.MapAttribute{
				Description: "Tags associated with the share.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"account_name": schema.StringAttribute{
				Description: "The display name of an account.",
				Computed:    true,
			},
			"associated_recipient_count": schema.Int64Attribute{
				Description: "The number of recipients in the 'associated' state. This field is only included when requested via the 'include_recipient_counts' parameter.",
				Computed:    true,
			},
			"associating_recipient_count": schema.Int64Attribute{
				Description: "The number of recipients in the 'associating' state. This field is only included when requested via the 'include_recipient_counts' parameter.",
				Computed:    true,
			},
			"created": schema.StringAttribute{
				Description:   "When the object was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"disassociated_recipient_count": schema.Int64Attribute{
				Description: "The number of recipients in the 'disassociated' state. This field is only included when requested via the 'include_recipient_counts' parameter.",
				Computed:    true,
			},
			"disassociating_recipient_count": schema.Int64Attribute{
				Description: "The number of recipients in the 'disassociating' state. This field is only included when requested via the 'include_recipient_counts' parameter.",
				Computed:    true,
			},
			"kind": schema.StringAttribute{
				Description: `Available values: "sent", "received".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("sent", "received"),
				},
			},
			"modified": schema.StringAttribute{
				Description: "When the object was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"organization_id": schema.StringAttribute{
				Description: "Organization identifier.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "active", "deleting", "deleted".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"deleting",
						"deleted",
					),
				},
			},
			"target_type": schema.StringAttribute{
				Description: `Available values: "account", "organization".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("account", "organization"),
				},
			},
		},
	}
}

func (r *ShareResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ShareResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
