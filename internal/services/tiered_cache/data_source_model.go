// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tiered_cache

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TieredCacheResultDataSourceEnvelope struct {
	Result TieredCacheDataSourceModel `json:"result,computed"`
}

type TieredCacheDataSourceModel struct {
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id"`
	Editable   types.Bool        `tfsdk:"editable" json:"editable"`
	ID         types.String      `tfsdk:"id" json:"id"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on"`
	Value      types.String      `tfsdk:"value" json:"value"`
}
