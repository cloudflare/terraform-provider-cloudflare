// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*UserDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"country": schema.StringAttribute{
				Description: "The country in which the user lives.",
				Computed:    true,
			},
			"first_name": schema.StringAttribute{
				Description: "User's first name",
				Computed:    true,
			},
			"has_business_zones": schema.BoolAttribute{
				Description: "Indicates whether user has any business zones",
				Computed:    true,
			},
			"has_enterprise_zones": schema.BoolAttribute{
				Description: "Indicates whether user has any enterprise zones",
				Computed:    true,
			},
			"has_pro_zones": schema.BoolAttribute{
				Description: "Indicates whether user has any pro zones",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier of the user.",
				Computed:    true,
			},
			"last_name": schema.StringAttribute{
				Description: "User's last name",
				Computed:    true,
			},
			"suspended": schema.BoolAttribute{
				Description: "Indicates whether user has been suspended",
				Computed:    true,
			},
			"telephone": schema.StringAttribute{
				Description: "User's telephone number",
				Computed:    true,
			},
			"two_factor_authentication_enabled": schema.BoolAttribute{
				Description: "Indicates whether two-factor authentication is enabled for the user account. Does not apply to API authentication.",
				Computed:    true,
			},
			"two_factor_authentication_locked": schema.BoolAttribute{
				Description: "Indicates whether two-factor authentication is required by one of the accounts that the user is a member of.",
				Computed:    true,
			},
			"zipcode": schema.StringAttribute{
				Description: "The zipcode or postal code where the user lives.",
				Computed:    true,
			},
			"betas": schema.ListAttribute{
				Description: "Lists the betas that the user is participating in.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"organizations": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[UserOrganizationsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Organization name.",
							Computed:    true,
						},
						"permissions": schema.ListAttribute{
							Description: "Access permissions for this User.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"roles": schema.ListAttribute{
							Description: "List of roles that a user has within an organization.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"status": schema.StringAttribute{
							Description: "Whether the user is a member of the organization or has an invitation pending.\nAvailable values: \"member\", \"invited\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("member", "invited"),
							},
						},
					},
				},
			},
		},
	}
}

func (d *UserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *UserDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
