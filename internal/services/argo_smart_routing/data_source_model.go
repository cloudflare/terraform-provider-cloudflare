// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_smart_routing

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoSmartRoutingResultDataSourceEnvelope struct {
	Result ArgoSmartRoutingDataSourceModel `json:"result,computed"`
}

type ArgoSmartRoutingDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
