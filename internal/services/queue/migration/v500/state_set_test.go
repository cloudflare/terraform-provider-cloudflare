// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4QueueState = `{
	"id": "023e105f4ecef8ad9ca31a8372d0c353",
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"name": "my-queue"
}`

// The target model must convert into the live schema. consumers[].settings
// carried email, pagerduty and webhooks, which the schema does not have, so
// Set failed even with consumers null.
func TestQueueUpgradeFromLegacyV0SetsState(t *testing.T) {
	ctx := context.Background()
	sourceSchema := v500.SourceCloudflareQueueSchema()
	raw, err := (&tfprotov6.RawState{JSON: []byte(v4QueueState)}).Unmarshal(sourceSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("could not decode v4 state: %v", err)
	}

	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: queue.ResourceSchema(ctx)}}
	v500.UpgradeFromLegacyV0(ctx, resource.UpgradeStateRequest{State: &tfsdk.State{Raw: raw, Schema: sourceSchema}}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("UpgradeFromLegacyV0 failed: %v", resp.Diagnostics)
	}

	var got queue.QueueModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("could not decode upgraded state: %v", diags)
	}
	if !got.QueueName.Equal(types.StringValue("my-queue")) {
		t.Errorf("queue_name = %s, want \"my-queue\"", got.QueueName)
	}
	if !got.QueueID.Equal(types.StringValue("023e105f4ecef8ad9ca31a8372d0c353")) {
		t.Errorf("queue_id = %s", got.QueueID)
	}
	if !got.Consumers.IsNull() {
		t.Errorf("consumers = %s, want null", got.Consumers)
	}
}
