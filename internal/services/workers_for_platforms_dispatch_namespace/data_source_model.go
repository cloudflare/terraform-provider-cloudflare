// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsDispatchNamespaceResultDataSourceEnvelope struct {
	Result WorkersForPlatformsDispatchNamespaceDataSourceModel `json:"result,computed"`
}

type WorkersForPlatformsDispatchNamespaceDataSourceModel struct {
	ID                types.String      `tfsdk:"id" path:"dispatch_namespace,computed"`
	DispatchNamespace types.String      `tfsdk:"dispatch_namespace" path:"dispatch_namespace,optional"`
	AccountID         types.String      `tfsdk:"account_id" path:"account_id,required"`
	CreatedBy         types.String      `tfsdk:"created_by" json:"created_by,computed"`
	CreatedOn         timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedBy        types.String      `tfsdk:"modified_by" json:"modified_by,computed"`
	ModifiedOn        timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	NamespaceID       types.String      `tfsdk:"namespace_id" json:"namespace_id,computed"`
	NamespaceName     types.String      `tfsdk:"namespace_name" json:"namespace_name,computed"`
	ScriptCount       types.Int64       `tfsdk:"script_count" json:"script_count,computed"`
}

func (m *WorkersForPlatformsDispatchNamespaceDataSourceModel) toReadParams(_ context.Context) (params workers_for_platforms.DispatchNamespaceGetParams, diags diag.Diagnostics) {
	params = workers_for_platforms.DispatchNamespaceGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}
