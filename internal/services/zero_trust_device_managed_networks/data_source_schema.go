// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
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
				Description: "The type of device managed network.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("tls"),
				},
			},
			"config": schema.StringAttribute{
				Description: "The configuration object containing information for the WARP client to detect the managed network.",
				Computed:    true,
				CustomType:  jsontypes.NormalizedType{},
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
