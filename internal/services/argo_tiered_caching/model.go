// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoTieredCachingResultEnvelope struct {
	Result ArgoTieredCachingModel `json:"result"`
}

type ArgoTieredCachingModel struct {
	ID         types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id"`
	Value      types.String      `tfsdk:"value" json:"value"`
	Editable   types.Bool        `tfsdk:"editable" json:"editable,computed"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}
