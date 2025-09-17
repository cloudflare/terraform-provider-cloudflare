package workers_script

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func writeFileBytes(partName string, filename string, contentType string, content io.Reader, writer *multipart.Writer) error {
	h := make(textproto.MIMEHeader)
	header := "form-data"

	if escapeQuotes(partName) != "" {
		header += fmt.Sprintf(`; name="%s"`, escapeQuotes(partName))
	}

	if escapeQuotes(filename) != "" {
		header += fmt.Sprintf(`; filename="%s"`, escapeQuotes(filename))
	}

	h.Set("Content-Disposition", header)
	h.Set("Content-Type", contentType)
	filewriter, err := writer.CreatePart(h)
	if err != nil {
		return err
	}
	_, err = io.Copy(filewriter, content)
	return err
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

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
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func calculateStringHash(content string) (string, error) {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:]), nil
}

var _ validator.String = &contentSHA256Validator{}

type contentSHA256Validator struct {
	ContentPath     string
	ContentFilePath string
}

func (v contentSHA256Validator) Description(_ context.Context) string {
	return fmt.Sprintf("Validates that the provided value matches the SHA-256 hash of content in either `content` or `content_file`.")
}

func (v contentSHA256Validator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v contentSHA256Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	providedHash := req.ConfigValue.ValueString()

	var config WorkersScriptModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var hasContent, hasContentFile bool

	if !config.Content.IsNull() {
		hasContent = true
	}

	if !config.ContentFile.IsNull() {
		hasContentFile = true
	}

	var actualHash string
	var err error

	if hasContent {
		actualHash, err = calculateStringHash(config.Content.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Hash Calculation Error",
				fmt.Sprintf("Failed to calculate SHA-256 hash of content: %s", err.Error()),
			)
			return
		}
	} else if hasContentFile {
		actualHash, err = calculateFileHash(config.ContentFile.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Hash Calculation Error",
				fmt.Sprintf("Failed to calculate SHA-256 hash of file '%s': %s", config.ContentFile.ValueString(), err.Error()),
			)
			return
		}
	}

	if providedHash != actualHash {
		var source string
		if hasContent {
			source = "content"
		} else if hasContentFile {
			source = fmt.Sprintf("content_file (%s)", config.ContentFile.ValueString())
		}

		resp.Diagnostics.AddAttributeError(
			req.Path,
			"SHA-256 Hash Mismatch",
			fmt.Sprintf("The provided SHA-256 hash '%s' does not match the actual hash '%s' of %s",
				providedHash, actualHash, source),
		)
	}
}

func ValidateContentSHA256() validator.String {
	return contentSHA256Validator{
		ContentPath:     "content",
		ContentFilePath: "content_file",
	}
}

func UpdateSecretTextsFromState[T any](
	ctx context.Context,
	refreshed customfield.NestedObjectList[T],
	state customfield.NestedObjectList[T],
) (customfield.NestedObjectList[T], diag.Diagnostics) {
	var diags diag.Diagnostics

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

	// Iterate through state elements first to preserve ordering
	for _, stateVal := range stateElems {
		stateObj, ok := stateVal.(basetypes.ObjectValue)
		if !ok {
			updatedElems = append(updatedElems, stateVal)
			continue
		}

		stateAttrs := stateObj.Attributes()
		typeAttr := stateAttrs["type"]
		nameAttr := stateAttrs["name"]

		if typeAttr.IsNull() || nameAttr.IsNull() {
			updatedElems = append(updatedElems, stateVal)
			continue
		}

		if typeAttr.(types.String).ValueString() != "secret_text" {
			updatedElems = append(updatedElems, stateVal)
			continue
		}

		name := nameAttr.(types.String).ValueString()

		// Find matching element in refreshed data
		var refreshedObj basetypes.ObjectValue
		var foundInRefreshed bool
		for _, refreshedVal := range refreshedElems {
			refreshedObjCandidate, ok := refreshedVal.(basetypes.ObjectValue)
			if !ok {
				continue
			}
			refreshedAttrs := refreshedObjCandidate.Attributes()
			if refreshedAttrs["type"].(types.String).ValueString() == "secret_text" &&
				refreshedAttrs["name"].(types.String).ValueString() == name {
				refreshedObj = refreshedObjCandidate
				foundInRefreshed = true
				break
			}
		}

		if !foundInRefreshed {
			// Element was removed from API, skip it
			continue
		}

		// Preserve original secret text from state
		originalText := stateAttrs["text"]
		if originalText != nil && !originalText.IsNull() && !originalText.IsUnknown() {
			refreshedAttrs := refreshedObj.Attributes()
			refreshedAttrs["text"] = originalText

			newObj, d := basetypes.NewObjectValue(attrTypes, refreshedAttrs)
			diags.Append(d...)
			refreshedObj = newObj
		}

		updatedElems = append(updatedElems, refreshedObj)
	}

	// Add any new elements from API that weren't in state
	for _, refreshedVal := range refreshedElems {
		refreshedObj, ok := refreshedVal.(basetypes.ObjectValue)
		if !ok {
			continue
		}

		refreshedAttrs := refreshedObj.Attributes()
		typeAttr := refreshedAttrs["type"]
		nameAttr := refreshedAttrs["name"]

		if typeAttr.IsNull() || nameAttr.IsNull() {
			continue
		}

		name := nameAttr.(types.String).ValueString()

		// Check if this element was already processed from state
		var foundInState bool
		for _, stateVal := range stateElems {
			stateObj, ok := stateVal.(basetypes.ObjectValue)
			if !ok {
				continue
			}
			stateAttrs := stateObj.Attributes()
			if stateAttrs["type"].(types.String).ValueString() == refreshedAttrs["type"].(types.String).ValueString() &&
				stateAttrs["name"].(types.String).ValueString() == name {
				foundInState = true
				break
			}
		}

		if !foundInState {
			// This is a new element from the API, add it
			updatedElems = append(updatedElems, refreshedObj)
		}
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
	m.planModify(ctx, req.Path, req.ConfigValue, req.StateValue, req.Plan, &resp.Diagnostics, func(knownValue attr.Value) {
		resp.PlanValue = knownValue.(types.String)
	})
}

func (m UnknownOnlyIfModifier) PlanModifyInt64(ctx context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	m.planModify(ctx, req.Path, req.ConfigValue, req.StateValue, req.Plan, &resp.Diagnostics, func(knownValue attr.Value) {
		resp.PlanValue = knownValue.(types.Int64)
	})
}

func (m UnknownOnlyIfModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	m.planModify(ctx, req.Path, req.ConfigValue, req.StateValue, req.Plan, &resp.Diagnostics, func(knownValue attr.Value) {
		resp.PlanValue = knownValue.(types.Bool)
	})
}

func (m UnknownOnlyIfModifier) planModify(ctx context.Context, attrPath path.Path, configValue attr.Value, stateValue attr.Value, plan tfsdk.Plan, diags *diag.Diagnostics, setPlanValue func(attr.Value)) {
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
