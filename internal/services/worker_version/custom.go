package worker_version

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func readFile(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		dirname, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("could not expand home directory in path %s: %w", path, err)
		}
		path = filepath.Join(dirname, path[2:])
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read file %s: %w", path, err)
	}

	return string(content), nil
}

func calculateFileHash(filePath string) (string, error) {
	if strings.HasPrefix(filePath, "~/") {
		dirname, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("could not expand home directory in path %s: %w", filePath, err)
		}
		filePath = filepath.Join(dirname, filePath[2:])
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func calculateStringHash(content string) (string, error) {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:]), nil
}

func ComputeSHA256HashOfContent() planmodifier.String {
	return computeSHA256HashOfContentModifier{}
}

var _ planmodifier.String = &computeSHA256HashOfContentModifier{}

type computeSHA256HashOfContentModifier struct{}

func (c computeSHA256HashOfContentModifier) Description(_ context.Context) string {
	return "Calculates the SHA-256 hash of the provided module content."
}

func (c computeSHA256HashOfContentModifier) MarkdownDescription(ctx context.Context) string {
	return c.Description(ctx)
}

func (c computeSHA256HashOfContentModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.Config.Raw.IsNull() {
		return
	}

	contentFilePath := req.Path.ParentPath().AtName("content_file")
	contentBase64Path := req.Path.ParentPath().AtName("content_base64")

	var contentFile types.String
	var contentBase64 types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, contentFilePath, &contentFile)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, contentBase64Path, &contentBase64)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !contentFile.IsNull() && !contentFile.IsUnknown() {
		contentSHA256, err := calculateFileHash(contentFile.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(req.Path, "Error computing SHA-256 hash", err.Error())
			return
		}
		resp.PlanValue = types.StringValue(contentSHA256)
	} else if !contentBase64.IsNull() && !contentBase64.IsUnknown() {
		content, err := base64.StdEncoding.DecodeString(contentBase64.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(req.Path, "Error decoding base64", err.Error())
			return
		}
		contentSHA256, err := calculateStringHash(string(content))
		if err != nil {
			resp.Diagnostics.AddAttributeError(req.Path, "Error computing SHA-256 hash", err.Error())
			return
		}
		resp.PlanValue = types.StringValue(contentSHA256)
	}
}

