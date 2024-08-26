// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*MTLSCertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"mtls_certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate expires.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Computed:    true,
			},
			"serial_number": schema.StringAttribute{
				Description: "The certificate serial number.",
				Computed:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the certificate.",
				Computed:    true,
			},
			"ca": schema.BoolAttribute{
				Description: "Indicates whether the certificate is a CA or leaf certificate.",
				Computed:    true,
				Optional:    true,
			},
			"certificates": schema.StringAttribute{
				Description: "The uploaded root CA certificate.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Optional unique name for the certificate. Only used for human readability.",
				Computed:    true,
				Optional:    true,
			},
			"uploaded_on": schema.StringAttribute{
				Description: "This is the time the certificate was uploaded.",
				Computed:    true,
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *MTLSCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MTLSCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("mtls_certificate_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("mtls_certificate_id")),
	}
}
