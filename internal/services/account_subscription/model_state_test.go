// This file is intentionally NOT code-generated.

package account_subscription_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_subscription"
)

// Every tfsdk tag on the model must be a schema attribute. A stray zone_id
// field made every Read fail with a Value Conversion Error (APIX-1934), and
// the parity test did not catch it.
func TestAccountSubscriptionModelFitsState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	s := account_subscription.ResourceSchema(ctx)
	state := tfsdk.State{Schema: s, Raw: tftypes.NewValue(s.Type().TerraformType(ctx), nil)}
	if diags := state.Set(ctx, &account_subscription.AccountSubscriptionModel{}); diags.HasError() {
		t.Fatalf("model does not fit the resource schema: %v", diags)
	}
}
