package migrations

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	// ErrMoveStateNilSourceStateTitle is the diagnostic title used when
	// req.SourceState is nil during a MoveState operation.
	ErrMoveStateNilSourceStateTitle = "Unable to Read Source State"

	// ErrMoveStateNilSourceStateDetail is the diagnostic detail format string.
	// Use with fmt.Sprintf, passing the source type name.
	ErrMoveStateNilSourceStateDetail = "The source state for %s could not be decoded. " +
		"This typically occurs when the state file uses the legacy flatmap format " +
		"from Terraform versions prior to 0.12. Run 'terraform apply -refresh-only' " +
		"with the v4 provider to upgrade the state format, then retry the v5 migration."
)

// DiagnoseMoveStateNilSourceState checks whether req.SourceState is nil and,
// if so, appends an error diagnostic to resp and returns true. Callers should
// return early when this returns true.
func DiagnoseMoveStateNilSourceState(req resource.MoveStateRequest, resp *resource.MoveStateResponse) bool {
	if req.SourceState != nil {
		return false
	}
	resp.Diagnostics.Append(diag.NewErrorDiagnostic(
		ErrMoveStateNilSourceStateTitle,
		fmt.Sprintf(ErrMoveStateNilSourceStateDetail, req.SourceTypeName),
	))
	return true
}

func FalseyStringToNull(v types.String) types.String {
	if v.IsNull() || v.IsUnknown() {
		return v
	}
	if v.ValueString() == "" {
		return types.StringNull()
	}
	return v
}

func FalseyBoolToNull(v types.Bool) types.Bool {
	if v.IsNull() || v.IsUnknown() {
		return v
	}
	if !v.ValueBool() {
		return types.BoolNull()
	}
	return v
}
