// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moq_relay

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/moq"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MoQRelaysResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MoQRelaysResultDataSourceModel] `json:"result,computed"`
}

type MoQRelaysDataSourceModel struct {
	AccountID     types.String                                                 `tfsdk:"account_id" path:"account_id,required"`
	CreatedAfter  timetypes.RFC3339                                            `tfsdk:"created_after" query:"created_after,optional" format:"date-time"`
	CreatedBefore timetypes.RFC3339                                            `tfsdk:"created_before" query:"created_before,optional" format:"date-time"`
	PerPage       types.Int64                                                  `tfsdk:"per_page" query:"per_page,optional"`
	Asc           types.Bool                                                   `tfsdk:"asc" query:"asc,computed_optional"`
	MaxItems      types.Int64                                                  `tfsdk:"max_items"`
	Result        customfield.NestedObjectList[MoQRelaysResultDataSourceModel] `tfsdk:"result"`
}

func (m *MoQRelaysDataSourceModel) toListParams(_ context.Context) (params moq.RelayListParams, diags diag.Diagnostics) {
	mCreatedAfter, errs := m.CreatedAfter.ValueRFC3339Time()
	diags.Append(errs...)
	mCreatedBefore, errs := m.CreatedBefore.ValueRFC3339Time()
	diags.Append(errs...)

	params = moq.RelayListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Asc.IsNull() {
		params.Asc = cloudflare.F(m.Asc.ValueBool())
	}
	if !m.CreatedAfter.IsNull() {
		params.CreatedAfter = cloudflare.F(mCreatedAfter)
	}
	if !m.CreatedBefore.IsNull() {
		params.CreatedBefore = cloudflare.F(mCreatedBefore)
	}
	if !m.PerPage.IsNull() {
		params.PerPage = cloudflare.F(m.PerPage.ValueInt64())
	}

	return
}

type MoQRelaysResultDataSourceModel struct {
	ID       types.String      `tfsdk:"id" json:"uid,computed"`
	Created  timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Name     types.String      `tfsdk:"name" json:"name,computed"`
	UID      types.String      `tfsdk:"uid" json:"uid,computed"`
}
