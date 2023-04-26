package flatteners

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Bool accepts a `*bool` and returns a `basetypes.BoolValue`. The
// return type automatically handles `BoolNull` should the boolean not be
// initialised.
//
// This flattener saves you repeating code that looks like the following when
// saving to state.
//
//	var enabled *bool
//	if !schema.Enabled.IsNull() {
//	    requestPayload.Enabled = types.BoolValue(enabled)
//	} else {
//	    requestPayload.Enabled = types.BoolNull()
//	}
//
// nolint: contextcheck
func Bool(in *bool) basetypes.BoolValue {
	if reflect.ValueOf(in).IsNil() {
		return types.BoolNull()
	}
	return types.BoolValue(cloudflare.Bool(in))
}
