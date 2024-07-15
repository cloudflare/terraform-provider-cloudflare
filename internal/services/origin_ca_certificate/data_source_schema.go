// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &OriginCACertificateDataSource{}
var _ datasource.DataSourceWithValidateConfig = &OriginCACertificateDataSource{}

func (r OriginCACertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"csr": schema.StringAttribute{
				Description: "The Certificate Signing Request (CSR). Must be newline-encoded.",
				Optional:    true,
			},
			"hostnames": schema.ListAttribute{
				Description: "Array of hostnames or wildcard names (e.g., *.example.com) bound to the certificate.",
				Optional:    true,
				ElementType: jsontypes.NewNormalizedNull().Type(ctx),
			},
			"request_type": schema.StringAttribute{
				Description: "Signature type desired on certificate (\"origin-rsa\" (rsa), \"origin-ecc\" (ecdsa), or \"keyless-certificate\" (for Keyless SSL servers).",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("origin-rsa", "origin-ecc", "keyless-certificate"),
				},
			},
			"requested_validity": schema.Float64Attribute{
				Description: "The number of days for which the certificate should be valid.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.OneOf(7, 30, 90, 365, 730, 1095, 5475),
				},
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "The Origin CA certificate. Will be newline-encoded.",
				Optional:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate will expire.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *OriginCACertificateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *OriginCACertificateDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
