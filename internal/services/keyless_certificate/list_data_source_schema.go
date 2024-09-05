// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*KeylessCertificatesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
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
				CustomType:  customfield.NewNestedObjectListType[KeylessCertificatesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Keyless certificate identifier tag.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the Keyless SSL was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether or not the Keyless SSL is on or off.",
							Computed:    true,
						},
						"host": schema.StringAttribute{
							Description: "The keyless SSL name.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the Keyless SSL was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "The keyless SSL name.",
							Computed:    true,
						},
						"permissions": schema.ListAttribute{
							Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"port": schema.Float64Attribute{
							Description: "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Status of the Keyless SSL.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("active", "deleted"),
							},
						},
						"tunnel": schema.SingleNestedAttribute{
							Description: "Configuration for using Keyless SSL through a Cloudflare Tunnel",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[KeylessCertificatesTunnelDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"private_ip": schema.StringAttribute{
									Description: "Private IP of the Key Server Host",
									Computed:    true,
								},
								"vnet_id": schema.StringAttribute{
									Description: "Cloudflare Tunnel Virtual Network ID",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *KeylessCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *KeylessCertificatesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
