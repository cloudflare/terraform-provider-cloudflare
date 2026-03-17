package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts v4 SDKv2 state to v5 Plugin Framework state.
// This handles the major structural transformation from:
// - v4: List fields with multiple string values (e.g., email = ["a@b.com", "c@d.com"])
// - v5: Each value becomes a separate object (e.g., two separate email blocks)
func Transform(ctx context.Context, source SourceV4ZeroTrustAccessGroupModel) (TargetV5ZeroTrustAccessGroupModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var target TargetV5ZeroTrustAccessGroupModel

	// Copy top-level fields directly
	target.ID = source.ID
	target.AccountID = source.AccountID
	target.ZoneID = source.ZoneID
	target.Name = source.Name

	// is_default is new in v5 - set to null for migrated resources
	target.IsDefault = types.BoolNull()

	// Transform include/exclude/require blocks
	// These require "exploding" lists into multiple objects
	if len(source.Include) > 0 {
		includeResult := transformRuleBlocks(ctx, source.Include, &diags)
		target.Include = &includeResult
	}

	if len(source.Exclude) > 0 {
		excludeResult := transformRuleBlocks(ctx, source.Exclude, &diags)
		target.Exclude = &excludeResult
	}

	if len(source.Require) > 0 {
		requireResult := transformRuleBlocks(ctx, source.Require, &diags)
		target.Require = &requireResult
	}

	return target, diags
}

// transformRuleBlocks transforms a v4 rule block (include/exclude/require) to v5 format.
// In v4, each block can have multiple fields, and many fields are lists.
// In v5, each selector becomes a separate object, and lists are "exploded" into multiple objects.
func transformRuleBlocks(ctx context.Context, sourceBlocks []SourceV4AccessGroupOptionModel, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	var result []*TargetV5ZeroTrustAccessGroupIncludeModel

	for _, sourceBlock := range sourceBlocks {
		// Process each selector type in the source block
		// Each selector can generate multiple v5 objects

		// Simple list selectors: Split each list item into a separate v5 object
		result = append(result, transformEmailList(ctx, sourceBlock.Email, diags)...)
		result = append(result, transformEmailDomainList(ctx, sourceBlock.EmailDomain, diags)...)
		result = append(result, transformEmailListList(ctx, sourceBlock.EmailList, diags)...)
		result = append(result, transformIPList(ctx, sourceBlock.IP, diags)...)
		result = append(result, transformIPListList(ctx, sourceBlock.IPList, diags)...)
		result = append(result, transformServiceTokenList(ctx, sourceBlock.ServiceToken, diags)...)
		result = append(result, transformGroupList(ctx, sourceBlock.Group, diags)...)
		result = append(result, transformDevicePostureList(ctx, sourceBlock.DevicePosture, diags)...)
		result = append(result, transformLoginMethodList(ctx, sourceBlock.LoginMethod, diags)...)
		result = append(result, transformGeoList(ctx, sourceBlock.Geo, diags)...)

		// String scalar selectors: Wrap in object
		if !sourceBlock.CommonName.IsNull() && !sourceBlock.CommonName.IsUnknown() {
			if strings.TrimSpace(sourceBlock.CommonName.ValueString()) != "" {
				result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
					CommonName: &TargetV5CommonNameModel{
						CommonName: sourceBlock.CommonName,
					},
				})
			}
		}

		if !sourceBlock.AuthMethod.IsNull() && !sourceBlock.AuthMethod.IsUnknown() {
			if strings.TrimSpace(sourceBlock.AuthMethod.ValueString()) != "" {
				result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
					AuthMethod: &TargetV5AuthMethodModel{
						AuthMethod: sourceBlock.AuthMethod,
					},
				})
			}
		}

		// Boolean selectors: Convert to empty objects
		if !sourceBlock.Everyone.IsNull() && !sourceBlock.Everyone.IsUnknown() && sourceBlock.Everyone.ValueBool() {
			result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
				Everyone: &TargetV5EveryoneModel{},
			})
		}

		if !sourceBlock.Certificate.IsNull() && !sourceBlock.Certificate.IsUnknown() && sourceBlock.Certificate.ValueBool() {
			result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
				Certificate: &TargetV5CertificateModel{},
			})
		}

		if !sourceBlock.AnyValidServiceToken.IsNull() && !sourceBlock.AnyValidServiceToken.IsUnknown() && sourceBlock.AnyValidServiceToken.ValueBool() {
			result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
				AnyValidServiceToken: &TargetV5AnyValidServiceTokenModel{},
			})
		}

		// Special case: common_names overflow array (split into multiple common_name objects)
		if !sourceBlock.CommonNames.IsNull() && !sourceBlock.CommonNames.IsUnknown() {
			var commonNames []string
			d := sourceBlock.CommonNames.ElementsAs(ctx, &commonNames, false)
			diags.Append(d...)
			if !diags.HasError() {
				for _, cn := range commonNames {
					result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
						CommonName: &TargetV5CommonNameModel{
							CommonName: types.StringValue(cn),
						},
					})
				}
			}
		}

		// Complex nested object selectors (SDKv2 stores these as arrays with MaxItems:1)
		// GitHub: Rename to github_organization and handle teams list
		if len(sourceBlock.GitHub) > 0 {
			result = append(result, transformGitHub(ctx, sourceBlock.GitHub[0], diags)...)
		}

		// GSuite: Unwrap email list
		if len(sourceBlock.GSuite) > 0 {
			result = append(result, transformGSuite(ctx, sourceBlock.GSuite[0], diags)...)
		}

		// Azure: Rename to azure_ad and unwrap id list
		if len(sourceBlock.Azure) > 0 {
			result = append(result, transformAzure(ctx, sourceBlock.Azure[0], diags)...)
		}

		// Okta: Unwrap name list
		if len(sourceBlock.Okta) > 0 {
			result = append(result, transformOkta(ctx, sourceBlock.Okta[0], diags)...)
		}

		// SAML: No transformation needed (same structure)
		if len(sourceBlock.SAML) > 0 {
			result = append(result, transformSAML(sourceBlock.SAML[0]))
		}

		// ExternalEvaluation: No transformation needed (same structure)
		if len(sourceBlock.ExternalEvaluation) > 0 {
			result = append(result, transformExternalEvaluation(sourceBlock.ExternalEvaluation[0]))
		}

		// AuthContext: No transformation needed (same structure)
		if len(sourceBlock.AuthContext) > 0 {
			result = append(result, transformAuthContext(sourceBlock.AuthContext[0]))
		}
	}

	// Filter out any selectors that are effectively empty (all fields are nil/null).
	// This ensures the migrated state matches what the provider will return on Read,
	// preventing drift from empty selector objects being normalized to nil.
	filtered := make([]*TargetV5ZeroTrustAccessGroupIncludeModel, 0, len(result))
	for _, selector := range result {
		if !isSelectorEmpty(selector) {
			filtered = append(filtered, selector)
		}
	}

	return filtered
}

