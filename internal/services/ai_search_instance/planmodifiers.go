package ai_search_instance

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// useStateForUnknownIncludingNull returns a plan modifier that preserves the
// prior state value for computed attributes even when the state value is null.
//
// The built-in UseStateForUnknown skips null state values, which causes
// non-empty plans for computed fields where the API returns null. This modifier
// handles that case by also preserving null state values for existing resources.

func useStateForUnknownIncludingNullString() planmodifier.String {
	return useStateForUnknownIncludingNullStringModifier{}
}

type useStateForUnknownIncludingNullStringModifier struct{}

func (m useStateForUnknownIncludingNullStringModifier) Description(_ context.Context) string {
	return "Preserves the prior state value (including null) for computed attributes."
}

func (m useStateForUnknownIncludingNullStringModifier) MarkdownDescription(_ context.Context) string {
	return "Preserves the prior state value (including null) for computed attributes."
}

func (m useStateForUnknownIncludingNullStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// During Create there is no prior state — let the provider compute the value.
	if req.State.Raw.IsNull() {
		return
	}

	// If the plan already has a known value, do not override it.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Preserve the prior state value even when it is null.
	resp.PlanValue = req.StateValue
}

func useStateForUnknownIncludingNullObject() planmodifier.Object {
	return useStateForUnknownIncludingNullObjectModifier{}
}

type useStateForUnknownIncludingNullObjectModifier struct{}

func (m useStateForUnknownIncludingNullObjectModifier) Description(_ context.Context) string {
	return "Preserves the prior state value (including null) for computed attributes."
}

func (m useStateForUnknownIncludingNullObjectModifier) MarkdownDescription(_ context.Context) string {
	return "Preserves the prior state value (including null) for computed attributes."
}

func (m useStateForUnknownIncludingNullObjectModifier) PlanModifyObject(_ context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// During Create there is no prior state — let the provider compute the value.
	if req.State.Raw.IsNull() {
		return
	}

	// If the plan already has a known value, do not override it.
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Preserve the prior state value even when it is null.
	resp.PlanValue = req.StateValue
}
