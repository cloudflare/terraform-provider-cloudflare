// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessPolicyResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustAccessPolicyResource)(nil)

func (r *ZeroTrustAccessPolicyResource) MoveState(ctx context.Context) []resource.StateMover {
	v4Schema := V4AccessPolicySchema()

	return []resource.StateMover{
		{
			SourceSchema: &v4Schema,
			StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
				if req.SourceTypeName != "cloudflare_access_policy" {
					return
				}

				var v4 V4AccessPolicyModel
				resp.Diagnostics.Append(req.SourceState.Get(ctx, &v4)...)
				if resp.Diagnostics.HasError() {
					return
				}

				v5, diags := transformV4ToV5AccessPolicy(ctx, v4)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.TargetState.Set(ctx, v5)...)
			},
		},
	}
}

func (r *ZeroTrustAccessPolicyResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v4Schema := V4AccessPolicySchema()

	return map[int64]resource.StateUpgrader{
		// Version 0: Handle v4 SDK state that was stored as version 0
		0: {
			PriorSchema: &v4Schema,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var v4 V4AccessPolicyModel
				resp.Diagnostics.Append(req.State.Get(ctx, &v4)...)
				if resp.Diagnostics.HasError() {
					return
				}

				v5, diags := transformV4ToV5AccessPolicy(ctx, v4)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, v5)...)
			},
		},
	}
}

// transformV4ToV5AccessPolicy transforms a v4 cloudflare_access_policy state to v5 format.
func transformV4ToV5AccessPolicy(ctx context.Context, v4 V4AccessPolicyModel) (*ZeroTrustAccessPolicyModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	v5 := &ZeroTrustAccessPolicyModel{
		// Direct copies
		ID:                           v4.ID,
		AccountID:                    v4.AccountID,
		Name:                         v4.Name,
		Decision:                     v4.Decision,
		SessionDuration:              v4.SessionDuration,
		IsolationRequired:            v4.IsolationRequired,
		PurposeJustificationRequired: v4.PurposeJustificationRequired,
		PurposeJustificationPrompt:   v4.PurposeJustificationPrompt,
		ApprovalRequired:             v4.ApprovalRequired,
	}

	// Transform approval_group -> approval_groups
	if len(v4.ApprovalGroup) > 0 {
		approvalGroups := make([]*ZeroTrustAccessPolicyApprovalGroupsModel, len(v4.ApprovalGroup))
		for i, ag := range v4.ApprovalGroup {
			// Convert email_addresses from types.List to *[]types.String
			var emailAddresses *[]types.String
			if !ag.EmailAddresses.IsNull() && !ag.EmailAddresses.IsUnknown() {
				var emails []types.String
				diags.Append(ag.EmailAddresses.ElementsAs(ctx, &emails, false)...)
				if len(emails) > 0 {
					emailAddresses = &emails
				}
			}

			approvalGroups[i] = &ZeroTrustAccessPolicyApprovalGroupsModel{
				ApprovalsNeeded: types.Float64Value(float64(ag.ApprovalsNeeded.ValueInt64())),
				EmailAddresses:  emailAddresses,
				EmailListUUID:   ag.EmailListUUID,
			}
		}
		v5.ApprovalGroups = &approvalGroups
	}

	// Transform include conditions
	includeConditions, d := transformV4ConditionsToV5Include(ctx, v4.Include)
	diags.Append(d...)
	v5.Include = includeConditions

	// Transform exclude conditions
	excludeConditions, d := transformV4ConditionsToV5Exclude(ctx, v4.Exclude)
	diags.Append(d...)
	v5.Exclude = excludeConditions

	// Transform require conditions
	requireConditions, d := transformV4ConditionsToV5Require(ctx, v4.Require)
	diags.Append(d...)
	v5.Require = requireConditions

	return v5, diags
}

