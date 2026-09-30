// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_permission_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/iam"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountPermissionGroupResultDataSourceEnvelope struct {
	Result AccountPermissionGroupDataSourceModel `json:"result,computed"`
}

type AccountPermissionGroupDataSourceModel struct {
	AccountID         types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	PermissionGroupID types.String                                                        `tfsdk:"permission_group_id" path:"permission_group_id,required"`
	ID                types.String                                                        `tfsdk:"id" json:"id,computed"`
	Name              types.String                                                        `tfsdk:"name" json:"name,computed"`
	Meta              customfield.NestedObject[AccountPermissionGroupMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
}

func (m *AccountPermissionGroupDataSourceModel) toReadParams(_ context.Context) (params iam.PermissionGroupGetParams, diags diag.Diagnostics) {
	params = iam.PermissionGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type AccountPermissionGroupMetaDataSourceModel struct {
	Category    types.String      `tfsdk:"category" json:"category,computed"`
	Deprecated  types.String      `tfsdk:"deprecated" json:"deprecated,computed"`
	Description types.String      `tfsdk:"description" json:"description,computed"`
	Editable    types.String      `tfsdk:"editable" json:"editable,computed"`
	EolAt       timetypes.RFC3339 `tfsdk:"eol_at" json:"eol_at,computed" format:"date-time"`
	Label       types.String      `tfsdk:"label" json:"label,computed"`
	Scopes      types.String      `tfsdk:"scopes" json:"scopes,computed"`
	Visibility  types.String      `tfsdk:"visibility" json:"visibility,computed"`
}
