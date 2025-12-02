package pages_project

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// handleDeploymentConfigsRemoval handles the case where deployment_configs is removed from config.
// When deployment_configs is not specified (null/unknown in config), but exists in state,
// we need to construct a plan that clears the optional-only fields (env_vars, bindings, etc.)
// while preserving computed fields. This ensures the PATCH request properly removes these values.
func handleDeploymentConfigsRemoval(
	ctx context.Context,
	config PagesProjectModel,
	plan PagesProjectModel,
	state PagesProjectModel,
) (PagesProjectModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if state.DeploymentConfigs.IsNull() {
		return plan, diags
	}

	planDC := plan.DeploymentConfigs
	stateDC := state.DeploymentConfigs

	// Check if plan deployment_configs is unknown (not specified in config)
	// and state has deployment_configs
	if planDC.IsUnknown() && !stateDC.IsNull() {
		stateValue, d := stateDC.Value(ctx)
		diags.Append(d...)
		if diags.HasError() {
			return plan, diags
		}

		newDCValue := PagesProjectDeploymentConfigsModel{}

		if !stateValue.Preview.IsNull() {
			previewState, d := stateValue.Preview.Value(ctx)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}

			newPreview := clearOptionalBindings(*previewState)
			newDCValue.Preview, d = customfield.NewObject(ctx, &newPreview)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}
		}

		if !stateValue.Production.IsNull() {
			productionState, d := stateValue.Production.Value(ctx)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}

			newProduction := clearOptionalBindingsProduction(*productionState)
			newDCValue.Production, d = customfield.NewObject(ctx, &newProduction)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}
		}

		plan.DeploymentConfigs, d = customfield.NewObject(ctx, &newDCValue)
		diags.Append(d...)
		return plan, diags
	}

	configDC := config.DeploymentConfigs
	if !configDC.IsNull() && !configDC.IsUnknown() && !stateDC.IsNull() {
		configValue, d := configDC.Value(ctx)
		diags.Append(d...)
		if diags.HasError() || configValue == nil {
			return plan, diags
		}

		stateValue, d := stateDC.Value(ctx)
		diags.Append(d...)
		if diags.HasError() || stateValue == nil {
			return plan, diags
		}

		planValue, d := planDC.Value(ctx)
		diags.Append(d...)
		if diags.HasError() || planValue == nil {
			return plan, diags
		}

		needsUpdate := false

		if configValue.Preview.IsNull() && !stateValue.Preview.IsNull() && planValue.Preview.IsUnknown() {
			previewState, d := stateValue.Preview.Value(ctx)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}

			newPreview := clearOptionalBindings(*previewState)
			planValue.Preview, d = customfield.NewObject(ctx, &newPreview)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}
			needsUpdate = true
		}

		if configValue.Production.IsNull() && !stateValue.Production.IsNull() && planValue.Production.IsUnknown() {
			productionState, d := stateValue.Production.Value(ctx)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}

			newProduction := clearOptionalBindingsProduction(*productionState)
			planValue.Production, d = customfield.NewObject(ctx, &newProduction)
			diags.Append(d...)
			if diags.HasError() {
				return plan, diags
			}
			needsUpdate = true
		}

		if needsUpdate {
			plan.DeploymentConfigs, d = customfield.NewObject(ctx, planValue)
			diags.Append(d...)
		}
	}

	return plan, diags
}

