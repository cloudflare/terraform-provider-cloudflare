// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_pattern

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityAllowPatternsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSecurityAllowPatternsResultDataSourceModel] `json:"result,computed"`
}

type EmailSecurityAllowPatternsDataSourceModel struct {
	AccountID    types.String                                                                  `tfsdk:"account_id" path:"account_id,required"`
	Direction    types.String                                                                  `tfsdk:"direction" query:"direction,optional"`
	IsRecipient  types.Bool                                                                    `tfsdk:"is_recipient" query:"is_recipient,optional"`
	IsSender     types.Bool                                                                    `tfsdk:"is_sender" query:"is_sender,optional"`
	IsSpoof      types.Bool                                                                    `tfsdk:"is_spoof" query:"is_spoof,optional"`
	Order        types.String                                                                  `tfsdk:"order" query:"order,optional"`
	PatternType  types.String                                                                  `tfsdk:"pattern_type" query:"pattern_type,optional"`
	Search       types.String                                                                  `tfsdk:"search" query:"search,optional"`
	VerifySender types.Bool                                                                    `tfsdk:"verify_sender" query:"verify_sender,optional"`
	MaxItems     types.Int64                                                                   `tfsdk:"max_items"`
	Result       customfield.NestedObjectList[EmailSecurityAllowPatternsResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailSecurityAllowPatternsDataSourceModel) toListParams(_ context.Context) (params email_security.SettingAllowPatternListParams, diags diag.Diagnostics) {
	params = email_security.SettingAllowPatternListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingAllowPatternListParamsDirection(m.Direction.ValueString()))
	}
	if !m.IsRecipient.IsNull() {
		params.IsRecipient = cloudflare.F(m.IsRecipient.ValueBool())
	}
	if !m.IsSender.IsNull() {
		params.IsSender = cloudflare.F(m.IsSender.ValueBool())
	}
	if !m.IsSpoof.IsNull() {
		params.IsSpoof = cloudflare.F(m.IsSpoof.ValueBool())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingAllowPatternListParamsOrder(m.Order.ValueString()))
	}
	if !m.PatternType.IsNull() {
		params.PatternType = cloudflare.F(email_security.SettingAllowPatternListParamsPatternType(m.PatternType.ValueString()))
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}
	if !m.VerifySender.IsNull() {
		params.VerifySender = cloudflare.F(m.VerifySender.ValueBool())
	}

	return
}

type EmailSecurityAllowPatternsResultDataSourceModel struct {
	ID           types.Int64       `tfsdk:"id" json:"id,computed"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsRecipient  types.Bool        `tfsdk:"is_recipient" json:"is_recipient,computed"`
	IsRegex      types.Bool        `tfsdk:"is_regex" json:"is_regex,computed"`
	IsSender     types.Bool        `tfsdk:"is_sender" json:"is_sender,computed"`
	IsSpoof      types.Bool        `tfsdk:"is_spoof" json:"is_spoof,computed"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Pattern      types.String      `tfsdk:"pattern" json:"pattern,computed"`
	PatternType  types.String      `tfsdk:"pattern_type" json:"pattern_type,computed"`
	VerifySender types.Bool        `tfsdk:"verify_sender" json:"verify_sender,computed"`
	Comments     types.String      `tfsdk:"comments" json:"comments,computed"`
}
