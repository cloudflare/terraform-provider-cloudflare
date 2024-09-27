// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestResultEnvelope struct {
	Result CloudforceOneRequestModel `json:"result"`
}

type CloudforceOneRequestModel struct {
	ID                types.String                                                    `tfsdk:"id" json:"id,computed"`
	AccountIdentifier types.String                                                    `tfsdk:"account_identifier" path:"account_identifier,required"`
	Priority          types.String                                                    `tfsdk:"priority" json:"priority,optional"`
	RequestType       types.String                                                    `tfsdk:"request_type" json:"request_type,optional"`
	Content           types.String                                                    `tfsdk:"content" json:"content,computed_optional"`
	Summary           types.String                                                    `tfsdk:"summary" json:"summary,computed_optional"`
	Tlp               types.String                                                    `tfsdk:"tlp" json:"tlp,computed_optional"`
	Completed         timetypes.RFC3339                                               `tfsdk:"completed" json:"completed,computed" format:"date-time"`
	Created           timetypes.RFC3339                                               `tfsdk:"created" json:"created,computed" format:"date-time"`
	MessageTokens     types.Int64                                                     `tfsdk:"message_tokens" json:"message_tokens,computed"`
	ReadableID        types.String                                                    `tfsdk:"readable_id" json:"readable_id,computed"`
	Request           types.String                                                    `tfsdk:"request" json:"request,computed"`
	Status            types.String                                                    `tfsdk:"status" json:"status,computed"`
	Success           types.Bool                                                      `tfsdk:"success" json:"success,computed"`
	Tokens            types.Int64                                                     `tfsdk:"tokens" json:"tokens,computed"`
	Updated           timetypes.RFC3339                                               `tfsdk:"updated" json:"updated,computed" format:"date-time"`
	Errors            customfield.NestedObjectList[CloudforceOneRequestErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages          customfield.NestedObjectList[CloudforceOneRequestMessagesModel] `tfsdk:"messages" json:"messages,computed"`
}

type CloudforceOneRequestErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CloudforceOneRequestMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}
