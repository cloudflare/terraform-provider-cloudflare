package zero_trust_access_application

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Object = targetRuleSelectorsValidator{}

// Selector size and value validation belongs to the child attributes. This
// validator reports a missing selector at the rule itself.
type targetRuleSelectorsValidator struct{}

func (targetRuleSelectorsValidator) Description(_ context.Context) string {
	return "must contain target_attributes or tags"
}

func (v targetRuleSelectorsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (targetRuleSelectorsValidator) ValidateObject(_ context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	attributes := req.ConfigValue.Attributes()
	for _, name := range []string{"target_attributes", "tags"} {
		if value := attributes[name]; value != nil && !value.IsNull() {
			// Unknown selectors may be populated later, so defer validation.
			return
		}
	}

	resp.Diagnostics.AddAttributeError(req.Path, "Missing Target Rule Selector", "Each target rule must contain target_attributes or tags.")
}
