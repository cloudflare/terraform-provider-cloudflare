// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_cws_tenant_structure

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustCWSTenantStructureResultDataSourceEnvelope struct {
	Result ZeroTrustCWSTenantStructureDataSourceModel `json:"result,computed"`
}

type ZeroTrustCWSTenantStructureDataSourceModel struct {
	ID          types.String                                         `tfsdk:"id" path:"structure_id,computed"`
	StructureID types.String                                         `tfsdk:"structure_id" path:"structure_id,optional"`
	AccountID   types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	Application types.String                                         `tfsdk:"application" json:"application,computed"`
	CreatedAt   timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ImportPath  types.String                                         `tfsdk:"import_path" json:"import_path,computed"`
	Name        types.String                                         `tfsdk:"name" json:"name,computed"`
	UpdatedAt   timetypes.RFC3339                                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Filter      *ZeroTrustCWSTenantStructureFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustCWSTenantStructureDataSourceModel) toReadParams(_ context.Context) (params zero_trust.CWSTenantStructureGetParams, diags diag.Diagnostics) {
	params = zero_trust.CWSTenantStructureGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustCWSTenantStructureDataSourceModel) toListParams(_ context.Context) (params zero_trust.CWSTenantStructureListParams, diags diag.Diagnostics) {
	params = zero_trust.CWSTenantStructureListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(zero_trust.CWSTenantStructureListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.ImportPath.IsNull() {
		params.ImportPath = cloudflare.F(m.Filter.ImportPath.ValueString())
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(zero_trust.CWSTenantStructureListParamsOrder(m.Filter.Order.ValueString()))
	}

	return
}

type ZeroTrustCWSTenantStructureFindOneByDataSourceModel struct {
	Direction  types.String `tfsdk:"direction" query:"direction,computed_optional"`
	ImportPath types.String `tfsdk:"import_path" query:"import_path,optional"`
	Name       types.String `tfsdk:"name" query:"name,optional"`
	Order      types.String `tfsdk:"order" query:"order,computed_optional"`
}
