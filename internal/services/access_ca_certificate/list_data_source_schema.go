// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_ca_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AccessCACertificatesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessCACertificatesDataSource{}

func (r AccessCACertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the CA.",
							Computed:    true,
						},
						"aud": schema.StringAttribute{
							Description: "The Application Audience (AUD) tag. Identifies the application associated with the CA.",
							Computed:    true,
						},
						"public_key": schema.StringAttribute{
							Description: "The public key to add to your SSH server configuration.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *AccessCACertificatesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessCACertificatesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
