package dns_record

import (
	"context"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// normalizeContentPlanModifier implements a plan modifier that normalizes DNS record content
// by handling trailing dots and IPv6 address normalization in addition to case normalization.
type normalizeContentPlanModifier struct{}

// NormalizeContent returns a plan modifier that handles comprehensive content normalization.
// It prevents drift when the API returns normalized values (trailing dots, IPv6 formats, case normalization).
func NormalizeContent() planmodifier.String {
	return normalizeContentPlanModifier{}
}

func (m normalizeContentPlanModifier) Description(ctx context.Context) string {
	return "Normalizes DNS record content to handle case, trailing dots and IPv6 address formats"
}

func (m normalizeContentPlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Normalizes DNS record content to handle case, trailing dots and IPv6 address formats"
}

func (m normalizeContentPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Nothing to do if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// Nothing to do if there is no planned value.
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	// Nothing to do if values are already equal
	if req.PlanValue.Equal(req.StateValue) {
		return
	}

	stateValue := req.StateValue.ValueString()
	planValue := req.PlanValue.ValueString()

	// First, check IPv6 normalization
	if suppressMatchingIPv6(stateValue, planValue) {
		resp.PlanValue = req.StateValue
		return
	}

	// Then check trailing dot normalization
	if suppressTrailingDots(stateValue, planValue) {
		resp.PlanValue = req.StateValue
		return
	}

	// Finally, check case normalization for DNS records (our original logic)
	// Get the record type to determine normalization rules
	var recordType types.String
	diags := req.Plan.GetAttribute(ctx, path.Root("type"), &recordType)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Apply normalization based on record type
	// For CNAME, NS, MX (content part), and PTR records, the API lowercases domain names
	switch strings.ToUpper(recordType.ValueString()) {
	case "CNAME", "NS", "PTR":
		// If the lowercase version of planned matches lowercase state, keep state value
		// This handles cases where API partially normalizes case
		if strings.ToLower(planValue) == strings.ToLower(stateValue) {
			resp.PlanValue = req.StateValue
			return
		}
	case "MX":
		// MX records have priority and domain, only domain part is lowercased
		// Format is typically "10 mail.example.com"
		plannedParts := strings.SplitN(planValue, " ", 2)
		stateParts := strings.SplitN(stateValue, " ", 2)

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

// fqdnNormalizePlanModifier implements a plan modifier that normalizes DNS record names
// by comparing the configured name (potentially a subdomain) with the state name (FQDN from API).
type fqdnNormalizePlanModifier struct{}

// FQDNNormalizePlanModifier returns a plan modifier that handles FQDN normalization for DNS record names.
// It prevents drift when the API returns FQDN but the configuration uses subdomain.
func FQDNNormalizePlanModifier() planmodifier.String {
	return fqdnNormalizePlanModifier{}
}

func (m fqdnNormalizePlanModifier) Description(ctx context.Context) string {
	return "Normalizes DNS record name to handle FQDN vs subdomain differences"
}

func (m fqdnNormalizePlanModifier) MarkdownDescription(ctx context.Context) string {
	return "Normalizes DNS record name to handle FQDN vs subdomain differences"
}

func (m fqdnNormalizePlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Nothing to do if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// Nothing to do if there is no planned value.
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	// Nothing to do if values are already equal
	if req.PlanValue.Equal(req.StateValue) {
		return
	}

	stateValue := req.StateValue.ValueString()
	planValue := req.PlanValue.ValueString()

	// Check if the state value is an FQDN that ends with a zone suffix
	// and the plan value is the same without the zone suffix
	if strings.HasSuffix(stateValue, ".") {
		stateValue = strings.TrimSuffix(stateValue, ".")
	}
	if strings.HasSuffix(planValue, ".") {
		planValue = strings.TrimSuffix(planValue, ".")
	}

	// If the plan value is "@" (apex), it should match the zone name in the state
	if planValue == "@" {
		// The state would have the full zone name
		// Since we can't easily get the zone name here without an API call,
		// we'll check if the state has more parts than just @
		if stateValue != "@" && strings.Contains(stateValue, ".") {
			// Keep the state value to prevent drift
			resp.PlanValue = req.StateValue
			return
		}
	}

	// Check if state is FQDN and plan is subdomain of the same record
	// For example: state="static.example.com.terraform.cfapi.net" plan="static.example.com"
	if strings.HasPrefix(stateValue, planValue+".") {
		// The plan value is a prefix of the state value with a domain suffix
		// This is the expected case - keep the state value to prevent drift
		resp.PlanValue = req.StateValue
		return
	}

	// Check for the case where both have the same subdomain but state has zone suffix
	// Extract the first part (subdomain) from both values
	stateParts := strings.Split(stateValue, ".")
	planParts := strings.Split(planValue, ".")

	// If plan has fewer parts and matches the beginning of state, it's likely the subdomain
	if len(planParts) < len(stateParts) {
		matches := true
		for i, part := range planParts {
			if i >= len(stateParts) || part != stateParts[i] {
				matches = false
				break
			}
		}
		if matches {
			// Plan is subdomain, state is FQDN - keep state value
			resp.PlanValue = req.StateValue
			return
		}
	}

	// If the values are semantically the same (just different FQDN format), keep state
	// This handles edge cases we might not have covered above
	if isSameDNSRecord(planValue, stateValue) {
		resp.PlanValue = req.StateValue
		return
	}
}

// isSameDNSRecord checks if two DNS names refer to the same record,
// accounting for FQDN vs subdomain differences
func isSameDNSRecord(plan, state string) bool {
	// Normalize by removing trailing dots
	plan = strings.TrimSuffix(plan, ".")
	state = strings.TrimSuffix(state, ".")

	// If one is clearly an FQDN of the other
	if strings.HasPrefix(state, plan+".") || strings.HasPrefix(plan, state+".") {
		return true
	}

	// Check if they have the same first label (subdomain)
	planFirst := strings.Split(plan, ".")[0]
	stateFirst := strings.Split(state, ".")[0]

	// If the first parts match and one has more parts (zone suffix), they're likely the same
	if planFirst == stateFirst && len(strings.Split(plan, ".")) != len(strings.Split(state, ".")) {
		return true
	}

	return false
}

// suppressTrailingDots checks if two values are the same after removing trailing dots
func suppressTrailingDots(old, new string) bool {
	newTrimmed := strings.TrimSuffix(new, ".")

	// Ensure to distinguish values consists of dots only.
	if newTrimmed == "" {
		return old == new
	}

	return strings.TrimSuffix(old, ".") == newTrimmed
}

// suppressMatchingIPv6 checks if two IPv6 addresses are equivalent
func suppressMatchingIPv6(old, new string) bool {
	oldIPv6 := net.ParseIP(old)
	if oldIPv6 == nil || oldIPv6.To16() == nil {
		return false
	}
	newIPv6 := net.ParseIP(new)
	if newIPv6 == nil || newIPv6.To16() == nil {
		return false
	}
	return oldIPv6.Equal(newIPv6)
}
