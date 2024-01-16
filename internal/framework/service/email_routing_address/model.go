package email_routing_address

import "github.com/hashicorp/terraform-plugin-framework/types"

type EmailRoutingAddressModel struct {
	AccountID types.String `tfsdk:"account_id"`
	ID        types.String `tfsdk:"id"`
	Tag       types.String `tfsdk:"tag"`
	Email     types.String `tfsdk:"email"`
	Verified  types.String `tfsdk:"verified"`
	Created   types.String `tfsdk:"created"`
	Modified  types.String `tfsdk:"modified"`
}
