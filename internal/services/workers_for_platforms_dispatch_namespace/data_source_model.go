// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsDispatchNamespaceResultDataSourceEnvelope struct {
	Result WorkersForPlatformsDispatchNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsDispatchNamespaceResultListDataSourceEnvelope struct {
	Result *[]*WorkersForPlatformsDispatchNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsDispatchNamespaceDataSourceModel struct {
	AccountID         types.String                                                  `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String                                                  `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	CreatedOn         timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn        timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed"`
	CreatedBy         types.String                                                  `tfsdk:"created_by" json:"created_by"`
	ModifiedBy        types.String                                                  `tfsdk:"modified_by" json:"modified_by"`
	NamespaceID       types.String                                                  `tfsdk:"namespace_id" json:"namespace_id"`
	NamespaceName     types.String                                                  `tfsdk:"namespace_name" json:"namespace_name"`
	ScriptCount       types.Int64                                                   `tfsdk:"script_count" json:"script_count"`
	Filter            *WorkersForPlatformsDispatchNamespaceFindOneByDataSourceModel `tfsdk:"filter"`
}

type WorkersForPlatformsDispatchNamespaceFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
