// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ShareDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Share identifier tag.",
				Computed:    true,
			},
			"share_id": schema.StringAttribute{
				Description: "Share identifier tag.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account identifier.",
				Required:    true,
			},
			"include_recipient_counts": schema.BoolAttribute{
				Description: "Include recipient counts in the response.",
				Optional:    true,
			},
			"include_resources": schema.BoolAttribute{
				Description: "Include resources in the response.",
				Optional:    true,
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
				Description: "When the share was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
				Description: "When the share was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "The name of the share.",
				Computed:    true,
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
			"resources": schema.ListNestedAttribute{
				Description: "A list of resources that are part of the share. This field is only included when requested via the 'include_resources' parameter.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ShareResourcesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Share Resource identifier.",
							Computed:    true,
						},
						"created": schema.StringAttribute{
							Description: "When the share was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"meta": schema.StringAttribute{
							Description: "Resource Metadata.",
							Computed:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"modified": schema.StringAttribute{
							Description: "When the share was modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"resource_account_id": schema.StringAttribute{
							Description: "Account identifier.",
							Computed:    true,
						},
						"resource_id": schema.StringAttribute{
							Description: "Share Resource identifier.",
							Computed:    true,
						},
						"resource_type": schema.StringAttribute{
							Description: "Resource Type.\nAvailable values: \"custom-ruleset\", \"gateway-policy\", \"gateway-destination-ip\", \"gateway-block-page-settings\", \"gateway-extended-email-matching\", \"idp-federation-grant\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"custom-ruleset",
									"gateway-policy",
									"gateway-destination-ip",
									"gateway-block-page-settings",
									"gateway-extended-email-matching",
									"idp-federation-grant",
								),
							},
						},
						"resource_version": schema.Int64Attribute{
							Description: "Resource Version.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Resource Status.\nAvailable values: \"active\", \"deleting\", \"deleted\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"active",
									"deleting",
									"deleted",
								),
							},
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description: "Direction to sort objects.\nAvailable values: \"asc\", \"desc\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("asc", "desc"),
						},
					},
					"kind": schema.StringAttribute{
						Description: "Filter shares by kind.\nAvailable values: \"sent\", \"received\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("sent", "received"),
						},
					},
					"order": schema.StringAttribute{
						Description: "Order shares by values in the given field.\nAvailable values: \"name\", \"created\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("name", "created"),
						},
					},
					"resource_types": schema.ListAttribute{
						Description: "Filter share resources by resource_types.",
						Optional:    true,
						Validators: []validator.List{
							listvalidator.ValueStringsAre(
								stringvalidator.OneOfCaseInsensitive(
									"custom-ruleset",
									"gateway-policy",
									"gateway-destination-ip",
									"gateway-block-page-settings",
									"gateway-extended-email-matching",
									"idp-federation-grant",
								),
							),
						},
						ElementType: types.StringType,
					},
					"status": schema.StringAttribute{
						Description: "Filter shares by status.\nAvailable values: \"active\", \"deleting\", \"deleted\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"active",
								"deleting",
								"deleted",
							),
						},
					},
					"tag": schema.ListAttribute{
						Description: "Filter shares by tag. Each value is either `key=value` (matches shares whose tags contain that key/value pair) or `key` alone (matches shares that have any value for that key). May be repeated; multiple `tag` parameters are ANDed together. Maximum 20 `tag` parameters per request.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"target_type": schema.StringAttribute{
						Description: "Filter shares by target_type.\nAvailable values: \"account\", \"organization\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("account", "organization"),
						},
					},
				},
			},
		},
	}
}

func (d *ShareDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ShareDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("share_id"), path.MatchRoot("filter")),
	}
}
