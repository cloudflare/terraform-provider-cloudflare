// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSTenantStructuresResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustCWSTenantStructuresResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustCWSTenantStructuresDataSourceModel struct {
	AccountID  types.String                                                                    `tfsdk:"account_id" path:"account_id,required"`
	ImportPath types.String                                                                    `tfsdk:"import_path" query:"import_path,optional"`
	Name       types.String                                                                    `tfsdk:"name" query:"name,optional"`
	Direction  types.String                                                                    `tfsdk:"direction" query:"direction,computed_optional"`
	Order      types.String                                                                    `tfsdk:"order" query:"order,computed_optional"`
	MaxItems   types.Int64                                                                     `tfsdk:"max_items"`
	Result     customfield.NestedObjectList[ZeroTrustCWSTenantStructuresResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustCWSTenantStructuresDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSTenantStructureListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSTenantStructureListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSTenantStructureListParamsDirection(m.Direction.ValueString()))
	}
	if !m.ImportPath.IsNull() {
		params.ImportPath = cloudflare.F(m.ImportPath.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSTenantStructureListParamsOrder(m.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSTenantStructuresResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Application types.String      `tfsdk:"application" json:"application,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name        types.String      `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ImportPath  types.String      `tfsdk:"import_path" json:"import_path,computed"`
}
