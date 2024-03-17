package workers_for_platforms

import "github.com/hashicorp/terraform-plugin-framework/types"

type WorkersForPlatformsNamespaceModel struct {
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
	ID        types.String `tfsdk:"id"`
}
