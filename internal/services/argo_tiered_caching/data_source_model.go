// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoTieredCachingResultDataSourceEnvelope struct {
	Result ArgoTieredCachingDataSourceModel `json:"result,computed"`
}

type ArgoTieredCachingDataSourceModel struct {
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id"`
	ID         types.String      `tfsdk:"id" json:"id"`
	Editable   types.Bool        `tfsdk:"editable" json:"editable"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on"`
	Value      types.String      `tfsdk:"value" json:"value"`
}
