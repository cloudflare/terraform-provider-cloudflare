package zero_trust_access_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IsNull interface {
	IsNull() bool
}

func normalizeEmptyAndNullSlice[T any](data **[]T, stateData *[]T) {
	if (*data != nil && len(**data) != 0) || (stateData != nil && len(*stateData) != 0) {
		return
	}
	*data = stateData
}

func isNullOrUnknown(s types.String) bool {
	return s.IsNull() || s.IsUnknown()
}

// nullifyEmptySelectors sets pointer-to-struct fields to nil when all their string
// fields are null/unknown. This prevents drift from the API returning empty nested
// objects like common_name:{common_name:null} instead of common_name:null.
func nullifyEmptySelectors(item *ZeroTrustAccessGroupIncludeModel) {
	if item.Group != nil && isNullOrUnknown(item.Group.ID) {
		item.Group = nil
	}
	if item.AuthContext != nil && isNullOrUnknown(item.AuthContext.ID) && isNullOrUnknown(item.AuthContext.AcID) && isNullOrUnknown(item.AuthContext.IdentityProviderID) {
		item.AuthContext = nil
	}
	if item.AuthMethod != nil && isNullOrUnknown(item.AuthMethod.AuthMethod) {
		item.AuthMethod = nil
	}
	if item.AzureAD != nil && isNullOrUnknown(item.AzureAD.ID) && isNullOrUnknown(item.AzureAD.IdentityProviderID) {
		item.AzureAD = nil
	}
	if item.CommonName != nil && isNullOrUnknown(item.CommonName.CommonName) {
		item.CommonName = nil
	}
	if item.Geo != nil && isNullOrUnknown(item.Geo.CountryCode) {
		item.Geo = nil
	}
	if item.DevicePosture != nil && isNullOrUnknown(item.DevicePosture.IntegrationUID) {
		item.DevicePosture = nil
	}
	if item.EmailDomain != nil && isNullOrUnknown(item.EmailDomain.Domain) {
		item.EmailDomain = nil
	}
	if item.EmailList != nil && isNullOrUnknown(item.EmailList.ID) {
		item.EmailList = nil
	}
	if item.Email != nil && isNullOrUnknown(item.Email.Email) {
		item.Email = nil
	}
	if item.ExternalEvaluation != nil && isNullOrUnknown(item.ExternalEvaluation.EvaluateURL) && isNullOrUnknown(item.ExternalEvaluation.KeysURL) {
		item.ExternalEvaluation = nil
	}
	if item.GitHubOrganization != nil && isNullOrUnknown(item.GitHubOrganization.IdentityProviderID) && isNullOrUnknown(item.GitHubOrganization.Name) {
		item.GitHubOrganization = nil
	}
	if item.GSuite != nil && isNullOrUnknown(item.GSuite.Email) && isNullOrUnknown(item.GSuite.IdentityProviderID) {
		item.GSuite = nil
	}
	if item.LoginMethod != nil && isNullOrUnknown(item.LoginMethod.ID) {
		item.LoginMethod = nil
	}
	if item.IPList != nil && isNullOrUnknown(item.IPList.ID) {
		item.IPList = nil
	}
	if item.IP != nil && isNullOrUnknown(item.IP.IP) {
		item.IP = nil
	}
	if item.Okta != nil && isNullOrUnknown(item.Okta.IdentityProviderID) && isNullOrUnknown(item.Okta.Name) {
		item.Okta = nil
	}
	if item.SAML != nil && isNullOrUnknown(item.SAML.AttributeName) && isNullOrUnknown(item.SAML.AttributeValue) && isNullOrUnknown(item.SAML.IdentityProviderID) {
		item.SAML = nil
	}
	if item.OIDC != nil && isNullOrUnknown(item.OIDC.ClaimName) && isNullOrUnknown(item.OIDC.ClaimValue) && isNullOrUnknown(item.OIDC.IdentityProviderID) {
		item.OIDC = nil
	}
	if item.ServiceToken != nil && isNullOrUnknown(item.ServiceToken.TokenID) {
		item.ServiceToken = nil
	}
	if item.LinkedAppToken != nil && isNullOrUnknown(item.LinkedAppToken.AppUID) {
		item.LinkedAppToken = nil
	}
	// AnyValidServiceToken, Certificate, Everyone are empty structs — non-nil IS the value
}

