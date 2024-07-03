// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &KeylessCertificateDataSource{}
var _ datasource.DataSourceWithValidateConfig = &KeylessCertificateDataSource{}

func (r KeylessCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"keyless_certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Keyless certificate identifier tag.",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the Keyless SSL was created.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether or not the Keyless SSL is on or off.",
				Optional:    true,
			},
			"host": schema.StringAttribute{
				Description: "The keyless SSL name.",
				Optional:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the Keyless SSL was last modified.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The keyless SSL name.",
				Optional:    true,
			},
			"permissions": schema.StringAttribute{
				Description: "Available permissions for the Keyless SSL for the current user requesting the item.",
				Optional:    true,
			},
			"port": schema.Float64Attribute{
				Description: "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
				Computed:    true,
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the Keyless SSL.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("active", "deleted"),
				},
			},
			"tunnel": schema.SingleNestedAttribute{
				Description: "Configuration for using Keyless SSL through a Cloudflare Tunnel",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"private_ip": schema.StringAttribute{
						Description: "Private IP of the Key Server Host",
						Required:    true,
					},
					"vnet_id": schema.StringAttribute{
						Description: "Cloudflare Tunnel Virtual Network ID",
						Required:    true,
					},
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *KeylessCertificateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *KeylessCertificateDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
