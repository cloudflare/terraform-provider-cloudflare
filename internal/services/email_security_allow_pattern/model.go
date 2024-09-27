// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security_allow_pattern

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailSecurityAllowPatternResultEnvelope struct {
	Result EmailSecurityAllowPatternModel `json:"result"`
}

type EmailSecurityAllowPatternModel struct {
	ID           types.Int64                            `tfsdk:"id" json:"id,computed"`
	AccountID    types.String                           `tfsdk:"account_id" path:"account_id,required"`
	Body         *[]*EmailSecurityAllowPatternBodyModel `tfsdk:"body" json:"body,optional"`
	Comments     types.String                           `tfsdk:"comments" json:"comments,computed_optional"`
	IsRecipient  types.Bool                             `tfsdk:"is_recipient" json:"is_recipient,computed_optional"`
	IsRegex      types.Bool                             `tfsdk:"is_regex" json:"is_regex,computed_optional"`
	IsSender     types.Bool                             `tfsdk:"is_sender" json:"is_sender,computed_optional"`
	IsSpoof      types.Bool                             `tfsdk:"is_spoof" json:"is_spoof,computed_optional"`
	Pattern      types.String                           `tfsdk:"pattern" json:"pattern,computed_optional"`
	PatternType  types.String                           `tfsdk:"pattern_type" json:"pattern_type,computed_optional"`
	VerifySender types.Bool                             `tfsdk:"verify_sender" json:"verify_sender,computed_optional"`
	CreatedAt    timetypes.RFC3339                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastModified timetypes.RFC3339                      `tfsdk:"last_modified" json:"last_modified,computed" format:"date-time"`
}

type EmailSecurityAllowPatternBodyModel struct {
	IsRecipient  types.Bool   `tfsdk:"is_recipient" json:"is_recipient,required"`
	IsRegex      types.Bool   `tfsdk:"is_regex" json:"is_regex,required"`
	IsSender     types.Bool   `tfsdk:"is_sender" json:"is_sender,required"`
	IsSpoof      types.Bool   `tfsdk:"is_spoof" json:"is_spoof,required"`
	Pattern      types.String `tfsdk:"pattern" json:"pattern,required"`
	PatternType  types.String `tfsdk:"pattern_type" json:"pattern_type,required"`
	VerifySender types.Bool   `tfsdk:"verify_sender" json:"verify_sender,required"`
	Comments     types.String `tfsdk:"comments" json:"comments,computed_optional"`
}
