// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &AccessMutualTLSCertificatesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessMutualTLSCertificatesDataSource{}

func (r AccessMutualTLSCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the application that will use this certificate.",
							Computed:    true,
							Optional:    true,
						},
						"associated_hostnames": schema.ListAttribute{
							Description: "The hostnames of the applications that will use this certificate.",
							Computed:    true,
							Optional:    true,
							ElementType: types.StringType,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"expires_on": schema.StringAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"fingerprint": schema.StringAttribute{
							Description: "The MD5 fingerprint of the certificate.",
							Computed:    true,
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the certificate.",
							Computed:    true,
							Optional:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (r *AccessMutualTLSCertificatesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessMutualTLSCertificatesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
