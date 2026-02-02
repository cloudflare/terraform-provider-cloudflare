package utils

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// RequiresReplaceIfNotCertificateSemantic returns a plan modifier that conditionally requires
// resource replacement if:
//
//   - The resource is planned for update.
//   - The plan and state values are not semantically equal.
//   - The configuration value is not null.
func RequiresReplaceIfNotCertificateSemantic() planmodifier.String {
	return stringplanmodifier.RequiresReplaceIf(
		func(_ context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
			if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() && !req.StateValue.IsNull() && !req.StateValue.IsUnknown() {
				configNormalized := strings.TrimRight(req.ConfigValue.ValueString(), "\n")
				stateNormalized := strings.TrimRight(req.StateValue.ValueString(), "\n")

				resp.RequiresReplace = configNormalized != stateNormalized
			} else {
				resp.RequiresReplace = true
			}
		},
		"Certificate change requires replacement",
		"Certificate change requires replacement",
	)
}