func nullifyEmptySelectorsExclude(item *ZeroTrustAccessGroupExcludeModel) {
	if item.Group != nil && isNullOrUnknown(item.Group.ID) {
		item.Group = nil
	}
	if item.AuthContext != nil && isNullOrUnknown(item.AuthContext.ID) && isNullOrUnknown(item.AuthContext.AcID) && isNullOrUnknown(item.AuthContext.IdentityProviderID) {
		item.AuthContext = nil
	}
	if item.AuthMethod != nil && isNullOrUnknown(item.AuthMethod.AuthMethod) {
		item.AuthMethod = nil
	}
	if item.AzureAD != nil && isNullOrUnknown(item.AzureAD.ID) && isNullOrUnknown(item.AzureAD.IdentityProviderID) {
		item.AzureAD = nil
	}
	if item.CommonName != nil && isNullOrUnknown(item.CommonName.CommonName) {
		item.CommonName = nil
	}
	if item.Geo != nil && isNullOrUnknown(item.Geo.CountryCode) {
		item.Geo = nil
	}
	if item.DevicePosture != nil && isNullOrUnknown(item.DevicePosture.IntegrationUID) {
		item.DevicePosture = nil
	}
	if item.EmailDomain != nil && isNullOrUnknown(item.EmailDomain.Domain) {
		item.EmailDomain = nil
	}
	if item.EmailList != nil && isNullOrUnknown(item.EmailList.ID) {
		item.EmailList = nil
	}
	if item.Email != nil && isNullOrUnknown(item.Email.Email) {
		item.Email = nil
	}
	if item.ExternalEvaluation != nil && isNullOrUnknown(item.ExternalEvaluation.EvaluateURL) && isNullOrUnknown(item.ExternalEvaluation.KeysURL) {
		item.ExternalEvaluation = nil
	}
	if item.GitHubOrganization != nil && isNullOrUnknown(item.GitHubOrganization.IdentityProviderID) && isNullOrUnknown(item.GitHubOrganization.Name) {
		item.GitHubOrganization = nil
	}
	if item.GSuite != nil && isNullOrUnknown(item.GSuite.Email) && isNullOrUnknown(item.GSuite.IdentityProviderID) {
		item.GSuite = nil
	}
	if item.LoginMethod != nil && isNullOrUnknown(item.LoginMethod.ID) {
		item.LoginMethod = nil
	}
	if item.IPList != nil && isNullOrUnknown(item.IPList.ID) {
		item.IPList = nil
	}
	if item.IP != nil && isNullOrUnknown(item.IP.IP) {
		item.IP = nil
	}
	if item.Okta != nil && isNullOrUnknown(item.Okta.IdentityProviderID) && isNullOrUnknown(item.Okta.Name) {
		item.Okta = nil
	}
	if item.SAML != nil && isNullOrUnknown(item.SAML.AttributeName) && isNullOrUnknown(item.SAML.AttributeValue) && isNullOrUnknown(item.SAML.IdentityProviderID) {
		item.SAML = nil
	}
	if item.OIDC != nil && isNullOrUnknown(item.OIDC.ClaimName) && isNullOrUnknown(item.OIDC.ClaimValue) && isNullOrUnknown(item.OIDC.IdentityProviderID) {
		item.OIDC = nil
	}
	if item.ServiceToken != nil && isNullOrUnknown(item.ServiceToken.TokenID) {
		item.ServiceToken = nil
	}
	if item.LinkedAppToken != nil && isNullOrUnknown(item.LinkedAppToken.AppUID) {
		item.LinkedAppToken = nil
	}
}

