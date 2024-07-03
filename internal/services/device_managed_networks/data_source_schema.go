// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package device_managed_networks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &DeviceManagedNetworksDataSource{}
var _ datasource.DataSourceWithValidateConfig = &DeviceManagedNetworksDataSource{}

func (r DeviceManagedNetworksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"network_id": schema.StringAttribute{
				Description: "API UUID.",
				Optional:    true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "The configuration object containing information for the WARP client to detect the managed network.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"tls_sockaddr": schema.StringAttribute{
						Description: "A network address of the form \"host:port\" that the WARP client will use to detect the presence of a TLS host.",
						Required:    true,
					},
					"sha256": schema.StringAttribute{
						Description: "The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.",
						Optional:    true,
					},
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the device managed network. This name must be unique.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of device managed network.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tls"),
				},
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (r *DeviceManagedNetworksDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *DeviceManagedNetworksDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
