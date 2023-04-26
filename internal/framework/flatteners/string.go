package flatteners

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// String accepts a `string` and returns a `basetypes.StringValue`. The
// return type automatically handles `StringNull` should the string be empty.
//
// Removes the need for the following code when saving to state.
//
//	if response.MyField == "" {
//	    state.MyField = types.StringValue(response.MyField)
//	} else {
//	    state.MyField = types.StringNull()
//	}
//
// Not recommended if you care about returning an empty string for the state.
//
// nolint: contextcheck
func String(in string) basetypes.StringValue {
	if in == "" {
		return types.StringNull()
	}
	return types.StringValue(in)
}
