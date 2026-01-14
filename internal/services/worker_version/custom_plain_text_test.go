package worker_version

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Regression test: plain_text bindings should preserve text values from state.
func TestUpdateSecretTextsFromState_PlainTextBindings(t *testing.T) {
	ctx := context.Background()

	bindingAttrTypes := map[string]attr.Type{
		"name": types.StringType,
		"type": types.StringType,
		"text": types.StringType,
	}
	bindingObjType := types.ObjectType{AttrTypes: bindingAttrTypes}

	// API response (text is null)
	refreshedBinding, diags := types.ObjectValue(bindingAttrTypes, map[string]attr.Value{
		"name": types.StringValue("BLOCKED_CITIES"),
		"type": types.StringValue("plain_text"),
		"text": types.StringNull(),
	})
	if diags.HasError() {
		t.Fatalf("Failed to create refreshed binding: %v", diags)
	}

	refreshedList, diags := types.ListValue(bindingObjType, []attr.Value{refreshedBinding})
	if diags.HasError() {
		t.Fatalf("Failed to create refreshed list: %v", diags)
	}
	refreshedBindings := customfield.NestedObjectList[WorkerVersionBindingsModel]{
		ListValue: refreshedList,
	}

	// State (has text value)
	stateBinding, diags := types.ObjectValue(bindingAttrTypes, map[string]attr.Value{
		"name": types.StringValue("BLOCKED_CITIES"),
		"type": types.StringValue("plain_text"),
		"text": types.StringValue("kyiv,kharkiv"),
	})
	if diags.HasError() {
		t.Fatalf("Failed to create state binding: %v", diags)
	}

	stateList, diags := types.ListValue(bindingObjType, []attr.Value{stateBinding})
	if diags.HasError() {
		t.Fatalf("Failed to create state list: %v", diags)
	}
	stateBindings := customfield.NestedObjectList[WorkerVersionBindingsModel]{
		ListValue: stateList,
	}

	result, diags := UpdateSecretTextsFromState(ctx, refreshedBindings, stateBindings)
	if diags.HasError() {
		t.Fatalf("UpdateSecretTextsFromState returned errors: %v", diags)
	}

	resultElems := result.Elements()
	if len(resultElems) != 1 {
		t.Fatalf("Expected 1 binding, got %d", len(resultElems))
	}

	resultObj := resultElems[0].(basetypes.ObjectValue)
	textAttr := resultObj.Attributes()["text"]
	if textAttr.IsNull() {
		t.Errorf("Text should be preserved from state, but was null")
	} else if textAttr.(types.String).ValueString() != "kyiv,kharkiv" {
		t.Errorf("Expected 'kyiv,kharkiv', got '%s'", textAttr.(types.String).ValueString())
	}
}

