// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_ip_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceIPProfileDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"profile_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"per_page": schema.Int64Attribute{
						Description: "The number of IP profiles to return per page.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.Between(1, 100),
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceIPProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceIPProfileDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("profile_id"), path.MatchRoot("filter")),
	}
}
