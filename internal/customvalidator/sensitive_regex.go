package customvalidator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// SensitiveRegexMatches validates that a string attribute matches the given
// regex pattern, but unlike stringvalidator.RegexMatches, it does NOT include
// the raw attribute value in validation error messages. This is critical for
// sensitive fields like API keys and tokens where leaking the value in error
// output would be a security concern.
type SensitiveRegexMatches struct {
	regexp  *regexp.Regexp
	message string
}

func (v SensitiveRegexMatches) Description(ctx context.Context) string {
	if v.message != "" {
		return v.message
	}
	return fmt.Sprintf("value must match regular expression '%s'", v.regexp)
}

func (v SensitiveRegexMatches) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v SensitiveRegexMatches) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if !v.regexp.MatchString(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Attribute Value",
			fmt.Sprintf("Attribute %s %s", req.Path, v.Description(ctx)),
		)
	}
}

// NewSensitiveRegexMatchesValidator returns a validator that checks the string
// matches the given regex, without leaking the raw value in error messages.
func NewSensitiveRegexMatchesValidator(r *regexp.Regexp, message string) validator.String {
	return SensitiveRegexMatches{
		regexp:  r,
		message: message,
	}
}
