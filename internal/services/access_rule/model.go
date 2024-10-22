// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRuleResultEnvelope struct {
	Result AccessRuleModel `json:"result"`
}

type AccessRuleModel struct {
	ID            types.String                                   `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID        types.String                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	Mode          types.String                                   `tfsdk:"mode" json:"mode,required"`
	Configuration *AccessRuleConfigurationModel                  `tfsdk:"configuration" json:"configuration,required"`
	Notes         types.String                                   `tfsdk:"notes" json:"notes,optional"`
	CreatedOn     timetypes.RFC3339                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn    timetypes.RFC3339                              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	AllowedModes  customfield.List[types.String]                 `tfsdk:"allowed_modes" json:"allowed_modes,computed"`
	Scope         customfield.NestedObject[AccessRuleScopeModel] `tfsdk:"scope" json:"scope,computed"`
}

func (m AccessRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessRuleModel) MarshalJSONForUpdate(state AccessRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccessRuleConfigurationModel struct {
	Target types.String `tfsdk:"target" json:"target,optional"`
	Value  types.String `tfsdk:"value" json:"value,optional"`
}

type AccessRuleScopeModel struct {
	ID    types.String `tfsdk:"id" json:"id,computed"`
	Email types.String `tfsdk:"email" json:"email,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
}