func clearOptionalBindings(state PagesProjectDeploymentConfigsPreviewModel) PagesProjectDeploymentConfigsPreviewModel {
	emptyFlags := []types.String{}
	emptyEnvVars := make(map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel)
	emptyAIBindings := make(map[string]PagesProjectDeploymentConfigsPreviewAIBindingsModel)
	emptyAnalyticsEngine := make(map[string]PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel)
	emptyBrowsers := make(map[string]PagesProjectDeploymentConfigsPreviewBrowsersModel)
	emptyD1 := make(map[string]PagesProjectDeploymentConfigsPreviewD1DatabasesModel)
	emptyDO := make(map[string]PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel)
	emptyHyperdrive := make(map[string]PagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel)
	emptyKV := make(map[string]PagesProjectDeploymentConfigsPreviewKVNamespacesModel)
	emptyMTLS := make(map[string]PagesProjectDeploymentConfigsPreviewMTLSCertificatesModel)
	emptyQueueProducers := make(map[string]PagesProjectDeploymentConfigsPreviewQueueProducersModel)
	emptyR2 := make(map[string]PagesProjectDeploymentConfigsPreviewR2BucketsModel)
	emptyServices := make(map[string]PagesProjectDeploymentConfigsPreviewServicesModel)
	emptyVectorize := make(map[string]PagesProjectDeploymentConfigsPreviewVectorizeBindingsModel)

	return PagesProjectDeploymentConfigsPreviewModel{
		AlwaysUseLatestCompatibilityDate: state.AlwaysUseLatestCompatibilityDate,
		BuildImageMajorVersion:           state.BuildImageMajorVersion,
		CompatibilityDate:                state.CompatibilityDate,
		FailOpen:                         state.FailOpen,
		UsageModel:                       state.UsageModel,
		CompatibilityFlags:               &emptyFlags,
		// Use empty maps instead of nil to trigger deletion
		AIBindings:              &emptyAIBindings,
		AnalyticsEngineDatasets: &emptyAnalyticsEngine,
		Browsers:                &emptyBrowsers,
		D1Databases:             &emptyD1,
		DurableObjectNamespaces: &emptyDO,
		EnvVars:                 &emptyEnvVars,
		HyperdriveBindings:      &emptyHyperdrive,
		KVNamespaces:            &emptyKV,
		MTLSCertificates:        &emptyMTLS,
		QueueProducers:          &emptyQueueProducers,
		R2Buckets:               &emptyR2,
		Services:                &emptyServices,
		VectorizeBindings:       &emptyVectorize,
		Limits:                  nil,
		Placement:               nil,
		WranglerConfigHash:      state.WranglerConfigHash,
	}
}

func clearOptionalBindingsProduction(state PagesProjectDeploymentConfigsProductionModel) PagesProjectDeploymentConfigsProductionModel {
	emptyFlags := []types.String{}
	emptyEnvVars := make(map[string]PagesProjectDeploymentConfigsProductionEnvVarsModel)
	emptyAIBindings := make(map[string]PagesProjectDeploymentConfigsProductionAIBindingsModel)
	emptyAnalyticsEngine := make(map[string]PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel)
	emptyBrowsers := make(map[string]PagesProjectDeploymentConfigsProductionBrowsersModel)
	emptyD1 := make(map[string]PagesProjectDeploymentConfigsProductionD1DatabasesModel)
	emptyDO := make(map[string]PagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel)
	emptyHyperdrive := make(map[string]PagesProjectDeploymentConfigsProductionHyperdriveBindingsModel)
	emptyKV := make(map[string]PagesProjectDeploymentConfigsProductionKVNamespacesModel)
	emptyMTLS := make(map[string]PagesProjectDeploymentConfigsProductionMTLSCertificatesModel)
	emptyQueueProducers := make(map[string]PagesProjectDeploymentConfigsProductionQueueProducersModel)
	emptyR2 := make(map[string]PagesProjectDeploymentConfigsProductionR2BucketsModel)
	emptyServices := make(map[string]PagesProjectDeploymentConfigsProductionServicesModel)
	emptyVectorize := make(map[string]PagesProjectDeploymentConfigsProductionVectorizeBindingsModel)

	return PagesProjectDeploymentConfigsProductionModel{
		AlwaysUseLatestCompatibilityDate: state.AlwaysUseLatestCompatibilityDate,
		BuildImageMajorVersion:           state.BuildImageMajorVersion,
		CompatibilityDate:                state.CompatibilityDate,
		FailOpen:                         state.FailOpen,
		UsageModel:                       state.UsageModel,
		CompatibilityFlags:               &emptyFlags,
		// Use empty maps instead of nil to trigger deletion
		AIBindings:              &emptyAIBindings,
		AnalyticsEngineDatasets: &emptyAnalyticsEngine,
		Browsers:                &emptyBrowsers,
		D1Databases:             &emptyD1,
		DurableObjectNamespaces: &emptyDO,
		EnvVars:                 &emptyEnvVars,
		HyperdriveBindings:      &emptyHyperdrive,
		KVNamespaces:            &emptyKV,
		MTLSCertificates:        &emptyMTLS,
		QueueProducers:          &emptyQueueProducers,
		R2Buckets:               &emptyR2,
		Services:                &emptyServices,
		VectorizeBindings:       &emptyVectorize,
		Limits:                  nil,
		Placement:               nil,
		WranglerConfigHash:      state.WranglerConfigHash,
	}
}

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

