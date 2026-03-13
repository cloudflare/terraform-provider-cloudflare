package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function is shared by both UpgradeFromV0 and MoveState handlers.
func Transform(ctx context.Context, source SourceDevicePostureRuleModel) (*TargetDevicePostureRuleModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	tflog.Debug(ctx, "Transforming device_posture_rule state from v4 to v5",
		map[string]interface{}{
			"source_type": source.Type.ValueString(),
			"source_name": source.Name.ValueString(),
		})

	// Direct copies for top-level fields, converting falsey values to null
	target := &TargetDevicePostureRuleModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		Name:        nullifyEmptyString(source.Name),
		Type:        source.Type,
		Description: nullifyEmptyString(source.Description),
		Expiration:  nullifyEmptyString(source.Expiration),
		Schedule:    nullifyEmptyString(source.Schedule),
	}

	// Transform input: []SourceInputModel (list with MaxItems:1) → *TargetInputModel
	if len(source.Input) > 0 {
		inputTarget, inputDiags := transformInput(ctx, source.Input[0])
		diags.Append(inputDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.Input = inputTarget
	}

	// Transform match: []SourceMatchModel → *[]*TargetMatchModel
	if len(source.Match) > 0 {
		matchTarget := transformMatch(source.Match)
		target.Match = matchTarget
	}

	return target, diags
}

