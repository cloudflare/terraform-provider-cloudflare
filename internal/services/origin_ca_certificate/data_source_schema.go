// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_ca_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*OriginCACertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The x509 serial number of the Origin CA certificate.",
				Computed:    true,
			},
			"certificate_id": schema.StringAttribute{
				Description: "The x509 serial number of the Origin CA certificate.",
				Optional:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "The Origin CA certificate. Will be newline-encoded.",
				Computed:    true,
			},
			"csr": schema.StringAttribute{
				Description: "The Certificate Signing Request (CSR). Must be newline-encoded.",
				Computed:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate will expire.",
				Computed:    true,
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
				Description: "The number of days for which the certificate was requested to be valid. Only returned in the Create Certificate response, where it echoes the requested value. Not returned when getting or listing certificates.\nAvailable values: 7, 30, 90, 365, 730, 1095, 5475.",
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
			"revoked_at": schema.StringAttribute{
				Description: "When the certificate was revoked.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"hostnames": schema.ListAttribute{
				Description: "Array of hostnames or wildcard names bound to the certificate.\nHostnames must be fully qualified domain names (FQDNs) belonging to zones on your account (e.g., `example.com` or `sub.example.com`). Wildcards are supported only as a `*.` prefix for a single level (e.g., `*.example.com`). Double wildcards (`*.*.example.com`) and interior wildcards (`foo.*.example.com`) are not allowed. The wildcard suffix must be a multi-label domain (`*.example.com` is valid, but `*.com` is not). Unicode/IDN hostnames are accepted and automatically converted to punycode.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *OriginCACertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OriginCACertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("certificate_id"), path.MatchRoot("filter")),
	}
}
