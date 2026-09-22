// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_tracing

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneTracingResultEnvelope struct {
	Result ZoneTracingModel `json:"result"`
}

type ZoneTracingModel struct {
	ID                types.String    `tfsdk:"id" json:"-,computed"`
	ZoneID            types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled           types.Bool      `tfsdk:"enabled" json:"enabled,optional"`
	ForwardContext    types.Bool      `tfsdk:"forward_context" json:"forward_context,optional"`
	Persist           types.Bool      `tfsdk:"persist" json:"persist,optional"`
	PropagationPolicy types.String    `tfsdk:"propagation_policy" json:"propagation_policy,optional"`
	SamplingRatio     types.Float64   `tfsdk:"sampling_ratio" json:"sampling_ratio,optional"`
	Destinations      *[]types.String `tfsdk:"destinations" json:"destinations,optional"`
}

func (m ZoneTracingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneTracingModel) MarshalJSONForUpdate(state ZoneTracingModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
