// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4/cloudforce_one"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CloudforceOneRequestMessageResultDataSourceEnvelope struct {
	Result CloudforceOneRequestMessageDataSourceModel `json:"result,computed"`
}

type CloudforceOneRequestMessageDataSourceModel struct {
	AccountIdentifier types.String `tfsdk:"account_identifier" path:"account_identifier,required"`
	RequestIdentifier types.String `tfsdk:"request_identifier" path:"request_identifier,required"`
}

func (m *CloudforceOneRequestMessageDataSourceModel) toReadParams(_ context.Context) (params cloudforce_one.RequestMessageGetParams, diags diag.Diagnostics) {
	params = cloudforce_one.RequestMessageGetParams{}

	return
}
