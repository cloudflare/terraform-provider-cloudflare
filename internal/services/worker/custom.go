package worker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NormalizeFloat64() planmodifier.Float64 {
	return normalizeFloat64Modifier{}
}

type normalizeFloat64Modifier struct{}

func (m normalizeFloat64Modifier) Description(_ context.Context) string {
	return "Normalize float64 values to prevent spurious changes"
}

func (m normalizeFloat64Modifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// For some reason sometimes req.StateValue.Equal(req.PlanValue) is false, but
// req.StateValue.ValueFloat64() == req.PlanValue.ValueFloat64() is true. This
// appears to be caused by a difference in how float64 values are parsed in
// different places. This plan modifier normalizes the float64 value in the plan
// such that req.StateValue.Equal(req.PlanValue) is true in this case.
func (m normalizeFloat64Modifier) PlanModifyFloat64(ctx context.Context, req planmodifier.Float64Request, resp *planmodifier.Float64Response) {
	if req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		return
	}
	resp.PlanValue = types.Float64Value(req.PlanValue.ValueFloat64())
}
