// This file is intentionally NOT code-generated.

package list_item_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item"
)

// The Lists API stores IPs and CIDRs in normalized form, so an unnormalized
// value such as 192.168.1.1/32 used to fail after create with "list item
// pagination did not return a matching list item". The ip attribute shares
// cloudflare_list's validator, which rejects such values at plan time.
func TestResourceSchemaIPRejectsUnnormalizedValues(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	attribute, ok := list_item.ResourceSchema(ctx).Attributes["ip"].(schema.StringAttribute)
	if !ok {
		t.Fatal("ip is not a string attribute")
	}

	tests := map[string]bool{
		"192.168.1.1":     true,
		"10.1.1.0/24":     true,
		"2001:db8::1":     true,
		"2001:db8::/48":   true,
		"192.168.1.1/32":  false,
		"2001:db8::1/128": false,
		"10.1.1.5/24":     false,
		"2001:DB8::1":     false,
	}

	for value, valid := range tests {
		t.Run(value, func(t *testing.T) {
			t.Parallel()
			hasError := false
			for _, v := range attribute.Validators {
				resp := &validator.StringResponse{}
				v.ValidateString(ctx, validator.StringRequest{
					Path:        path.Root("ip"),
					ConfigValue: types.StringValue(value),
				}, resp)
				hasError = hasError || resp.Diagnostics.HasError()
			}
			if hasError == valid {
				t.Fatalf("ip %q: validation error = %t, want %t", value, hasError, !valid)
			}
		})
	}
}
