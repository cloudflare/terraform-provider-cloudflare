// File generated for StateUpgrader migration from v4 to v5

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceWorkersKVNamespaceModel represents the v4 (SDKv2) state structure
type SourceWorkersKVNamespaceModel struct {
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`
	Title     types.String `tfsdk:"title"`
}

// TargetWorkersKVNamespaceModel represents the v5 (Framework) state structure
type TargetWorkersKVNamespaceModel struct {
	ID                  types.String `tfsdk:"id"`
	AccountID           types.String `tfsdk:"account_id"`
	Title               types.String `tfsdk:"title"`
	SupportsURLEncoding types.Bool   `tfsdk:"supports_url_encoding"`
}
