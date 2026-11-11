package worker_version_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_version"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The exports schema must carry every field of WorkerVersionExportsModel, or
// State.Set fails for any config that sets exports (#7416).
func TestWorkerVersionExportsStateRoundTrip(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	exports := map[string]worker_version.WorkerVersionExportsModel{
		"default": {Type: types.StringValue("worker")},
		"Counter": {
			Type:      types.StringValue("durable-object"),
			Storage:   types.StringValue("sqlite"),
			Container: types.StringValue("my-container"),
			State:     types.StringValue("created"),
		},
	}
	model := worker_version.WorkerVersionModel{
		AccountID:             types.StringValue("account-123"),
		WorkerID:              types.StringValue("worker-123"),
		Exports:               &exports,
		CompatibilityFlags:    customfield.NullSet[types.String](ctx),
		Annotations:           customfield.NullObject[worker_version.WorkerVersionAnnotationsModel](ctx),
		Assets:                customfield.NullObject[worker_version.WorkerVersionAssetsModel](ctx),
		Bindings:              customfield.NullObjectList[worker_version.WorkerVersionBindingsModel](ctx),
		CacheOptions:          customfield.NullObject[worker_version.WorkerVersionCacheOptionsModel](ctx),
		Limits:                customfield.NullObject[worker_version.WorkerVersionLimitsModel](ctx),
		URLs:                  customfield.NullList[types.String](ctx),
		ExportsReconciliation: customfield.NullObject[worker_version.WorkerVersionExportsReconciliationModel](ctx),
	}

	state := tfsdk.State{Schema: worker_version.ResourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("State.Set: %v", diags)
	}

	var got worker_version.WorkerVersionModel
	if diags := state.Get(ctx, &got); diags.HasError() {
		t.Fatalf("State.Get: %v", diags)
	}
	if got.Exports == nil {
		t.Fatal("exports is nil after round trip")
	}
	if c := (*got.Exports)["Counter"].Container.ValueString(); c != "my-container" {
		t.Errorf("exports.Counter.container = %q, want %q", c, "my-container")
	}
	if typ := (*got.Exports)["default"].Type.ValueString(); typ != "worker" {
		t.Errorf("exports.default.type = %q, want %q", typ, "worker")
	}
}
