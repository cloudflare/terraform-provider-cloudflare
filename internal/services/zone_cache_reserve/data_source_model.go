// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneCacheReserveResultDataSourceEnvelope struct {
	Result ZoneCacheReserveDataSourceModel `json:"result,computed"`
}

type ZoneCacheReserveDataSourceModel struct {
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id"`
	ID         types.String      `tfsdk:"id" json:"id"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on"`
	Value      types.String      `tfsdk:"value" json:"value"`
}
