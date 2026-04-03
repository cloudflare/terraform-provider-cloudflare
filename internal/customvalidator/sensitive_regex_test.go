package customvalidator

import (
	"context"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestSensitiveRegexMatches_ValidateString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		value       types.String
		regex       *regexp.Regexp
		message     string
		expectError bool
	}{
		{
			name:        "matching value passes",
			value:       types.StringValue("abc123"),
			regex:       regexp.MustCompile(`^[a-z0-9]+$`),
			message:     "must be lowercase alphanumeric",
			expectError: false,
		},
		{
			name:        "non-matching value fails",
			value:       types.StringValue("ABC!@#"),
			regex:       regexp.MustCompile(`^[a-z0-9]+$`),
			message:     "must be lowercase alphanumeric",
			expectError: true,
		},
		{
			name:        "null value skips validation",
			value:       types.StringNull(),
			regex:       regexp.MustCompile(`^[a-z]+$`),
			message:     "must be lowercase",
			expectError: false,
		},
		{
			name:        "unknown value skips validation",
			value:       types.StringUnknown(),
			regex:       regexp.MustCompile(`^[a-z]+$`),
			message:     "must be lowercase",
			expectError: false,
		},
		{
			name:        "old format api key passes",
			value:       types.StringValue("fake0key000000000000000000000000000000"),
			regex:       regexp.MustCompile(`^[0-9A-Za-z\-_]{37,60}$`),
			message:     "API key must only contain characters 0-9, a-z, A-Z, hyphens and underscores",
			expectError: false,
		},
		{
			name:        "new prefixed api key passes",
			value:       types.StringValue("test_fake0key0aaaa0bbbb0cccc0dddd0eeee0ffff0gggg0hhhh"),
			regex:       regexp.MustCompile(`^[0-9A-Za-z\-_]{37,60}$`),
			message:     "API key must only contain characters 0-9, a-z, A-Z, hyphens and underscores",
			expectError: false,
		},
		{
			name:        "api key too short fails",
			value:       types.StringValue("tooshort"),
			regex:       regexp.MustCompile(`^[0-9A-Za-z\-_]{37,60}$`),
			message:     "API key must only contain characters 0-9, a-z, A-Z, hyphens and underscores",
			expectError: true,
		},
		{
			name:        "api key with invalid characters fails",
			value:       types.StringValue("this-has-special!chars@that#fail$regex00000000000"),
			regex:       regexp.MustCompile(`^[0-9A-Za-z\-_]{37,60}$`),
			message:     "API key must only contain characters 0-9, a-z, A-Z, hyphens and underscores",
			expectError: true,
		},
		{
			name:        "old format api token passes",
			value:       types.StringValue("fake0token0FAKE0TOKEN0aaa0bbb0ccc0ddd0eee"),
			regex:       regexp.MustCompile(`^[0-9A-Za-z\-_]{40,80}$`),
			message:     "API tokens must only contain characters a-z, A-Z, 0-9, hyphens and underscores",
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			v := NewSensitiveRegexMatchesValidator(tc.regex, tc.message)

			req := validator.StringRequest{
				Path:        path.Root("test_attr"),
				ConfigValue: tc.value,
			}
			resp := &validator.StringResponse{}

			v.ValidateString(context.Background(), req, resp)

			if tc.expectError && !resp.Diagnostics.HasError() {
				t.Error("expected validation error but got none")
			}
			if !tc.expectError && resp.Diagnostics.HasError() {
				t.Errorf("expected no error but got: %s", resp.Diagnostics.Errors())
			}
		})
	}
}

func TestSensitiveRegexMatches_ErrorDoesNotContainValue(t *testing.T) {
	t.Parallel()

	sensitiveValue := "test_FAKE_sensitive_value_0000000000000000"
	v := NewSensitiveRegexMatchesValidator(
		regexp.MustCompile(`^[a-z]{10}$`),
		"must be exactly 10 lowercase letters",
	)

	req := validator.StringRequest{
		Path:        path.Root("api_key"),
		ConfigValue: types.StringValue(sensitiveValue),
	}
	resp := &validator.StringResponse{}

	v.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected validation error but got none")
	}

	for _, d := range resp.Diagnostics.Errors() {
		if strings.Contains(d.Summary(), sensitiveValue) {
			t.Errorf("error summary contains the sensitive value: %s", d.Summary())
		}
		if strings.Contains(d.Detail(), sensitiveValue) {
			t.Errorf("error detail contains the sensitive value: %s", d.Detail())
		}
	}
}

func TestSensitiveRegexMatches_Description(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("custom message", func(t *testing.T) {
		t.Parallel()

		v := SensitiveRegexMatches{
			regexp:  regexp.MustCompile(`^[a-z]+$`),
			message: "must be lowercase letters",
		}

		expected := "must be lowercase letters"
		if got := v.Description(ctx); got != expected {
			t.Errorf("Description() = %q, want %q", got, expected)
		}
		if got := v.MarkdownDescription(ctx); got != expected {
			t.Errorf("MarkdownDescription() = %q, want %q", got, expected)
		}
	})

	t.Run("default message", func(t *testing.T) {
		t.Parallel()

		v := SensitiveRegexMatches{
			regexp:  regexp.MustCompile(`^[a-z]+$`),
			message: "",
		}

		expected := "value must match regular expression '^[a-z]+$'"
		if got := v.Description(ctx); got != expected {
			t.Errorf("Description() = %q, want %q", got, expected)
		}
	})
}
