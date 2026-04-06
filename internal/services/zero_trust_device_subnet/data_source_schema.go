// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_subnet

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceSubnetDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The UUID of the subnet.",
				Computed:    true,
			},
			"subnet_id": schema.StringAttribute{
				Description: "The UUID of the subnet.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Required:    true,
			},
			"comment": schema.StringAttribute{
				Description: "An optional description of the subnet.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"deleted_at": schema.StringAttribute{
				Description: "Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"is_default_network": schema.BoolAttribute{
				Description: "If `true`, this is the default subnet for the account. There can only be one default subnet per account.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for the subnet.",
				Computed:    true,
			},
			"network": schema.StringAttribute{
				Description: "The private IPv4 or IPv6 range defining the subnet, in CIDR notation.",
				Computed:    true,
			},
			"subnet_type": schema.StringAttribute{
				Description: "The type of subnet.\nAvailable values: \"cloudflare_source\", \"warp\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("cloudflare_source", "warp"),
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceSubnetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceSubnetDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
