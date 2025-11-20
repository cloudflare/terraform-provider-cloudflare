package r2_bucket_event_notification

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

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
