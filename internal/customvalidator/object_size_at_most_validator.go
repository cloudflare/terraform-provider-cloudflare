package customvalidator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type objectSizeAtMostValidator struct {
	max int
}

func (v objectSizeAtMostValidator) Description(_ context.Context) string {
	return fmt.Sprintf("must contain at most %d elements", v.max)
}

func (v objectSizeAtMostValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v objectSizeAtMostValidator) ValidateObject(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	nonNullCount := 0
	for _, attr := range req.ConfigValue.Attributes() {
		if !attr.IsNull() && !attr.IsUnknown() {
			nonNullCount++
		}
	}
	if nonNullCount > v.max {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			v.Description(ctx),
			fmt.Sprintf("%d", nonNullCount),
		))
	}
}

// ObjectSizeAtMost returns an AttributeValidator which ensures that any configured
// attribute or function parameter value:
//
//   - Is an object.
//   - Contains at most max elements.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func ObjectSizeAtMost(maxVal int) objectSizeAtMostValidator {
	return objectSizeAtMostValidator{
		max: maxVal,
	}
}
