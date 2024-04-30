// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel_indicator_feed

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IntelIndicatorFeedResultEnvelope struct {
	Result IntelIndicatorFeedModel `json:"result,computed"`
}

type IntelIndicatorFeedModel struct {
	ID          types.Int64  `tfsdk:"id" json:"id,computed"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	Description types.String `tfsdk:"description" json:"description"`
	Name        types.String `tfsdk:"name" json:"name"`
}
