package v500

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles the version 0 collision between V4 and V5 state.
//
// Both V4 and V5 used schema_version=0 for cloudflare_workers_script, but with
// incompatible state formats (e.g. placement as array vs object).
//
// Detection: V4 state has "name" field (non-null), V5 state has "script_name" instead.
// We inspect req.State.Raw (tftypes.Value) directly to avoid needing a full union model.
//
// Paths:
//   - V4 state (Path B): Full V4→V5 transform via Transform()
//   - V5 state (Path C): Pass through raw state
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	isV4, err := detectV4State(req.State.Raw)
	if err != nil {
		resp.Diagnostics.AddError("Failed to detect state format", fmt.Sprintf("Could not determine V4 vs V5 state: %s", err))
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected V4 state (name field present), performing full V4→V5 transformation")
		upgradeFromV4(ctx, req, resp)
	} else {
		tflog.Info(ctx, "Detected V5 state (script_name field present), passing through")
		resp.State.Raw = req.State.Raw
	}
}

// detectV4State inspects raw tftypes.Value to check if "name" attribute is present and non-null.
// V4 state has "name", V5 state has "script_name". These are mutually exclusive.
func detectV4State(raw tftypes.Value) (bool, error) {
	var rawState map[string]tftypes.Value
	if err := raw.As(&rawState); err != nil {
		return false, fmt.Errorf("failed to read raw state as object: %w", err)
	}

	nameVal, hasName := rawState["name"]
	if hasName && nameVal.IsKnown() && !nameVal.IsNull() {
		return true, nil
	}

	return false, nil
}

// upgradeFromV4 handles V4 state at version 0: extracts V4 fields from raw state and transforms to V5.
//
// We can't use req.State.Get() because the raw tftypes.Value carries the union schema's type
// structure, and tfsdk.State.Get() requires an exact struct match. Instead, we extract
// V4 fields directly from the raw tftypes map.
func upgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var rawState map[string]tftypes.Value
	if err := req.State.Raw.As(&rawState); err != nil {
		resp.Diagnostics.AddError("Failed to read V4 state", fmt.Sprintf("Could not read raw state as map: %s", err))
		return
	}

	v4State := SourceWorkerScriptModel{
		ID:                extractString(rawState, "id"),
		Name:              extractString(rawState, "name"),
		AccountID:         extractString(rawState, "account_id"),
		Content:           extractString(rawState, "content"),
		Module:            extractBool(rawState, "module"),
		DispatchNamespace: extractString(rawState, "dispatch_namespace"),
	}

	// Extract binding lists
	v4State.PlainTextBinding = extractBindings(rawState, "plain_text_binding", func(m map[string]tftypes.Value) SourcePlainTextBindingModel {
		return SourcePlainTextBindingModel{Name: extractString(m, "name"), Text: extractString(m, "text")}
	})
	v4State.SecretTextBinding = extractBindings(rawState, "secret_text_binding", func(m map[string]tftypes.Value) SourceSecretTextBindingModel {
		return SourceSecretTextBindingModel{Name: extractString(m, "name"), Text: extractString(m, "text")}
	})
	v4State.KVNamespaceBinding = extractBindings(rawState, "kv_namespace_binding", func(m map[string]tftypes.Value) SourceKVNamespaceBindingModel {
		return SourceKVNamespaceBindingModel{Name: extractString(m, "name"), NamespaceID: extractString(m, "namespace_id")}
	})
	v4State.WebassemblyBinding = extractBindings(rawState, "webassembly_binding", func(m map[string]tftypes.Value) SourceWebassemblyBindingModel {
		return SourceWebassemblyBindingModel{Name: extractString(m, "name"), Module: extractString(m, "module")}
	})
	v4State.ServiceBinding = extractBindings(rawState, "service_binding", func(m map[string]tftypes.Value) SourceServiceBindingModel {
		return SourceServiceBindingModel{Name: extractString(m, "name"), Service: extractString(m, "service"), Environment: extractString(m, "environment")}
	})
	v4State.R2BucketBinding = extractBindings(rawState, "r2_bucket_binding", func(m map[string]tftypes.Value) SourceR2BucketBindingModel {
		return SourceR2BucketBindingModel{Name: extractString(m, "name"), BucketName: extractString(m, "bucket_name")}
	})
	v4State.AnalyticsEngineBinding = extractBindings(rawState, "analytics_engine_binding", func(m map[string]tftypes.Value) SourceAnalyticsEngineBindingModel {
		return SourceAnalyticsEngineBindingModel{Name: extractString(m, "name"), Dataset: extractString(m, "dataset")}
	})
	v4State.QueueBinding = extractBindings(rawState, "queue_binding", func(m map[string]tftypes.Value) SourceQueueBindingModel {
		return SourceQueueBindingModel{Binding: extractString(m, "binding"), Queue: extractString(m, "queue")}
	})
	v4State.D1DatabaseBinding = extractBindings(rawState, "d1_database_binding", func(m map[string]tftypes.Value) SourceD1DatabaseBindingModel {
		return SourceD1DatabaseBindingModel{Name: extractString(m, "name"), DatabaseID: extractString(m, "database_id")}
	})
	v4State.HyperdriveConfigBinding = extractBindings(rawState, "hyperdrive_config_binding", func(m map[string]tftypes.Value) SourceHyperdriveConfigBindingModel {
		return SourceHyperdriveConfigBindingModel{Binding: extractString(m, "binding"), ID: extractString(m, "id")}
	})

	// Note: placement is not in the union schema (deleted due to array/object type conflict).
	// V4 placement data is restored from config on next terraform apply.

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "V4→V5 state upgrade completed successfully")
}

// tftypes extraction helpers

func extractString(m map[string]tftypes.Value, key string) types.String {
	val, ok := m[key]
	if !ok || val.IsNull() || !val.IsKnown() {
		return types.StringNull()
	}
	var s string
	if err := val.As(&s); err != nil {
		return types.StringNull()
	}
	return types.StringValue(s)
}

func extractBool(m map[string]tftypes.Value, key string) types.Bool {
	val, ok := m[key]
	if !ok || val.IsNull() || !val.IsKnown() {
		return types.BoolNull()
	}
	var b bool
	if err := val.As(&b); err != nil {
		return types.BoolNull()
	}
	return types.BoolValue(b)
}

func extractBindings[T any](m map[string]tftypes.Value, key string, build func(map[string]tftypes.Value) T) []T {
	val, ok := m[key]
	if !ok || val.IsNull() || !val.IsKnown() {
		return nil
	}
	var items []tftypes.Value
	if err := val.As(&items); err != nil {
		return nil
	}
	result := make([]T, 0, len(items))
	for _, item := range items {
		var itemMap map[string]tftypes.Value
		if err := item.As(&itemMap); err != nil {
			continue
		}
		result = append(result, build(itemMap))
	}
	return result
}

// UpgradeFromV1 handles version 1 state — always a no-op.
//
// Version 1 state is either:
//   - V5 state after the run_worker_first V0→V1 upgrade
//   - tf-migrate output (V4 state transformed to V5 format, kept at version 1)
//
// Both are already in V5 format, so no transformation needed.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading workers_script state from version=1 to current (no-op)")
	resp.State.Raw = req.State.Raw
}
