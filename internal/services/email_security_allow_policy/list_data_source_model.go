// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityAllowPoliciesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailSecurityAllowPoliciesResultDataSourceModel] `json:"result,computed"`
}

type EmailSecurityAllowPoliciesDataSourceModel struct {
	AccountID          types.String                                                                  `tfsdk:"account_id" path:"account_id,required"`
	Direction          types.String                                                                  `tfsdk:"direction" query:"direction,optional"`
	IsAcceptableSender types.Bool                                                                    `tfsdk:"is_acceptable_sender" query:"is_acceptable_sender,optional"`
	IsExemptRecipient  types.Bool                                                                    `tfsdk:"is_exempt_recipient" query:"is_exempt_recipient,optional"`
	IsTrustedSender    types.Bool                                                                    `tfsdk:"is_trusted_sender" query:"is_trusted_sender,optional"`
	Order              types.String                                                                  `tfsdk:"order" query:"order,optional"`
	Pattern            types.String                                                                  `tfsdk:"pattern" query:"pattern,optional"`
	PatternType        types.String                                                                  `tfsdk:"pattern_type" query:"pattern_type,optional"`
	Search             types.String                                                                  `tfsdk:"search" query:"search,optional"`
	VerifySender       types.Bool                                                                    `tfsdk:"verify_sender" query:"verify_sender,optional"`
	MaxItems           types.Int64                                                                   `tfsdk:"max_items"`
	Result             customfield.NestedObjectList[EmailSecurityAllowPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *EmailSecurityAllowPoliciesDataSourceModel) toListParams(_ context.Context) (params email_security.SettingAllowPolicyListParams, diags diag.Diagnostics) {
	params = email_security.SettingAllowPolicyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingAllowPolicyListParamsDirection(m.Direction.ValueString()))
	}
	if !m.IsAcceptableSender.IsNull() {
		params.IsAcceptableSender = cloudflare.F(m.IsAcceptableSender.ValueBool())
	}
	if !m.IsExemptRecipient.IsNull() {
		params.IsExemptRecipient = cloudflare.F(m.IsExemptRecipient.ValueBool())
	}
	if !m.IsTrustedSender.IsNull() {
		params.IsTrustedSender = cloudflare.F(m.IsTrustedSender.ValueBool())
	}
	if !m.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingAllowPolicyListParamsOrder(m.Order.ValueString()))
	}
	if !m.Pattern.IsNull() {
		params.Pattern = cloudflare.F(m.Pattern.ValueString())
	}
	if !m.PatternType.IsNull() {
		params.PatternType = cloudflare.F(email_security.SettingAllowPolicyListParamsPatternType(m.PatternType.ValueString()))
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}
	if !m.VerifySender.IsNull() {
		params.VerifySender = cloudflare.F(m.VerifySender.ValueBool())
	}

	return
}

type EmailSecurityAllowPoliciesResultDataSourceModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastModified       timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	Comments           types.String      `tfsdk:"comments" json:"comments,computed"`
	IsAcceptableSender types.Bool        `tfsdk:"is_acceptable_sender" json:"is_acceptable_sender,computed"`
	IsExemptRecipient  types.Bool        `tfsdk:"is_exempt_recipient" json:"is_exempt_recipient,computed"`
	IsRecipient        types.Bool        `tfsdk:"is_recipient" json:"is_recipient,computed"`
	IsRegex            types.Bool        `tfsdk:"is_regex" json:"is_regex,computed"`
	IsSender           types.Bool        `tfsdk:"is_sender" json:"is_sender,computed"`
	IsSpoof            types.Bool        `tfsdk:"is_spoof" json:"is_spoof,computed"`
	IsTrustedSender    types.Bool        `tfsdk:"is_trusted_sender" json:"is_trusted_sender,computed"`
	ModifiedAt         timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Pattern            types.String      `tfsdk:"pattern" json:"pattern,computed"`
	PatternType        types.String      `tfsdk:"pattern_type" json:"pattern_type,computed"`
	VerifySender       types.Bool        `tfsdk:"verify_sender" json:"verify_sender,computed"`
}
