package user

import "github.com/hashicorp/terraform-plugin-framework/types"

type CloudflareUserDataSourceModel struct {
	ID       types.String `tfsdk:"id"`
	Email    types.String `tfsdk:"email"`
	Username types.String `tfsdk:"username"`
}
