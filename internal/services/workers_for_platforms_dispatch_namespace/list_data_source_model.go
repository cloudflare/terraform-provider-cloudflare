// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsDispatchNamespacesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersForPlatformsDispatchNamespacesResultDataSourceModel] `json:"result,computed"`
}

type WorkersForPlatformsDispatchNamespacesDataSourceModel struct {
	AccountID types.String                                                                             `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[WorkersForPlatformsDispatchNamespacesResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersForPlatformsDispatchNamespacesDataSourceModel) toListParams(_ context.Context) (params workers_for_platforms.DispatchNamespaceListParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type WorkersForPlatformsDispatchNamespacesResultDataSourceModel struct {
	CreatedBy     types.String      `tfsdk:"created_by" json:"created_by,computed"`
	CreatedOn     timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedBy    types.String      `tfsdk:"modified_by" json:"modified_by,computed"`
	ModifiedOn    timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	NamespaceID   types.String      `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName types.String      `tfsdk:"namespace_name" json:"namespace_name,computed"`
	ScriptCount   types.Int64       `tfsdk:"script_count" json:"script_count,computed"`
}
