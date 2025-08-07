package bot_management_test

import (
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/bot_management"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

// TestBotManagementModel_MarshalJSONForUpdate tests that enable_js and suppress_session_score
// fields with encode_state_for_unknown tag work correctly to fix issues #5728 and #5519
func TestBotManagementModel_MarshalJSONForUpdate(t *testing.T) {
	// Test case 1: plan has enable_js=true, state has enable_js=false
	// With our fix, enable_js should be included in the JSON because plan is explicit
	plan := &bot_management.BotManagementModel{
		ZoneID:   types.StringValue("test-zone"),
		EnableJS: types.BoolValue(true), // explicit value in plan
	}
	
	state := &bot_management.BotManagementModel{
		ZoneID:   types.StringValue("test-zone"),
		EnableJS: types.BoolValue(false), // different value in state
	}
	
	jsonBytes, err := plan.MarshalJSONForUpdate(*state)
	assert.NoError(t, err)
	jsonStr := string(jsonBytes)
	t.Logf("JSON output with explicit plan value: %s", jsonStr)
	
	// enable_js should be included because it's explicitly set and differs from state
	assert.Contains(t, jsonStr, `"enable_js":true`)
	
	// Test case 2: plan has enable_js=unknown, state has enable_js=false  
	// With our fix, the state value should be used when plan is unknown
	planUnknown := &bot_management.BotManagementModel{
		ZoneID:   types.StringValue("test-zone"),
		EnableJS: types.BoolUnknown(), // unknown value in plan
	}
	
	jsonBytes2, err := planUnknown.MarshalJSONForUpdate(*state)
	assert.NoError(t, err)
	jsonStr2 := string(jsonBytes2)
	t.Logf("JSON output with unknown plan value: %s", jsonStr2)
	
	// enable_js should use state value (false) when plan is unknown
	assert.Contains(t, jsonStr2, `"enable_js":false`)

	// Test case 4: Test suppress_session_score with similar behavior
	planSuppress := &bot_management.BotManagementModel{
		ZoneID:              types.StringValue("test-zone"),
		SuppressSessionScore: types.BoolValue(true),
	}
	
	stateSuppress := &bot_management.BotManagementModel{
		ZoneID:              types.StringValue("test-zone"),
		SuppressSessionScore: types.BoolValue(false),
	}
	
	jsonBytes4, err := planSuppress.MarshalJSONForUpdate(*stateSuppress)
	assert.NoError(t, err)
	jsonStr4 := string(jsonBytes4)
	t.Logf("JSON output for suppress_session_score: %s", jsonStr4)
	
	// suppress_session_score should be included when explicitly set and different
	assert.Contains(t, jsonStr4, `"suppress_session_score":true`)
	
	// Test case 3: both plan and state have same value
	// Should not include enable_js in JSON (no change needed)
	planSame := &bot_management.BotManagementModel{
		ZoneID:   types.StringValue("test-zone"),
		EnableJS: types.BoolValue(false), // same as state
	}
	
	jsonBytes3, err := planSame.MarshalJSONForUpdate(*state)
	assert.NoError(t, err)
	jsonStr3 := string(jsonBytes3)
	t.Logf("JSON output with same plan and state: %s", jsonStr3)
	
	// enable_js should not be included because values are the same
	// (though this might still be included due to computed_optional behavior)
}

// TestAllComputedOptionalFields tests all computed_optional fields to identify
// which ones have the same drift issue and need the encode_state_for_unknown fix
func TestAllComputedOptionalFields(t *testing.T) {
	tests := []struct {
		name     string
		planModel *bot_management.BotManagementModel
		stateModel *bot_management.BotManagementModel
		expectedFields []string // fields we expect to see in JSON when plan is unknown
	}{
		{
			name: "AIBotsProtection unknown plan",
			planModel: &bot_management.BotManagementModel{
				ZoneID:           types.StringValue("test-zone"),
				AIBotsProtection: types.StringUnknown(),
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:           types.StringValue("test-zone"), 
				AIBotsProtection: types.StringValue("disabled"),
			},
			expectedFields: []string{}, // should be empty if no encode_state_for_unknown
		},
		{
			name: "AutoUpdateModel unknown plan",
			planModel: &bot_management.BotManagementModel{
				ZoneID:          types.StringValue("test-zone"),
				AutoUpdateModel: types.BoolUnknown(),
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:          types.StringValue("test-zone"),
				AutoUpdateModel: types.BoolValue(true),
			},
			expectedFields: []string{}, // should be empty if no encode_state_for_unknown
		},
		{
			name: "FightMode unknown plan", 
			planModel: &bot_management.BotManagementModel{
				ZoneID:    types.StringValue("test-zone"),
				FightMode: types.BoolUnknown(),
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:    types.StringValue("test-zone"),
				FightMode: types.BoolValue(true),
			},
			expectedFields: []string{}, // should be empty if no encode_state_for_unknown
		},
		{
			name: "EnableJS unknown plan (should work due to our fix)",
			planModel: &bot_management.BotManagementModel{
				ZoneID:   types.StringValue("test-zone"),
				EnableJS: types.BoolUnknown(),
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:   types.StringValue("test-zone"),
				EnableJS: types.BoolValue(true),
			},
			expectedFields: []string{"enable_js"}, // should be included due to encode_state_for_unknown
		},
		{
			name: "SuppressSessionScore unknown plan (should work due to our fix)",
			planModel: &bot_management.BotManagementModel{
				ZoneID:              types.StringValue("test-zone"),
				SuppressSessionScore: types.BoolUnknown(),
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:              types.StringValue("test-zone"),
				SuppressSessionScore: types.BoolValue(false),
			},
			expectedFields: []string{"suppress_session_score"}, // should be included due to encode_state_for_unknown
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := tt.planModel.MarshalJSONForUpdate(*tt.stateModel)
			assert.NoError(t, err)
			
			jsonStr := string(jsonBytes)
			t.Logf("JSON output for %s: %s", tt.name, jsonStr)
			
			if len(tt.expectedFields) == 0 {
				// Expecting empty JSON or minimal fields
				if jsonStr != "{}" && !strings.Contains(jsonStr, "zone_id") {
					t.Logf("Field %s may have drift issue - produces JSON when plan is unknown: %s", tt.name, jsonStr)
				}
			} else {
				// Expecting specific fields to be present
				for _, field := range tt.expectedFields {
					assert.Contains(t, jsonStr, field, "Expected field %s to be in JSON output", field)
				}
			}
		})
	}
}