// transformV4ConditionsToV5Include transforms v4 condition groups to v5 Include format
func transformV4ConditionsToV5Include(ctx context.Context, v4Conditions []V4ConditionGroupModel) (customfield.NestedObjectSet[ZeroTrustAccessPolicyIncludeModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(v4Conditions) == 0 {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyIncludeModel](ctx), nil
	}

	var v5Conditions []ZeroTrustAccessPolicyIncludeModel

	for _, condGroup := range v4Conditions {
		// Transform boolean conditions
		if !condGroup.Everyone.IsNull() && condGroup.Everyone.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
				Everyone: &ZeroTrustAccessPolicyIncludeEveryoneModel{},
			})
		}

		if !condGroup.Certificate.IsNull() && condGroup.Certificate.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
				Certificate: &ZeroTrustAccessPolicyIncludeCertificateModel{},
			})
		}

		if !condGroup.AnyValidServiceToken.IsNull() && condGroup.AnyValidServiceToken.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
				AnyValidServiceToken: &ZeroTrustAccessPolicyIncludeAnyValidServiceTokenModel{},
			})
		}

		// Transform array conditions - each element becomes a separate condition object
		// Email
		if !condGroup.Email.IsNull() && !condGroup.Email.IsUnknown() {
			var emails []string
			diags.Append(condGroup.Email.ElementsAs(ctx, &emails, false)...)
			for _, email := range emails {
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					Email: &ZeroTrustAccessPolicyIncludeEmailModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					Group: &ZeroTrustAccessPolicyIncludeGroupModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					IP: &ZeroTrustAccessPolicyIncludeIPModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					EmailDomain: &ZeroTrustAccessPolicyIncludeEmailDomainModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					Geo: &ZeroTrustAccessPolicyIncludeGeoModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					LoginMethod: &ZeroTrustAccessPolicyIncludeLoginMethodModel{
						ID: types.StringValue(method),
					},
				})
			}
		}

		// Common name (single value)
		if !condGroup.CommonName.IsNull() && !condGroup.CommonName.IsUnknown() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
				CommonName: &ZeroTrustAccessPolicyIncludeCommonNameModel{
					CommonName: condGroup.CommonName,
				},
			})
		}

		// Auth method (single value)
		if !condGroup.AuthMethod.IsNull() && !condGroup.AuthMethod.IsUnknown() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
				AuthMethod: &ZeroTrustAccessPolicyIncludeAuthMethodModel{
					AuthMethod: condGroup.AuthMethod,
				},
			})
		}

		// Device posture (simple string list in v4)
		if !condGroup.DevicePosture.IsNull() && !condGroup.DevicePosture.IsUnknown() {
			var postures []string
			diags.Append(condGroup.DevicePosture.ElementsAs(ctx, &postures, false)...)
			for _, posture := range postures {
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					DevicePosture: &ZeroTrustAccessPolicyIncludeDevicePostureModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					EmailList: &ZeroTrustAccessPolicyIncludeEmailListModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					IPList: &ZeroTrustAccessPolicyIncludeIPListModel{
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
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyIncludeModel{
					ServiceToken: &ZeroTrustAccessPolicyIncludeServiceTokenModel{
						TokenID: types.StringValue(token),
					},
				})
			}
		}
	}

	if len(v5Conditions) == 0 {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyIncludeModel](ctx), diags
	}

	result, d := customfield.NewObjectSet(ctx, v5Conditions)
	diags.Append(d...)
	if diags.HasError() {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyIncludeModel](ctx), diags
	}

	return result, diags
}

