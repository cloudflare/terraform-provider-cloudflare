// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_tag

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessTagResultEnvelope struct {
	Result AccessTagModel `json:"result,computed"`
}

type AccessTagModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Name      types.String `tfsdk:"name" json:"name"`
}
