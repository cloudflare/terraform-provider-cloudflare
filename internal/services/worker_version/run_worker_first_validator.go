package worker_version

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ validator.Dynamic = runWorkerFirstValidator{}

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

	if _, isBool := valueType.(basetypes.BoolType); isBool {
		return
	}

	if listType, isList := valueType.(basetypes.ListType); isList {
		if _, isStringElement := listType.ElemType.(basetypes.StringType); isStringElement {
			return
		}
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid list element type",
			"When using a list, all elements must be strings containing path rules")
		return
	}

	if tupleType, isTuple := valueType.(basetypes.TupleType); isTuple {
		for i, elemType := range tupleType.ElemTypes {
			if _, isString := elemType.(basetypes.StringType); !isString {
				resp.Diagnostics.AddAttributeError(req.Path, "Invalid tuple element type",
					fmt.Sprintf("Element at index %d must be a string, got %T", i, elemType))
				return
			}
		}
		return
	}

	resp.Diagnostics.AddAttributeError(req.Path, "Invalid type for run_worker_first",
		fmt.Sprintf("Expected boolean or list/tuple of strings, got %T", valueType))
}
