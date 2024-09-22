package zero_trust_access_groups

import "github.com/hashicorp/terraform-plugin-framework/types"

// ZeroTrustAccessGroupModel describes the data source data model for groups
type ZeroTrustAccessGroupsModel struct {
	AccountID types.String                `tfsdk:"account_id"`
	Groups    []ZeroTrustAccessGroupModel `tfsdk:"groups"`
}

// ZeroTrustAccessGroupModel describes the data source data model for a group
type ZeroTrustAccessGroupModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}
