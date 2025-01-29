// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_managed_networks

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceManagedNetworksListDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceManagedNetworksListResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"config": schema.SingleNestedAttribute{
							Description: "The configuration object containing information for the WARP client to detect the managed network.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustDeviceManagedNetworksListConfigDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"tls_sockaddr": schema.StringAttribute{
									Description: "A network address of the form \"host:port\" that the WARP client will use to detect the presence of a TLS host.",
									Computed:    true,
								},
								"sha256": schema.StringAttribute{
									Description: "The SHA-256 hash of the TLS certificate presented by the host found at tls_sockaddr. If absent, regular certificate verification (trusted roots, valid timestamp, etc) will be used to validate the certificate.",
									Computed:    true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the device managed network. This name must be unique.",
							Computed:    true,
						},
						"network_id": schema.StringAttribute{
							Description: "API UUID.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of device managed network.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("tls"),
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceManagedNetworksListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceManagedNetworksListDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
