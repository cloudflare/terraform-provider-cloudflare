// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type R2BucketResultEnvelope struct {
	Result R2BucketModel `json:"result,computed"`
}

type R2BucketModel struct {
	Name         types.String `tfsdk:"name" json:"name,computed"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	LocationHint types.String `tfsdk:"locationhint" json:"locationHint"`
}
