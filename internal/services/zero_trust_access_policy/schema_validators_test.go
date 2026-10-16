package zero_trust_access_policy_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_policy"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestZeroTrustAccessPolicyRDPClipboardFormatsAcceptFile(t *testing.T) {
	t.Parallel()
	s := zero_trust_access_policy.ResourceSchema(context.TODO())
	rules := s.Attributes["connection_rules"].(schema.SingleNestedAttribute)
	rdp := rules.Attributes["rdp"].(schema.SingleNestedAttribute)

	for _, name := range []string{"allowed_clipboard_local_to_remote_formats", "allowed_clipboard_remote_to_local_formats"} {
		list := rdp.Attributes[name].(schema.ListAttribute)
		val := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("text"), types.StringValue("file")})
		for _, v := range list.Validators {
			resp := &validator.ListResponse{}
			v.ValidateList(context.TODO(), validator.ListRequest{ConfigValue: val}, resp)
			if resp.Diagnostics.HasError() {
				t.Errorf("%s: unexpected error: %v", name, resp.Diagnostics)
			}
		}
	}
}