// NormalizeBuildConfig converts an empty build_config (all fields null/unknown) to nil
// This handles the case where the API returns {} for build_config but the user didn't configure it
func NormalizeBuildConfig(data *PagesProjectModel) {
	if data.BuildConfig == nil {
		return
	}

	bc := data.BuildConfig
	// Check if all fields are null/unknown - if so, set build_config to nil
	if (bc.BuildCaching.IsNull() || bc.BuildCaching.IsUnknown()) &&
		(bc.BuildCommand.IsNull() || bc.BuildCommand.IsUnknown()) &&
		(bc.DestinationDir.IsNull() || bc.DestinationDir.IsUnknown()) &&
		(bc.RootDir.IsNull() || bc.RootDir.IsUnknown()) &&
		(bc.WebAnalyticsTag.IsNull() || bc.WebAnalyticsTag.IsUnknown()) &&
		(bc.WebAnalyticsToken.IsNull() || bc.WebAnalyticsToken.IsUnknown()) {
		data.BuildConfig = nil
	}
}

// NormalizeDeploymentConfigsBindings converts empty binding maps to nil
// This handles the case where the API returns {} for bindings that weren't configured
func NormalizeDeploymentConfigsBindings(ctx context.Context, data *PagesProjectModel) diag.Diagnostics {
	var diags diag.Diagnostics

	if data.DeploymentConfigs.IsNull() || data.DeploymentConfigs.IsUnknown() {
		return diags
	}

	dcValue, d := data.DeploymentConfigs.Value(ctx)
	diags.Append(d...)
	if diags.HasError() || dcValue == nil {
		return diags
	}

	needsUpdate := false

	// Normalize preview bindings
	if !dcValue.Preview.IsNull() && !dcValue.Preview.IsUnknown() {
		previewValue, d := dcValue.Preview.Value(ctx)
		diags.Append(d...)
		if diags.HasError() || previewValue == nil {
			return diags
		}

		normalizedPreview := normalizeBindingsToNil(previewValue)
		if normalizedPreview != previewValue {
			dcValue.Preview, d = customfield.NewObject(ctx, normalizedPreview)
			diags.Append(d...)
			if diags.HasError() {
				return diags
			}
			needsUpdate = true
		}
	}

	// Normalize production bindings
	if !dcValue.Production.IsNull() && !dcValue.Production.IsUnknown() {
		productionValue, d := dcValue.Production.Value(ctx)
		diags.Append(d...)
		if diags.HasError() || productionValue == nil {
			return diags
		}

		normalizedProduction := normalizeProductionBindingsToNil(productionValue)
		if normalizedProduction != productionValue {
			dcValue.Production, d = customfield.NewObject(ctx, normalizedProduction)
			diags.Append(d...)
			if diags.HasError() {
				return diags
			}
			needsUpdate = true
		}
	}

	if needsUpdate {
		data.DeploymentConfigs, d = customfield.NewObject(ctx, dcValue)
		diags.Append(d...)
	}

	return diags
}