// transformInput converts a single v4 input block to v5 input nested attribute.
// Falsey values (empty strings, false bools, zero floats) from API defaults are
// converted to null to match what the v5 Read path produces via normalizeReadData.
func transformInput(ctx context.Context, source SourceInputModel) (*TargetInputModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetInputModel{
		// String fields: convert empty strings to null
		OperatingSystem:  nullifyEmptyString(source.OperatingSystem),
		Path:             nullifyEmptyString(source.Path),
		Sha256:           nullifyEmptyString(source.Sha256),
		Thumbprint:       nullifyEmptyString(source.Thumbprint),
		ID:               nullifyEmptyString(source.ID),
		Domain:           nullifyEmptyString(source.Domain),
		Operator:         nullifyEmptyString(source.Operator),
		Version:          nullifyEmptyString(source.Version),
		OSDistroName:     nullifyEmptyString(source.OSDistroName),
		OSDistroRevision: nullifyEmptyString(source.OSDistroRevision),
		OSVersionExtra:   nullifyEmptyString(source.OSVersionExtra),
		CertificateID:    nullifyEmptyString(source.CertificateID),
		Cn:               nullifyEmptyString(source.Cn),
		ComplianceStatus: nullifyEmptyString(source.ComplianceStatus),
		ConnectionID:     nullifyEmptyString(source.ConnectionID),
		LastSeen:         nullifyEmptyString(source.LastSeen),
		OS:               nullifyEmptyString(source.OS),
		Overall:          nullifyEmptyString(source.Overall),
		SensorConfig:     nullifyEmptyString(source.SensorConfig),
		State:            nullifyEmptyString(source.State),
		VersionOperator:  nullifyEmptyString(source.VersionOperator),
		CountOperator:    nullifyEmptyString(source.CountOperator),
		IssueCount:       nullifyEmptyString(source.IssueCount),
		EidLastSeen:      nullifyEmptyString(source.EidLastSeen),
		RiskLevel:        nullifyEmptyString(source.RiskLevel),
		ScoreOperator:    nullifyEmptyString(source.ScoreOperator),
		NetworkStatus:    nullifyEmptyString(source.NetworkStatus),
		OperationalState: nullifyEmptyString(source.OperationalState),

		// Bool fields: convert false to null (API defaults)
		Exists:          nullifyFalseBool(source.Exists),
		Enabled:         nullifyFalseBool(source.Enabled),
		RequireAll:      nullifyFalseBool(source.RequireAll),
		CheckPrivateKey: nullifyFalseBool(source.CheckPrivateKey),
		Infected:        nullifyFalseBool(source.Infected),
		IsActive:        nullifyFalseBool(source.IsActive),

		// Float64 fields: convert zero to null (API defaults)
		UpdateWindowDays: nullifyZeroFloat64(source.UpdateWindowDays),
		TotalScore:       nullifyZeroFloat64(source.TotalScore),
		ActiveThreats:    nullifyZeroFloat64(source.ActiveThreats),
		Score:            nullifyZeroFloat64(source.Score),

		// Note: source.Running is intentionally NOT copied (removed in v5)
	}

	// Convert check_disks: types.Set → *[]types.String
	if !source.CheckDisks.IsNull() && !source.CheckDisks.IsUnknown() {
		var elements []types.String
		diags.Append(source.CheckDisks.ElementsAs(ctx, &elements, false)...)
		if diags.HasError() {
			return nil, diags
		}
		if len(elements) > 0 {
			target.CheckDisks = &elements
		}
	}

	// Convert extended_key_usage: types.List → *[]types.String
	if !source.ExtendedKeyUsage.IsNull() && !source.ExtendedKeyUsage.IsUnknown() {
		var elements []types.String
		diags.Append(source.ExtendedKeyUsage.ElementsAs(ctx, &elements, false)...)
		if diags.HasError() {
			return nil, diags
		}
		if len(elements) > 0 {
			target.ExtendedKeyUsage = &elements
		}
	}

	// Convert subject_alternative_names: types.List → *[]types.String
	if !source.SubjectAlternativeNames.IsNull() && !source.SubjectAlternativeNames.IsUnknown() {
		var elements []types.String
		diags.Append(source.SubjectAlternativeNames.ElementsAs(ctx, &elements, false)...)
		if diags.HasError() {
			return nil, diags
		}
		target.SubjectAlternativeNames = &elements
	}

	// Convert locations: []SourceLocationsModel → *TargetLocationsModel
	if len(source.Locations) > 0 {
		locTarget, locDiags := transformLocations(ctx, source.Locations[0])
		diags.Append(locDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.Locations = locTarget
	}

	return target, diags
}

// transformLocations converts a v4 locations block to v5 locations nested attribute.
func transformLocations(ctx context.Context, source SourceLocationsModel) (*TargetLocationsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetLocationsModel{}

	// Convert paths: types.List → *[]types.String
	if !source.Paths.IsNull() && !source.Paths.IsUnknown() {
		var elements []types.String
		diags.Append(source.Paths.ElementsAs(ctx, &elements, false)...)
		if diags.HasError() {
			return nil, diags
		}
		target.Paths = &elements
	}

	// Convert trust_stores: types.List → *[]types.String
	if !source.TrustStores.IsNull() && !source.TrustStores.IsUnknown() {
		var elements []types.String
		diags.Append(source.TrustStores.ElementsAs(ctx, &elements, false)...)
		if diags.HasError() {
			return nil, diags
		}
		target.TrustStores = &elements
	}

	return target, diags
}

// transformMatch converts v4 match blocks to v5 match list.
func transformMatch(source []SourceMatchModel) *[]*TargetMatchModel {
	matches := make([]*TargetMatchModel, len(source))
	for i, m := range source {
		matches[i] = &TargetMatchModel{
			Platform: m.Platform,
		}
	}
	return &matches
}

// nullifyEmptyString converts an empty string to null.
// The v4 API returns "" for unset optional string fields; v5 represents these as null.
func nullifyEmptyString(val types.String) types.String {
	if val.IsNull() || val.IsUnknown() {
		return val
	}
	if val.ValueString() == "" {
		return types.StringNull()
	}
	return val
}

// nullifyFalseBool converts false to null.
// The v4 API returns false for unset optional bool fields; v5 represents these as null.
func nullifyFalseBool(val types.Bool) types.Bool {
	if val.IsNull() || val.IsUnknown() {
		return val
	}
	if !val.ValueBool() {
		return types.BoolNull()
	}
	return val
}

// nullifyZeroFloat64 converts zero to null.
// The v4 API returns 0 for unset optional numeric fields; v5 represents these as null.
func nullifyZeroFloat64(val types.Float64) types.Float64 {
	if val.IsNull() || val.IsUnknown() {
		return val
	}
	if val.ValueFloat64() == 0 {
		return types.Float64Null()
	}
	return val
}
