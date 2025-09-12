// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyResultEnvelope struct {
	Result AccessPolicyModel `json:"result"`
}

type AccessPolicyModel struct {
	ID                           types.String `tfsdk:"id" json:"id,computed"`
	AccountID                    types.String `tfsdk:"account_id" path:"account_id,required"`
	ZoneID                       types.String `tfsdk:"zone_id" path:"zone_id,required"`
	ApplicationID                types.String `tfsdk:"application_id" json:"application_id,required"`
	Decision                     types.String `tfsdk:"decision" json:"decision,required"`
	Name                         types.String `tfsdk:"name" json:"name,required"`
	Precedence                   types.Int64  `tfsdk:"precedence" json:"precedence,required"`
	ApprovalRequired             types.Bool   `tfsdk:"approval_required" json:"approval_required,optional"`
	IsolationRequired            types.Bool   `tfsdk:"isolation_required" json:"isolation_required,optional"`
	PurposeJustificationPrompt   types.String `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,optional"`
	PurposeJustificationRequired types.Bool   `tfsdk:"purpose_justification_required" json:"purpose_justification_required,optional"`
	SessionDuration              types.String `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	Exclude                      types.List   `tfsdk:"exclude" json:"exclude,optional"`
	Include                      types.List   `tfsdk:"include" json:"include,optional"`
	Require                      types.List   `tfsdk:"require" json:"require,optional"`
}
