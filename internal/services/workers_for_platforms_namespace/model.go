// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsNamespaceResultEnvelope struct {
	Result WorkersForPlatformsNamespaceModel `json:"result,computed"`
}

type WorkersForPlatformsNamespaceResultDataSourceEnvelope struct {
	Result WorkersForPlatformsNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsNamespacesResultDataSourceEnvelope struct {
	Result WorkersForPlatformsNamespacesDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsNamespaceModel struct {
	ID            types.String `tfsdk:"id" json:"-,computed"`
	NamespaceID   types.String `tfsdk:"namespace_id" json:"namespace_id"`
	AccountID     types.String `tfsdk:"account_id" path:"account_id"`
	Name          types.String `tfsdk:"name" json:"name"`
	CreatedBy     types.String `tfsdk:"created_by" json:"created_by,computed"`
	CreatedOn     types.String `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedBy    types.String `tfsdk:"modified_by" json:"modified_by,computed"`
	ModifiedOn    types.String `tfsdk:"modified_on" json:"modified_on,computed"`
	NamespaceName types.String `tfsdk:"namespace_name" json:"namespace_name,computed"`
	ScriptCount   types.Int64  `tfsdk:"script_count" json:"script_count,computed"`
}

type WorkersForPlatformsNamespaceDataSourceModel struct {
}

type WorkersForPlatformsNamespacesDataSourceModel struct {
}
