// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*OrganizationProfileDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				Required: true,
			},
			"business_address": schema.StringAttribute{
				Computed: true,
			},
			"business_email": schema.StringAttribute{
				Computed: true,
			},
			"business_name": schema.StringAttribute{
				Computed: true,
			},
			"business_phone": schema.StringAttribute{
				Computed: true,
			},
			"external_metadata": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *OrganizationProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OrganizationProfileDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
