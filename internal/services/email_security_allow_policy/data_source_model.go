// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityAllowPolicyResultDataSourceEnvelope struct {
	Result EmailSecurityAllowPolicyDataSourceModel `json:"result,computed"`
}

type EmailSecurityAllowPolicyDataSourceModel struct {
	ID                 types.String                                      `tfsdk:"id" path:"policy_id,computed"`
	PolicyID           types.String                                      `tfsdk:"policy_id" path:"policy_id,optional"`
	AccountID          types.String                                      `tfsdk:"account_id" path:"account_id,required"`
	Comments           types.String                                      `tfsdk:"comments" json:"comments,computed"`
	CreatedAt          timetypes.RFC3339                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	IsAcceptableSender types.Bool                                        `tfsdk:"is_acceptable_sender" json:"is_acceptable_sender,computed"`
	IsExemptRecipient  types.Bool                                        `tfsdk:"is_exempt_recipient" json:"is_exempt_recipient,computed"`
	IsRecipient        types.Bool                                        `tfsdk:"is_recipient" json:"is_recipient,computed"`
	IsRegex            types.Bool                                        `tfsdk:"is_regex" json:"is_regex,computed"`
	IsSender           types.Bool                                        `tfsdk:"is_sender" json:"is_sender,computed"`
	IsSpoof            types.Bool                                        `tfsdk:"is_spoof" json:"is_spoof,computed"`
	IsTrustedSender    types.Bool                                        `tfsdk:"is_trusted_sender" json:"is_trusted_sender,computed"`
	LastModified       timetypes.RFC3339                                 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	ModifiedAt         timetypes.RFC3339                                 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
	Pattern            types.String                                      `tfsdk:"pattern" json:"pattern,computed"`
	PatternType        types.String                                      `tfsdk:"pattern_type" json:"pattern_type,computed"`
	VerifySender       types.Bool                                        `tfsdk:"verify_sender" json:"verify_sender,computed"`
	Filter             *EmailSecurityAllowPolicyFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *EmailSecurityAllowPolicyDataSourceModel) toReadParams(_ context.Context) (params email_security.SettingAllowPolicyGetParams, diags diag.Diagnostics) {
	params = email_security.SettingAllowPolicyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *EmailSecurityAllowPolicyDataSourceModel) toListParams(_ context.Context) (params email_security.SettingAllowPolicyListParams, diags diag.Diagnostics) {
	params = email_security.SettingAllowPolicyListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(email_security.SettingAllowPolicyListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.IsAcceptableSender.IsNull() {
		params.IsAcceptableSender = cloudflare.F(m.Filter.IsAcceptableSender.ValueBool())
	}
	if !m.Filter.IsExemptRecipient.IsNull() {
		params.IsExemptRecipient = cloudflare.F(m.Filter.IsExemptRecipient.ValueBool())
	}
	if !m.Filter.IsTrustedSender.IsNull() {
		params.IsTrustedSender = cloudflare.F(m.Filter.IsTrustedSender.ValueBool())
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(email_security.SettingAllowPolicyListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Pattern.IsNull() {
		params.Pattern = cloudflare.F(m.Filter.Pattern.ValueString())
	}
	if !m.Filter.PatternType.IsNull() {
		params.PatternType = cloudflare.F(email_security.SettingAllowPolicyListParamsPatternType(m.Filter.PatternType.ValueString()))
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}
	if !m.Filter.VerifySender.IsNull() {
		params.VerifySender = cloudflare.F(m.Filter.VerifySender.ValueBool())
	}

	return
}

type EmailSecurityAllowPolicyFindOneByDataSourceModel struct {
	Direction          types.String `tfsdk:"direction" query:"direction,optional"`
	IsAcceptableSender types.Bool   `tfsdk:"is_acceptable_sender" query:"is_acceptable_sender,optional"`
	IsExemptRecipient  types.Bool   `tfsdk:"is_exempt_recipient" query:"is_exempt_recipient,optional"`
	IsTrustedSender    types.Bool   `tfsdk:"is_trusted_sender" query:"is_trusted_sender,optional"`
	Order              types.String `tfsdk:"order" query:"order,optional"`
	Pattern            types.String `tfsdk:"pattern" query:"pattern,optional"`
	PatternType        types.String `tfsdk:"pattern_type" query:"pattern_type,optional"`
	Search             types.String `tfsdk:"search" query:"search,optional"`
	VerifySender       types.Bool   `tfsdk:"verify_sender" query:"verify_sender,optional"`
}
