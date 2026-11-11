package zero_trust_access_application_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_application"
)

// This resource's schema is hand-maintained, so API enum changes are not picked up by
// codegen. The values mirror the authenticators the Access API accepts; which values
// each application type accepts is enforced by the API. ssh_piv_key was renamed to
// piv_key and the API rejects it.
func TestMFAAllowedAuthenticatorsMatchAPI(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	s := zero_trust_access_application.ResourceSchema(ctx)
	policies, ok := s.Attributes["policies"].(schema.ListNestedAttribute)
	if !ok {
		t.Fatal("policies is not a ListNestedAttribute")
	}

	for name, attrs := range map[string]map[string]schema.Attribute{
		"mfa_config":            s.Attributes,
		"policies[].mfa_config": policies.NestedObject.Attributes,
	} {
		t.Run(name, func(t *testing.T) {
			mfa, ok := attrs["mfa_config"].(schema.SingleNestedAttribute)
			if !ok {
				t.Fatalf("%s is not a SingleNestedAttribute", name)
			}
			list, ok := mfa.Attributes["allowed_authenticators"].(schema.ListAttribute)
			if !ok {
				t.Fatalf("%s.allowed_authenticators is not a ListAttribute", name)
			}

			for value, want := range map[string]bool{
				"totp":          true,
				"biometrics":    true,
				"security_key":  true,
				"piv_key":       true,
				"ssh_fido2_key": true,
				"ssh_piv_key":   false,
			} {
				config := types.ListValueMust(types.StringType, []attr.Value{types.StringValue(value)})
				var hasError bool
				for _, v := range list.Validators {
					resp := &validator.ListResponse{}
					v.ValidateList(ctx, validator.ListRequest{Path: path.Root("allowed_authenticators"), ConfigValue: config}, resp)
					hasError = hasError || resp.Diagnostics.HasError()
				}
				if hasError == want {
					t.Errorf("%s.allowed_authenticators accepts %q = %t, want %t", name, value, !hasError, want)
				}
			}
		})
	}
}
