package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts v4 cloudflare_worker_script state to v5 cloudflare_workers_script state.
//
// Key transformations:
// - name → script_name
// - module (bool) → main_module/body_part (string)
// - 10 separate binding arrays → unified bindings list with type discriminator
// - placement array → placement object
// - tags → removed (deprecated in v5)
// - dispatch_namespace → removed (not supported in v5)
// - Computed fields set to null (refreshed from API on next plan/apply)
func Transform(ctx context.Context, source SourceWorkerScriptModel) (*TargetWorkersScriptModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError("Missing required field", "account_id is required for workers_script migration.")
		return nil, diags
	}

	target := &TargetWorkersScriptModel{
		// Required fields
		AccountID:  source.AccountID,
		ScriptName: source.Name, // name → script_name
		ID:         source.Name, // id matches script_name

		// Content
		Content: source.Content,

		// Computed fields — set to null, will refresh from API
		ContentFile:       types.StringNull(),
		ContentSHA256:     types.StringNull(),
		ContentType:       types.StringNull(),
		CreatedOn:         timetypes.NewRFC3339Null(),
		Etag:              types.StringNull(),
		HasAssets:         types.BoolNull(),
		HasModules:        types.BoolNull(),
		LastDeployedFrom:  types.StringNull(),
		MigrationTag:      types.StringNull(),
		ModifiedOn:        timetypes.NewRFC3339Null(),
		PlacementMode:     types.StringNull(),
		PlacementStatus:   types.StringNull(),
		StartupTimeMs:     types.Int64Null(),
		CompatibilityDate: types.StringNull(),
		Logpush:           types.BoolNull(),
		UsageModel:        types.StringNull(),
	}

	// module (bool) → main_module/body_part (string)
	transformModule(source, target)

	// Consolidate 10 binding arrays → unified bindings list
	bindings := transformBindings(source)
	target.Bindings = newBindingsList(ctx, bindings)

	// placement array → placement object
	transformPlacement(ctx, source, target)

	return target, diags
}

// transformModule converts v4 module boolean to v5 main_module/body_part strings.
func transformModule(source SourceWorkerScriptModel, target *TargetWorkersScriptModel) {
	if source.Module.IsNull() || source.Module.IsUnknown() {
		target.MainModule = types.StringNull()
		target.BodyPart = types.StringNull()
		return
	}

	if source.Module.ValueBool() {
		// module = true → ES module syntax
		target.MainModule = types.StringValue("worker.js")
		target.BodyPart = types.StringNull()
	} else {
		// module = false → service worker syntax
		target.MainModule = types.StringNull()
		target.BodyPart = types.StringValue("worker.js")
	}
}

// transformBindings consolidates 10 v4 binding block arrays into a single v5 bindings list.
func transformBindings(source SourceWorkerScriptModel) []*TargetBindingsModel {
	var bindings []*TargetBindingsModel

	for _, b := range source.PlainTextBinding {
		bindings = append(bindings, newBinding("plain_text", b.Name, func(tb *TargetBindingsModel) {
			tb.Text = b.Text
		}))
	}

	for _, b := range source.SecretTextBinding {
		bindings = append(bindings, newBinding("secret_text", b.Name, func(tb *TargetBindingsModel) {
			tb.Text = b.Text
		}))
	}

	for _, b := range source.KVNamespaceBinding {
		bindings = append(bindings, newBinding("kv_namespace", b.Name, func(tb *TargetBindingsModel) {
			tb.NamespaceID = b.NamespaceID
		}))
	}

	for _, b := range source.WebassemblyBinding {
		bindings = append(bindings, newBinding("wasm_module", b.Name, func(tb *TargetBindingsModel) {
			tb.Part = b.Module // module → part
		}))
	}

	for _, b := range source.ServiceBinding {
		bindings = append(bindings, newBinding("service", b.Name, func(tb *TargetBindingsModel) {
			tb.Service = b.Service
			tb.Environment = b.Environment
		}))
	}

	for _, b := range source.R2BucketBinding {
		bindings = append(bindings, newBinding("r2_bucket", b.Name, func(tb *TargetBindingsModel) {
			tb.BucketName = b.BucketName
		}))
	}

	for _, b := range source.AnalyticsEngineBinding {
		bindings = append(bindings, newBinding("analytics_engine", b.Name, func(tb *TargetBindingsModel) {
			tb.Dataset = b.Dataset
		}))
	}

	// queue_binding: binding→name, queue→queue_name
	for _, b := range source.QueueBinding {
		bindings = append(bindings, newBinding("queue", b.Binding, func(tb *TargetBindingsModel) {
			tb.QueueName = b.Queue
		}))
	}

	// d1_database_binding: database_id→id
	for _, b := range source.D1DatabaseBinding {
		bindings = append(bindings, newBinding("d1", b.Name, func(tb *TargetBindingsModel) {
			tb.ID = b.DatabaseID
		}))
	}

	// hyperdrive_config_binding: binding→name
	for _, b := range source.HyperdriveConfigBinding {
		bindings = append(bindings, newBinding("hyperdrive", b.Binding, func(tb *TargetBindingsModel) {
			tb.ID = b.ID
		}))
	}

	return bindings
}

// newBinding creates a new TargetBindingsModel with the given type, name, and field setter.
// All optional fields default to null.
func newBinding(bindingType string, name types.String, setFields func(*TargetBindingsModel)) *TargetBindingsModel {
	b := &TargetBindingsModel{
		Type: types.StringValue(bindingType),
		Name: name,
		// All other fields default to null
	}
	setFields(b)
	return b
}

// transformPlacement converts v4 placement array to v5 placement object.
func transformPlacement(ctx context.Context, source SourceWorkerScriptModel, target *TargetWorkersScriptModel) {
	if len(source.Placement) == 0 {
		return
	}
	p := source.Placement[0]
	target.Placement = newPlacementObject(ctx, p.Mode)
}

// newBindingsList constructs a customfield.NestedObjectList from a slice of bindings.
func newBindingsList(ctx context.Context, bindings []*TargetBindingsModel) customfield.NestedObjectList[TargetBindingsModel] {
	if len(bindings) == 0 {
		return customfield.NullObjectList[TargetBindingsModel](ctx)
	}
	// Dereference pointers for NewObjectListMust
	values := make([]TargetBindingsModel, len(bindings))
	for i, b := range bindings {
		values[i] = *b
	}
	return customfield.NewObjectListMust(ctx, values)
}

// newPlacementObject constructs a customfield.NestedObject for placement.
func newPlacementObject(ctx context.Context, mode types.String) customfield.NestedObject[TargetPlacementModel] {
	obj, _ := customfield.NewObject(ctx, &TargetPlacementModel{
		Mode:           mode,
		LastAnalyzedAt: timetypes.NewRFC3339Null(),
		Status:         types.StringNull(),
		Region:         types.StringNull(),
		Hostname:       types.StringNull(),
		Host:           types.StringNull(),
	})
	return obj
}
