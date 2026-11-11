// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security_custom_topics

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISecurityCustomTopicsResultEnvelope struct {
	Result AISecurityCustomTopicsModel `json:"result"`
}

type AISecurityCustomTopicsModel struct {
	ID     types.String                          `tfsdk:"id" json:"-,computed"`
	ZoneID types.String                          `tfsdk:"zone_id" path:"zone_id,required"`
	Topics *[]*AISecurityCustomTopicsTopicsModel `tfsdk:"topics" json:"topics,optional"`
}

func (m AISecurityCustomTopicsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AISecurityCustomTopicsModel) MarshalJSONForUpdate(state AISecurityCustomTopicsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AISecurityCustomTopicsTopicsModel struct {
	Label types.String `tfsdk:"label" json:"label,required"`
	Topic types.String `tfsdk:"topic" json:"topic,required"`
}
