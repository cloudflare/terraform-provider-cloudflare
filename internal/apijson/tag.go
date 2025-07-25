package apijson

import (
	"reflect"
	"strings"
)

const jsonStructTag = "json"
const formatStructTag = "format"

type parsedStructTag struct {
	name              string
	extras            bool
	metadata          bool
	inline            bool
	required          bool
	optional          bool
	computed          bool
	computed_optional bool
	noRefresh         bool
	// Don't skip this value, even if it's computed (no-op for computed optional fields)
	// If encodeStateForUnknown is set on a computed field, this flag should also be set;
	// otherwise this flag will have no effect
	// NOTE: won't work if update behavior is 'patch'
	forceEncode bool
	// If the value in the plan is unknown,
	// encode the value from the state instead
	// This is similar to the UseStateForUnknown plan modifier,
	// but it only impacts serialization of request bodies, not planning.
	// NOTE #1: only use this for computed/computed_optional values that may be changed by the server;
	// otherwise just use the UseStateForUnknown plan modifier
	// NOTE #2: won't work if update behavior is 'patch'
	encodeStateValueWhenPlanUnknown bool
}

func parseJSONStructTag(field reflect.StructField) (tag parsedStructTag, ok bool) {
	raw, ok := field.Tag.Lookup(jsonStructTag)
	if !ok {
		return
	}
	parts := strings.Split(raw, ",")
	if len(parts) == 0 {
		return tag, false
	}
	tag.name = parts[0]
	for _, part := range parts[1:] {
		switch part {
		case "extras":
			tag.extras = true
		case "metadata":
			tag.metadata = true
		case "inline":
			tag.inline = true
		case "required":
			tag.required = true
		case "optional":
			tag.optional = true
		case "computed":
			tag.computed = true
		case "computed_optional":
			tag.computed_optional = true
		case "no_refresh":
			tag.noRefresh = true
		case "encode_state_for_unknown":
			tag.encodeStateValueWhenPlanUnknown = true
		case "force_encode":
			tag.forceEncode = true
		}
	}
	return
}

func parseFormatStructTag(field reflect.StructField) (format string, ok bool) {
	format, ok = field.Tag.Lookup(formatStructTag)
	return
}
