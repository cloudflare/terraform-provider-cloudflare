package gateway_categories

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GatewayCategoriesModel struct {
	AccountID  types.String           `tfsdk:"account_id"`
	Categories []GatewayCategoryModel `tfsdk:"categories"`
}

type GatewayCategoryModel struct {
	ID            types.Int64               `tfsdk:"id"`
	Name          types.String              `tfsdk:"name"`
	Description   types.String              `tfsdk:"description"`
	Class         types.String              `tfsdk:"class"`
	Beta          types.Bool                `tfsdk:"beta"`
	Subcategories []GatewaySubCategoryModel `tfsdk:"subcategories"`
}

type GatewaySubCategoryModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Class       types.String `tfsdk:"class"`
	Beta        types.Bool   `tfsdk:"beta"`
}
