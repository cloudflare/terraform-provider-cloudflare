// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoTieredCachingResultEnvelope struct {
	Result ArgoTieredCachingModel `json:"result,computed"`
}

type ArgoTieredCachingModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Value  types.String `tfsdk:"value" json:"value"`
}
