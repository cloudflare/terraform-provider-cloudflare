package workers_script

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUpgradeStateFromV0BindingsInherit(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name  string
		value types.String
	}{
		{name: "null", value: types.StringNull()},
		{name: "known", value: types.StringValue("strict")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			priorSchema := resourceSchemaV0(ctx)
			priorType := priorSchema.Type().TerraformType(ctx).(tftypes.Object)
			nullAttributes := func(objectType tftypes.Object) map[string]tftypes.Value {
				values := make(map[string]tftypes.Value, len(objectType.AttributeTypes))
				for name, attributeType := range objectType.AttributeTypes {
					values[name] = tftypes.NewValue(attributeType, nil)
				}
				return values
			}
			values := nullAttributes(priorType)
			inheritValue, err := tc.value.ToTerraformValue(ctx)
			if err != nil {
				t.Fatalf("encode bindings_inherit: %v", err)
			}
			values["bindings_inherit"] = inheritValue
			assetsType := priorType.AttributeTypes["assets"].(tftypes.Object)
			assets := nullAttributes(assetsType)
			configType := assetsType.AttributeTypes["config"].(tftypes.Object)
			config := nullAttributes(configType)
			config["run_worker_first"] = tftypes.NewValue(tftypes.Bool, false)
			assets["config"] = tftypes.NewValue(configType, config)
			values["assets"] = tftypes.NewValue(assetsType, assets)

			req := resource.UpgradeStateRequest{State: &tfsdk.State{
				Schema: *priorSchema,
				Raw:    tftypes.NewValue(priorType, values),
			}}
			resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
			upgradeStateFromV0(ctx, req, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("upgrade failed: %v", resp.Diagnostics)
			}

			var upgraded WorkersScriptModel
			if diags := resp.State.Get(ctx, &upgraded); diags.HasError() {
				t.Fatalf("decode upgraded state: %v", diags)
			}
			if !upgraded.BindingsInherit.Equal(tc.value) {
				t.Errorf("bindings_inherit = %v, want %v", upgraded.BindingsInherit, tc.value)
			}
			if upgraded.Assets == nil || upgraded.Assets.Config == nil {
				t.Fatal("assets config was not preserved")
			}
			if value := upgraded.Assets.Config.RunWorkerFirst.UnderlyingValue(); value == nil || !value.Equal(types.BoolValue(false)) {
				t.Errorf("run_worker_first = %v, want dynamic false", value)
			}
		})
	}
}
