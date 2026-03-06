// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_ip_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceIPProfilesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"per_page": schema.Int64Attribute{
				Description: "The number of IP profiles to return per page.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceIPProfilesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Device IP profile.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "The RFC3339Nano timestamp when the Device IP profile was created.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An optional description of the Device IP profile.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Device IP profile is enabled.",
							Computed:    true,
						},
						"match": schema.StringAttribute{
							Description: `The wirefilter expression to match registrations. Available values: "identity.name", "identity.email", "identity.groups.id", "identity.groups.name", "identity.groups.email", "identity.saml_attributes".`,
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "A user-friendly name for the Device IP profile.",
							Computed:    true,
						},
						"precedence": schema.Int64Attribute{
							Description: "The precedence of the Device IP profile. Lower values indicate higher precedence. Device IP profile will be evaluated in ascending order of this field.",
							Computed:    true,
						},
						"subnet_id": schema.StringAttribute{
							Description: "The ID of the Subnet.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Description: "The RFC3339Nano timestamp when the Device IP profile was last updated.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceIPProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceIPProfilesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
