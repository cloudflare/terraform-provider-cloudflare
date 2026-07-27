// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group_members

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*UserGroupMembersDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Settings Read",
				"Account Settings Write",
				"SCIM Provisioning",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "User Group identifier tag.",
				Computed:    true,
			},
			"user_group_id": schema.StringAttribute{
				Description: "User Group identifier tag.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"fuzzy_email": schema.StringAttribute{
				Description: "A string used for filtering members by partial email match.",
				Optional:    true,
			},
			"direction": schema.StringAttribute{
				Description: "The sort order of returned user group members by email.\nAvailable values: \"asc\", \"desc\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"email": schema.StringAttribute{
				Description: "The contact email address of the user.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The member's status in the account.\nAvailable values: \"accepted\", \"pending\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("accepted", "pending"),
				},
			},
		},
	}
}

func (d *UserGroupMembersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *UserGroupMembersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
