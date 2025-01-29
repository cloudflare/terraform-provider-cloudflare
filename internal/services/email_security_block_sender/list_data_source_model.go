// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_block_sender

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityBlockSendersResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSecurityBlockSendersResultDataSourceModel] `json:"result,computed"`
}

type EmailSecurityBlockSendersDataSourceModel struct {
	AccountID   types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	Direction   types.String                                                                 `tfsdk:"direction" query:"direction,optional"`
	Order       types.String                                                                 `tfsdk:"order" query:"order,optional"`
	PatternType types.String                                                                 `tfsdk:"pattern_type" query:"pattern_type,optional"`
	Search      types.String                                                                 `tfsdk:"search" query:"search,optional"`
	MaxItems    types.Int64                                                                  `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[EmailSecurityBlockSendersResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailSecurityBlockSendersDataSourceModel) toListParams(_ context.Context) (params email_security.SettingBlockSenderListParams, diags diag.Diagnostics) {
	params = email_security.SettingBlockSenderListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingBlockSenderListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingBlockSenderListParamsOrder(m.Order.ValueString()))
	}
	if !m.PatternType.IsNull() {
		params.PatternType = cloudflare.F(email_security.SettingBlockSenderListParamsPatternType(m.PatternType.ValueString()))
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	return
}

type EmailSecurityBlockSendersResultDataSourceModel struct {
	ID           types.Int64       `tfsdk:"id" json:"id,computed"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsRegex      types.Bool        `tfsdk:"is_regex" json:"is_regex,computed"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Pattern      types.String      `tfsdk:"pattern" json:"pattern,computed"`
	PatternType  types.String      `tfsdk:"pattern_type" json:"pattern_type,computed"`
	Comments     types.String      `tfsdk:"comments" json:"comments,computed"`
}
