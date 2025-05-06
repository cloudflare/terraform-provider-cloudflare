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
		}
	}
	return
}

func parseFormatStructTag(field reflect.StructField) (format string, ok bool) {
	format, ok = field.Tag.Lookup(formatStructTag)
	return
}
