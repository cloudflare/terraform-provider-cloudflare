// This file is intentionally NOT code-generated.

package zero_trust_access_policy_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_policy"
)

// app_count changes whenever an application attaches or detaches the policy,
// outside this resource, so it must stay out of the resource schema (#5608).
// It remains in the data sources.
func TestResourceSchemaOmitsAppCount(t *testing.T) {
	t.Parallel()

	if _, ok := zero_trust_access_policy.ResourceSchema(context.Background()).Attributes["app_count"]; ok {
		t.Fatal("app_count must not be in the zero_trust_access_policy resource schema (#5608)")
	}
}

// The API accepts "file" as an RDP clipboard format. A reseal once reverted the
// resource validator to "text" only.
func TestResourceSchemaAcceptsRDPFileClipboardFormat(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	connectionRules, ok := zero_trust_access_policy.ResourceSchema(ctx).Attributes["connection_rules"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatal("connection_rules is not a single nested attribute")
	}
	rdp, ok := connectionRules.Attributes["rdp"].(schema.SingleNestedAttribute)
	if !ok {
		t.Fatal("connection_rules.rdp is not a single nested attribute")
	}

	value := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("text"), types.StringValue("file")})

	for _, name := range []string{"allowed_clipboard_local_to_remote_formats", "allowed_clipboard_remote_to_local_formats"} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			attribute, ok := rdp.Attributes[name].(schema.ListAttribute)
			if !ok {
				t.Fatalf("%s is not a list attribute", name)
			}
			if len(attribute.Validators) == 0 {
				t.Fatal("expected a validator on the clipboard formats")
			}
			for _, v := range attribute.Validators {
				resp := &validator.ListResponse{}
				v.ValidateList(ctx, validator.ListRequest{
					Path:        path.Root("connection_rules").AtName("rdp").AtName(name),
					ConfigValue: value,
				}, resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf(`validator rejects ["text", "file"]: %v`, resp.Diagnostics)
				}
			}

			invalid := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("image")})
			rejected := false
			for _, v := range attribute.Validators {
				resp := &validator.ListResponse{}
				v.ValidateList(ctx, validator.ListRequest{
					Path:        path.Root("connection_rules").AtName("rdp").AtName(name),
					ConfigValue: invalid,
				}, resp)
				rejected = rejected || resp.Diagnostics.HasError()
			}
			if !rejected {
				t.Fatal("validator accepts an unknown clipboard format")
			}
		})
	}
}