// ============================================================================
// Simple List Transformers
// ============================================================================

func transformEmailList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var emails []string
	d := sourceList.ElementsAs(ctx, &emails, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, email := range emails {
		if strings.TrimSpace(email) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			Email: &TargetV5EmailModel{
				Email: types.StringValue(email),
			},
		})
	}
	return result
}

func transformEmailDomainList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var domains []string
	d := sourceList.ElementsAs(ctx, &domains, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, domain := range domains {
		if strings.TrimSpace(domain) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			EmailDomain: &TargetV5EmailDomainModel{
				Domain: types.StringValue(domain),
			},
		})
	}
	return result
}

func transformEmailListList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var ids []string
	d := sourceList.ElementsAs(ctx, &ids, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, id := range ids {
		if strings.TrimSpace(id) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			EmailList: &TargetV5EmailListModel{
				ID: types.StringValue(id),
			},
		})
	}
	return result
}

func transformIPList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var ips []string
	d := sourceList.ElementsAs(ctx, &ips, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, ip := range ips {
		if strings.TrimSpace(ip) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			IP: &TargetV5IPModel{
				IP: types.StringValue(ip),
			},
		})
	}
	return result
}

func transformIPListList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var ids []string
	d := sourceList.ElementsAs(ctx, &ids, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, id := range ids {
		if strings.TrimSpace(id) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			IPList: &TargetV5IPListModel{
				ID: types.StringValue(id),
			},
		})
	}
	return result
}

func transformServiceTokenList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var tokenIDs []string
	d := sourceList.ElementsAs(ctx, &tokenIDs, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, tokenID := range tokenIDs {
		if strings.TrimSpace(tokenID) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			ServiceToken: &TargetV5ServiceTokenModel{
				TokenID: types.StringValue(tokenID),
			},
		})
	}
	return result
}

func transformGroupList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var ids []string
	d := sourceList.ElementsAs(ctx, &ids, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, id := range ids {
		if strings.TrimSpace(id) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			Group: &TargetV5GroupModel{
				ID: types.StringValue(id),
			},
		})
	}
	return result
}

