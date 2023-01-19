package defaults

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Set = setDefaultModifier{}

func DefaultSet(elements []attr.Value) setDefaultModifier {
	return setDefaultModifier{Elements: elements}
}

type setDefaultModifier struct {
	Elements []attr.Value
}

// Description returns a plain text description of the validator's behavior,
// suitable for a practitioner to understand its impact.
func (m setDefaultModifier) Description(_ context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to %s", m.Elements)
}

// MarkdownDescription returns a markdown formatted description of the
// validator's behavior, suitable for a practitioner to understand its impact.
func (m setDefaultModifier) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m.Elements)
}

// PlanModifySet updates the planned value with the default if its not null.
func (m setDefaultModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// If the attribute configuration is not null, we are done here
	if !req.ConfigValue.IsNull() {
		return
	}

	// If the attribute plan is "known" and "not null", then a previous plan
	// modifier in the sequence has already been applied, and we don't want to
	// interfere.
	if !req.PlanValue.IsUnknown() && !req.PlanValue.IsNull() {
		return
	}
	resp.PlanValue, resp.Diagnostics = types.SetValue(req.PlanValue.ElementType(ctx), m.Elements)
}
