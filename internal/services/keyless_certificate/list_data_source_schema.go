// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &KeylessCertificatesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &KeylessCertificatesDataSource{}

func (r KeylessCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Keyless certificate identifier tag.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the Keyless SSL was created.",
							Computed:    true,
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
						},
						"name": schema.StringAttribute{
							Description: "The keyless SSL name.",
							Computed:    true,
						},
						"permissions": schema.ListAttribute{
							Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
							Computed:    true,
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
					},
				},
			},
		},
	}
}

func (r *KeylessCertificatesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *KeylessCertificatesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
