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
