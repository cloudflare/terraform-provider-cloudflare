// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_security_custom_topics

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/ai_security"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AISecurityCustomTopicsResultDataSourceEnvelope struct {
	Result AISecurityCustomTopicsDataSourceModel `json:"result,computed"`
}

type AISecurityCustomTopicsDataSourceModel struct {
	ID     types.String                                                              `tfsdk:"id" path:"zone_id,computed"`
	ZoneID types.String                                                              `tfsdk:"zone_id" path:"zone_id,required"`
	Topics customfield.NestedObjectList[AISecurityCustomTopicsTopicsDataSourceModel] `tfsdk:"topics" json:"topics,computed"`
}

func (m *AISecurityCustomTopicsDataSourceModel) toReadParams(_ context.Context) (params ai_security.CustomTopicGetParams, diags diag.Diagnostics) {
	params = ai_security.CustomTopicGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type AISecurityCustomTopicsTopicsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label,computed"`
	Topic types.String `tfsdk:"topic" json:"topic,computed"`
}