func nullifyEmptySelectorsRequire(item *ZeroTrustAccessGroupRequireModel) {
	if item.Group != nil && isNullOrUnknown(item.Group.ID) {
		item.Group = nil
	}
	if item.AuthContext != nil && isNullOrUnknown(item.AuthContext.ID) && isNullOrUnknown(item.AuthContext.AcID) && isNullOrUnknown(item.AuthContext.IdentityProviderID) {
		item.AuthContext = nil
	}
	if item.AuthMethod != nil && isNullOrUnknown(item.AuthMethod.AuthMethod) {
		item.AuthMethod = nil
	}
	if item.AzureAD != nil && isNullOrUnknown(item.AzureAD.ID) && isNullOrUnknown(item.AzureAD.IdentityProviderID) {
		item.AzureAD = nil
	}
	if item.CommonName != nil && isNullOrUnknown(item.CommonName.CommonName) {
		item.CommonName = nil
	}
	if item.Geo != nil && isNullOrUnknown(item.Geo.CountryCode) {
		item.Geo = nil
	}
	if item.DevicePosture != nil && isNullOrUnknown(item.DevicePosture.IntegrationUID) {
		item.DevicePosture = nil
	}
	if item.EmailDomain != nil && isNullOrUnknown(item.EmailDomain.Domain) {
		item.EmailDomain = nil
	}
	if item.EmailList != nil && isNullOrUnknown(item.EmailList.ID) {
		item.EmailList = nil
	}
	if item.Email != nil && isNullOrUnknown(item.Email.Email) {
		item.Email = nil
	}
	if item.ExternalEvaluation != nil && isNullOrUnknown(item.ExternalEvaluation.EvaluateURL) && isNullOrUnknown(item.ExternalEvaluation.KeysURL) {
		item.ExternalEvaluation = nil
	}
	if item.GitHubOrganization != nil && isNullOrUnknown(item.GitHubOrganization.IdentityProviderID) && isNullOrUnknown(item.GitHubOrganization.Name) {
		item.GitHubOrganization = nil
	}
	if item.GSuite != nil && isNullOrUnknown(item.GSuite.Email) && isNullOrUnknown(item.GSuite.IdentityProviderID) {
		item.GSuite = nil
	}
	if item.LoginMethod != nil && isNullOrUnknown(item.LoginMethod.ID) {
		item.LoginMethod = nil
	}
	if item.IPList != nil && isNullOrUnknown(item.IPList.ID) {
		item.IPList = nil
	}
	if item.IP != nil && isNullOrUnknown(item.IP.IP) {
		item.IP = nil
	}
	if item.Okta != nil && isNullOrUnknown(item.Okta.IdentityProviderID) && isNullOrUnknown(item.Okta.Name) {
		item.Okta = nil
	}
	if item.SAML != nil && isNullOrUnknown(item.SAML.AttributeName) && isNullOrUnknown(item.SAML.AttributeValue) && isNullOrUnknown(item.SAML.IdentityProviderID) {
		item.SAML = nil
	}
	if item.OIDC != nil && isNullOrUnknown(item.OIDC.ClaimName) && isNullOrUnknown(item.OIDC.ClaimValue) && isNullOrUnknown(item.OIDC.IdentityProviderID) {
		item.OIDC = nil
	}
	if item.ServiceToken != nil && isNullOrUnknown(item.ServiceToken.TokenID) {
		item.ServiceToken = nil
	}
	if item.LinkedAppToken != nil && isNullOrUnknown(item.LinkedAppToken.AppUID) {
		item.LinkedAppToken = nil
	}
}

func normalizeIncludeItems(items *[]*ZeroTrustAccessGroupIncludeModel) {
	if items == nil {
		return
	}
	for _, item := range *items {
		if item != nil {
			nullifyEmptySelectors(item)
		}
	}
}

func normalizeExcludeItems(items *[]*ZeroTrustAccessGroupExcludeModel) {
	if items == nil {
		return
	}
	for _, item := range *items {
		if item != nil {
			nullifyEmptySelectorsExclude(item)
		}
	}
}

func normalizeRequireItems(items *[]*ZeroTrustAccessGroupRequireModel) {
	if items == nil {
		return
	}
	for _, item := range *items {
		if item != nil {
			nullifyEmptySelectorsRequire(item)
		}
	}
}

// Normalizing function to ensure consistency between the state/plan and the meaning of the API response.
// Alters the API response before applying it to the state by laxing equalities between null & zero-value
// for some attributes, and nullifies fields that terraform should not be saving in the state.
func normalizeReadZeroTrustAccessGroupAPIData(ctx context.Context, data, sourceData *ZeroTrustAccessGroupModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	normalizeIncludeItems(data.Include)
	normalizeExcludeItems(data.Exclude)
	normalizeRequireItems(data.Require)

	normalizeEmptyAndNullSlice(&data.Include, sourceData.Include)
	normalizeEmptyAndNullSlice(&data.Require, sourceData.Require)
	normalizeEmptyAndNullSlice(&data.Exclude, sourceData.Exclude)

	return diags
}

func normalizeImportZeroTrustAccessGroupAPIData(ctx context.Context, data *ZeroTrustAccessGroupModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	normalizeIncludeItems(data.Include)
	normalizeExcludeItems(data.Exclude)
	normalizeRequireItems(data.Require)

	if data.Include != nil && len(*data.Include) == 0 {
		data.Include = nil
	}

	if data.Require != nil && len(*data.Require) == 0 {
		data.Require = nil
	}

	if data.Exclude != nil && len(*data.Exclude) == 0 {
		data.Exclude = nil
	}

	return diags
}
