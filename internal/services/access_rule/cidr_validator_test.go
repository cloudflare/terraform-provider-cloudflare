package access_rule

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestCIDRValidator(t *testing.T) {
	testCases := map[string]bool{
		"192.168.0.1/32":           false,
		"192.168.0.1/24":           true,
		"192.168.0.1/31":           false,
		"192.168.0.1/16":           true,
		"fd82:0f75:cf0d:d7b3::/64": true,
		"fd82:0f75:cf0d:d7b3::/48": true,
		"fd82:0f75:cf0d:d7b3::/32": true,
		"fd82:0f75:cf0d:d7b3::/63": false,
		"fd82:0f75:cf0d:d7b3::/16": false,
	}

	v := cidrValidator()

	for cidr, expectValid := range testCases {
		t.Run(cidr, func(t *testing.T) {
			req := validator.StringRequest{
				Path:        path.Root("test"),
				ConfigValue: types.StringValue(cidr),
			}
			resp := &validator.StringResponse{}

			v.ValidateString(context.Background(), req, resp)

			hasError := resp.Diagnostics.HasError()
			if expectValid && hasError {
				t.Errorf("expected %q to be valid, but got error: %s", cidr, resp.Diagnostics.Errors())
			}
			if !expectValid && !hasError {
				t.Errorf("expected %q to be invalid, but no error was returned", cidr)
			}
		})
	}
}
