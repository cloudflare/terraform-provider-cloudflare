package pages_project

import (
	"context"

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

// PrepareForUpdate converts nil binding maps to empty maps in the plan model.
// This is needed because the Cloudflare Pages API requires empty objects {} to delete
// bindings, not null values. When a binding field is not specified in the Terraform config,
// it will be nil in the plan, but we need to send {} to the API to delete it if it exists
// in the current state.
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

// convertNilToEmptyMapsPreview converts nil binding maps to empty maps for preview config
func convertNilToEmptyMapsPreview(plan *PagesProjectDeploymentConfigsPreviewModel, state *PagesProjectDeploymentConfigsPreviewModel) (*PagesProjectDeploymentConfigsPreviewModel, bool) {
	if plan == nil || state == nil {
		return plan, false
	}

	modified := false

	// EnvVars
	if plan.EnvVars == nil && state.EnvVars != nil && len(*state.EnvVars) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewEnvVarsModel)
		plan.EnvVars = &emptyMap
		modified = true
	}

	// KV Namespaces
	if plan.KVNamespaces == nil && state.KVNamespaces != nil && len(*state.KVNamespaces) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewKVNamespacesModel)
		plan.KVNamespaces = &emptyMap
		modified = true
	}

	// D1 Databases
	if plan.D1Databases == nil && state.D1Databases != nil && len(*state.D1Databases) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewD1DatabasesModel)
		plan.D1Databases = &emptyMap
		modified = true
	}

	// R2 Buckets
	if plan.R2Buckets == nil && state.R2Buckets != nil && len(*state.R2Buckets) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewR2BucketsModel)
		plan.R2Buckets = &emptyMap
		modified = true
	}

	// Durable Object Namespaces
	if plan.DurableObjectNamespaces == nil && state.DurableObjectNamespaces != nil && len(*state.DurableObjectNamespaces) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel)
		plan.DurableObjectNamespaces = &emptyMap
		modified = true
	}

	// AI Bindings
	if plan.AIBindings == nil && state.AIBindings != nil && len(*state.AIBindings) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewAIBindingsModel)
		plan.AIBindings = &emptyMap
		modified = true
	}

	// Analytics Engine Datasets
	if plan.AnalyticsEngineDatasets == nil && state.AnalyticsEngineDatasets != nil && len(*state.AnalyticsEngineDatasets) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel)
		plan.AnalyticsEngineDatasets = &emptyMap
		modified = true
	}

	// Browsers
	if plan.Browsers == nil && state.Browsers != nil && len(*state.Browsers) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewBrowsersModel)
		plan.Browsers = &emptyMap
		modified = true
	}

	// Hyperdrive Bindings
	if plan.HyperdriveBindings == nil && state.HyperdriveBindings != nil && len(*state.HyperdriveBindings) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel)
		plan.HyperdriveBindings = &emptyMap
		modified = true
	}

	// MTLS Certificates
	if plan.MTLSCertificates == nil && state.MTLSCertificates != nil && len(*state.MTLSCertificates) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewMTLSCertificatesModel)
		plan.MTLSCertificates = &emptyMap
		modified = true
	}

	// Queue Producers
	if plan.QueueProducers == nil && state.QueueProducers != nil && len(*state.QueueProducers) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewQueueProducersModel)
		plan.QueueProducers = &emptyMap
		modified = true
	}

	// Services
	if plan.Services == nil && state.Services != nil && len(*state.Services) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewServicesModel)
		plan.Services = &emptyMap
		modified = true
	}

	// Vectorize Bindings
	if plan.VectorizeBindings == nil && state.VectorizeBindings != nil && len(*state.VectorizeBindings) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsPreviewVectorizeBindingsModel)
		plan.VectorizeBindings = &emptyMap
		modified = true
	}

	return plan, modified
}

// convertNilToEmptyMapsProduction converts nil binding maps to empty maps for production config
func convertNilToEmptyMapsProduction(plan *PagesProjectDeploymentConfigsProductionModel, state *PagesProjectDeploymentConfigsProductionModel) (*PagesProjectDeploymentConfigsProductionModel, bool) {
	if plan == nil || state == nil {
		return plan, false
	}

	modified := false

	// EnvVars
	if plan.EnvVars == nil && state.EnvVars != nil && len(*state.EnvVars) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionEnvVarsModel)
		plan.EnvVars = &emptyMap
		modified = true
	}

	// KV Namespaces
	if plan.KVNamespaces == nil && state.KVNamespaces != nil && len(*state.KVNamespaces) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionKVNamespacesModel)
		plan.KVNamespaces = &emptyMap
		modified = true
	}

	// D1 Databases
	if plan.D1Databases == nil && state.D1Databases != nil && len(*state.D1Databases) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionD1DatabasesModel)
		plan.D1Databases = &emptyMap
		modified = true
	}

	// R2 Buckets
	if plan.R2Buckets == nil && state.R2Buckets != nil && len(*state.R2Buckets) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionR2BucketsModel)
		plan.R2Buckets = &emptyMap
		modified = true
	}

	// Durable Object Namespaces
	if plan.DurableObjectNamespaces == nil && state.DurableObjectNamespaces != nil && len(*state.DurableObjectNamespaces) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel)
		plan.DurableObjectNamespaces = &emptyMap
		modified = true
	}

	// AI Bindings
	if plan.AIBindings == nil && state.AIBindings != nil && len(*state.AIBindings) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionAIBindingsModel)
		plan.AIBindings = &emptyMap
		modified = true
	}

	// Analytics Engine Datasets
	if plan.AnalyticsEngineDatasets == nil && state.AnalyticsEngineDatasets != nil && len(*state.AnalyticsEngineDatasets) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel)
		plan.AnalyticsEngineDatasets = &emptyMap
		modified = true
	}

	// Browsers
	if plan.Browsers == nil && state.Browsers != nil && len(*state.Browsers) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionBrowsersModel)
		plan.Browsers = &emptyMap
		modified = true
	}

	// Hyperdrive Bindings
	if plan.HyperdriveBindings == nil && state.HyperdriveBindings != nil && len(*state.HyperdriveBindings) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionHyperdriveBindingsModel)
		plan.HyperdriveBindings = &emptyMap
		modified = true
	}

	// MTLS Certificates
	if plan.MTLSCertificates == nil && state.MTLSCertificates != nil && len(*state.MTLSCertificates) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionMTLSCertificatesModel)
		plan.MTLSCertificates = &emptyMap
		modified = true
	}

	// Queue Producers
	if plan.QueueProducers == nil && state.QueueProducers != nil && len(*state.QueueProducers) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionQueueProducersModel)
		plan.QueueProducers = &emptyMap
		modified = true
	}

	// Services
	if plan.Services == nil && state.Services != nil && len(*state.Services) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionServicesModel)
		plan.Services = &emptyMap
		modified = true
	}

	// Vectorize Bindings
	if plan.VectorizeBindings == nil && state.VectorizeBindings != nil && len(*state.VectorizeBindings) > 0 {
		emptyMap := make(map[string]PagesProjectDeploymentConfigsProductionVectorizeBindingsModel)
		plan.VectorizeBindings = &emptyMap
		modified = true
	}

	return plan, modified
}
