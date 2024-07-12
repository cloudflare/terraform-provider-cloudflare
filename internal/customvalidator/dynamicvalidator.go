package customvalidator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Dynamic = subtypesValidator{}

type subtypesValidator struct {
	allowedTypes []attr.Type
}

func (v subtypesValidator) ValidateDynamic(ctx context.Context, req validator.DynamicRequest, resp *validator.DynamicResponse) {
	dynamic := req.ConfigValue
	if dynamic.IsNull() || dynamic.IsUnknown() || dynamic.IsUnderlyingValueNull() || dynamic.IsUnderlyingValueUnknown() {
		return
	}
	value := dynamic.UnderlyingValue()
	for _, t := range v.allowedTypes {
		if value.Type(ctx) == t {
			return
		}
	}
	resp.Diagnostics.AddAttributeError(req.Path, "Value is not one of the allowed types", v.Description(ctx))
}

func (v subtypesValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v subtypesValidator) MarkdownDescription(_ context.Context) string {
	var s []string
	for _, t := range v.allowedTypes {
		s = append(s, t.String())
	}
	return fmt.Sprintf("The following types are allowed: %s", strings.Join(s, ", "))
}

func AllowedSubtypes(types ...attr.Type) validator.Dynamic {
	return subtypesValidator{
		allowedTypes: types,
	}
}
