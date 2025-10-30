package utils

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ConvertTerraformTagsToAPI converts Terraform tags (map[string]types.String)
// to API tags (map[string]string) for use in API requests.
// Returns an empty map if input is nil.
func ConvertTerraformTagsToAPI(tags *map[string]types.String) map[string]string {
	result := make(map[string]string)
	if tags == nil {
		return result
	}

	for k, v := range *tags {
		result[k] = v.ValueString()
	}
	return result
}

// ConvertAPITagsToTerraform converts API tags (map[string]string)
// to Terraform tags (map[string]types.String) for storing in state.
// Returns nil if input map is empty.
func ConvertAPITagsToTerraform(tags map[string]string) *map[string]types.String {
	if len(tags) == 0 {
		return nil
	}

	result := make(map[string]types.String, len(tags))
	for k, v := range tags {
		result[k] = types.StringValue(v)
	}
	return &result
}

// TagsChanged checks if Terraform tags have changed between planned and state.
// Returns true if tags are different, false if they're the same.
func TagsChanged(planned, state *map[string]types.String) bool {
	// Check for nil/existence changes
	if (planned == nil && state != nil) ||
		(planned != nil && state == nil) {
		return true
	}

	// Both nil = no change
	if planned == nil && state == nil {
		return false
	}

	// Check length difference
	if len(*planned) != len(*state) {
		return true
	}

	// Check each key-value pair
	for k, v := range *planned {
		stateV, exists := (*state)[k]
		if !exists || stateV.ValueString() != v.ValueString() {
			return true
		}
	}

	return false
}
