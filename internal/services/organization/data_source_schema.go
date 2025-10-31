// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organization

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*OrganizationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				Required: true,
			},
			"create_time": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"meta": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[OrganizationMetaDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"flags": schema.SingleNestedAttribute{
						Description: "Organization flags for feature enablement",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[OrganizationMetaFlagsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"account_creation": schema.StringAttribute{
								Computed: true,
							},
							"account_deletion": schema.StringAttribute{
								Computed: true,
							},
							"account_migration": schema.StringAttribute{
								Computed: true,
							},
							"account_mobility": schema.StringAttribute{
								Computed: true,
							},
							"sub_org_creation": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"managed_by": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"parent": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[OrganizationParentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"profile": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[OrganizationProfileDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
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
			},
		},
	}
}

func (d *OrganizationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OrganizationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
