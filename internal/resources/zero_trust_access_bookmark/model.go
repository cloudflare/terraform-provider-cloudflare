// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_bookmark

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessBookmarkResultEnvelope struct {
	Result ZeroTrustAccessBookmarkModel `json:"result,computed"`
}

type ZeroTrustAccessBookmarkModel struct {
	ID         types.String `tfsdk:"id" json:"id,computed"`
	Identifier types.String `tfsdk:"identifier" path:"identifier"`
	UUID       types.String `tfsdk:"uuid" path:"uuid"`
}
