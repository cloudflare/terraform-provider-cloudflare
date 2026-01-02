package pages_project

import (
	"context"
	"reflect"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// PreserveSecretEnvVars returns the source (plan/state) deployment configs if non-secret
// values match, otherwise merges secret values from source into destination.
// Returns source directly when possible to avoid "inconsistent result" errors.
func PreserveSecretEnvVars(
	ctx context.Context,
	sourceConfigs customfield.NestedObject[PagesProjectDeploymentConfigsModel],
	destConfigs customfield.NestedObject[PagesProjectDeploymentConfigsModel],
) (customfield.NestedObject[PagesProjectDeploymentConfigsModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	if destConfigs.IsNull() {
		return destConfigs, diags
	}

	if sourceConfigs.IsNull() || sourceConfigs.IsUnknown() {
		return destConfigs, diags
	}

	// If non-secret values match, return source directly to preserve object identity.
	match, d := deploymentConfigsMatchIgnoringSecrets(ctx, sourceConfigs, destConfigs)
	diags.Append(d...)
	if diags.HasError() {
		return destConfigs, diags
	}

	if match {
		return sourceConfigs, diags
	}

	// Non-secret values differ, merge secrets into dest
	return mergeDeploymentConfigsWithSecrets(ctx, sourceConfigs, destConfigs)
}

// deploymentConfigsMatchIgnoringSecrets returns true if all non-secret values match.
func deploymentConfigsMatchIgnoringSecrets(
	ctx context.Context,
	source customfield.NestedObject[PagesProjectDeploymentConfigsModel],
	dest customfield.NestedObject[PagesProjectDeploymentConfigsModel],
) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	sourceValue, d := source.Value(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	destValue, d := dest.Value(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	// Compare Preview configs
	if !previewConfigsMatchIgnoringSecrets(ctx, sourceValue.Preview, destValue.Preview) {
		return false, diags
	}

	// Compare Production configs
	if !productionConfigsMatchIgnoringSecrets(ctx, sourceValue.Production, destValue.Production) {
		return false, diags
	}

	return true, diags
}

// previewConfigsMatchIgnoringSecrets compares preview configs, ignoring secret values.
func previewConfigsMatchIgnoringSecrets(
	ctx context.Context,
	source customfield.NestedObject[PagesProjectDeploymentConfigsPreviewModel],
	dest customfield.NestedObject[PagesProjectDeploymentConfigsPreviewModel],
) bool {
	// Both null or both unknown means they match
	if source.IsNull() && dest.IsNull() {
		return true
	}
	if source.IsUnknown() && dest.IsUnknown() {
		return true
	}
	// One null/unknown and other isn't means they don't match
	if source.IsNull() || source.IsUnknown() || dest.IsNull() || dest.IsUnknown() {
		return false
	}

	sourceValue, d := source.Value(ctx)
	if d.HasError() {
		return false
	}
	destValue, d := dest.Value(ctx)
	if d.HasError() {
		return false
	}

	// Compare non-secret fields
	if !sourceValue.AlwaysUseLatestCompatibilityDate.Equal(destValue.AlwaysUseLatestCompatibilityDate) {
		return false
	}
	if !sourceValue.BuildImageMajorVersion.Equal(destValue.BuildImageMajorVersion) {
		return false
	}
	if !sourceValue.CompatibilityDate.Equal(destValue.CompatibilityDate) {
		return false
	}
	if !sourceValue.FailOpen.Equal(destValue.FailOpen) {
		return false
	}
	if !sourceValue.UsageModel.Equal(destValue.UsageModel) {
		return false
	}
	if !sourceValue.WranglerConfigHash.Equal(destValue.WranglerConfigHash) {
		return false
	}

	// Compare compatibility flags
	if !compatibilityFlagsEqual(sourceValue.CompatibilityFlags, destValue.CompatibilityFlags) {
		return false
	}

	// Compare env vars (ignoring secret values)
	if !envVarsMatchIgnoringSecrets(sourceValue.EnvVars, destValue.EnvVars) {
		return false
	}

	return true
}

// productionConfigsMatchIgnoringSecrets compares production configs, ignoring secret values.
func productionConfigsMatchIgnoringSecrets(
	ctx context.Context,
	source customfield.NestedObject[PagesProjectDeploymentConfigsProductionModel],
	dest customfield.NestedObject[PagesProjectDeploymentConfigsProductionModel],
) bool {
	// Both null or both unknown means they match
	if source.IsNull() && dest.IsNull() {
		return true
	}
	if source.IsUnknown() && dest.IsUnknown() {
		return true
	}
	// One null/unknown and other isn't means they don't match
	if source.IsNull() || source.IsUnknown() || dest.IsNull() || dest.IsUnknown() {
		return false
	}

	sourceValue, d := source.Value(ctx)
	if d.HasError() {
		return false
	}
	destValue, d := dest.Value(ctx)
	if d.HasError() {
		return false
	}

	// Compare non-secret fields
	if !sourceValue.AlwaysUseLatestCompatibilityDate.Equal(destValue.AlwaysUseLatestCompatibilityDate) {
		return false
	}
	if !sourceValue.BuildImageMajorVersion.Equal(destValue.BuildImageMajorVersion) {
		return false
	}
	if !sourceValue.CompatibilityDate.Equal(destValue.CompatibilityDate) {
		return false
	}
	if !sourceValue.FailOpen.Equal(destValue.FailOpen) {
		return false
	}
	if !sourceValue.UsageModel.Equal(destValue.UsageModel) {
		return false
	}
	if !sourceValue.WranglerConfigHash.Equal(destValue.WranglerConfigHash) {
		return false
	}

	// Compare compatibility flags
	if !compatibilityFlagsEqual(sourceValue.CompatibilityFlags, destValue.CompatibilityFlags) {
		return false
	}

	// Compare env vars (ignoring secret values)
	if !envVarsMatchIgnoringSecretsProduction(sourceValue.EnvVars, destValue.EnvVars) {
		return false
	}

	return true
}

// compatibilityFlagsEqual compares compatibility flag slices.
func compatibilityFlagsEqual(source, dest *[]types.String) bool {
	if source == nil && dest == nil {
		return true
	}
	if source == nil || dest == nil {
		// One nil, one not - check if the non-nil one is empty
		if source == nil {
			return len(*dest) == 0
		}
		return len(*source) == 0
	}
	if len(*source) != len(*dest) {
		return false
	}
	for i := range *source {
		if !(*source)[i].Equal((*dest)[i]) {
			return false
		}
	}
	return true
}

// envVarsMatchIgnoringSecrets compares env vars, ignoring secret_text values.
func envVarsMatchIgnoringSecrets(
	source *map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel,
	dest *map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel,
) bool {
	if source == nil && dest == nil {
		return true
	}
	if source == nil || dest == nil {
		// One nil, one not - check if the non-nil one is empty
		if source == nil {
			return len(*dest) == 0
		}
		return len(*source) == 0
	}
	if len(*source) != len(*dest) {
		return false
	}
	for name, sourceVar := range *source {
		destVar, exists := (*dest)[name]
		if !exists {
			return false
		}
		// Compare type
		if !sourceVar.Type.Equal(destVar.Type) {
			return false
		}
		// Only compare values for plain_text
		if sourceVar.Type.ValueString() == "plain_text" {
			if !sourceVar.Value.Equal(destVar.Value) {
				return false
			}
		}
	}
	return true
}

// envVarsMatchIgnoringSecretsProduction compares production env vars, ignoring secret_text values.
func envVarsMatchIgnoringSecretsProduction(
	source *map[string]PagesProjectDeploymentConfigsProductionEnvVarsModel,
	dest *map[string]PagesProjectDeploymentConfigsProductionEnvVarsModel,
) bool {
	if source == nil && dest == nil {
		return true
	}
	if source == nil || dest == nil {
		// One nil, one not - check if the non-nil one is empty
		if source == nil {
			return len(*dest) == 0
		}
		return len(*source) == 0
	}
	if len(*source) != len(*dest) {
		return false
	}
	for name, sourceVar := range *source {
		destVar, exists := (*dest)[name]
		if !exists {
			return false
		}
		// Compare type
		if !sourceVar.Type.Equal(destVar.Type) {
			return false
		}
		// Only compare values for plain_text
		if sourceVar.Type.ValueString() == "plain_text" {
			if !sourceVar.Value.Equal(destVar.Value) {
				return false
			}
		}
	}
	return true
}

// mergeDeploymentConfigsWithSecrets merges secret env var values from source into dest.
func mergeDeploymentConfigsWithSecrets(
	ctx context.Context,
	sourceConfigs customfield.NestedObject[PagesProjectDeploymentConfigsModel],
	destConfigs customfield.NestedObject[PagesProjectDeploymentConfigsModel],
) (customfield.NestedObject[PagesProjectDeploymentConfigsModel], diag.Diagnostics) {
	var diags diag.Diagnostics

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
