// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsDispatchNamespaceResultDataSourceEnvelope struct {
	Result WorkersForPlatformsDispatchNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsDispatchNamespaceResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersForPlatformsDispatchNamespaceDataSourceModel] `json:"result,computed"`
}

type WorkersForPlatformsDispatchNamespaceDataSourceModel struct {
	AccountID         types.String                                                  `tfsdk:"account_id" path:"account_id"`
	DispatchNamespace types.String                                                  `tfsdk:"dispatch_namespace" path:"dispatch_namespace"`
	CreatedOn         timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed"`
	ModifiedOn        timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed"`
	CreatedBy         types.String                                                  `tfsdk:"created_by" json:"created_by,computed_optional"`
	ModifiedBy        types.String                                                  `tfsdk:"modified_by" json:"modified_by,computed_optional"`
	NamespaceID       types.String                                                  `tfsdk:"namespace_id" json:"namespace_id,computed_optional"`
	NamespaceName     types.String                                                  `tfsdk:"namespace_name" json:"namespace_name,computed_optional"`
	ScriptCount       types.Int64                                                   `tfsdk:"script_count" json:"script_count,computed_optional"`
	Filter            *WorkersForPlatformsDispatchNamespaceFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *WorkersForPlatformsDispatchNamespaceDataSourceModel) toReadParams() (params workers_for_platforms.DispatchNamespaceGetParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkersForPlatformsDispatchNamespaceDataSourceModel) toListParams() (params workers_for_platforms.DispatchNamespaceListParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type WorkersForPlatformsDispatchNamespaceFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
