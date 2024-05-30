// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type D1DatabaseResultEnvelope struct {
	Result D1DatabaseModel `json:"result,computed"`
}

type D1DatabaseModel struct {
	UUID      types.String `tfsdk:"uuid" json:"uuid"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Name      types.String `tfsdk:"name" json:"name"`
}
