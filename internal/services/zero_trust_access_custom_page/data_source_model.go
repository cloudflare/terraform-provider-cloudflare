// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPageResultDataSourceEnvelope struct {
	Result ZeroTrustAccessCustomPageDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessCustomPageResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessCustomPageDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessCustomPageDataSourceModel struct {
	AccountID    types.String                                       `tfsdk:"account_id" path:"account_id,optional"`
	CustomPageID types.String                                       `tfsdk:"custom_page_id" path:"custom_page_id,optional"`
	CustomHTML   types.String                                       `tfsdk:"custom_html" json:"custom_html,optional"`
	AppCount     types.Int64                                        `tfsdk:"app_count" json:"app_count,computed"`
	CreatedAt    timetypes.RFC3339                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name         types.String                                       `tfsdk:"name" json:"name,computed"`
	Type         types.String                                       `tfsdk:"type" json:"type,computed"`
	UID          types.String                                       `tfsdk:"uid" json:"uid,computed"`
	UpdatedAt    timetypes.RFC3339                                  `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Filter       *ZeroTrustAccessCustomPageFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustAccessCustomPageDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessCustomPageGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCustomPageGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustAccessCustomPageDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessCustomPageListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCustomPageListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustAccessCustomPageFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
