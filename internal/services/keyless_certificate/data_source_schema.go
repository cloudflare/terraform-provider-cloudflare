// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*KeylessCertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"keyless_certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
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
			"port": schema.Float64Attribute{
				Description: "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the Keyless SSL.\navailable values: \"active\", \"deleted\"",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "deleted"),
				},
			},
			"permissions": schema.ListAttribute{
				Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"tunnel": schema.SingleNestedAttribute{
				Description: "Configuration for using Keyless SSL through a Cloudflare Tunnel",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[KeylessCertificateTunnelDataSourceModel](ctx),
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
	}
}

func (d *KeylessCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *KeylessCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