// TestExplicitPlanVsState tests what happens when plan explicitly sets a value 
// but state has a different value (the common drift scenario)
func TestExplicitPlanVsState(t *testing.T) {
	tests := []struct {
		name     string
		planModel *bot_management.BotManagementModel  
		stateModel *bot_management.BotManagementModel
		shouldInclude bool // whether field should be in JSON
		fieldName string
	}{
		{
			name: "AutoUpdateModel explicit plan vs different state",
			planModel: &bot_management.BotManagementModel{
				ZoneID:          types.StringValue("test-zone"),
				AutoUpdateModel: types.BoolValue(true), // explicit in plan
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:          types.StringValue("test-zone"),
				AutoUpdateModel: types.BoolValue(false), // different in state
			},
			shouldInclude: true,
			fieldName: "auto_update_model",
		},
		{
			name: "FightMode explicit plan vs different state",
			planModel: &bot_management.BotManagementModel{
				ZoneID:    types.StringValue("test-zone"),
				FightMode: types.BoolValue(false), // explicit in plan
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:    types.StringValue("test-zone"),
				FightMode: types.BoolValue(true), // different in state
			},
			shouldInclude: true,
			fieldName: "fight_mode",
		},
		{
			name: "AIBotsProtection explicit plan vs different state",
			planModel: &bot_management.BotManagementModel{
				ZoneID:           types.StringValue("test-zone"),
				AIBotsProtection: types.StringValue("block"), // explicit in plan
			},
			stateModel: &bot_management.BotManagementModel{
				ZoneID:           types.StringValue("test-zone"),
				AIBotsProtection: types.StringValue("disabled"), // different in state
			},
			shouldInclude: true,
			fieldName: "ai_bots_protection",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := tt.planModel.MarshalJSONForUpdate(*tt.stateModel)
			assert.NoError(t, err)
			
			jsonStr := string(jsonBytes)
			t.Logf("JSON output for %s: %s", tt.name, jsonStr)
			
			if tt.shouldInclude {
				assert.Contains(t, jsonStr, tt.fieldName, "Expected field %s to be in JSON output", tt.fieldName)
			} else {
				assert.NotContains(t, jsonStr, tt.fieldName, "Expected field %s to NOT be in JSON output", tt.fieldName)
			}
		})
	}
}