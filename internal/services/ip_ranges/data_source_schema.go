// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ip_ranges

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*IPRangesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"networks": schema.StringAttribute{
				Description: "Specified as `jdcloud` to list IPs used by JD Cloud data centers.",
				Optional:    true,
			},
			"etag": schema.StringAttribute{
				Description: "A digest of the IP data. Useful for determining if the data has changed.",
				Computed:    true,
			},
			"ipv4_cidrs": schema.ListAttribute{
				Description: "List of Cloudflare IPv4 CIDR addresses.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"ipv6_cidrs": schema.ListAttribute{
				Description: "List of Cloudflare IPv6 CIDR addresses.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"jdcloud_cidrs": schema.ListAttribute{
				Description: "List IPv4 and IPv6 CIDRs, only populated if `?networks=jdcloud` is used.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *IPRangesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *IPRangesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