// normalizeBindingsToNil converts empty binding maps to nil for preview config
func normalizeBindingsToNil(preview *PagesProjectDeploymentConfigsPreviewModel) *PagesProjectDeploymentConfigsPreviewModel {
	if preview == nil {
		return nil
	}

	// Create a copy to avoid modifying the original
	result := *preview

	if preview.AIBindings != nil && len(*preview.AIBindings) == 0 {
		result.AIBindings = nil
	}
	if preview.AnalyticsEngineDatasets != nil && len(*preview.AnalyticsEngineDatasets) == 0 {
		result.AnalyticsEngineDatasets = nil
	}
	if preview.Browsers != nil && len(*preview.Browsers) == 0 {
		result.Browsers = nil
	}
	if preview.D1Databases != nil && len(*preview.D1Databases) == 0 {
		result.D1Databases = nil
	}
	if preview.DurableObjectNamespaces != nil && len(*preview.DurableObjectNamespaces) == 0 {
		result.DurableObjectNamespaces = nil
	}
	if preview.EnvVars != nil && len(*preview.EnvVars) == 0 {
		result.EnvVars = nil
	}
	if preview.HyperdriveBindings != nil && len(*preview.HyperdriveBindings) == 0 {
		result.HyperdriveBindings = nil
	}
	if preview.KVNamespaces != nil && len(*preview.KVNamespaces) == 0 {
		result.KVNamespaces = nil
	}
	if preview.MTLSCertificates != nil && len(*preview.MTLSCertificates) == 0 {
		result.MTLSCertificates = nil
	}
	if preview.QueueProducers != nil && len(*preview.QueueProducers) == 0 {
		result.QueueProducers = nil
	}
	if preview.R2Buckets != nil && len(*preview.R2Buckets) == 0 {
		result.R2Buckets = nil
	}
	if preview.Services != nil && len(*preview.Services) == 0 {
		result.Services = nil
	}
	if preview.VectorizeBindings != nil && len(*preview.VectorizeBindings) == 0 {
		result.VectorizeBindings = nil
	}

	return &result
}

// normalizeProductionBindingsToNil converts empty binding maps to nil for production config
func normalizeProductionBindingsToNil(production *PagesProjectDeploymentConfigsProductionModel) *PagesProjectDeploymentConfigsProductionModel {
	if production == nil {
		return nil
	}

	// Create a copy to avoid modifying the original
	result := *production

	if production.AIBindings != nil && len(*production.AIBindings) == 0 {
		result.AIBindings = nil
	}
	if production.AnalyticsEngineDatasets != nil && len(*production.AnalyticsEngineDatasets) == 0 {
		result.AnalyticsEngineDatasets = nil
	}
	if production.Browsers != nil && len(*production.Browsers) == 0 {
		result.Browsers = nil
	}
	if production.D1Databases != nil && len(*production.D1Databases) == 0 {
		result.D1Databases = nil
	}
	if production.DurableObjectNamespaces != nil && len(*production.DurableObjectNamespaces) == 0 {
		result.DurableObjectNamespaces = nil
	}
	if production.EnvVars != nil && len(*production.EnvVars) == 0 {
		result.EnvVars = nil
	}
	if production.HyperdriveBindings != nil && len(*production.HyperdriveBindings) == 0 {
		result.HyperdriveBindings = nil
	}
	if production.KVNamespaces != nil && len(*production.KVNamespaces) == 0 {
		result.KVNamespaces = nil
	}
	if production.MTLSCertificates != nil && len(*production.MTLSCertificates) == 0 {
		result.MTLSCertificates = nil
	}
	if production.QueueProducers != nil && len(*production.QueueProducers) == 0 {
		result.QueueProducers = nil
	}
	if production.R2Buckets != nil && len(*production.R2Buckets) == 0 {
		result.R2Buckets = nil
	}
	if production.Services != nil && len(*production.Services) == 0 {
		result.Services = nil
	}
	if production.VectorizeBindings != nil && len(*production.VectorizeBindings) == 0 {
		result.VectorizeBindings = nil
	}

	return &result
}
