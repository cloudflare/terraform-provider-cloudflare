package pages_project

import (
	"context"
	"reflect"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// PreserveSecretEnvVars returns a copy of the destination deployment configs with
// secret_text environment variable values carried over from the source deployment configs,
// since the API returns empty strings for secret environment variable values.
// This prevents false positive drift detection for secret environment variables.
func PreserveSecretEnvVars(
	ctx context.Context,
	sourceConfigs customfield.NestedObject[PagesProjectDeploymentConfigsModel],
	destConfigs customfield.NestedObject[PagesProjectDeploymentConfigsModel],
) (customfield.NestedObject[PagesProjectDeploymentConfigsModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	if destConfigs.IsNull() || sourceConfigs.IsNull() || sourceConfigs.IsUnknown() {
		return destConfigs, diags
	}

	destConfigsValue, d := destConfigs.Value(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return destConfigs, diags
	}

	sourceConfigsValue, d := sourceConfigs.Value(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return destConfigs, diags
	}

	// Process Preview environment
	if !destConfigsValue.Preview.IsNull() && !sourceConfigsValue.Preview.IsNull() {
		destPreview, d := destConfigsValue.Preview.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return destConfigs, diags
		}

		sourcePreview, d := sourceConfigsValue.Preview.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return destConfigs, diags
		}

		if destPreview.EnvVars != nil && sourcePreview.EnvVars != nil {
			destPreview.EnvVars = preserveSecretEnvVarsInMap(
				*sourcePreview.EnvVars,
				*destPreview.EnvVars,
			)
		}

		destConfigsValue.Preview, d = customfield.NewObject(ctx, destPreview)
		diags.Append(d...)
		if diags.HasError() {
			return destConfigs, diags
		}
	}

	// Process Production environment
	if !destConfigsValue.Production.IsNull() && !sourceConfigsValue.Production.IsNull() {
		destProduction, d := destConfigsValue.Production.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return destConfigs, diags
		}

		sourceProduction, d := sourceConfigsValue.Production.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return destConfigs, diags
		}

		if destProduction.EnvVars != nil && sourceProduction.EnvVars != nil {
			destProduction.EnvVars = preserveSecretEnvVarsInMap(
				*sourceProduction.EnvVars,
				*destProduction.EnvVars,
			)
		}

		destConfigsValue.Production, d = customfield.NewObject(ctx, destProduction)
		diags.Append(d...)
		if diags.HasError() {
			return destConfigs, diags
		}
	}

	result, d := customfield.NewObject(ctx, destConfigsValue)
	diags.Append(d...)
	return result, diags
}

type EnvVarModel interface {
	PagesProjectDeploymentConfigsPreviewEnvVarsModel | PagesProjectDeploymentConfigsProductionEnvVarsModel
}

func preserveSecretEnvVarsInMap[T EnvVarModel](
	sourceEnvVars map[string]T,
	destEnvVars map[string]T,
) *map[string]T {
	updatedEnvVars := make(map[string]T)

	for name, destVar := range destEnvVars {
		var varType, varValue, sourceValue string

		switch v := any(destVar).(type) {
		case PagesProjectDeploymentConfigsPreviewEnvVarsModel:
			varType = v.Type.ValueString()
			varValue = v.Value.ValueString()
			if sourceVar, exists := sourceEnvVars[name]; exists {
				sourceValue = any(sourceVar).(PagesProjectDeploymentConfigsPreviewEnvVarsModel).Value.ValueString()
				if varType == "secret_text" && varValue == "" && sourceValue != "" {
					if sv := any(sourceVar).(PagesProjectDeploymentConfigsPreviewEnvVarsModel).Value; !sv.IsNull() && !sv.IsUnknown() {
						v.Value = sv
						destVar = any(v).(T)
					}
				}
			}
		case PagesProjectDeploymentConfigsProductionEnvVarsModel:
			varType = v.Type.ValueString()
			varValue = v.Value.ValueString()
			if sourceVar, exists := sourceEnvVars[name]; exists {
				sourceValue = any(sourceVar).(PagesProjectDeploymentConfigsProductionEnvVarsModel).Value.ValueString()
				if varType == "secret_text" && varValue == "" && sourceValue != "" {
					if sv := any(sourceVar).(PagesProjectDeploymentConfigsProductionEnvVarsModel).Value; !sv.IsNull() && !sv.IsUnknown() {
						v.Value = sv
						destVar = any(v).(T)
					}
				}
			}
		}

		updatedEnvVars[name] = destVar
	}

	return &updatedEnvVars
}