func transformDevicePostureList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var uids []string
	d := sourceList.ElementsAs(ctx, &uids, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, uid := range uids {
		if strings.TrimSpace(uid) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			DevicePosture: &TargetV5DevicePostureModel{
				IntegrationUID: types.StringValue(uid),
			},
		})
	}
	return result
}

func transformLoginMethodList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var ids []string
	d := sourceList.ElementsAs(ctx, &ids, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, id := range ids {
		if strings.TrimSpace(id) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			LoginMethod: &TargetV5LoginMethodModel{
				ID: types.StringValue(id),
			},
		})
	}
	return result
}

func transformGeoList(ctx context.Context, sourceList types.List, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	if sourceList.IsNull() || sourceList.IsUnknown() {
		return nil
	}

	var countryCodes []string
	d := sourceList.ElementsAs(ctx, &countryCodes, false)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}

	var result []*TargetV5ZeroTrustAccessGroupIncludeModel
	for _, countryCode := range countryCodes {
		if strings.TrimSpace(countryCode) == "" {
			continue
		}
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			Geo: &TargetV5GeoModel{
				CountryCode: types.StringValue(countryCode),
			},
		})
	}
	return result
}

// ============================================================================
// Complex Nested Object Transformers
// ============================================================================

// transformGitHub handles the GitHub -> GitHubOrganization rename and teams list explosion
func transformGitHub(ctx context.Context, source SourceV4GitHubModel, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	var result []*TargetV5ZeroTrustAccessGroupIncludeModel

	// Extract teams list
	var teams []string
	if !source.Teams.IsNull() && !source.Teams.IsUnknown() {
		d := source.Teams.ElementsAs(ctx, &teams, false)
		diags.Append(d...)
		if diags.HasError() {
			return nil
		}
	}

	// If teams list is empty, create one object without team field
	if len(teams) == 0 {
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			GitHubOrganization: &TargetV5GitHubOrganizationModel{
				Name:               source.Name,
				IdentityProviderID: source.IdentityProviderID,
				Team:               types.StringNull(), // No team specified
			},
		})
	} else {
		// Create one object per team
		for _, team := range teams {
			result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
				GitHubOrganization: &TargetV5GitHubOrganizationModel{
					Name:               source.Name,
					IdentityProviderID: source.IdentityProviderID,
					Team:               types.StringValue(team),
				},
			})
		}
	}

	return result
}

// transformGSuite unwraps the email list (takes first element)
func transformGSuite(ctx context.Context, source SourceV4GSuiteModel, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	var result []*TargetV5ZeroTrustAccessGroupIncludeModel

	// Extract email list
	var emails []string
	if !source.Email.IsNull() && !source.Email.IsUnknown() {
		d := source.Email.ElementsAs(ctx, &emails, false)
		diags.Append(d...)
		if diags.HasError() {
			return nil
		}
	}

	// Create one object per email
	for _, email := range emails {
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			GSuite: &TargetV5GSuiteModel{
				Email:              types.StringValue(email),
				IdentityProviderID: source.IdentityProviderID,
			},
		})
	}

	return result
}

// transformAzure handles the Azure -> AzureAD rename and id list explosion
func transformAzure(ctx context.Context, source SourceV4AzureModel, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	var result []*TargetV5ZeroTrustAccessGroupIncludeModel

	// Extract ID list
	var ids []string
	if !source.ID.IsNull() && !source.ID.IsUnknown() {
		d := source.ID.ElementsAs(ctx, &ids, false)
		diags.Append(d...)
		if diags.HasError() {
			return nil
		}
	}

	// Create one object per ID
	for _, id := range ids {
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			AzureAD: &TargetV5AzureADModel{
				ID:                 types.StringValue(id),
				IdentityProviderID: source.IdentityProviderID,
			},
		})
	}

	return result
}

// transformOkta unwraps the name list
func transformOkta(ctx context.Context, source SourceV4OktaModel, diags *diag.Diagnostics) []*TargetV5ZeroTrustAccessGroupIncludeModel {
	var result []*TargetV5ZeroTrustAccessGroupIncludeModel

	// Extract name list
	var names []string
	if !source.Name.IsNull() && !source.Name.IsUnknown() {
		d := source.Name.ElementsAs(ctx, &names, false)
		diags.Append(d...)
		if diags.HasError() {
			return nil
		}
	}

	// Create one object per name
	for _, name := range names {
		result = append(result, &TargetV5ZeroTrustAccessGroupIncludeModel{
			Okta: &TargetV5OktaModel{
				Name:               types.StringValue(name),
				IdentityProviderID: source.IdentityProviderID,
			},
		})
	}

	return result
}

