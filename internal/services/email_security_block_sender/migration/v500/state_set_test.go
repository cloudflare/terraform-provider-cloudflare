// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_block_sender"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_block_sender/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const priorBlockSenderState = `{
	"id": 12345,
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"is_regex": false,
	"pattern": "spam@example.com",
	"pattern_type": "EMAIL",
	"comments": "known spammer",
	"created_at": "2024-01-01T00:00:00Z",
	"last_modified": "2024-01-02T00:00:00Z"
}`

// The target model must convert into the live schema, which gained expires_at,
// or UpgradeState fails with a Value Conversion Error.
func TestEmailSecurityBlockSenderUpgradeFromV0SetsState(t *testing.T) {
	ctx := context.Background()
	priorSchema := v500.PriorSchema()
	raw, err := (&tfprotov6.RawState{JSON: []byte(priorBlockSenderState)}).Unmarshal(priorSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("could not decode prior state: %v", err)
	}

	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: email_security_block_sender.ResourceSchema(ctx)}}
	v500.UpgradeFromV0(ctx, resource.UpgradeStateRequest{State: &tfsdk.State{Raw: raw, Schema: priorSchema}}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("UpgradeFromV0 failed: %v", resp.Diagnostics)
	}

	var got email_security_block_sender.EmailSecurityBlockSenderModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("could not decode upgraded state: %v", diags)
	}
	if !got.ID.Equal(types.StringValue("12345")) {
		t.Errorf("id = %s, want \"12345\"", got.ID)
	}
	if !got.Pattern.Equal(types.StringValue("spam@example.com")) {
		t.Errorf("pattern = %s", got.Pattern)
	}
	if !got.ExpiresAt.IsNull() {
		t.Errorf("expires_at = %s, want null", got.ExpiresAt)
	}
}
