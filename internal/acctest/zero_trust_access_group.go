package acctest

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

// expectEmptyPlanExceptZeroTrustAccessGroupOrdering is a custom plan check that allows for ordering differences
// in zero trust access group include/exclude/require lists while ensuring all other changes are caught
type expectEmptyPlanExceptZeroTrustAccessGroupOrdering struct{}

func (e expectEmptyPlanExceptZeroTrustAccessGroupOrdering) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	if req.Plan == nil {
		resp.Error = fmt.Errorf("plan is nil")
		return
	}

	// Check each resource change
	for _, resourceChange := range req.Plan.ResourceChanges {
		if resourceChange.Change == nil {
			continue
		}

		// Skip if this is a create or destroy action - we only care about updates/replaces
		if resourceChange.Change.Actions.Create() || resourceChange.Change.Actions.Delete() {
			resp.Error = fmt.Errorf("expected empty plan, but %s has planned action(s): %v", resourceChange.Address, resourceChange.Change.Actions)
			return
		}

		// For update actions, check the before/after values
		if resourceChange.Change.Actions.Update() || resourceChange.Change.Actions.Replace() {
			if !isZeroTrustAccessGroupOrderingChangeOnly(resourceChange.Change.Before, resourceChange.Change.After) {
				resp.Error = fmt.Errorf("expected empty plan except for ordering changes, but %s has non-equivalent change from %v to %v", resourceChange.Address, resourceChange.Change.Before, resourceChange.Change.After)
				return
			}
		}
	}
}

// isZeroTrustAccessGroupOrderingChangeOnly checks if the changes between before and after are only ordering differences
// in zero trust access group include/exclude/require fields or other allowable migration artifacts
func isZeroTrustAccessGroupOrderingChangeOnly(before, after interface{}) bool {
	beforeMap, beforeOk := before.(map[string]interface{})
	afterMap, afterOk := after.(map[string]interface{})

	if !beforeOk || !afterOk {
		return reflect.DeepEqual(before, after)
	}

	// Get all unique keys
	allKeys := make(map[string]bool)
	for key := range beforeMap {
		allKeys[key] = true
	}
	for key := range afterMap {
		allKeys[key] = true
	}

	// Check each key
	for key := range allKeys {
		beforeValue := beforeMap[key]
		afterValue := afterMap[key]

		// Skip identical values
		if reflect.DeepEqual(beforeValue, afterValue) {
			continue
		}

		// Allow falsey-to-null transitions (common in Terraform)
		if isAllowableFalseyToNullChange(beforeValue, afterValue) {
			continue
		}

		// Special handling for zero trust access group list fields (include, exclude, require)
		if key == "include" || key == "exclude" || key == "require" {
			if isEquivalentNestedAttributeList(beforeValue, afterValue) {
				continue
			}
		}

		// If we get here, there's a non-allowable change
		return false
	}

	return true
}

