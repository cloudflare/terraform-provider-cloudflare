// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2_bucket_event_notification

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type R2BucketEventNotificationResultEnvelope struct {
	Result R2BucketEventNotificationModel `json:"result"`
}

type R2BucketEventNotificationModel struct {
	AccountID    types.String                            `tfsdk:"account_id" path:"account_id,required"`
	BucketName   types.String                            `tfsdk:"bucket_name" path:"bucket_name,required"`
	QueueID      types.String                            `tfsdk:"queue_id" path:"queue_id,required"`
	Rules        *[]*R2BucketEventNotificationRulesModel `tfsdk:"rules" json:"rules,optional"`
	QueueName    types.String                            `tfsdk:"queue_name" json:"queueName,computed"`
	Jurisdiction types.String                            `tfsdk:"jurisdiction" json:"-,computed_optional,no_refresh"`
}

func (m R2BucketEventNotificationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m R2BucketEventNotificationModel) MarshalJSONForUpdate(state R2BucketEventNotificationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type R2BucketEventNotificationRulesModel struct {
	Actions     *[]types.String `tfsdk:"actions" json:"actions,required"`
	Description types.String    `tfsdk:"description" json:"description,optional"`
	Prefix      types.String    `tfsdk:"prefix" json:"prefix,optional"`
	Suffix      types.String    `tfsdk:"suffix" json:"suffix,optional"`
}

func (enRules1 R2BucketEventNotificationRulesModel) Equal(enRules2 R2BucketEventNotificationRulesModel) bool {
	if !stringEqualNullOrEmpty(enRules1.Prefix, enRules2.Prefix) {
		return false
	}

	if !stringEqualNullOrEmpty(enRules1.Suffix, enRules2.Suffix) {
		return false
	}

	if !stringEqualNullOrEmpty(enRules1.Description, enRules2.Description) {
		return false
	}

	return actionsEqual(enRules1.Actions, enRules2.Actions)
}

// String comaprison function that treats empty strings as equal to null
func stringEqualNullOrEmpty(s1, s2 types.String) bool {
	if s1.Equal(s2) {
		return true
	}

	if s1.IsNull() && s2.Equal(basetypes.NewStringValue("")) {
		return true
	}
	if s2.IsNull() && s1.Equal(basetypes.NewStringValue("")) {
		return true
	}

	return false
}

// actionsEqual does an order-independent comparison of 2 actions arrays
func actionsEqual(a1, a2 *[]types.String) bool {
	if a1 == a2 {
		return true
	}
	if a1 == nil || a2 == nil {
		return false
	}

	// Check lengths
	if len(*a1) != len(*a2) {
		return false
	}

	counts1 := make(map[string]int)
	for _, action := range *a1 {
		counts1[action.ValueString()]++
	}

	counts2 := make(map[string]int)
	for _, action := range *a2 {
		counts2[action.ValueString()]++
	}

	if len(counts1) != len(counts2) {
		return false
	}

	for action, count := range counts1 {
		if counts2[action] != count {
			return false
		}
	}

	return true
}

// RulesEqual compares two Rules arrays for equality
func RulesEqual(planRules, stateRules *[]*R2BucketEventNotificationRulesModel) bool {
	if planRules == stateRules {
		return true
	}
	if planRules == nil || stateRules == nil {
		return false
	}

	// Check lengths
	if len(*planRules) != len(*stateRules) {
		return false
	}

	// Compare each rule element by element
	for i := range *planRules {
		rule1 := (*planRules)[i]
		rule2 := (*stateRules)[i]

		if !rule1.Equal(*rule2) {
			return false
		}
	}

	return true
}
