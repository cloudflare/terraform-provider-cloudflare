// This file is intentionally NOT code-generated.

package worker_version_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_version"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// A v0 state with assets.config set must decode with the v0 prior schema and
// upgrade into the live schema. The v0 schema used to carry the latest
// assets.config type, with a dynamic run_worker_first, which did not match
// resourceModelV0.
func TestUpgradeWorkerVersionFromV0WithAssetsConfig(t *testing.T) {
	for _, tc := range []struct {
		name           string
		runWorkerFirst string
		want           []string
	}{
		{name: "run_worker_first_list", runWorkerFirst: `["/api/*", "!/api/static/*"]`, want: []string{"/api/*", "!/api/static/*"}},
		{name: "run_worker_first_null", runWorkerFirst: `null`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			upgrader := worker_version.NewResource().(resource.ResourceWithUpgradeState).UpgradeState(ctx)[0]
			rawState := &tfprotov6.RawState{JSON: []byte(fmt.Sprintf(`{
				"id": "11111111-1111-4111-8111-111111111111",
				"account_id": "f037e56e89293a057740de681ac9abbe",
				"worker_id": "my-worker",
				"main_module": "index.js",
				"compatibility_date": "2025-01-01",
				"assets": {
					"config": {
						"html_handling": "auto-trailing-slash",
						"not_found_handling": "single-page-application",
						"run_worker_first": %s
					},
					"jwt": null,
					"directory": "./public",
					"asset_manifest_sha256": "abc123"
				},
				"number": 1,
				"source": "terraform"
			}`, tc.runWorkerFirst))}
			raw, err := rawState.Unmarshal(upgrader.PriorSchema.Type().TerraformType(ctx))
			if err != nil {
				t.Fatalf("could not decode v0 state with the v0 prior schema: %v", err)
			}

			resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: worker_version.ResourceSchema(ctx)}}
			upgrader.StateUpgrader(ctx, resource.UpgradeStateRequest{
				RawState: rawState,
				State:    &tfsdk.State{Raw: raw, Schema: *upgrader.PriorSchema},
			}, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("upgrade failed: %v", resp.Diagnostics)
			}

			var got worker_version.WorkerVersionModel
			if diags := resp.State.Get(ctx, &got); diags.HasError() {
				t.Fatalf("could not decode upgraded state: %v", diags)
			}
			assets, diags := got.Assets.Value(ctx)
			if diags.HasError() || assets == nil {
				t.Fatalf("assets missing after upgrade: %v", diags)
			}
			if !assets.Directory.Equal(types.StringValue("./public")) {
				t.Errorf("assets.directory = %s, want \"./public\"", assets.Directory)
			}
			config, diags := assets.Config.Value(ctx)
			if diags.HasError() || config == nil {
				t.Fatalf("assets.config missing after upgrade: %v", diags)
			}
			if !config.HTMLHandling.Equal(types.StringValue("auto-trailing-slash")) {
				t.Errorf("assets.config.html_handling = %s", config.HTMLHandling)
			}
			if !config.BasePath.IsNull() {
				t.Errorf("assets.config.base_path = %s, want null", config.BasePath)
			}

			value := config.RunWorkerFirst.UnderlyingValue()
			if tc.want == nil {
				if value != nil && !value.IsNull() {
					t.Errorf("assets.config.run_worker_first = %s, want null", value)
				}
				return
			}
			list, ok := value.(types.List)
			if !ok {
				t.Fatalf("assets.config.run_worker_first = %T %v, want a list", value, value)
			}
			var rules []string
			if diags := list.ElementsAs(ctx, &rules, false); diags.HasError() {
				t.Fatalf("could not read run_worker_first: %v", diags)
			}
			if fmt.Sprint(rules) != fmt.Sprint(tc.want) {
				t.Errorf("assets.config.run_worker_first = %v, want %v", rules, tc.want)
			}
		})
	}
}
