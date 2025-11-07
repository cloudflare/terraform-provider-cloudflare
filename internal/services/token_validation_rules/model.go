// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_rules

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TokenValidationRulesResultEnvelope struct {
	Result TokenValidationRulesModel `json:"result"`
}

type TokenValidationRulesModel struct {
	ID          types.String                       `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String                       `tfsdk:"zone_id" path:"zone_id,required"`
	Action      types.String                       `tfsdk:"action" json:"action,required"`
	Description types.String                       `tfsdk:"description" json:"description,required"`
	Enabled     types.Bool                         `tfsdk:"enabled" json:"enabled,required"`
	Expression  types.String                       `tfsdk:"expression" json:"expression,required"`
	Title       types.String                       `tfsdk:"title" json:"title,required"`
	Selector    *TokenValidationRulesSelectorModel `tfsdk:"selector" json:"selector,required"`
	Position    *TokenValidationRulesPositionModel `tfsdk:"position" json:"position,optional,no_refresh"`
	CreatedAt   timetypes.RFC3339                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastUpdated timetypes.RFC3339                  `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
}

func (m TokenValidationRulesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TokenValidationRulesModel) MarshalJSONForUpdate(state TokenValidationRulesModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type TokenValidationRulesSelectorModel struct {
	Exclude *[]*TokenValidationRulesSelectorExcludeModel `tfsdk:"exclude" json:"exclude,optional"`
	Include *[]*TokenValidationRulesSelectorIncludeModel `tfsdk:"include" json:"include,optional"`
}

type TokenValidationRulesSelectorExcludeModel struct {
	OperationIDs *[]types.String `tfsdk:"operation_ids" json:"operation_ids,optional"`
}

type TokenValidationRulesSelectorIncludeModel struct {
	Host *[]types.String `tfsdk:"host" json:"host,optional"`
}

type TokenValidationRulesPositionModel struct {
	Index  types.Int64  `tfsdk:"index" json:"index,optional"`
	Before types.String `tfsdk:"before" json:"before,optional"`
	After  types.String `tfsdk:"after" json:"after,optional"`
}
