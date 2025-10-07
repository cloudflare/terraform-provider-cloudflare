package worker_version

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ validator.Dynamic = runWorkerFirstValidator{}

// runWorkerFirstValidator validates that a dynamic value is either:
// - A boolean (true/false)
// - A list/tuple of strings (for path rules)
type runWorkerFirstValidator struct{}

func (v runWorkerFirstValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v runWorkerFirstValidator) MarkdownDescription(_ context.Context) string {
	return "value must be either a boolean or a list of strings"
}

func (v runWorkerFirstValidator) ValidateDynamic(ctx context.Context, req validator.DynamicRequest, resp *validator.DynamicResponse) {
	dynamic := req.ConfigValue
	if dynamic.IsNull() || dynamic.IsUnknown() || dynamic.IsUnderlyingValueNull() || dynamic.IsUnderlyingValueUnknown() {
		return
	}

	value := dynamic.UnderlyingValue()
	valueType := value.Type(ctx)

	// Check if it's a boolean
	if _, isBool := valueType.(basetypes.BoolType); isBool {
		return // Boolean is valid
	}

	// Check if it's a list type
	if listType, isList := valueType.(basetypes.ListType); isList {
		// Verify all elements are strings
		if _, isStringElement := listType.ElemType.(basetypes.StringType); isStringElement {
			return // List of strings is valid
		}
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid list element type",
			"When using a list, all elements must be strings containing path rules")
		return
	}

	// Check if it's a tuple type (HCL list syntax like ["/*", "!/api/*"])
	if tupleType, isTuple := valueType.(basetypes.TupleType); isTuple {
		// Verify all elements are strings
		for i, elemType := range tupleType.ElemTypes {
			if _, isString := elemType.(basetypes.StringType); !isString {
				resp.Diagnostics.AddAttributeError(req.Path, "Invalid tuple element type",
					fmt.Sprintf("Element at index %d must be a string, got %T", i, elemType))
				return
			}
		}
		return // Tuple of strings is valid
	}

	// If we get here, the type is not supported
	resp.Diagnostics.AddAttributeError(req.Path, "Invalid type for run_worker_first",
		fmt.Sprintf("Expected boolean or list/tuple of strings, got %T", valueType))
}
