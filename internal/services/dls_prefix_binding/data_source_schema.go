// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls_prefix_binding

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*DLSPrefixBindingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Unique identifier for the prefix binding.",
				Computed:    true,
			},
			"binding_id": schema.StringAttribute{
				Description: "Unique identifier for the prefix binding.",
				Required:    true,
			},
			"account_id": schema.Int64Attribute{
				Required: true,
			},
			"cidr": schema.StringAttribute{
				Description: "The CIDR that is bound.",
				Computed:    true,
			},
			"prefix_id": schema.StringAttribute{
				Description: "The ID of the parent prefix.",
				Computed:    true,
			},
			"region_key": schema.StringAttribute{
				Description: "The region key used for the binding.",
				Computed:    true,
			},
		},
	}
}

func (d *DLSPrefixBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DLSPrefixBindingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