// transformV4ConditionsToV5Exclude transforms v4 condition groups to v5 Exclude format
func transformV4ConditionsToV5Exclude(ctx context.Context, v4Conditions []V4ConditionGroupModel) (customfield.NestedObjectSet[ZeroTrustAccessPolicyExcludeModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(v4Conditions) == 0 {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyExcludeModel](ctx), nil
	}

	var v5Conditions []ZeroTrustAccessPolicyExcludeModel

	for _, condGroup := range v4Conditions {
		// Transform boolean conditions
		if !condGroup.Everyone.IsNull() && condGroup.Everyone.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyExcludeModel{
				Everyone: &ZeroTrustAccessPolicyExcludeEveryoneModel{},
			})
		}

		if !condGroup.Certificate.IsNull() && condGroup.Certificate.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyExcludeModel{
				Certificate: &ZeroTrustAccessPolicyExcludeCertificateModel{},
			})
		}

		if !condGroup.AnyValidServiceToken.IsNull() && condGroup.AnyValidServiceToken.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyExcludeModel{
				AnyValidServiceToken: &ZeroTrustAccessPolicyExcludeAnyValidServiceTokenModel{},
			})
		}

		// Transform array conditions
		if !condGroup.Email.IsNull() && !condGroup.Email.IsUnknown() {
			var emails []string
			diags.Append(condGroup.Email.ElementsAs(ctx, &emails, false)...)
			for _, email := range emails {
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyExcludeModel{
					Email: &ZeroTrustAccessPolicyExcludeEmailModel{
						Email: types.StringValue(email),
					},
				})
			}
		}

		if !condGroup.Group.IsNull() && !condGroup.Group.IsUnknown() {
			var groups []string
			diags.Append(condGroup.Group.ElementsAs(ctx, &groups, false)...)
			for _, group := range groups {
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyExcludeModel{
					Group: &ZeroTrustAccessPolicyExcludeGroupModel{
						ID: types.StringValue(group),
					},
				})
			}
		}
	}

	if len(v5Conditions) == 0 {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyExcludeModel](ctx), diags
	}

	result, d := customfield.NewObjectSet(ctx, v5Conditions)
	diags.Append(d...)
	if diags.HasError() {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyExcludeModel](ctx), diags
	}

	return result, diags
}

// transformV4ConditionsToV5Require transforms v4 condition groups to v5 Require format
func transformV4ConditionsToV5Require(ctx context.Context, v4Conditions []V4ConditionGroupModel) (customfield.NestedObjectSet[ZeroTrustAccessPolicyRequireModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(v4Conditions) == 0 {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyRequireModel](ctx), nil
	}

	var v5Conditions []ZeroTrustAccessPolicyRequireModel

	for _, condGroup := range v4Conditions {
		// Transform boolean conditions
		if !condGroup.Everyone.IsNull() && condGroup.Everyone.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyRequireModel{
				Everyone: &ZeroTrustAccessPolicyRequireEveryoneModel{},
			})
		}

		if !condGroup.Certificate.IsNull() && condGroup.Certificate.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyRequireModel{
				Certificate: &ZeroTrustAccessPolicyRequireCertificateModel{},
			})
		}

		if !condGroup.AnyValidServiceToken.IsNull() && condGroup.AnyValidServiceToken.ValueBool() {
			v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyRequireModel{
				AnyValidServiceToken: &ZeroTrustAccessPolicyRequireAnyValidServiceTokenModel{},
			})
		}

		// Transform array conditions
		if !condGroup.Email.IsNull() && !condGroup.Email.IsUnknown() {
			var emails []string
			diags.Append(condGroup.Email.ElementsAs(ctx, &emails, false)...)
			for _, email := range emails {
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyRequireModel{
					Email: &ZeroTrustAccessPolicyRequireEmailModel{
						Email: types.StringValue(email),
					},
				})
			}
		}

		if !condGroup.Group.IsNull() && !condGroup.Group.IsUnknown() {
			var groups []string
			diags.Append(condGroup.Group.ElementsAs(ctx, &groups, false)...)
			for _, group := range groups {
				v5Conditions = append(v5Conditions, ZeroTrustAccessPolicyRequireModel{
					Group: &ZeroTrustAccessPolicyRequireGroupModel{
						ID: types.StringValue(group),
					},
				})
			}
		}
	}

	if len(v5Conditions) == 0 {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyRequireModel](ctx), diags
	}

	result, d := customfield.NewObjectSet(ctx, v5Conditions)
	diags.Append(d...)
	if diags.HasError() {
		return customfield.NullObjectSet[ZeroTrustAccessPolicyRequireModel](ctx), diags
	}

	return result, diags
}

func init() {
	// Ensure at compile time that the struct satisfies the interface
	var _ resource.ResourceWithMoveState = &ZeroTrustAccessPolicyResource{}
	var _ resource.ResourceWithUpgradeState = &ZeroTrustAccessPolicyResource{}
}
