// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersForPlatformsDispatchNamespaceResultEnvelope struct {
	Result WorkersForPlatformsDispatchNamespaceModel `json:"result"`
}

type WorkersForPlatformsDispatchNamespaceModel struct {
	ID            types.String      `tfsdk:"id" json:"-,computed"`
	NamespaceID   types.String      `tfsdk:"namespace_id" json:"namespace_id,computed"`
	AccountID     types.String      `tfsdk:"account_id" path:"account_id"`
	Name          types.String      `tfsdk:"name" json:"name"`
	CreatedBy     types.String      `tfsdk:"created_by" json:"created_by,computed"`
	CreatedOn     timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedBy    types.String      `tfsdk:"modified_by" json:"modified_by,computed"`
	ModifiedOn    timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	NamespaceName types.String      `tfsdk:"namespace_name" json:"namespace_name,computed"`
	ScriptCount   types.Int64       `tfsdk:"script_count" json:"script_count,computed"`
}