// PrepareForUpdate converts nil binding maps to empty maps so the API deletes them.
func PrepareForUpdate(ctx context.Context, plan *PagesProjectModel, state *PagesProjectModel) (*PagesProjectModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if plan == nil || state == nil {
		return plan, diags
	}

	// Only process if both plan and state have deployment configs
	if plan.DeploymentConfigs.IsNull() || plan.DeploymentConfigs.IsUnknown() ||
		state.DeploymentConfigs.IsNull() || state.DeploymentConfigs.IsUnknown() {
		return plan, diags
	}

	planConfigs, d := plan.DeploymentConfigs.Value(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return plan, diags
	}

	stateConfigs, d := state.DeploymentConfigs.Value(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return plan, diags
	}

	modified := false

	// Process Preview config
	if !planConfigs.Preview.IsNull() && !planConfigs.Preview.IsUnknown() &&
		!stateConfigs.Preview.IsNull() && !stateConfigs.Preview.IsUnknown() {
		planPreview, d := planConfigs.Preview.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return plan, diags
		}

		statePreview, d := stateConfigs.Preview.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return plan, diags
		}

		previewModified := false
		planPreview, previewModified = convertNilToEmptyMapsPreview(planPreview, statePreview)
		if previewModified {
			planConfigs.Preview, d = customfield.NewObject(ctx, planPreview)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}
			modified = true
		}
	}

	// Process Production config
	if !planConfigs.Production.IsNull() && !planConfigs.Production.IsUnknown() &&
		!stateConfigs.Production.IsNull() && !stateConfigs.Production.IsUnknown() {
		planProduction, d := planConfigs.Production.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return plan, diags
		}

		stateProduction, d := stateConfigs.Production.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return plan, diags
		}

		productionModified := false
		planProduction, productionModified = convertNilToEmptyMapsProduction(planProduction, stateProduction)
		if productionModified {
			planConfigs.Production, d = customfield.NewObject(ctx, planProduction)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}
			modified = true
		}
	}

	if modified {
		plan.DeploymentConfigs, d = customfield.NewObject(ctx, planConfigs)
		diags.Append(d...)
	}

	return plan, diags
}

// bindingFields lists the struct field names that contain binding maps.
var bindingFields = []string{
	"EnvVars", "KVNamespaces", "D1Databases", "R2Buckets", "DurableObjectNamespaces",
	"AIBindings", "AnalyticsEngineDatasets", "Browsers", "HyperdriveBindings",
	"MTLSCertificates", "QueueProducers", "Services", "VectorizeBindings",
}

// normalizeEmptyMaps sets empty binding map pointers to nil using reflection.
func normalizeEmptyMaps(v any) bool {
	if v == nil {
		return false
	}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	modified := false
	for _, name := range bindingFields {
		field := rv.FieldByName(name)
		if !field.IsValid() || field.Kind() != reflect.Ptr || field.IsNil() {
			continue
		}
		mapVal := field.Elem()
		if mapVal.Kind() == reflect.Map && mapVal.Len() == 0 {
			field.Set(reflect.Zero(field.Type()))
			modified = true
		}
	}
	return modified
}

// convertNilToEmptyMaps sets nil binding maps to empty maps when state has values.
func convertNilToEmptyMaps(plan, state any) bool {
	if plan == nil || state == nil {
		return false
	}
	planVal := reflect.ValueOf(plan)
	stateVal := reflect.ValueOf(state)
	if planVal.Kind() == reflect.Ptr {
		planVal = planVal.Elem()
	}
	if stateVal.Kind() == reflect.Ptr {
		stateVal = stateVal.Elem()
	}
	if planVal.Kind() != reflect.Struct || stateVal.Kind() != reflect.Struct {
		return false
	}
	modified := false
	for _, name := range bindingFields {
		planField := planVal.FieldByName(name)
		stateField := stateVal.FieldByName(name)
		if !planField.IsValid() || !stateField.IsValid() {
			continue
		}
		if planField.Kind() != reflect.Ptr || stateField.Kind() != reflect.Ptr {
			continue
		}
		if planField.IsNil() && !stateField.IsNil() {
			stateMap := stateField.Elem()
			if stateMap.Kind() == reflect.Map && stateMap.Len() > 0 {
				// Create empty map of the same type
				emptyMap := reflect.MakeMap(stateMap.Type())
				newPtr := reflect.New(stateMap.Type())
				newPtr.Elem().Set(emptyMap)
				planField.Set(newPtr)
				modified = true
			}
		}
	}
	return modified
}

// Typed wrappers for the reflection-based helpers.
func normalizeEmptyMapsPreview(v *PagesProjectDeploymentConfigsPreviewModel) bool {
	return normalizeEmptyMaps(v)
}

func normalizeEmptyMapsProduction(v *PagesProjectDeploymentConfigsProductionModel) bool {
	return normalizeEmptyMaps(v)
}

func convertNilToEmptyMapsPreview(plan, state *PagesProjectDeploymentConfigsPreviewModel) (*PagesProjectDeploymentConfigsPreviewModel, bool) {
	return plan, convertNilToEmptyMaps(plan, state)
}

func convertNilToEmptyMapsProduction(plan, state *PagesProjectDeploymentConfigsProductionModel) (*PagesProjectDeploymentConfigsProductionModel, bool) {
	return plan, convertNilToEmptyMaps(plan, state)
}
