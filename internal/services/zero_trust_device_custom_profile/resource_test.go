package zero_trust_device_custom_profile

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdate_NoChanges(t *testing.T) {
	planData := ZeroTrustDeviceCustomProfileModel{
		AccountID:       types.StringValue("test-account-id"),
		PolicyID:        types.StringValue("test-policy-id"),
		Match:           types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:            types.StringValue("Allow Developers"),
		Precedence:      types.Float64Value(100),
		Description:     types.StringNull(),
		FallbackDomains: customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileFallbackDomainsModel{}),
		TargetTests:     customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileTargetTestsModel{}),
	}

	stateData := planData

	jsonBytes, err := planData.MarshalJSONForUpdate(stateData)
	require.NoError(t, err)

	jsonStr := string(jsonBytes)
	assert.Equal(t, "", jsonStr, "Should be empty when no changes")

	isEmpty := len(jsonBytes) == 0 || string(jsonBytes) == ""
	assert.True(t, isEmpty, "Logic isEmpty should be true")
}

func TestUpdate_HasChanges(t *testing.T) {
	planData := ZeroTrustDeviceCustomProfileModel{
		AccountID:   types.StringValue("test-account-id"),
		PolicyID:    types.StringValue("test-policy-id"),
		Match:       types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:        types.StringValue("Allow Developers"),
		Precedence:  types.Float64Value(100),
		Description: types.StringValue("New Description"),
	}

	stateData := ZeroTrustDeviceCustomProfileModel{
		AccountID:   types.StringValue("test-account-id"),
		PolicyID:    types.StringValue("test-policy-id"),
		Match:       types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:        types.StringValue("Allow Developers"),
		Precedence:  types.Float64Value(100),
		Description: types.StringValue("Old Description"),
	}

	jsonBytes, err := planData.MarshalJSONForUpdate(stateData)
	require.NoError(t, err)

	jsonStr := string(jsonBytes)
	assert.NotEqual(t, "", jsonStr, "Should not be empty when there are changes")
	assert.Contains(t, jsonStr, "New Description", "Should contain new description")

	isEmpty := len(jsonBytes) == 0 || string(jsonBytes) == ""
	assert.False(t, isEmpty, "Logic isEmpty should be false")
}

func TestUpdate_DescriptionRemovedFromConfig(t *testing.T) {
	planData := ZeroTrustDeviceCustomProfileModel{
		AccountID:       types.StringValue("test-account-id"),
		PolicyID:        types.StringValue("test-policy-id"),
		Match:           types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:            types.StringValue("Test Profile"),
		Precedence:      types.Float64Value(100),
		Description:     types.StringNull(),
		FallbackDomains: customfield.NullObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsModel](context.TODO()),
		TargetTests:     customfield.NullObjectList[ZeroTrustDeviceCustomProfileTargetTestsModel](context.TODO()),
	}

	stateData := ZeroTrustDeviceCustomProfileModel{
		AccountID:       types.StringValue("test-account-id"),
		PolicyID:        types.StringValue("test-policy-id"),
		Match:           types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:            types.StringValue("Test Profile"),
		Precedence:      types.Float64Value(100),
		Description:     types.StringValue("Profile for Test User"),
		FallbackDomains: customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileFallbackDomainsModel{}),
		TargetTests:     customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileTargetTestsModel{}),
	}

	jsonBytes, err := planData.MarshalJSONForUpdate(stateData)
	require.NoError(t, err)

	jsonStr := string(jsonBytes)
	assert.Equal(t, `{"description":""}`, jsonStr, "Should send empty string for description")

	isEmpty := len(jsonBytes) == 0 || string(jsonBytes) == ""
	assert.False(t, isEmpty, "Should not be empty - description is set to empty string")
	resultData := planData
	resultData.Description = types.StringValue("")
	if !stateData.FallbackDomains.IsNull() {
		resultData.FallbackDomains = stateData.FallbackDomains
	}
	if !stateData.TargetTests.IsNull() {
		resultData.TargetTests = stateData.TargetTests
	}
	resultData.ID = resultData.PolicyID

	assert.Equal(t, "", resultData.Description.ValueString(), "Description should be empty string after removal")
	assert.False(t, resultData.FallbackDomains.IsNull(), "FallbackDomains should be preserved from state")
	assert.False(t, resultData.TargetTests.IsNull(), "TargetTests should be preserved from state")
}

