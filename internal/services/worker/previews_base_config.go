package worker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func previewsBaseConfigChanged(ctx context.Context, plan, state, config attr.Value) (bool, error) {
	p, err := plan.ToTerraformValue(ctx)
	if err != nil {
		return false, err
	}
	s, err := state.ToTerraformValue(ctx)
	if err != nil {
		return false, err
	}
	c, err := config.ToTerraformValue(ctx)
	if err != nil {
		return false, err
	}
	return previewValueChanged(p, s, c)
}

// Compare known preview values recursively. Unknown Optional+Computed values
// may be stabilized, but unknown values supplied by configuration must survive
// planning so that Terraform can resolve them before the update.
func previewValueChanged(plan, state, config tftypes.Value) (bool, error) {
	if !plan.IsKnown() {
		return !config.IsNull(), nil
	}
	if plan.Equal(state) {
		return false, nil
	}
	if plan.IsNull() || state.IsNull() || !state.IsKnown() {
		return true, nil
	}

	switch plan.Type().(type) {
	case tftypes.Object, tftypes.Map:
		var planned, prior, configured map[string]tftypes.Value
		if err := plan.As(&planned); err != nil {
			return false, err
		}
		if err := state.As(&prior); err != nil {
			return false, err
		}
		if config.IsKnown() && !config.IsNull() {
			if err := config.As(&configured); err != nil {
				return false, err
			}
		}
		if len(planned) != len(prior) {
			return true, nil
		}
		for name, value := range planned {
			old, ok := prior[name]
			if !ok {
				return true, nil
			}
			configuredValue, ok := configured[name]
			if !ok {
				configuredValue = tftypes.NewValue(value.Type(), nil)
				if !config.IsKnown() {
					configuredValue = tftypes.NewValue(value.Type(), tftypes.UnknownValue)
				}
			}
			changed, err := previewValueChanged(value, old, configuredValue)
			if err != nil || changed {
				return changed, err
			}
		}
		return false, nil
	default:
		// Partially unknown collections supplied by configuration cannot safely
		// be replaced with state. Fully known collections compare normally.
		return plan.IsFullyKnown() || !config.IsNull(), nil
	}
}
