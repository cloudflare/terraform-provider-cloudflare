// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*OriginCACertificatesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"limit": schema.Int64Attribute{
				Description: "Limit to the number of records returned.",
				Optional:    true,
			},
			"offset": schema.Int64Attribute{
				Description: "Offset the results",
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
				CustomType:  customfield.NewNestedObjectListType[OriginCACertificatesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"csr": schema.StringAttribute{
							Description: "The Certificate Signing Request (CSR). Must be newline-encoded.",
							Computed:    true,
						},
						"hostnames": schema.ListAttribute{
							Description: "Array of hostnames or wildcard names (e.g., *.example.com) bound to the certificate.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"request_type": schema.StringAttribute{
							Description: "Signature type desired on certificate (\"origin-rsa\" (rsa), \"origin-ecc\" (ecdsa), or \"keyless-certificate\" (for Keyless SSL servers).\nAvailable values: \"origin-rsa\", \"origin-ecc\", \"keyless-certificate\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"origin-rsa",
									"origin-ecc",
									"keyless-certificate",
								),
							},
						},
						"requested_validity": schema.Float64Attribute{
							Description: "The number of days for which the certificate should be valid.\nAvailable values: 7, 30, 90, 365, 730, 1095, 5475.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.OneOf(
									7,
									30,
									90,
									365,
									730,
									1095,
									5475,
								),
							},
						},
						"id": schema.StringAttribute{
							Description: "Identifier.",
							Computed:    true,
						},
						"certificate": schema.StringAttribute{
							Description: "The Origin CA certificate. Will be newline-encoded.",
							Computed:    true,
						},
						"expires_on": schema.StringAttribute{
							Description: "When the certificate will expire.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *OriginCACertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *OriginCACertificatesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
