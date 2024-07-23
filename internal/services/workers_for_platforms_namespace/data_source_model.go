// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsNamespaceResultDataSourceEnvelope struct {
	Result WorkersForPlatformsNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsNamespaceResultListDataSourceEnvelope struct {
	Result *[]*WorkersForPlatformsNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsNamespaceDataSourceModel struct {
	AccountID         types.String                                          `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String                                          `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	CreatedBy         types.String                                          `tfsdk:"created_by" json:"created_by"`
	CreatedOn         types.String                                          `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedBy        types.String                                          `tfsdk:"modified_by" json:"modified_by"`
	ModifiedOn        types.String                                          `tfsdk:"modified_on" json:"modified_on,computed"`
	NamespaceID       types.String                                          `tfsdk:"namespace_id" json:"namespace_id"`
	NamespaceName     types.String                                          `tfsdk:"namespace_name" json:"namespace_name"`
	ScriptCount       types.Int64                                           `tfsdk:"script_count" json:"script_count"`
	FindOneBy         *WorkersForPlatformsNamespaceFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type WorkersForPlatformsNamespaceFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
