package v500

import (
	"context"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeBoolFalseToNull converts false boolean values to null.
// The v5 provider's API treats false and null as equivalent for these optional boolean fields.
// By normalizing false to null during migration, we prevent drift after the v5 provider refreshes state.
func normalizeBoolFalseToNull(b types.Bool) types.Bool {
	if b.IsNull() || b.IsUnknown() {
		return b
	}
	if !b.ValueBool() {
		// false -> null (they are semantically equivalent)
		return types.BoolNull()
	}
	return b
}

// Transform converts a v4 cloudflare_access_policy state to v5 cloudflare_zero_trust_access_policy state.
func Transform(ctx context.Context, v4 SourceAccessPolicyModel) (*TargetAccessPolicyModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &TargetAccessPolicyModel{
		ID:                           v4.ID,
		AccountID:                    v4.AccountID,
		Name:                         v4.Name,
		Decision:                     v4.Decision,
		SessionDuration:              v4.SessionDuration,
		IsolationRequired:            normalizeBoolFalseToNull(v4.IsolationRequired),
		PurposeJustificationRequired: normalizeBoolFalseToNull(v4.PurposeJustificationRequired),
		PurposeJustificationPrompt:   migrations.FalseyStringToNull(v4.PurposeJustificationPrompt),
		ApprovalRequired:             normalizeBoolFalseToNull(v4.ApprovalRequired),
	}

	// Transform approval_group -> approval_groups
	if len(v4.ApprovalGroup) > 0 {
		approvalGroups := make([]*TargetApprovalGroupsModel, len(v4.ApprovalGroup))
		for i, ag := range v4.ApprovalGroup {
			var emailAddresses *[]types.String
			if !ag.EmailAddresses.IsNull() && !ag.EmailAddresses.IsUnknown() {
				var emails []types.String
				diags.Append(ag.EmailAddresses.ElementsAs(ctx, &emails, false)...)
				if len(emails) > 0 {
					emailAddresses = &emails
				}
			}

			approvalGroups[i] = &TargetApprovalGroupsModel{
				ApprovalsNeeded: types.Float64Value(float64(ag.ApprovalsNeeded.ValueInt64())),
				EmailAddresses:  emailAddresses,
				EmailListUUID:   ag.EmailListUUID,
			}
		}
		v5.ApprovalGroups = &approvalGroups
	}

	// Transform include conditions
	includeConditions, d := transformConditions(ctx, v4.Include)
	diags.Append(d...)
	v5.Include = includeConditions

	// Transform exclude conditions
	excludeConditions, d := transformConditions(ctx, v4.Exclude)
	diags.Append(d...)
	v5.Exclude = excludeConditions

	// Transform require conditions
	requireConditions, d := transformConditions(ctx, v4.Require)
	diags.Append(d...)
	v5.Require = requireConditions

	return v5, diags
}

// transformConditions transforms v4 condition groups to v5 format
func transformConditions(ctx context.Context, v4Conditions []SourceConditionGroupModel) ([]TargetConditionModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(v4Conditions) == 0 {
		return nil, nil
	}

	var v5Conditions []TargetConditionModel

	for _, condGroup := range v4Conditions {
		// Transform boolean conditions
		if !condGroup.Everyone.IsNull() && condGroup.Everyone.ValueBool() {
			v5Conditions = append(v5Conditions, TargetConditionModel{
				Everyone: &TargetEveryoneModel{},
			})
		}

		if !condGroup.Certificate.IsNull() && condGroup.Certificate.ValueBool() {
			v5Conditions = append(v5Conditions, TargetConditionModel{
				Certificate: &TargetCertificateModel{},
			})
		}

		if !condGroup.AnyValidServiceToken.IsNull() && condGroup.AnyValidServiceToken.ValueBool() {
			v5Conditions = append(v5Conditions, TargetConditionModel{
				AnyValidServiceToken: &TargetAnyValidServiceTokenModel{},
			})
		}

		// Transform array conditions - each element becomes a separate condition object
		// Email
		if !condGroup.Email.IsNull() && !condGroup.Email.IsUnknown() {
			var emails []string
			diags.Append(condGroup.Email.ElementsAs(ctx, &emails, false)...)
			for _, email := range emails {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					Email: &TargetEmailModel{
						Email: types.StringValue(email),
					},
				})
			}
		}

		// Group
		if !condGroup.Group.IsNull() && !condGroup.Group.IsUnknown() {
			var groups []string
			diags.Append(condGroup.Group.ElementsAs(ctx, &groups, false)...)
			for _, group := range groups {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					Group: &TargetGroupModel{
						ID: types.StringValue(group),
					},
				})
			}
		}

		// IP
		if !condGroup.IP.IsNull() && !condGroup.IP.IsUnknown() {
			var ips []string
			diags.Append(condGroup.IP.ElementsAs(ctx, &ips, false)...)
			for _, ip := range ips {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					IP: &TargetIPModel{
						IP: types.StringValue(ip),
					},
				})
			}
		}

		// Email domain
		if !condGroup.EmailDomain.IsNull() && !condGroup.EmailDomain.IsUnknown() {
			var domains []string
			diags.Append(condGroup.EmailDomain.ElementsAs(ctx, &domains, false)...)
			for _, domain := range domains {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					EmailDomain: &TargetEmailDomainModel{
						Domain: types.StringValue(domain),
					},
				})
			}
		}

		// Geo
		if !condGroup.Geo.IsNull() && !condGroup.Geo.IsUnknown() {
			var geos []string
			diags.Append(condGroup.Geo.ElementsAs(ctx, &geos, false)...)
			for _, geo := range geos {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					Geo: &TargetGeoModel{
						CountryCode: types.StringValue(geo),
					},
				})
			}
		}

		// Login Method
		if !condGroup.LoginMethod.IsNull() && !condGroup.LoginMethod.IsUnknown() {
			var methods []string
			diags.Append(condGroup.LoginMethod.ElementsAs(ctx, &methods, false)...)
			for _, method := range methods {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					LoginMethod: &TargetLoginMethodModel{
						ID: types.StringValue(method),
					},
				})
			}
		}

		// Common name (single value)
		if !condGroup.CommonName.IsNull() && !condGroup.CommonName.IsUnknown() {
			commonName := strings.TrimSpace(condGroup.CommonName.ValueString())
			if commonName != "" {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					CommonName: &TargetCommonNameModel{
						CommonName: types.StringValue(commonName),
					},
				})
			}
		}

		// Auth method (single value)
		if !condGroup.AuthMethod.IsNull() && !condGroup.AuthMethod.IsUnknown() {
			authMethod := strings.TrimSpace(condGroup.AuthMethod.ValueString())
			if authMethod != "" {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					AuthMethod: &TargetAuthMethodModel{
						AuthMethod: types.StringValue(authMethod),
					},
				})
			}
		}

		// Device posture (simple string list in v4)
		if !condGroup.DevicePosture.IsNull() && !condGroup.DevicePosture.IsUnknown() {
			var postures []string
			diags.Append(condGroup.DevicePosture.ElementsAs(ctx, &postures, false)...)
			for _, posture := range postures {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					DevicePosture: &TargetDevicePostureModel{
						IntegrationUID: types.StringValue(posture),
					},
				})
			}
		}

		// Email list (simple string list in v4)
		if !condGroup.EmailList.IsNull() && !condGroup.EmailList.IsUnknown() {
			var lists []string
			diags.Append(condGroup.EmailList.ElementsAs(ctx, &lists, false)...)
			for _, list := range lists {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					EmailList: &TargetEmailListModel{
						ID: types.StringValue(list),
					},
				})
			}
		}

		// IP list (simple string list in v4)
		if !condGroup.IPList.IsNull() && !condGroup.IPList.IsUnknown() {
			var lists []string
			diags.Append(condGroup.IPList.ElementsAs(ctx, &lists, false)...)
			for _, list := range lists {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					IPList: &TargetIPListModel{
						ID: types.StringValue(list),
					},
				})
			}
		}

		// Service token (simple string list in v4)
		if !condGroup.ServiceToken.IsNull() && !condGroup.ServiceToken.IsUnknown() {
			var tokens []string
			diags.Append(condGroup.ServiceToken.ElementsAs(ctx, &tokens, false)...)
			for _, token := range tokens {
				v5Conditions = append(v5Conditions, TargetConditionModel{
					ServiceToken: &TargetServiceTokenModel{
						TokenID: types.StringValue(token),
					},
				})
			}
		}

		// SAML blocks
		for _, saml := range condGroup.SAML {
			v5Conditions = append(v5Conditions, TargetConditionModel{
				SAML: &TargetSAMLModel{
					AttributeName:      saml.AttributeName,
					AttributeValue:     saml.AttributeValue,
					IdentityProviderID: saml.IdentityProviderID,
				},
			})
		}

		// OIDC blocks
		for _, oidc := range condGroup.OIDC {
			v5Conditions = append(v5Conditions, TargetConditionModel{
				OIDC: &TargetOIDCModel{
					ClaimName:          oidc.ClaimName,
					ClaimValue:         oidc.ClaimValue,
					IdentityProviderID: oidc.IdentityProviderID,
				},
			})
		}

		// Azure AD blocks - expand array to individual conditions
		for _, azure := range condGroup.AzureAD {
			if !azure.ID.IsNull() && !azure.ID.IsUnknown() {
				var ids []string
				diags.Append(azure.ID.ElementsAs(ctx, &ids, false)...)
				for _, id := range ids {
					v5Conditions = append(v5Conditions, TargetConditionModel{
						AzureAD: &TargetAzureADModel{
							ID:                 types.StringValue(id),
							IdentityProviderID: azure.IdentityProviderID,
						},
					})
				}
			}
		}

		// Okta blocks - expand array to individual conditions
		for _, okta := range condGroup.Okta {
			if !okta.Name.IsNull() && !okta.Name.IsUnknown() {
				var names []string
				diags.Append(okta.Name.ElementsAs(ctx, &names, false)...)
				for _, name := range names {
					v5Conditions = append(v5Conditions, TargetConditionModel{
						Okta: &TargetOktaModel{
							Name:               types.StringValue(name),
							IdentityProviderID: okta.IdentityProviderID,
						},
					})
				}
			}
		}

		// GSuite blocks - expand array to individual conditions
		for _, gsuite := range condGroup.GSuite {
			if !gsuite.Email.IsNull() && !gsuite.Email.IsUnknown() {
				var emails []string
				diags.Append(gsuite.Email.ElementsAs(ctx, &emails, false)...)
				for _, email := range emails {
					v5Conditions = append(v5Conditions, TargetConditionModel{
						GSuite: &TargetGSuiteModel{
							Email:              types.StringValue(email),
							IdentityProviderID: gsuite.IdentityProviderID,
						},
					})
				}
			}
		}

		// GitHub blocks - expand teams to individual conditions
		for _, github := range condGroup.GitHub {
			if !github.Teams.IsNull() && !github.Teams.IsUnknown() {
				var teams []string
				diags.Append(github.Teams.ElementsAs(ctx, &teams, false)...)
				for _, team := range teams {
					v5Conditions = append(v5Conditions, TargetConditionModel{
						GitHubOrganization: &TargetGitHubOrganizationModel{
							Name:               github.Name,
							Team:               types.StringValue(team),
							IdentityProviderID: github.IdentityProviderID,
						},
					})
				}
			} else {
				// No teams, just the org
				v5Conditions = append(v5Conditions, TargetConditionModel{
					GitHubOrganization: &TargetGitHubOrganizationModel{
						Name:               github.Name,
						IdentityProviderID: github.IdentityProviderID,
					},
				})
			}
		}

		// External evaluation blocks
		for _, ext := range condGroup.ExternalEvaluation {
			v5Conditions = append(v5Conditions, TargetConditionModel{
				ExternalEvaluation: &TargetExternalEvaluationModel{
					EvaluateURL: ext.EvaluateURL,
					KeysURL:     ext.KeysURL,
				},
			})
		}

		// Auth context blocks
		for _, ac := range condGroup.AuthContext {
			v5Conditions = append(v5Conditions, TargetConditionModel{
				AuthContext: &TargetAuthContextModel{
					ID:                 ac.ID,
					AcID:               ac.AcID,
					IdentityProviderID: ac.IdentityProviderID,
				},
			})
		}
	}

	return v5Conditions, diags
}
