// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package client_certificate

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ClientCertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"client_certificate_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "The Client Certificate PEM",
				Computed:    true,
			},
			"common_name": schema.StringAttribute{
				Description: "Common Name of the Client Certificate",
				Computed:    true,
			},
			"country": schema.StringAttribute{
				Description: "Country, provided by the CSR",
				Computed:    true,
			},
			"csr": schema.StringAttribute{
				Description: "The Certificate Signing Request (CSR). Must be newline-encoded.",
				Computed:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "Date that the Client Certificate expires",
				Computed:    true,
			},
			"fingerprint_sha256": schema.StringAttribute{
				Description: "Unique identifier of the Client Certificate",
				Computed:    true,
			},
			"issued_on": schema.StringAttribute{
				Description: "Date that the Client Certificate was issued by the Certificate Authority",
				Computed:    true,
			},
			"location": schema.StringAttribute{
				Description: "Location, provided by the CSR",
				Computed:    true,
			},
			"organization": schema.StringAttribute{
				Description: "Organization, provided by the CSR",
				Computed:    true,
			},
			"organizational_unit": schema.StringAttribute{
				Description: "Organizational Unit, provided by the CSR",
				Computed:    true,
			},
			"serial_number": schema.StringAttribute{
				Description: "The serial number on the created Client Certificate.",
				Computed:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the Client Certificate..",
				Computed:    true,
			},
			"ski": schema.StringAttribute{
				Description: "Subject Key Identifier",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "State, provided by the CSR",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Client Certificates may be active or revoked, and the pending_reactivation or pending_revocation represent in-progress asynchronous transitions\nAvailable values: \"active\", \"pending_reactivation\", \"pending_revocation\", \"revoked\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"active",
						"pending_reactivation",
						"pending_revocation",
						"revoked",
					),
				},
			},
			"validity_days": schema.Int64Attribute{
				Description: "The number of days the Client Certificate will be valid after the issued_on date",
				Computed:    true,
			},
			"certificate_authority": schema.SingleNestedAttribute{
				Description: "Certificate Authority used to issue the Client Certificate",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ClientCertificateCertificateAuthorityDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"limit": schema.Int64Attribute{
						Description: "Limit to the number of records returned.",
						Optional:    true,
					},
					"offset": schema.Int64Attribute{
						Description: "Offset the results",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "Client Certitifcate Status to filter results by.\nAvailable values: \"all\", \"active\", \"pending_reactivation\", \"pending_revocation\", \"revoked\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"all",
								"active",
								"pending_reactivation",
								"pending_revocation",
								"revoked",
							),
						},
					},
				},
			},
		},
	}
}

func (d *ClientCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ClientCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("client_certificate_id"), path.MatchRoot("filter")),
	}
}