// Binding types with sensitive text fields not returned by the API.
var bindingTypesWithSensitiveText = map[string]bool{
	"secret_text": true,
	"plain_text":  true,
}

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

		bindingType := typeAttr.(types.String).ValueString()
		if !bindingTypesWithSensitiveText[bindingType] {
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
			stateType := stateAttrs["type"].(types.String).ValueString()
			stateName := stateAttrs["name"].(types.String).ValueString()
			// Match by name and binding type (both secret_text and plain_text have sensitive text)
			if bindingTypesWithSensitiveText[stateType] && stateName == name {
				originalText = stateAttrs["text"]
				foundInState = true
				break
			}
		}

		if !foundInState {
			updatedElems = append(updatedElems, refreshedObj)
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

// Sorts the list of refreshed bindings to match the order they appear in state
// (or planned state), with refreshed bindings that don't appear in state
// ordered last.
func SortRefreshedBindingsToMatchPrevious[T any](
	ctx context.Context,
	refreshedBindings customfield.NestedObjectList[T],
	previousBindings customfield.NestedObjectList[T],
) (customfield.NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics
	if refreshedBindings.IsNull() {
		return refreshedBindings, diags
	}

	refreshedBindingElements := refreshedBindings.Elements()
	// Mapping of binding name to refreshed binding value.
	refreshedBindingsByName := make(map[string]attr.Value, len(refreshedBindingElements))
	for _, val := range refreshedBindingElements {
		refreshedObj, ok := val.(basetypes.ObjectValue)
		if !ok {
			diags.AddError("Invalid element type", "Expected element type to be basetypes.ObjectType.")
			return refreshedBindings, diags
		}

		refreshedAttrs := refreshedObj.Attributes()
		nameAttr, ok := refreshedAttrs["name"]
		if !ok {
			diags.AddError("Invalid element", "Missing 'name' attribute")
			return refreshedBindings, diags
		}

		nameString, ok := nameAttr.(types.String)
		if !ok {
			diags.AddError("Invalid element", "'name' attribute must be a string")
			return refreshedBindings, diags
		}
		refreshedBindingsByName[nameString.ValueString()] = refreshedObj
	}

	// Refreshed bindings sorted to match the order they appear in state (or
	// planned state), with refreshed bindings that don't appear in state
	// ordered last.
	sortedBindings := make([]attr.Value, 0, len(refreshedBindingElements))
	for _, val := range previousBindings.Elements() {
		stateObj, ok := val.(basetypes.ObjectValue)
		if !ok {
			diags.AddError("Invalid element type", "Expected element type to be basetypes.ObjectType.")
			return refreshedBindings, diags
		}

		stateAttrs := stateObj.Attributes()
		nameAttr, ok := stateAttrs["name"]
		if !ok {
			diags.AddError("Invalid element", "Missing 'name' attribute")
			return refreshedBindings, diags
		}

		nameString, ok := nameAttr.(types.String)
		if !ok {
			diags.AddError("Invalid element", "'name' attribute must be a string")
			return refreshedBindings, diags
		}

		// Reorder refreshed bindings that exist in state (or planned state) to
		// match the order they appear in state.
		if refreshedBinding, ok := refreshedBindingsByName[nameString.ValueString()]; ok {
			sortedBindings = append(sortedBindings, refreshedBinding)
			// Binding names must be unique, the API will never return multiple
			// bindings with the same name.
			delete(refreshedBindingsByName, nameString.ValueString())
		}
	}

	// Add any refreshed bindings that don't exist in state (or planned state).
	for _, refreshedBinding := range refreshedBindingsByName {
		sortedBindings = append(sortedBindings, refreshedBinding)
	}

	value, d := types.ListValue(refreshedBindings.ElementType(ctx), sortedBindings)
	diags.Append(d...)

	return customfield.NestedObjectList[T]{
		ListValue: value,
	}, diags
}

// Sorts the given list of bindings in ascending order by name. This matches the
// order that the Workers API returns bindings.
func SortBindingsByName[T any](
	ctx context.Context,
	bindings customfield.NestedObjectList[T],
) (customfield.NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics
	if bindings.IsNull() {
		return bindings, diags
	}

	sortedBindings := bindings.Elements()
	slices.SortFunc(sortedBindings, func(a, b attr.Value) int {
		aObj, ok := a.(basetypes.ObjectValue)
		if !ok {
			diags.AddError("Invalid element type", "Expected element type to be basetypes.ObjectType.")
			return 0
		}
		aAttrs := aObj.Attributes()
		aNameAttr, ok := aAttrs["name"]
		if !ok {
			diags.AddError("Invalid element", "Missing 'name' attribute")
			return 0
		}
		aNameString, ok := aNameAttr.(types.String)
		if !ok {
			diags.AddError("Invalid element", "'name' attribute must be a string")
			return 0
		}

		bObj, ok := b.(basetypes.ObjectValue)
		if !ok {
			diags.AddError("Invalid element type", "Expected element type to be basetypes.ObjectType.")
			return 0
		}
		bAttrs := bObj.Attributes()
		bNameAttr, ok := bAttrs["name"]
		if !ok {
			diags.AddError("Invalid element", "Missing 'name' attribute")
			return 0
		}
		bNameString, ok := bNameAttr.(types.String)
		if !ok {
			diags.AddError("Invalid element", "'name' attribute must be a string")
			return 0
		}
		return strings.Compare(aNameString.ValueString(), bNameString.ValueString())
	})

	value, d := types.ListValue(bindings.ElementType(ctx), sortedBindings)
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

func RequiresReplaceIfStateValueExists() planmodifier.String {
	description := "Requires replacement if the state value is not null or unknown."
	return stringplanmodifier.RequiresReplaceIf(
		func(ctx context.Context, req planmodifier.StringRequest, res *stringplanmodifier.RequiresReplaceIfFuncResponse) {
			res.RequiresReplace = !req.StateValue.IsNull() && !req.StateValue.IsUnknown()
		},
		description,
		description,
	)
}

// UseStateForUnknownOrNullSensitive preserves state value when config is null/unknown.
func UseStateForUnknownOrNullSensitive() planmodifier.String {
	return useStateForUnknownOrNullSensitiveModifier{}
}

type useStateForUnknownOrNullSensitiveModifier struct{}

func (m useStateForUnknownOrNullSensitiveModifier) Description(_ context.Context) string {
	return "Preserves state value when config is null or unknown."
}

func (m useStateForUnknownOrNullSensitiveModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateForUnknownOrNullSensitiveModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	tflog.Debug(ctx, "UseStateForUnknownOrNullSensitive plan modifier called",
		map[string]interface{}{
			"path":         req.Path.String(),
			"state_value":  req.StateValue.String(),
			"config_value": req.ConfigValue.String(),
			"plan_value":   req.PlanValue.String(),
		})

	// If there's no prior state (e.g., during create), nothing to do
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		tflog.Debug(ctx, "UseStateForUnknownOrNullSensitive: state is null/unknown, skipping")
		return
	}

	// If config explicitly sets a value, use it
	if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
		tflog.Debug(ctx, "UseStateForUnknownOrNullSensitive: config has value, using config")
		return
	}

	// Config is null or unknown, but state has a value - use state value
	// This handles the import case where Terraform generates config without sensitive values
	tflog.Debug(ctx, "UseStateForUnknownOrNullSensitive: preserving state value",
		map[string]interface{}{
			"preserved_value": req.StateValue.String(),
		})
	resp.PlanValue = req.StateValue
}
