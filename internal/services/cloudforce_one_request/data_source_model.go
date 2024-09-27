// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3/cloudforce_one"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestResultDataSourceEnvelope struct {
	Result CloudforceOneRequestDataSourceModel `json:"result,computed"`
}

type CloudforceOneRequestResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CloudforceOneRequestDataSourceModel] `json:"result,computed"`
}

type CloudforceOneRequestDataSourceModel struct {
	AccountIdentifier types.String                                  `tfsdk:"account_identifier" path:"account_identifier,optional"`
	RequestIdentifier types.String                                  `tfsdk:"request_identifier" path:"request_identifier,optional"`
	Content           types.String                                  `tfsdk:"content" json:"content,optional"`
	Completed         timetypes.RFC3339                             `tfsdk:"completed" json:"completed,computed" format:"date-time"`
	Created           timetypes.RFC3339                             `tfsdk:"created" json:"created,computed" format:"date-time"`
	ID                types.String                                  `tfsdk:"id" json:"id,computed"`
	MessageTokens     types.Int64                                   `tfsdk:"message_tokens" json:"message_tokens,computed"`
	ReadableID        types.String                                  `tfsdk:"readable_id" json:"readable_id,computed"`
	Request           types.String                                  `tfsdk:"request" json:"request,computed"`
	Status            types.String                                  `tfsdk:"status" json:"status,computed"`
	Summary           types.String                                  `tfsdk:"summary" json:"summary,computed"`
	Tlp               types.String                                  `tfsdk:"tlp" json:"tlp,computed"`
	Tokens            types.Int64                                   `tfsdk:"tokens" json:"tokens,computed"`
	Updated           timetypes.RFC3339                             `tfsdk:"updated" json:"updated,computed" format:"date-time"`
	Priority          types.Dynamic                                 `tfsdk:"priority" json:"priority,computed"`
	Filter            *CloudforceOneRequestFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CloudforceOneRequestDataSourceModel) toListParams(_ context.Context) (params cloudforce_one.RequestListParams, diags diag.Diagnostics) {
	params = cloudforce_one.RequestListParams{}

	return
}

type CloudforceOneRequestFindOneByDataSourceModel struct {
	AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier,required"`
}
