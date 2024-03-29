// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PagerulesResultEnvelope struct {
	Result PagerulesModel `json:"result,computed"`
}

type PagerulesModel struct {
	ZoneID     types.String             `tfsdk:"zone_id" path:"zone_id"`
	PageruleID types.String             `tfsdk:"pagerule_id" path:"pagerule_id"`
	Actions    []*PagerulesActionsModel `tfsdk:"actions" json:"actions"`
	Targets    []*PagerulesTargetsModel `tfsdk:"targets" json:"targets"`
	Priority   types.Int64              `tfsdk:"priority" json:"priority"`
	Status     types.String             `tfsdk:"status" json:"status"`
}

type PagerulesActionsModel struct {
	ModifiedOn types.String                `tfsdk:"modified_on" json:"modified_on,computed"`
	Name       types.String                `tfsdk:"name" json:"name"`
	Value      *PagerulesActionsValueModel `tfsdk:"value" json:"value"`
}

type PagerulesActionsValueModel struct {
	Type types.String `tfsdk:"type" json:"type"`
	URL  types.String `tfsdk:"url" json:"url"`
}

type PagerulesTargetsModel struct {
	Constraint *PagerulesTargetsConstraintModel `tfsdk:"constraint" json:"constraint"`
	Target     types.String                     `tfsdk:"target" json:"target"`
}

type PagerulesTargetsConstraintModel struct {
	Operator types.String `tfsdk:"operator" json:"operator"`
	Value    types.String `tfsdk:"value" json:"value"`
}
