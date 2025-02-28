// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceManagedNetworksDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "API UUID.",
				Computed:    true,
			},
			"network_id": schema.StringAttribute{
				Description: "API UUID.",
				Computed:    true,
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the device managed network. This name must be unique.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of device managed network.\nAvailable values: \"tls\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tls"),
				},
			},
			"config": schema.SingleNestedAttribute{
				Description: "The configuration object containing information for the WARP client to detect the managed network.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDeviceManagedNetworksConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"tls_sockaddr": schema.StringAttribute{
						Description: `A network address of the form "host:port" that the WARP client will use to detect the presence of a TLS host.`,
						Computed:    true,
					},
					"sha256": schema.StringAttribute{
						Description: "The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceManagedNetworksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceManagedNetworksDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