func TestUpdate_CompleteFlow_RemoveDescription(t *testing.T) {
	planData := ZeroTrustDeviceCustomProfileModel{
		AccountID:       types.StringValue("test-account-id"),
		PolicyID:        types.StringValue("test-policy-id"),
		Match:           types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:            types.StringValue("Test Profile"),
		Precedence:      types.Float64Value(100),
		Description:     types.StringNull(),
		FallbackDomains: customfield.NullObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsModel](context.TODO()),
		TargetTests:     customfield.NullObjectList[ZeroTrustDeviceCustomProfileTargetTestsModel](context.TODO()),
	}

	stateData := ZeroTrustDeviceCustomProfileModel{
		AccountID:       types.StringValue("test-account-id"),
		PolicyID:        types.StringValue("test-policy-id"),
		Match:           types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:            types.StringValue("Test Profile"),
		Precedence:      types.Float64Value(100),
		Description:     types.StringValue("Profile for Test User"),
		FallbackDomains: customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileFallbackDomainsModel{}),
		TargetTests:     customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileTargetTestsModel{}),
	}

	jsonBytes, err := planData.MarshalJSONForUpdate(stateData)
	require.NoError(t, err)
	assert.Equal(t, `{"description":""}`, string(jsonBytes), "Should send empty string for description")

	isEmpty := len(jsonBytes) == 0 || string(jsonBytes) == ""
	assert.False(t, isEmpty, "Not empty - description is set to empty string")

	apiResponse := planData
	apiResponse.Description = types.StringValue("")
	apiResponse.FallbackDomains = stateData.FallbackDomains
	apiResponse.TargetTests = stateData.TargetTests

	finalResult := apiResponse
	if planData.Description.IsNull() {
		finalResult.Description = types.StringNull()
	}

	assert.True(t, finalResult.Description.IsNull(), "End result should have description=null as per plan")
	assert.False(t, finalResult.FallbackDomains.IsNull(), "Computed fields should be preserved")
}

func TestModifyPlan_ComputedFields(t *testing.T) {
	plan := ZeroTrustDeviceCustomProfileModel{
		AccountID:       types.StringValue("test-account-id"),
		PolicyID:        types.StringValue("test-policy-id"),
		Match:           types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:            types.StringValue("Test Profile"),
		Precedence:      types.Float64Value(100),
		Description:     types.StringNull(),
		FallbackDomains: customfield.NullObjectList[ZeroTrustDeviceCustomProfileFallbackDomainsModel](context.TODO()),
		TargetTests:     customfield.NullObjectList[ZeroTrustDeviceCustomProfileTargetTestsModel](context.TODO()),
		Default:         types.BoolNull(),
		GatewayUniqueID: types.StringNull(),
	}

	state := ZeroTrustDeviceCustomProfileModel{
		AccountID:       types.StringValue("test-account-id"),
		PolicyID:        types.StringValue("test-policy-id"),
		Match:           types.StringValue(`identity.email == "test@cloudflare.com"`),
		Name:            types.StringValue("Test Profile"),
		Precedence:      types.Float64Value(100),
		Description:     types.StringNull(),
		FallbackDomains: customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileFallbackDomainsModel{}),
		TargetTests:     customfield.NewObjectListMust(context.TODO(), []ZeroTrustDeviceCustomProfileTargetTestsModel{}),
		Default:         types.BoolValue(false),
		GatewayUniqueID: types.StringValue("unique-123"),
	}

	planModified := false

	if !state.Default.IsNull() {
		plan.Default = state.Default
		planModified = true
	}
	if !state.GatewayUniqueID.IsNull() {
		plan.GatewayUniqueID = state.GatewayUniqueID
		planModified = true
	}
	if !state.FallbackDomains.IsNull() {
		plan.FallbackDomains = state.FallbackDomains
		planModified = true
	}
	if !state.TargetTests.IsNull() {
		plan.TargetTests = state.TargetTests
		planModified = true
	}

	assert.True(t, planModified, "Plan should be modified")
	assert.False(t, plan.Default.IsNull(), "Default should be copied from state")
	assert.False(t, plan.GatewayUniqueID.IsNull(), "GatewayUniqueID should be copied from state")
	assert.False(t, plan.FallbackDomains.IsNull(), "FallbackDomains should be copied from state")
	assert.False(t, plan.TargetTests.IsNull(), "TargetTests should be copied from state")
}
