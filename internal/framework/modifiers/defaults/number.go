package defaults

import (
	"context"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Number = numberDefaultModifier{}

func DefaultNumber(s *big.Float) numberDefaultModifier {
	return numberDefaultModifier{Default: s}
}

type numberDefaultModifier struct {
	Default *big.Float
}

// Description returns a plain text description of the validator's behavior,
// suitable for a practitioner to understand its impact.
func (m numberDefaultModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to %s", m.Default)
}

// MarkdownDescription returns a markdown formatted description of the
// validator's behavior, suitable for a practitioner to understand its impact.
func (m numberDefaultModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("If value is not configured, defaults to `%s`", m.Default)
}

// PlanModifyNumber updates the planned value with the default if its not null.
func (m numberDefaultModifier) PlanModifyNumber(ctx context.Context, req planmodifier.NumberRequest, resp *planmodifier.NumberResponse) {
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
	resp.PlanValue = types.NumberValue(m.Default)
}
