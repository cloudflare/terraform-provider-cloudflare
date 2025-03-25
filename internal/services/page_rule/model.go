// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PageRuleResultEnvelope struct {
	Result PageRuleModel `json:"result"`
}

type PageRuleModel struct {
	ID         types.String      `tfsdk:"id" json:"id,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Priority   types.Int64       `tfsdk:"priority" json:"priority,computed_optional"`
	Status     types.String      `tfsdk:"status" json:"status,computed_optional"`
	CreatedOn  timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`

	Target  types.String          `tfsdk:"target" json:"target,required"`
	Actions *PageRuleActionsModel `tfsdk:"actions" json:"actions,required"`
}

func (m PageRuleModel) MarshalJSON() (data []byte, err error) {
	return m.marshalCustom()
}

func (m PageRuleModel) MarshalJSONForUpdate(state PageRuleModel) (data []byte, err error) {
	return m.marshalCustomForUpdate(state)
}
