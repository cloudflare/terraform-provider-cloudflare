package defaults

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Map = mapDefaultModifier{}

func DefaultMap(elements map[string]attr.Value) mapDefaultModifier {
	return mapDefaultModifier{Elements: elements}
}

type mapDefaultModifier struct {
	Elements map[string]attr.Value
}

// Description returns a plain text description of the validator's behavior,
// suitable for a practitioner to understand its impact.
func (m mapDefaultModifier) Description(_ context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to %s", m.Elements)
}

// MarkdownDescription returns a markdown formatted description of the
// validator's behavior, suitable for a practitioner to understand its impact.
func (m mapDefaultModifier) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m.Elements)
}

// PlanModifyMap updates the planned value with the default if its not null.
func (m mapDefaultModifier) PlanModifyMap(ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
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
	resp.PlanValue, resp.Diagnostics = types.MapValue(req.PlanValue.ElementType(ctx), m.Elements)
}
