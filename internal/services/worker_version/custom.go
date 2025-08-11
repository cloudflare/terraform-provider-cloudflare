package worker_version

import (
	"context"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func UpdateSecretTextsFromState[T any](
	ctx context.Context,
	refreshed customfield.NestedObjectList[T],
	state customfield.NestedObjectList[T],
) (customfield.NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	if refreshed.IsNull() {
		return refreshed, diags
	}

	refreshedElems := refreshed.Elements()
	stateElems := state.Elements()

	updatedElems := make([]attr.Value, 0, len(refreshedElems))

	elemType := refreshed.ElementType(ctx)

	objType, ok := elemType.(basetypes.ObjectType)
	if !ok {
		diags.AddError("Invalid element type", "Expected element type to be basetypes.ObjectType.")
		return refreshed, diags
	}

	attrTypes := objType.AttributeTypes()

	for _, val := range refreshedElems {
		refreshedObj, ok := val.(basetypes.ObjectValue)
		if !ok {
			updatedElems = append(updatedElems, val)
			continue
		}

		refreshedAttrs := refreshedObj.Attributes()
		typeAttr := refreshedAttrs["type"]
		nameAttr := refreshedAttrs["name"]

		if typeAttr.IsNull() || nameAttr.IsNull() {
			updatedElems = append(updatedElems, val)
			continue
		}

		if typeAttr.(types.String).ValueString() != "secret_text" {
			updatedElems = append(updatedElems, val)
			continue
		}

		name := nameAttr.(types.String).ValueString()

		var originalText attr.Value
		var foundInState bool
		for _, stateVal := range stateElems {
			stateObj, ok := stateVal.(basetypes.ObjectValue)
			if !ok {
				continue
			}
			stateAttrs := stateObj.Attributes()
			if stateAttrs["type"].(types.String).ValueString() == "secret_text" &&
				stateAttrs["name"].(types.String).ValueString() == name {
				originalText = stateAttrs["text"]
				foundInState = true
				break
			}
		}

		if !foundInState {
			continue
		}

		if originalText != nil && !originalText.IsNull() && !originalText.IsUnknown() {
			refreshedAttrs["text"] = originalText

			newObj, d := basetypes.NewObjectValue(attrTypes, refreshedAttrs)
			diags.Append(d...)
			refreshedObj = newObj
		}

		updatedElems = append(updatedElems, refreshedObj)
	}

	value, d := types.ListValue(refreshed.ElementType(ctx), updatedElems)
	diags.Append(d...)

	return customfield.NestedObjectList[T]{
		ListValue: value,
	}, diags
}

// UnknownOnlyIfModifier only allows a value to be marked as unknown if
// some other attribute is equal to a given value.
//
// This can be useful in cases where a collection type is polymorphic and a
// Computed nested attribute is causing unwanted plan diffs for unaffected element types.
// Essentially, this plan modifier is a workaround for the lack of support for
// discriminated unions in Terraform's resource schemas.
type UnknownOnlyIfModifier struct {
	conditionAttributeName string
	triggerValue           attr.Value
}

func (m UnknownOnlyIfModifier) Description(_ context.Context) string {
	return fmt.Sprintf("Marks attribute as known unless %s equals %s", m.conditionAttributeName, m.triggerValue.String())
}

func (m UnknownOnlyIfModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m UnknownOnlyIfModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	m.planModify(ctx, req.Path, req.ConfigValue, req.Plan, &resp.Diagnostics, func(knownValue attr.Value) {
		resp.PlanValue = knownValue.(types.String)
	})
}

func (m UnknownOnlyIfModifier) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	m.planModify(ctx, req.Path, req.ConfigValue, req.Plan, &resp.Diagnostics, func(knownValue attr.Value) {
		resp.PlanValue = knownValue.(types.Int64)
	})
}

func (m UnknownOnlyIfModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	m.planModify(ctx, req.Path, req.ConfigValue, req.Plan, &resp.Diagnostics, func(knownValue attr.Value) {
		resp.PlanValue = knownValue.(types.Bool)
	})
}

func (m UnknownOnlyIfModifier) planModify(ctx context.Context, attrPath path.Path, configValue attr.Value, plan tfsdk.Plan, diags *diag.Diagnostics, setPlanValue func(attr.Value)) {
	parentPath := attrPath.ParentPath()
	conditionPath := parentPath.AtName(m.conditionAttributeName)

	var planValue attr.Value
	var conditionValue attr.Value

	diags.Append(plan.GetAttribute(ctx, attrPath, &planValue)...)
	diags.Append(plan.GetAttribute(ctx, conditionPath, &conditionValue)...)

	if diags.HasError() {
		return
	}

	if conditionValue.Equal(m.triggerValue) {
		return
	}

	if planValue.IsUnknown() {
		setPlanValue(configValue)
	}
}

// UnknownOnlyIf creates a modifier that keeps an attribute from being marked as unknown unless a sibling attribute has a given value
func UnknownOnlyIf(siblingName string, triggerValue string) planmodifier.String {
	return UnknownOnlyIfModifier{
		conditionAttributeName: siblingName,
		triggerValue:           types.StringValue(triggerValue),
	}
}
