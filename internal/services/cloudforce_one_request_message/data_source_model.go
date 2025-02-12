// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4/cloudforce_one"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestMessageResultDataSourceEnvelope struct {
	Result CloudforceOneRequestMessageDataSourceModel `json:"result,computed"`
}

type CloudforceOneRequestMessageDataSourceModel struct {
	AccountIdentifier types.String      `tfsdk:"account_identifier" path:"account_identifier,required"`
	RequestIdentifier types.String      `tfsdk:"request_identifier" path:"request_identifier,required"`
	Author            types.String      `tfsdk:"author" json:"author,computed"`
	Content           types.String      `tfsdk:"content" json:"content,computed"`
	Created           timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	ID                types.Int64       `tfsdk:"id" json:"id,computed"`
	IsFollowOnRequest types.Bool        `tfsdk:"is_follow_on_request" json:"is_follow_on_request,computed"`
	Updated           timetypes.RFC3339 `tfsdk:"updated" json:"updated,computed" format:"date-time"`
}

func (m *CloudforceOneRequestMessageDataSourceModel) toReadParams(_ context.Context) (params cloudforce_one.RequestMessageGetParams, diags diag.Diagnostics) {
	params = cloudforce_one.RequestMessageGetParams{}

	return
}
