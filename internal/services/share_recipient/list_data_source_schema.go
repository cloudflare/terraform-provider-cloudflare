// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package share_recipient

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ShareRecipientsDataSource)(nil)

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
			"include_resources": schema.BoolAttribute{
				Description: "Include resources in the response.",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[ShareRecipientsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Share Recipient identifier tag.",
							Computed:    true,
						},
						"account_id": schema.StringAttribute{
							Description: "Account identifier.",
							Computed:    true,
						},
						"association_status": schema.StringAttribute{
							Description: "Share Recipient association status.\nAvailable values: \"associating\", \"associated\", \"disassociating\", \"disassociated\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"associating",
									"associated",
									"disassociating",
									"disassociated",
								),
							},
						},
						"created": schema.StringAttribute{
							Description: "When the share was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"modified": schema.StringAttribute{
							Description: "When the share was modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"resources": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[ShareRecipientsResourcesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"error": schema.StringAttribute{
										Description: "Share Recipient error message.",
										Computed:    true,
									},
									"resource_id": schema.StringAttribute{
										Description: "Share Resource identifier.",
										Computed:    true,
									},
									"resource_version": schema.Int64Attribute{
										Description: "Resource Version.",
										Computed:    true,
									},
									"terminal": schema.BoolAttribute{
										Description: "Whether the error is terminal or will be continually retried.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ShareRecipientsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ShareRecipientsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
