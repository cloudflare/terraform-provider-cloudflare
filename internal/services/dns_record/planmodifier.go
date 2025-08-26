package dns_record

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// normalizeContentPlanModifier implements a plan modifier that normalizes DNS record content
// based on the record type to match what the Cloudflare API returns.
type normalizeContentPlanModifier struct{}

// NormalizeContent returns a plan modifier that normalizes DNS record content.
// For CNAME, NS, MX, and PTR records, the content is lowercased to match API behavior.
func NormalizeContent() planmodifier.String {
	return normalizeContentPlanModifier{}
}

func (m normalizeContentPlanModifier) Description(ctx context.Context) string {
	return "Normalizes DNS record content to match Cloudflare API behavior"
}

func (m normalizeContentPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Normalizes DNS record content to match Cloudflare API behavior (e.g., lowercases CNAME content)"
}

func (m normalizeContentPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Nothing to do if there is no planned value
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	// Nothing to do if there is no state value (resource is being created)
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}

	// Get the record type to determine normalization rules
	var recordType types.String
	diags := req.Plan.GetAttribute(ctx, path.Root("type"), &recordType)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Apply normalization based on record type
	plannedContent := req.PlanValue.ValueString()
	stateContent := req.StateValue.ValueString()

	// For CNAME, NS, MX (content part), and PTR records, the API lowercases domain names
	switch strings.ToUpper(recordType.ValueString()) {
	case "CNAME", "NS", "PTR":
		// If the lowercase version of planned matches lowercase state, keep state value
		// This handles cases where API partially normalizes case
		if strings.ToLower(plannedContent) == strings.ToLower(stateContent) {
			resp.PlanValue = req.StateValue
			return
		}
	case "MX":
		// MX records have priority and domain, only domain part is lowercased
		// Format is typically "10 mail.example.com"
		plannedParts := strings.SplitN(plannedContent, " ", 2)
		stateParts := strings.SplitN(stateContent, " ", 2)

		if len(plannedParts) == 2 && len(stateParts) == 2 {
			// Compare priority and lowercased domain
			if plannedParts[0] == stateParts[0] &&
				strings.ToLower(plannedParts[1]) == strings.ToLower(stateParts[1]) {
				resp.PlanValue = req.StateValue
				return
			}
		}
	}

	// For other record types or if normalization doesn't match, keep planned value
}

// omitIfRelatedFieldNullPlanModifier implements a plan modifier that keeps the field
// null if a related field is null (e.g., comment_modified_on should be null if comment is null).
type omitIfRelatedFieldNullPlanModifier struct {
	relatedFieldPath path.Path
}

// OmitIfRelatedFieldNull returns a plan modifier that keeps the field null
// if the related field is null, preventing "known after apply" for conditional computed fields.
func OmitIfRelatedFieldNull(relatedFieldPath path.Path) planmodifier.String {
	return omitIfRelatedFieldNullPlanModifier{
		relatedFieldPath: relatedFieldPath,
	}
}

func (m omitIfRelatedFieldNullPlanModifier) Description(ctx context.Context) string {
	return "Keeps field null if related field is null"
}

func (m omitIfRelatedFieldNullPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Keeps field null if related field is null to prevent drift for conditional computed fields"
}

func (m omitIfRelatedFieldNullPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Check if the related field is null
	var relatedValue types.String
	diags := req.Plan.GetAttribute(ctx, m.relatedFieldPath, &relatedValue)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If related field is null, keep this field null (not unknown)
	if relatedValue.IsNull() {
		resp.PlanValue = types.StringNull()
		return
	}

	// Otherwise, use state value if it exists
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

// omitIfRelatedSetEmptyPlanModifier implements a plan modifier for fields that should be null
// when a related set field is empty (e.g., tags_modified_on should be null if tags is empty).
type omitIfRelatedSetEmptyPlanModifier struct {
	relatedFieldPath path.Path
}

// OmitIfRelatedSetEmpty returns a plan modifier that keeps the field null
// if the related set field is empty.
func OmitIfRelatedSetEmpty(relatedFieldPath path.Path) planmodifier.String {
	return omitIfRelatedSetEmptyPlanModifier{
		relatedFieldPath: relatedFieldPath,
	}
}

func (m omitIfRelatedSetEmptyPlanModifier) Description(ctx context.Context) string {
	return "Keeps field null if related set is empty"
}

func (m omitIfRelatedSetEmptyPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Keeps field null if related set is empty to prevent drift for conditional computed fields"
}

func (m omitIfRelatedSetEmptyPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Get the raw attribute value
	var attrValue attr.Value
	diags := req.Plan.GetAttribute(ctx, m.relatedFieldPath, &attrValue)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if it's null
	if attrValue.IsNull() {
		// If tags are null, keep this field null but use state value if available
		if !req.StateValue.IsNull() {
			resp.PlanValue = req.StateValue
		} else {
			resp.PlanValue = types.StringNull()
		}
		return
	}

	// Check if it's a standard types.Set
	if setVal, ok := attrValue.(types.Set); ok {
		if len(setVal.Elements()) == 0 {
			// If tags are empty, this field should be null but preserve state if available
			if !req.StateValue.IsNull() {
				resp.PlanValue = req.StateValue
			} else {
				resp.PlanValue = types.StringNull()
			}
			return
		}
	} else {
		// For custom set types (like customfield.Set), check the Terraform representation
		tfValue, err := attrValue.ToTerraformValue(ctx)
		if err != nil {
			// If we can't convert, assume non-empty to be safe
			if !req.StateValue.IsNull() {
				resp.PlanValue = req.StateValue
			}
			return
		}

		// If it's null in Terraform representation, treat as empty
		if tfValue.IsNull() {
			// If tags are null/empty, preserve state value
			if !req.StateValue.IsNull() {
				resp.PlanValue = req.StateValue
			} else {
				resp.PlanValue = types.StringNull()
			}
			return
		}

		// Try to extract as a set/list to count elements
		var elements []tftypes.Value
		if err := tfValue.As(&elements); err == nil {
			if len(elements) == 0 {
				// If tags are empty, preserve state value
				if !req.StateValue.IsNull() {
					resp.PlanValue = req.StateValue
				} else {
					resp.PlanValue = types.StringNull()
				}
				return
			}
		}
	}

	// If we get here, the set is not empty, so use state value if it exists
	if !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
	}
}

// normalizeEmptySetPlanModifier implements a plan modifier that treats empty sets
// and null values as equivalent to prevent drift when the API omits empty arrays.
type normalizeEmptySetPlanModifier struct{}

// NormalizeEmptySet returns a plan modifier that treats empty sets and null as equivalent.
// This prevents drift when the API returns null for empty arrays but Terraform config has [].
func NormalizeEmptySet() planmodifier.Set {
	return normalizeEmptySetPlanModifier{}
}

func (m normalizeEmptySetPlanModifier) Description(ctx context.Context) string {
	return "Treats empty sets and null as equivalent to prevent drift"
}

func (m normalizeEmptySetPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Treats empty sets and null as equivalent to prevent drift when API omits empty arrays"
}

func (m normalizeEmptySetPlanModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// Handle the case where config is null but state has empty array
	// This happens when user doesn't specify the field but API returns empty array
	if req.ConfigValue.IsNull() && !req.StateValue.IsNull() && len(req.StateValue.Elements()) == 0 {
		// Keep the state value (empty set) to prevent drift
		resp.PlanValue = req.StateValue
		return
	}

	// If plan value is null, nothing to do for other cases
	if req.PlanValue.IsNull() {
		return
	}

	// If config value is null (and we didn't handle it above), let it be null
	if req.ConfigValue.IsNull() {
		return
	}

	// If both state and plan are empty sets, keep the state value to prevent drift
	if !req.StateValue.IsNull() && len(req.StateValue.Elements()) == 0 && len(req.PlanValue.Elements()) == 0 {
		resp.PlanValue = req.StateValue
		return
	}

	// If state is null but plan is empty set (first apply), keep the plan value
	// This handles the case where user explicitly sets tags = []
}

// preserveStateIfConfigNullPlanModifier implements a plan modifier that keeps the state value
// when the config is null but the state has values. This prevents drift when the API
// provides default values that aren't specified in configuration.
type preserveStateIfConfigNullPlanModifier struct{}

// PreserveStateIfConfigNull returns a plan modifier that preserves state values
// when configuration is null but state contains API defaults.
func PreserveStateIfConfigNull() planmodifier.Object {
	return preserveStateIfConfigNullPlanModifier{}
}

func (m preserveStateIfConfigNullPlanModifier) Description(ctx context.Context) string {
	return "Preserves state values when config is null but state has API defaults"
}

func (m preserveStateIfConfigNullPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Preserves state values when config is null but state has API defaults to prevent drift"
}

func (m preserveStateIfConfigNullPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// If config is null but state has values, keep the state
	if req.ConfigValue.IsNull() && !req.StateValue.IsNull() {
		resp.PlanValue = req.StateValue
		return
	}

	// For other cases, use standard behavior
}
