package worker_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestWorkerModifyPlanPreviewsBaseConfig(t *testing.T) {
	ctx := context.Background()
	preview := func(logpush types.Bool, limits types.Int64) customfield.NestedObject[worker.WorkerPreviewsBaseConfigModel] {
		return customfield.NewObjectMust(ctx, &worker.WorkerPreviewsBaseConfigModel{
			Logpush: logpush,
			Limits:  customfield.NewObjectMust(ctx, &worker.WorkerPreviewsBaseConfigLimitsModel{CPUMs: limits}),
		})
	}
	prior := preview(types.BoolValue(false), types.Int64Value(100))
	changed := preview(types.BoolValue(true), types.Int64Value(100))
	null := customfield.NullObject[worker.WorkerPreviewsBaseConfigModel](ctx)
	unknown := customfield.UnknownObject[worker.WorkerPreviewsBaseConfigModel](ctx)
	computedModel, diags := prior.Value(ctx)
	if diags.HasError() {
		t.Fatal(diags)
	}
	computedModel.CacheOptions = customfield.UnknownObject[worker.WorkerPreviewsBaseConfigCacheOptionsModel](ctx)
	computedUnknown := customfield.NewObjectMust(ctx, computedModel)
	computedModel.Logpush = types.BoolValue(true)
	changedWithUnknown := customfield.NewObjectMust(ctx, computedModel)
	computedModel.Logpush = types.BoolValue(false)
	computedModel.CacheOptions = customfield.NewObjectMust(ctx, &worker.WorkerPreviewsBaseConfigCacheOptionsModel{
		Enabled: types.BoolValue(false), CrossVersionCache: types.BoolValue(false),
	})
	cachePrior := customfield.NewObjectMust(ctx, computedModel)
	computedModel.CacheOptions = customfield.NewObjectMust(ctx, &worker.WorkerPreviewsBaseConfigCacheOptionsModel{
		Enabled: types.BoolValue(false), CrossVersionCache: types.BoolUnknown(),
	})
	unknownCacheLeaf := customfield.NewObjectMust(ctx, computedModel)

	for _, tt := range []struct {
		name                string
		prior, plan, config customfield.NestedObject[worker.WorkerPreviewsBaseConfigModel]
		wantUpdate          bool
		checkPayload        bool
	}{
		{"preview-only edit", prior, changed, changed, true, true},
		{"nested limit edit", prior, preview(types.BoolValue(false), types.Int64Value(200)), preview(types.BoolValue(false), types.Int64Value(200)), true, true},
		{"add preview configuration", null, changed, changed, true, true},
		{"remove planned preview configuration", prior, null, null, true, false},
		{"unchanged configuration", prior, prior, prior, false, false},
		{"unknown configured block", prior, unknown, unknown, true, false},
		{"unknown configured leaf", prior, preview(types.BoolUnknown(), types.Int64Value(100)), preview(types.BoolUnknown(), types.Int64Value(100)), true, false},
		{"unknown computed block", prior, unknown, null, false, false},
		{"unknown computed descendant", prior, computedUnknown, prior, false, false},
		{"unknown computed descendant with populated state", cachePrior, computedUnknown, prior, false, false},
		{"unknown computed leaf with populated state", cachePrior, unknownCacheLeaf, prior, false, false},
		{"known edit with unknown computed descendant", prior, changedWithUnknown, changed, true, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			model := worker.WorkerModel{
				ID: types.StringValue("worker-id"), AccountID: types.StringValue("account-id"), Name: types.StringValue("worker"),
				Logpush: types.BoolValue(false), PreviewsBaseConfig: tt.prior,
				UpdatedOn: timetypes.NewRFC3339ValueMust("2026-10-08T10:00:00Z"),
			}
			schema := worker.ResourceSchema(ctx)
			state := tfsdk.State{Schema: schema}
			if diags := state.Set(ctx, &model); diags.HasError() {
				t.Fatalf("set state: %v", diags)
			}
			planned := model
			planned.PreviewsBaseConfig = tt.plan
			planned.UpdatedOn = timetypes.NewRFC3339Unknown()
			planned.References = customfield.UnknownObject[worker.WorkerReferencesModel](ctx)
			planned.Subdomain = customfield.UnknownObject[worker.WorkerSubdomainModel](ctx)
			planned.Observability = customfield.UnknownObject[worker.WorkerObservabilityModel](ctx)
			plan := tfsdk.Plan{Schema: schema}
			if diags := plan.Set(ctx, &planned); diags.HasError() {
				t.Fatalf("set plan: %v", diags)
			}
			configured := model
			configured.PreviewsBaseConfig = tt.config
			configState := tfsdk.State{Schema: schema}
			if diags := configState.Set(ctx, &configured); diags.HasError() {
				t.Fatalf("set config: %v", diags)
			}
			response := resource.ModifyPlanResponse{Plan: plan}
			worker.NewResource().(*worker.WorkerResource).ModifyPlan(ctx, resource.ModifyPlanRequest{
				Plan: plan, State: state, Config: tfsdk.Config{Schema: schema, Raw: configState.Raw},
			}, &response)
			if response.Diagnostics.HasError() {
				t.Fatalf("modify plan: %v", response.Diagnostics)
			}
			var result worker.WorkerModel
			if diags := response.Plan.Get(ctx, &result); diags.HasError() {
				t.Fatalf("get plan: %v", diags)
			}
			want := tt.prior
			if tt.wantUpdate {
				want = tt.plan
			}
			if !result.PreviewsBaseConfig.Equal(want) {
				t.Errorf("preview config was not preserved: got %s, want %s", result.PreviewsBaseConfig, want)
			}
			if result.UpdatedOn.IsUnknown() != tt.wantUpdate {
				t.Errorf("updated_on unknown = %v, want %v", result.UpdatedOn.IsUnknown(), tt.wantUpdate)
			}
			if !tt.wantUpdate && !response.Plan.Raw.Equal(state.Raw) {
				t.Error("unchanged worker still plans an update")
			}
			if tt.checkPayload {
				payload, err := result.MarshalJSONForUpdate(model)
				if err != nil {
					t.Fatalf("serialize update: %v", err)
				}
				var body map[string]json.RawMessage
				if err := json.Unmarshal(payload, &body); err != nil {
					t.Fatal(err)
				}
				previewJSON, ok := body["previews_base_config"]
				if !ok {
					t.Fatalf("preview-only update missing from payload: %s", payload)
				}
				var actual struct {
					Logpush bool `json:"logpush"`
					Limits  struct {
						CPUMs int64 `json:"cpu_ms"`
					} `json:"limits"`
				}
				if err := json.Unmarshal(previewJSON, &actual); err != nil {
					t.Fatal(err)
				}
				expected, diags := tt.plan.Value(ctx)
				if diags.HasError() {
					t.Fatal(diags)
				}
				limits, diags := expected.Limits.Value(ctx)
				if diags.HasError() {
					t.Fatal(diags)
				}
				if actual.Logpush != expected.Logpush.ValueBool() || actual.Limits.CPUMs != limits.CPUMs.ValueInt64() {
					t.Errorf("update payload did not retain planned values: %s", previewJSON)
				}
			}
		})
	}
}

