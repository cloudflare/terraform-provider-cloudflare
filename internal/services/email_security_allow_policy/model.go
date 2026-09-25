// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_policy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityAllowPolicyResultEnvelope struct {
	Result EmailSecurityAllowPolicyModel `json:"result"`
}

type EmailSecurityAllowPolicyModel struct {
	ID                 types.String      `tfsdk:"id" json:"id,computed"`
	AccountID          types.String      `tfsdk:"account_id" path:"account_id,required"`
	IsAcceptableSender types.Bool        `tfsdk:"is_acceptable_sender" json:"is_acceptable_sender,required"`
	IsExemptRecipient  types.Bool        `tfsdk:"is_exempt_recipient" json:"is_exempt_recipient,required"`
	IsRegex            types.Bool        `tfsdk:"is_regex" json:"is_regex,required"`
	IsTrustedSender    types.Bool        `tfsdk:"is_trusted_sender" json:"is_trusted_sender,required"`
	Pattern            types.String      `tfsdk:"pattern" json:"pattern,required"`
	PatternType        types.String      `tfsdk:"pattern_type" json:"pattern_type,required"`
	VerifySender       types.Bool        `tfsdk:"verify_sender" json:"verify_sender,required"`
	Comments           types.String      `tfsdk:"comments" json:"comments,optional"`
	IsRecipient        types.Bool        `tfsdk:"is_recipient" json:"is_recipient,optional"`
	IsSender           types.Bool        `tfsdk:"is_sender" json:"is_sender,optional"`
	IsSpoof            types.Bool        `tfsdk:"is_spoof" json:"is_spoof,optional"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastModified       timetypes.RFC3339 `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
	ModifiedAt         timetypes.RFC3339 `tfsdk:"modified_at" json:"modified_at,computed" format:"date-time"`
}

func (m EmailSecurityAllowPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailSecurityAllowPolicyModel) MarshalJSONForUpdate(state EmailSecurityAllowPolicyModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