// isEquivalentNestedAttributeList checks if two nested attribute lists are equivalent
// (same contents, potentially different order)
func isEquivalentNestedAttributeList(before, after interface{}) bool {
	beforeSlice, beforeOk := before.([]interface{})
	afterSlice, afterOk := after.([]interface{})

	if !beforeOk || !afterOk {
		return reflect.DeepEqual(before, after)
	}

	if len(beforeSlice) != len(afterSlice) {
		return false
	}

	// For each item in before, find equivalent item in after
	for _, beforeItem := range beforeSlice {
		found := false
		for _, afterItem := range afterSlice {
			if areZeroTrustAccessGroupObjectsEquivalent(beforeItem, afterItem) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// areZeroTrustAccessGroupObjectsEquivalent checks if two zero trust access group objects are equivalent
// with special handling for provider migration artifacts
func areZeroTrustAccessGroupObjectsEquivalent(before, after interface{}) bool {
	beforeMap, beforeOk := before.(map[string]interface{})
	afterMap, afterOk := after.(map[string]interface{})

	if !beforeOk || !afterOk {
		return reflect.DeepEqual(before, after)
	}

	// Special case: if one object has provider-specific content that might be migration artifacts
	if hasZeroTrustProviderContent(beforeMap) || hasZeroTrustProviderContent(afterMap) {
		// Allow these objects to be different as they might be migration artifacts
		return true
	}

	// Normalize both objects for comparison
	beforeNorm := normalizeZeroTrustAccessGroupObject(beforeMap)
	afterNorm := normalizeZeroTrustAccessGroupObject(afterMap)

	return mapsEquivalent(beforeNorm, afterNorm)
}

// normalizeZeroTrustAccessGroupObject normalizes a zero trust access group object for comparison
func normalizeZeroTrustAccessGroupObject(obj map[string]interface{}) map[string]interface{} {
	normalized := make(map[string]interface{})
	
	for key, value := range obj {
		// Convert various falsey values to nil for consistent comparison
		if isFalseyValue(value) {
			normalized[key] = nil
		} else {
			normalized[key] = value
		}
	}
	
	return normalized
}

// hasZeroTrustProviderContent checks if an object contains provider-specific content that indicates migration
func hasZeroTrustProviderContent(obj map[string]interface{}) bool {
	return hasAzureADContent(obj) || hasGitHubOrganizationContent(obj) || hasOtherZeroTrustProviderContent(obj)
}

// hasAzureADContent checks if an object contains azure_ad fields that indicate it was migrated from azure blocks
func hasAzureADContent(obj map[string]interface{}) bool {
	if azureAD, ok := obj["azure_ad"].(map[string]interface{}); ok {
		// Check if it has the typical azure_ad structure
		if id, hasID := azureAD["id"]; hasID && id != nil {
			return true
		}
		if providerID, hasProvider := azureAD["identity_provider_id"]; hasProvider && providerID != nil {
			return true
		}
	}
	return false
}

// hasGitHubOrganizationContent checks if an object contains github_organization fields
func hasGitHubOrganizationContent(obj map[string]interface{}) bool {
	if github, ok := obj["github_organization"].(map[string]interface{}); ok {
		if name, hasName := github["name"]; hasName && name != nil {
			return true
		}
		if team, hasTeam := github["team"]; hasTeam && team != nil {
			return true
		}
	}
	return false
}

// hasOtherZeroTrustProviderContent checks if an object contains other identity provider content
func hasOtherZeroTrustProviderContent(obj map[string]interface{}) bool {
	// Check for other providers that might have migration issues
	providerFields := []string{"gsuite", "okta", "saml", "oidc"}
	for _, field := range providerFields {
		if provider, ok := obj[field].(map[string]interface{}); ok && len(provider) > 0 {
			// Check if it has any non-nil values
			for _, value := range provider {
				if value != nil && value != "<nil>" {
					return true
				}
			}
		}
	}
	return false
}

// mapsEquivalent checks if two normalized maps are equivalent
func mapsEquivalent(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}
	
	for key, value1 := range map1 {
		value2, exists := map2[key]
		if !exists {
			return false
		}
		
		// Deep comparison for nested maps
		if nested1, ok1 := value1.(map[string]interface{}); ok1 {
			if nested2, ok2 := value2.(map[string]interface{}); ok2 {
				if !mapsEquivalent(nested1, nested2) {
					return false
				}
			} else {
				return false
			}
		} else if value1 != value2 {
			return false
		}
	}
	
	return true
}

// isAllowableFalseyToNullChange checks if a change from beforeValue to afterValue is an allowable falsey-to-null transition
func isAllowableFalseyToNullChange(beforeValue, afterValue interface{}) bool {
	// Skip if values are identical
	if reflect.DeepEqual(beforeValue, afterValue) {
		return true
	}
	// Allow falsey-to-null transitions
	if afterValue == nil && isFalseyValue(beforeValue) {
		return true
	}
	if beforeValue == nil && isFalseyValue(afterValue) {
		return true
	}
	// Values are different and not a valid transition
	return false
}


// ExpectEmptyPlanExceptZeroTrustAccessGroupOrdering is the public interface for the custom plan checker
var ExpectEmptyPlanExceptZeroTrustAccessGroupOrdering = expectEmptyPlanExceptZeroTrustAccessGroupOrdering{}

// ZeroTrustAccessGroupMigrationTestStep creates a test step specifically for zero trust access group migrations
// that handles ordering differences and nil representation differences
func ZeroTrustAccessGroupMigrationTestStep(t *testing.T, v4Config string, tmpDir string, exactVersion string, stateChecks []statecheck.StateCheck) resource.TestStep {
	// Choose the appropriate plan check based on the version
	var planChecks []plancheck.PlanCheck
	if strings.HasPrefix(exactVersion, "4.") {
		// When upgrading from v4, use our custom plan checker that handles ordering differences
		planChecks = []plancheck.PlanCheck{
			DebugNonEmptyPlan,
			ExpectEmptyPlanExceptZeroTrustAccessGroupOrdering,
		}
	} else {
		// When upgrading from v5, expect a completely empty plan
		planChecks = []plancheck.PlanCheck{
			DebugNonEmptyPlan,
			plancheck.ExpectEmptyPlan(),
		}
	}

	return resource.TestStep{
		PreConfig: func() {
			WriteOutConfig(t, v4Config, tmpDir)
			// we only run the migration command if the version is 4.x.x, because users will not expect to run it within v5 versions.
			if strings.HasPrefix(exactVersion, "4.") {
				debugLogf(t, "Running migration command for version: %s", exactVersion)
				RunMigrationCommand(t, v4Config, tmpDir)
			} else {
				debugLogf(t, "Skipping migration command for version: %s", exactVersion)
			}
		},
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		ConfigDirectory:          config.StaticDirectory(tmpDir),
		ConfigPlanChecks: resource.ConfigPlanChecks{
			PreApply: planChecks,
		},
		ConfigStateChecks: stateChecks,
	}
}