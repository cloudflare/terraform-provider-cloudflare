package customfield

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
)

// NullifyEmptyObject checks whether a NestedObject has all-null attributes and,
// if so, returns a null NestedObject instead. This is useful when an API returns
// empty objects (e.g. {"hostname": {}}) for fields the user did not set, which
// would otherwise cause Terraform set-correlation failures because a known object
// with all-null attributes is not structurally equal to a null object.
//
// IMPORTANT: Do not use this on empty marker structs (zero fields) where the
// mere presence of the object is the semantic value (e.g. everyone = {},
// certificate = {}, any_valid_service_token = {}). For those types, {} and null
// are semantically different. This function is safe for such types because it
// returns the original object unchanged when there are zero attributes.
func NullifyEmptyObject[T any](ctx context.Context, obj NestedObject[T]) NestedObject[T] {
	if obj.IsNull() || obj.IsUnknown() {
		return obj
	}

	attrs := obj.ObjectValue.Attributes()
	if len(attrs) == 0 {
		// Empty struct (zero fields) -- presence IS the value, do not nullify.
		return obj
	}

	for _, v := range attrs {
		if !attrIsNullOrUnknown(v) {
			return obj
		}
	}

	return NullObject[T](ctx)
}

// attrIsNullOrUnknown returns true if the given attr.Value is null or unknown.
// It handles both the attr.Value interface methods and nested null objects.
func attrIsNullOrUnknown(v attr.Value) bool {
	return v.IsNull() || v.IsUnknown()
}
