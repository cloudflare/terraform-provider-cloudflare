package workers_for_platforms_dispatch_namespace

import "github.com/hashicorp/terraform-plugin-framework/types"

type WorkersForPlatformsDispatchNamespaceModel struct {
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
	ID        types.String `tfsdk:"id"`
}