// transformSAML requires no transformation (same structure in v4 and v5)
func transformSAML(source SourceV4SAMLModel) *TargetV5ZeroTrustAccessGroupIncludeModel {
	return &TargetV5ZeroTrustAccessGroupIncludeModel{
		SAML: &TargetV5SAMLModel{
			AttributeName:      source.AttributeName,
			AttributeValue:     source.AttributeValue,
			IdentityProviderID: source.IdentityProviderID,
		},
	}
}

// transformExternalEvaluation requires no transformation (same structure in v4 and v5)
func transformExternalEvaluation(source SourceV4ExternalEvaluationModel) *TargetV5ZeroTrustAccessGroupIncludeModel {
	return &TargetV5ZeroTrustAccessGroupIncludeModel{
		ExternalEvaluation: &TargetV5ExternalEvaluationModel{
			EvaluateURL: source.EvaluateURL,
			KeysURL:     source.KeysURL,
		},
	}
}

// transformAuthContext requires no transformation (same structure in v4 and v5)
func transformAuthContext(source SourceV4AuthContextModel) *TargetV5ZeroTrustAccessGroupIncludeModel {
	return &TargetV5ZeroTrustAccessGroupIncludeModel{
		AuthContext: &TargetV5AuthContextModel{
			ID:                 source.ID,
			AcID:               source.AcID,
			IdentityProviderID: source.IdentityProviderID,
		},
	}
}

// ============================================================================
// Helper Functions
// ============================================================================

// isSelectorEmpty checks if a selector object is effectively empty (all fields are nil or contain only null values).
// This is used to filter out empty selectors that would be normalized to nil by the framework,
// preventing drift on the first plan after migration.
func isSelectorEmpty(selector *TargetV5ZeroTrustAccessGroupIncludeModel) bool {
	if selector == nil {
		return true
	}

	// Check if all selector fields are nil
	hasNonNilField := false

	// Check simple nested objects (boolean selectors - these are valid if non-nil)
	if selector.Everyone != nil || selector.Certificate != nil || selector.AnyValidServiceToken != nil {
		hasNonNilField = true
	}

	// Check selectors with string fields - these are only valid if the field has a non-null value
	if selector.CommonName != nil && hasNonEmptyString(selector.CommonName.CommonName) {
		hasNonNilField = true
	}
	if selector.AuthMethod != nil && hasNonEmptyString(selector.AuthMethod.AuthMethod) {
		hasNonNilField = true
	}
	if selector.Email != nil && hasNonEmptyString(selector.Email.Email) {
		hasNonNilField = true
	}
	if selector.EmailDomain != nil && hasNonEmptyString(selector.EmailDomain.Domain) {
		hasNonNilField = true
	}
	if selector.EmailList != nil && hasNonEmptyString(selector.EmailList.ID) {
		hasNonNilField = true
	}
	if selector.IP != nil && hasNonEmptyString(selector.IP.IP) {
		hasNonNilField = true
	}
	if selector.IPList != nil && hasNonEmptyString(selector.IPList.ID) {
		hasNonNilField = true
	}
	if selector.Geo != nil && hasNonEmptyString(selector.Geo.CountryCode) {
		hasNonNilField = true
	}
	if selector.Group != nil && hasNonEmptyString(selector.Group.ID) {
		hasNonNilField = true
	}
	if selector.DevicePosture != nil && hasNonEmptyString(selector.DevicePosture.IntegrationUID) {
		hasNonNilField = true
	}
	if selector.LoginMethod != nil && hasNonEmptyString(selector.LoginMethod.ID) {
		hasNonNilField = true
	}
	if selector.ServiceToken != nil && hasNonEmptyString(selector.ServiceToken.TokenID) {
		hasNonNilField = true
	}

	// Complex nested objects - these are valid if non-nil (we assume they have valid fields)
	if selector.AuthContext != nil || selector.AzureAD != nil || selector.ExternalEvaluation != nil ||
		selector.GitHubOrganization != nil || selector.GSuite != nil || selector.Okta != nil ||
		selector.SAML != nil || selector.OIDC != nil || selector.LinkedAppToken != nil {
		hasNonNilField = true
	}

	return !hasNonNilField
}

func hasNonEmptyString(v types.String) bool {
	if v.IsNull() || v.IsUnknown() {
		return false
	}
	return strings.TrimSpace(v.ValueString()) != ""
}