func TestWorkerModifyPlanCreateAndDestroy(t *testing.T) {
	ctx := context.Background()
	schema := worker.ResourceSchema(ctx)
	model := worker.WorkerModel{AccountID: types.StringValue("account-id"), Name: types.StringValue("worker")}
	plan := tfsdk.Plan{Schema: schema}
	if diags := plan.Set(ctx, &model); diags.HasError() {
		t.Fatal(diags)
	}
	state := tfsdk.State{Schema: schema, Raw: plan.Raw}
	nullPlan := tfsdk.Plan{Schema: schema}
	if diags := nullPlan.Set(ctx, (*worker.WorkerModel)(nil)); diags.HasError() {
		t.Fatal(diags)
	}
	for _, tt := range []struct {
		name  string
		plan  tfsdk.Plan
		state tfsdk.State
	}{
		{"create", plan, tfsdk.State{Schema: schema, Raw: nullPlan.Raw}},
		{"destroy", nullPlan, state},
	} {
		t.Run(tt.name, func(t *testing.T) {
			response := resource.ModifyPlanResponse{Plan: tt.plan}
			worker.NewResource().(*worker.WorkerResource).ModifyPlan(ctx, resource.ModifyPlanRequest{Plan: tt.plan, State: tt.state}, &response)
			if response.Diagnostics.HasError() {
				t.Fatal(response.Diagnostics)
			}
			if !response.Plan.Raw.Equal(tt.plan.Raw) {
				t.Fatal("create/destroy plan was modified")
			}
		})
	}
}
