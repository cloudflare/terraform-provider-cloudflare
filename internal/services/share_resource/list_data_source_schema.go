// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_resource

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ShareResourcesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier.",
				Required:    true,
			},
			"share_id": schema.StringAttribute{
				Description: "Share identifier tag.",
				Required:    true,
			},
			"resource_type": schema.StringAttribute{
				Description: "Filter share resources by resource_type.\nAvailable values: \"custom-ruleset\", \"gateway-policy\", \"gateway-destination-ip\", \"gateway-block-page-settings\", \"gateway-extended-email-matching\", \"idp-federation-grant\".",
				Optional:    true,
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
			"status": schema.StringAttribute{
				Description: "Filter share resources by status.\nAvailable values: \"active\", \"deleting\", \"deleted\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"deleting",
						"deleted",
					),
				},
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ShareResourcesResultDataSourceModel](ctx),
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
		},
	}
}

func (d *ShareResourcesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ShareResourcesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
