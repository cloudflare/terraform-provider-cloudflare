package d1

import "github.com/hashicorp/terraform-plugin-framework/types"

type DatabaseModel struct {
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
	ID        types.String `tfsdk:"id"`
	Version   types.String `tfsdk:"version"`
}
