package migrations

import "github.com/hashicorp/terraform-plugin-framework/types"

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
