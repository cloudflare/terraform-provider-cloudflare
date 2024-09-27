// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestMessageResultEnvelope struct {
	Result CloudforceOneRequestMessageModel `json:"result"`
}

type CloudforceOneRequestMessageModel struct {
	ID                types.Int64                                                            `tfsdk:"id" json:"id,computed"`
	AccountIdentifier types.String                                                           `tfsdk:"account_identifier" path:"account_identifier,required"`
	RequestIdentifier types.String                                                           `tfsdk:"request_identifier" path:"request_identifier,required"`
	Content           types.String                                                           `tfsdk:"content" json:"content,computed_optional"`
	Author            types.String                                                           `tfsdk:"author" json:"author,computed"`
	Created           timetypes.RFC3339                                                      `tfsdk:"created" json:"created,computed" format:"date-time"`
	IsFollowOnRequest types.Bool                                                             `tfsdk:"is_follow_on_request" json:"is_follow_on_request,computed"`
	Success           types.Bool                                                             `tfsdk:"success" json:"success,computed"`
	Updated           timetypes.RFC3339                                                      `tfsdk:"updated" json:"updated,computed" format:"date-time"`
	Errors            customfield.NestedObjectList[CloudforceOneRequestMessageErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages          customfield.NestedObjectList[CloudforceOneRequestMessageMessagesModel] `tfsdk:"messages" json:"messages,computed"`
}

type CloudforceOneRequestMessageErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CloudforceOneRequestMessageMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}
