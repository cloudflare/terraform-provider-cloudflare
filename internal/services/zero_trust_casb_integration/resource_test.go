package zero_trust_casb_integration

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const serviceAccountKey = `{
	"type":"service_account",
	"project_id":"project-id",
	"private_key":"private-key",
	"client_email":"service-account@example.com"
}`

var testSecretsSalt = []byte("0123456789abcdef0123456789abcdef")

var serviceAccountCredentialsPayload = map[string]any{
	"type":         "service_account",
	"project_id":   "project-id",
	"private_key":  "private-key",
	"client_email": "service-account@example.com",
}

// ── Credential payloads ─────────────────────────────────────────────────────

// TestNewParamsPerAuthMethod pins the exact wire payload for every supported
// vendor and auth method, and asserts that each argument lands on the correct
// side of the confidentiality split.
func TestNewParamsPerAuthMethod(t *testing.T) {
	tests := map[string]struct {
		model            ZeroTrustCasbIntegrationModel
		wantApplication  string
		wantAuthMethod   string
		wantConfidential map[string]any
		wantPublic       map[string]any
	}{
		"anthropic admin api key": {
			model: ZeroTrustCasbIntegrationModel{
				Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
					AdminAPIKey: &ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel{
						APIKey:   types.StringValue("sk-ant-admin01-key"),
						TenantID: types.StringValue("org_1234"),
					},
				},
			},
			wantApplication:  "ANTHROPIC",
			wantAuthMethod:   authMethodAnthropicAdminAPIKey,
			wantConfidential: map[string]any{"api_key": "sk-ant-admin01-key"},
			wantPublic:       map[string]any{"tenant_id": "org_1234"},
		},
		"anthropic admin api key without tenant": {
			model: ZeroTrustCasbIntegrationModel{
				Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
					AdminAPIKey: &ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel{
						APIKey:   types.StringValue("sk-ant-admin01-key"),
						TenantID: types.StringNull(),
					},
				},
			},
			wantApplication:  "ANTHROPIC",
			wantAuthMethod:   authMethodAnthropicAdminAPIKey,
			wantConfidential: map[string]any{"api_key": "sk-ant-admin01-key"},
			// An omitted optional argument must be absent from the payload
			// rather than sent as an empty string.
			wantPublic: map[string]any{},
		},
		"anthropic workspace api key": {
			model: ZeroTrustCasbIntegrationModel{
				Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
					WorkspaceAPIKey: &ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel{
						APIKey:   types.StringValue("sk-ant-workspace-key"),
						TenantID: types.StringValue("wrkspc_1234"),
					},
				},
			},
			wantApplication:  "ANTHROPIC",
			wantAuthMethod:   authMethodAnthropicWorkspaceAPIKey,
			wantConfidential: map[string]any{"api_key": "sk-ant-workspace-key"},
			wantPublic:       map[string]any{"tenant_id": "wrkspc_1234"},
		},
		"anthropic compliance api key": {
			model: ZeroTrustCasbIntegrationModel{
				Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
					ComplianceAPIKey: &ZeroTrustCasbIntegrationAnthropicComplianceAPIKeyModel{
						ComplianceAPIKey: types.StringValue("sk-ant-compliance-key"),
						TenantID:         types.StringNull(),
					},
				},
			},
			wantApplication:  "ANTHROPIC",
			wantAuthMethod:   authMethodAnthropicComplianceAPIKey,
			wantConfidential: map[string]any{"compliance_api_key": "sk-ant-compliance-key"},
			wantPublic:       map[string]any{},
		},
		// AWS role delegation has no secret at all: the external ID is
		// worthless without Cloudflare's AWS principal.
		"aws iam role": {
			model: ZeroTrustCasbIntegrationModel{
				AWS: &ZeroTrustCasbIntegrationAWSModel{
					IAMRole: &ZeroTrustCasbIntegrationAWSIAMRoleModel{
						RoleARN:    types.StringValue("arn:aws:iam::123456789012:role/Cloudflare_DSPM_Auditor"),
						ExternalID: types.StringValue("external-id"),
					},
				},
			},
			wantApplication:  "AWS",
			wantAuthMethod:   authMethodAWSIAMRole,
			wantConfidential: map[string]any{},
			wantPublic: map[string]any{
				"role_arn":    "arn:aws:iam::123456789012:role/Cloudflare_DSPM_Auditor",
				"external_id": "external-id",
			},
		},
		"box server authentication": {
			model: ZeroTrustCasbIntegrationModel{
				Box: &ZeroTrustCasbIntegrationBoxModel{
					ServerAuthentication: &ZeroTrustCasbIntegrationBoxServerAuthenticationModel{
						EnterpriseID: types.StringValue("1234567"),
					},
				},
			},
			wantApplication:  "BOX",
			wantAuthMethod:   authMethodBoxServerAuthentication,
			wantConfidential: map[string]any{},
			wantPublic:       map[string]any{"enterprise_id": "1234567"},
		},
		"google cloud platform service account": {
			model: ZeroTrustCasbIntegrationModel{
				GoogleCloudPlatform: &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
					ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{
						ServiceAccountKeyJSON: jsontypes.NewNormalizedValue(serviceAccountKey),
					},
				},
			},
			wantApplication: "GOOGLE_CLOUD_PLATFORM",
			wantAuthMethod:  authMethodGoogleCloudPlatformServiceAccount,
			wantConfidential: map[string]any{
				"service_account_credentials": serviceAccountCredentialsPayload,
			},
			wantPublic: map[string]any{},
		},
		"google workspace domain wide delegation": {
			model: ZeroTrustCasbIntegrationModel{
				GoogleWorkspace: &ZeroTrustCasbIntegrationGoogleWorkspaceModel{
					DomainWideDelegationServiceAccount: &ZeroTrustCasbIntegrationGoogleWorkspaceDomainWideDelegationServiceAccountModel{
						ServiceAccountKeyJSON: jsontypes.NewNormalizedValue(serviceAccountKey),
						AdministratorEmail:    types.StringValue("admin@example.com"),
					},
				},
			},
			wantApplication: "GOOGLE_WORKSPACE",
			wantAuthMethod:  authMethodGoogleWorkspaceDomainWideDelegation,
			wantConfidential: map[string]any{
				"service_account_credentials": serviceAccountCredentialsPayload,
			},
			wantPublic: map[string]any{"administrator_email": "admin@example.com"},
		},
		"openai standard api key": {
			model: ZeroTrustCasbIntegrationModel{
				OpenAI: &ZeroTrustCasbIntegrationOpenAIModel{
					StandardAPIKey: &ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel{
						AdminAPIKey:    types.StringValue("sk-admin-key"),
						OrganizationID: types.StringValue("org-1234"),
						ProjectAPIKey:  types.StringValue("sk-proj-key"),
						ProjectID:      types.StringValue("proj_1234"),
					},
				},
			},
			wantApplication: "OPENAI",
			wantAuthMethod:  authMethodOpenAIStandardAPIKey,
			wantConfidential: map[string]any{
				"admin_api_key":   "sk-admin-key",
				"project_api_key": "sk-proj-key",
			},
			wantPublic: map[string]any{
				"organization_id": "org-1234",
				"project_id":      "proj_1234",
			},
		},
		"openai standard api key without project": {
			model: ZeroTrustCasbIntegrationModel{
				OpenAI: &ZeroTrustCasbIntegrationOpenAIModel{
					StandardAPIKey: &ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel{
						AdminAPIKey:    types.StringValue("sk-admin-key"),
						OrganizationID: types.StringValue("org-1234"),
						ProjectAPIKey:  types.StringNull(),
						ProjectID:      types.StringNull(),
					},
				},
			},
			wantApplication:  "OPENAI",
			wantAuthMethod:   authMethodOpenAIStandardAPIKey,
			wantConfidential: map[string]any{"admin_api_key": "sk-admin-key"},
			wantPublic:       map[string]any{"organization_id": "org-1234"},
		},
		"openai compliance api key": {
			model: ZeroTrustCasbIntegrationModel{
				OpenAI: &ZeroTrustCasbIntegrationOpenAIModel{
					ComplianceAPIKey: &ZeroTrustCasbIntegrationOpenAIComplianceAPIKeyModel{
						AdminAPIKey:      types.StringValue("sk-admin-key"),
						OrganizationID:   types.StringValue("org-1234"),
						ComplianceAPIKey: types.StringValue("sk-compliance-key"),
						WorkspaceID:      types.StringValue("wrkspc-1234"),
						ProjectAPIKey:    types.StringNull(),
						ProjectID:        types.StringNull(),
					},
				},
			},
			wantApplication: "OPENAI",
			wantAuthMethod:  authMethodOpenAIComplianceAPIKey,
			wantConfidential: map[string]any{
				"admin_api_key":      "sk-admin-key",
				"compliance_api_key": "sk-compliance-key",
			},
			wantPublic: map[string]any{
				"organization_id": "org-1234",
				"workspace_id":    "wrkspc-1234",
			},
		},
	}

	ctx := context.Background()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			model := test.model
			model.AccountID = types.StringValue("account-id")
			model.Name = types.StringValue("integration")
			model.DLPProfiles = customfield.NullSet[types.String](ctx)
			model.Permissions = customfield.NullSet[types.String](ctx)
			model.UseCases = stringsToSet(ctx, []string{"casb"})

			resolved, err := model.resolveAuth()
			if err != nil {
				t.Fatalf("resolveAuth returned an error: %v", err)
			}
			if resolved.authMethod != test.wantAuthMethod {
				t.Errorf("unexpected auth method: got %q, want %q", resolved.authMethod, test.wantAuthMethod)
			}

			// The confidentiality split is the whole security model: anything
			// in confidentialArgs is covered by the digest and never stored,
			// anything in publicArgs is stored and diffed normally.
			confidential, err := resolved.builder.confidentialArgs()
			if err != nil {
				t.Fatalf("confidentialArgs returned an error: %v", err)
			}
			if !reflect.DeepEqual(confidential, test.wantConfidential) {
				t.Errorf("unexpected confidential arguments\n got: %#v\nwant: %#v",
					confidential, test.wantConfidential)
			}
			if public := resolved.builder.publicArgs(); !reflect.DeepEqual(public, test.wantPublic) {
				t.Errorf("unexpected public arguments\n got: %#v\nwant: %#v", public, test.wantPublic)
			}

			params, diags := model.newParams(ctx, &model)
			if diags.HasError() {
				t.Fatalf("newParams returned errors: %v", diags)
			}
			body, err := params.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON returned an error: %v", err)
			}

			var got map[string]any
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatalf("failed to decode request body: %v", err)
			}

			wantCredentials := map[string]any{}
			for key, value := range test.wantPublic {
				wantCredentials[key] = value
			}
			for key, value := range test.wantConfidential {
				wantCredentials[key] = value
			}
			want := map[string]any{
				"application": test.wantApplication,
				"auth_method": test.wantAuthMethod,
				"credentials": wantCredentials,
				"name":        "integration",
				"use_cases":   []any{"casb"},
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("unexpected request body\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}

// TestNewParamsIncludesCollections checks that the optional collections reach
// the wire, and that an unset collection is omitted rather than sent empty.
func TestNewParamsIncludesCollections(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		dlpProfiles customfield.Set[types.String]
		permissions customfield.Set[types.String]
		useCases    customfield.Set[types.String]
		want        map[string]any
	}{
		"all collections set": {
			dlpProfiles: stringsToSet(ctx, []string{"profile-a"}),
			permissions: stringsToSet(ctx, []string{"read"}),
			useCases:    stringsToSet(ctx, []string{"casb", "ces"}),
			want: map[string]any{
				"dlp_profiles": []any{"profile-a"},
				"permissions":  []any{"read"},
				"use_cases":    []any{"casb", "ces"},
			},
		},
		"only use cases set": {
			dlpProfiles: customfield.NullSet[types.String](ctx),
			permissions: customfield.NullSet[types.String](ctx),
			useCases:    stringsToSet(ctx, []string{"casb"}),
			want: map[string]any{
				"use_cases": []any{"casb"},
			},
		},
		// An unknown collection carries no value yet, so it must not be sent.
		"unknown collections omitted": {
			dlpProfiles: customfield.UnknownSet[types.String](ctx),
			permissions: customfield.UnknownSet[types.String](ctx),
			useCases:    stringsToSet(ctx, []string{"casb"}),
			want: map[string]any{
				"use_cases": []any{"casb"},
			},
		},
		"empty set is sent": {
			dlpProfiles: stringsToSet(ctx, []string{}),
			permissions: customfield.NullSet[types.String](ctx),
			useCases:    stringsToSet(ctx, []string{"casb"}),
			want: map[string]any{
				"dlp_profiles": []any{},
				"use_cases":    []any{"casb"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			model := testResourceModel(ctx, false)
			model.DLPProfiles = test.dlpProfiles
			model.Permissions = test.permissions
			model.UseCases = test.useCases

			params, diags := model.newParams(ctx, &model)
			if diags.HasError() {
				t.Fatalf("newParams returned errors: %v", diags)
			}
			body, err := params.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON returned an error: %v", err)
			}
			var got map[string]any
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatalf("failed to decode request body: %v", err)
			}

			for _, key := range []string{"dlp_profiles", "permissions", "use_cases"} {
				want, wanted := test.want[key]
				value, present := got[key]
				if wanted != present {
					t.Errorf("expected %q present=%t, got present=%t", key, wanted, present)
					continue
				}
				if wanted && !reflect.DeepEqual(value, want) {
					t.Errorf("unexpected %q\n got: %#v\nwant: %#v", key, value, want)
				}
			}
		})
	}
}

func TestCredentialsRejectsInvalidServiceAccountKeyJSON(t *testing.T) {
	tests := map[string]string{
		"invalid JSON":    `{`,
		"non-object JSON": `[]`,
		"null JSON":       `null`,
	}
	for name, jsonKey := range tests {
		t.Run(name, func(t *testing.T) {
			model := ZeroTrustCasbIntegrationModel{
				GoogleCloudPlatform: &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
					ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{
						ServiceAccountKeyJSON: jsontypes.NewNormalizedValue(jsonKey),
					},
				},
			}
			if _, err := model.credentials(); err == nil {
				t.Fatal("expected credentials to reject the service account key")
			}
			// The digest is derived from the same payload, so it has to fail
			// rather than hash an unusable value.
			if _, err := model.secretsDigest(testSecretsSalt); err == nil {
				t.Fatal("expected secretsDigest to reject the service account key")
			}
		})
	}
}

// ── Vendor and auth method selection ────────────────────────────────────────

// TestConfigValidatorsEnforceExactlyOne exercises the real ConfigValidators
// against a populated config.
func TestConfigValidatorsEnforceExactlyOne(t *testing.T) {
	gcp := func() *ZeroTrustCasbIntegrationGoogleCloudPlatformModel {
		return &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
			ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{
				ServiceAccountKeyJSON: jsontypes.NewNormalizedValue(serviceAccountKey),
			},
		}
	}
	adminKey := func() *ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel {
		return &ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel{
			APIKey:   types.StringValue("sk-ant-admin01-key"),
			TenantID: types.StringNull(),
		}
	}
	workspaceKey := func() *ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel {
		return &ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel{
			APIKey:   types.StringValue("sk-ant-workspace-key"),
			TenantID: types.StringValue("wrkspc_1234"),
		}
	}

	tests := map[string]struct {
		model   ZeroTrustCasbIntegrationModel
		wantErr bool
	}{
		"one vendor and one auth method": {
			model: ZeroTrustCasbIntegrationModel{GoogleCloudPlatform: gcp()},
		},
		"openai is selectable": {
			model: ZeroTrustCasbIntegrationModel{
				OpenAI: &ZeroTrustCasbIntegrationOpenAIModel{
					StandardAPIKey: &ZeroTrustCasbIntegrationOpenAIStandardAPIKeyModel{
						AdminAPIKey:    types.StringValue("sk-admin-key"),
						OrganizationID: types.StringValue("org-1234"),
						ProjectAPIKey:  types.StringNull(),
						ProjectID:      types.StringNull(),
					},
				},
			},
		},
		"box is selectable": {
			model: ZeroTrustCasbIntegrationModel{
				Box: &ZeroTrustCasbIntegrationBoxModel{
					ServerAuthentication: &ZeroTrustCasbIntegrationBoxServerAuthenticationModel{
						EnterpriseID: types.StringValue("1234567"),
					},
				},
			},
		},
		"no vendor": {
			model:   ZeroTrustCasbIntegrationModel{},
			wantErr: true,
		},
		"two vendors": {
			model: ZeroTrustCasbIntegrationModel{
				GoogleCloudPlatform: gcp(),
				Anthropic:           &ZeroTrustCasbIntegrationAnthropicModel{AdminAPIKey: adminKey()},
			},
			wantErr: true,
		},
		"one vendor with two auth methods": {
			model: ZeroTrustCasbIntegrationModel{
				Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
					AdminAPIKey:     adminKey(),
					WorkspaceAPIKey: workspaceKey(),
				},
			},
			wantErr: true,
		},
		"vendor without an auth method": {
			model:   ZeroTrustCasbIntegrationModel{Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{}},
			wantErr: true,
		},
	}

	ctx := context.Background()
	r := &ZeroTrustCasbIntegrationResource{}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			model := test.model
			model.AccountID = types.StringValue("account-id")
			model.Name = types.StringValue("integration")
			model.Paused = types.BoolValue(false)
			model.ID = types.StringValue("integration-id")
			model.SecretsDigest = types.StringValue("digest")
			model.DLPProfiles = customfield.NullSet[types.String](ctx)
			model.Permissions = customfield.NullSet[types.String](ctx)
			model.UseCases = stringsToSet(ctx, []string{"casb"})

			// tfsdk.Config has no Set, so build the raw value via State.
			schema := ResourceSchema(ctx)
			state := tfsdk.State{Schema: schema}
			if diags := state.Set(ctx, &model); diags.HasError() {
				t.Fatalf("failed to prepare config: %v", diags)
			}
			config := tfsdk.Config{Schema: schema, Raw: state.Raw}

			resp := &frameworkresource.ValidateConfigResponse{}
			for _, validator := range r.ConfigValidators(ctx) {
				validator.ValidateResource(ctx, frameworkresource.ValidateConfigRequest{Config: config}, resp)
			}

			if test.wantErr && !resp.Diagnostics.HasError() {
				t.Error("expected the configuration to be rejected")
			}
			if !test.wantErr && resp.Diagnostics.HasError() {
				t.Errorf("expected the configuration to be accepted: %v", resp.Diagnostics)
			}
		})
	}
}

// TestAuthMethodPathsCoverSupportedAuthMethods keeps the schema validators and
// the import allow-list from drifting apart: the validators are generated from
// authMethodPathsByVendor, while imports are checked against
// supportedAuthMethods.
func TestAuthMethodPathsCoverSupportedAuthMethods(t *testing.T) {
	applicationsByVendor := map[string]string{
		"anthropic":             "ANTHROPIC",
		"aws":                   "AWS",
		"box":                   "BOX",
		"google_cloud_platform": "GOOGLE_CLOUD_PLATFORM",
		"google_workspace":      "GOOGLE_WORKSPACE",
		"openai":                "OPENAI",
	}

	schemaMatrix := authMethodPathsByVendor()
	if len(schemaMatrix) != len(supportedAuthMethods) {
		t.Errorf("expected %d vendors in the schema matrix, got %d",
			len(supportedAuthMethods), len(schemaMatrix))
	}

	for vendor, authMethods := range schemaMatrix {
		application, ok := applicationsByVendor[vendor]
		if !ok {
			t.Errorf("unknown vendor %q in the schema matrix", vendor)
			continue
		}
		supported, ok := supportedAuthMethods[application]
		if !ok {
			t.Errorf("application %q is missing from supportedAuthMethods", application)
			continue
		}
		if len(authMethods) != len(supported) {
			t.Errorf("expected %s to expose %d auth methods, got %d",
				vendor, len(supported), len(authMethods))
		}
		for _, authMethod := range authMethods {
			if _, ok := supported[authMethod]; !ok {
				t.Errorf("auth method %q is not importable for %q", authMethod, application)
			}
		}
	}

	// The validators must reference every auth method, otherwise a
	// configuration could set two of them at once.
	if got, want := len(authMethodPaths()), 9; got != want {
		t.Errorf("expected %d auth method validator paths, got %d", want, got)
	}
	if got, want := len(vendorPaths()), len(schemaMatrix); got != want {
		t.Errorf("expected %d vendor validator paths, got %d", want, got)
	}
}

// TestResolveAuthRejectsAmbiguousModels covers the defensive checks in
// resolveAuth for models built outside the schema validators.
func TestResolveAuthRejectsAmbiguousModels(t *testing.T) {
	tests := map[string]ZeroTrustCasbIntegrationModel{
		"none": {},
		"two vendors": {
			GoogleCloudPlatform: &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
				ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{},
			},
			AWS: &ZeroTrustCasbIntegrationAWSModel{IAMRole: &ZeroTrustCasbIntegrationAWSIAMRoleModel{}},
		},
		"two auth methods": {
			Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
				AdminAPIKey:      &ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel{},
				ComplianceAPIKey: &ZeroTrustCasbIntegrationAnthropicComplianceAPIKeyModel{},
			},
		},
		"vendor without auth method": {
			Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{},
		},
	}
	for name, model := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := model.resolveAuth(); err == nil {
				t.Error("expected resolveAuth to reject the model")
			}
		})
	}
}

// TestSupportedAuthMethodsOmitsNonProductionAuthMethods guards the deliberate
// exclusions on the import path.
func TestSupportedAuthMethodsOmitsNonProductionAuthMethods(t *testing.T) {
	for application, authMethods := range supportedAuthMethods {
		for _, excluded := range []string{
			"google_workspace_workload_identity_federation",
			"github_app_installation",
			"oauth2_standard",
			"oauth2_tenant_id",
			"oauth2_domain",
			"servicenow_oauth",
			"salesforce_fedramp",
		} {
			if _, bad := authMethods[excluded]; bad {
				t.Errorf("%s must not offer %s", application, excluded)
			}
		}
	}
	for _, excluded := range []string{"GOOGLE_CHAT", "OKTA", "ONE_PASSWORD", "SLACK", "MICROSOFT_INTERNAL"} {
		if _, bad := supportedAuthMethods[excluded]; bad {
			t.Errorf("%s must not be supported", excluded)
		}
	}
}

// ── Replacement semantics ───────────────────────────────────────────────────

// TestRequiresReplaceOnAuthMethodChange pins the boundary between an in-place
// update and a replacement. Editing arguments inside a block updates in place;
// switching to a different block cannot, because the update endpoint carries no
// application or auth_method field.
func TestRequiresReplaceOnAuthMethodChange(t *testing.T) {
	ctx := context.Background()
	schema := ResourceSchema(ctx)
	modifier := requiresReplaceOnAuthMethodChange()
	gcpPath := path.Root("google_cloud_platform").AtName("google_cloud_platform_service_account")

	gcp := func(key string) ZeroTrustCasbIntegrationModel {
		model := testResourceModel(ctx, false)
		model.GoogleCloudPlatform.ServiceAccount.ServiceAccountKeyJSON = jsontypes.NewNormalizedValue(key)
		return model
	}
	// Terraform nulls write-only arguments out of state, so this is what the
	// block actually looks like once an apply has finished.
	gcpAfterApply := func() ZeroTrustCasbIntegrationModel {
		model := testResourceModel(ctx, false)
		model.GoogleCloudPlatform.ServiceAccount.ServiceAccountKeyJSON = jsontypes.NewNormalizedNull()
		return model
	}
	anthropic := func() ZeroTrustCasbIntegrationModel {
		model := testResourceModel(ctx, false)
		model.GoogleCloudPlatform = nil
		model.Anthropic = &ZeroTrustCasbIntegrationAnthropicModel{
			AdminAPIKey: &ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel{
				APIKey:   types.StringValue("sk-ant-admin01-key"),
				TenantID: types.StringNull(),
			},
		}
		return model
	}
	imported := func() ZeroTrustCasbIntegrationModel {
		model := testResourceModel(ctx, false)
		model.GoogleCloudPlatform = nil
		model.SecretsDigest = types.StringNull()
		return model
	}

	raw := func(model ZeroTrustCasbIntegrationModel) tftypes.Value {
		state := tfsdk.State{Schema: schema}
		if diags := state.Set(ctx, &model); diags.HasError() {
			t.Fatalf("failed to prepare raw value: %v", diags)
		}
		return state.Raw
	}
	objectAt := func(rawValue tftypes.Value) types.Object {
		var object types.Object
		data := tfsdk.State{Schema: schema, Raw: rawValue}
		if diags := data.GetAttribute(ctx, gcpPath, &object); diags.HasError() {
			t.Fatalf("failed to read %s: %v", gcpPath, diags)
		}
		return object
	}

	nullRaw := tftypes.NewValue(schema.Type().TerraformType(ctx), nil)

	tests := map[string]struct {
		state       tftypes.Value
		plan        tftypes.Value
		wantReplace bool
	}{
		// Switching vendor or auth method cannot be done in place: the update
		// endpoint carries no application or auth_method field.
		"auth method added": {
			state:       raw(anthropic()),
			plan:        raw(gcp(`{"type":"service_account"}`)),
			wantReplace: true,
		},
		"auth method removed": {
			state:       raw(gcpAfterApply()),
			plan:        raw(anthropic()),
			wantReplace: true,
		},
		// Editing arguments inside a block is an ordinary in-place update.
		"secret rotated in place": {
			state:       raw(gcpAfterApply()),
			plan:        raw(gcp(`{"type":"service_account","private_key":"rotated"}`)),
			wantReplace: false,
		},
		"unchanged": {
			state:       raw(gcpAfterApply()),
			plan:        raw(gcpAfterApply()),
			wantReplace: false,
		},
		"first apply after import": {
			state:       raw(imported()),
			plan:        raw(gcp(`{"type":"service_account"}`)),
			wantReplace: false,
		},
		// Creation and destruction are not replacements.
		"create": {
			state:       nullRaw,
			plan:        raw(gcp(`{"type":"service_account"}`)),
			wantReplace: false,
		},
		"destroy": {
			state:       raw(gcpAfterApply()),
			plan:        nullRaw,
			wantReplace: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := planmodifier.ObjectRequest{
				Path:  gcpPath,
				State: tfsdk.State{Schema: schema, Raw: test.state},
				Plan:  tfsdk.Plan{Schema: schema, Raw: test.plan},
			}
			if !test.state.IsNull() {
				req.StateValue = objectAt(test.state)
			} else {
				req.StateValue = types.ObjectNull(nil)
			}
			if !test.plan.IsNull() {
				req.PlanValue = objectAt(test.plan)
			} else {
				req.PlanValue = types.ObjectNull(nil)
			}

			resp := &planmodifier.ObjectResponse{PlanValue: req.PlanValue}
			modifier.PlanModifyObject(ctx, req, resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("plan modifier returned errors: %v", resp.Diagnostics)
			}
			if resp.RequiresReplace != test.wantReplace {
				t.Errorf("expected RequiresReplace=%t, got %t", test.wantReplace, resp.RequiresReplace)
			}
		})
	}
}

// ── Digest behaviour ────────────────────────────────────────────────────────

func TestSecretsDigestIsStableUnderReformatting(t *testing.T) {
	plan := ZeroTrustCasbIntegrationModel{
		GoogleCloudPlatform: &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
			ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{
				ServiceAccountKeyJSON: jsontypes.NewNormalizedValue(`{"project_id":"project-id","type":"service_account"}`),
			},
		},
	}
	state := ZeroTrustCasbIntegrationModel{
		GoogleCloudPlatform: &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
			ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{
				ServiceAccountKeyJSON: jsontypes.NewNormalizedValue(`{
					"type": "service_account",
					"project_id": "project-id"
				}`),
			},
		},
	}

	planDigest, err := plan.secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}
	stateDigest, err := state.secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}
	if planDigest != stateDigest {
		t.Error("expected semantically equivalent JSON keys to produce the same digest")
	}
	salt, ok := secretsDigestSalt(planDigest)
	if !ok {
		t.Fatalf("expected a valid salted digest, got %q", planDigest)
	}
	if !reflect.DeepEqual(salt, testSecretsSalt) {
		t.Errorf("unexpected digest salt\n got: %x\nwant: %x", salt, testSecretsSalt)
	}
}

// TestSecretsDigestCoversOnlyConfidentialArgs is the difference between this
// digest and a digest over the whole credential payload: a change to a plain
// argument must be diffed by Terraform in the normal way and must not perturb
// the digest.
func TestSecretsDigestCoversOnlyConfidentialArgs(t *testing.T) {
	model := func(tenantID, apiKey string) ZeroTrustCasbIntegrationModel {
		return ZeroTrustCasbIntegrationModel{
			Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
				WorkspaceAPIKey: &ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel{
					APIKey:   types.StringValue(apiKey),
					TenantID: types.StringValue(tenantID),
				},
			},
		}
	}

	base, err := model("wrkspc_1", "key").secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}
	publicChanged, err := model("wrkspc_2", "key").secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}
	secretChanged, err := model("wrkspc_1", "rotated").secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}

	if base != publicChanged {
		t.Error("expected a change to a non-confidential argument to leave the digest untouched")
	}
	if base == secretChanged {
		t.Error("expected a rotated secret to change the digest")
	}
}

func TestSecretsDigestUsesSalt(t *testing.T) {
	model := ZeroTrustCasbIntegrationModel{
		Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
			AdminAPIKey: &ZeroTrustCasbIntegrationAnthropicAdminAPIKeyModel{
				APIKey:   types.StringValue("same-key"),
				TenantID: types.StringValue("same-tenant"),
			},
		},
	}

	base, err := model.secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}
	otherSalt := []byte("fedcba9876543210fedcba9876543210")
	other, err := model.secretsDigest(otherSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}
	if base == other {
		t.Error("expected identical secrets with different salts to produce different digests")
	}
}

func TestSecretsDigestSaltRejectsMalformedValues(t *testing.T) {
	for _, value := range []string{
		"",
		strings.Repeat("a", 64),
		"v2:" + strings.Repeat("a", 64) + ":" + strings.Repeat("b", 64),
		"v1:abcd:" + strings.Repeat("b", 64),
		"v1:" + strings.Repeat("a", 64) + ":abcd",
	} {
		if _, ok := secretsDigestSalt(value); ok {
			t.Errorf("expected %q to be rejected", value)
		}
	}
}

// TestCredentialsChanged covers both halves of the comparison: the digest
// stands in for the write-only arguments, and the plain arguments are compared
// directly because they are present in state.
func TestCredentialsChanged(t *testing.T) {
	anthropic := func(tenantID, apiKey string) ZeroTrustCasbIntegrationModel {
		return ZeroTrustCasbIntegrationModel{
			Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
				WorkspaceAPIKey: &ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel{
					APIKey:   types.StringValue(apiKey),
					TenantID: types.StringValue(tenantID),
				},
			},
		}
	}
	// stateAfterApply is what Terraform actually persists: the write-only
	// secret is null and the digest is the only remaining record of it.
	stateAfterApply := func(tenantID string, applied ZeroTrustCasbIntegrationModel) ZeroTrustCasbIntegrationModel {
		digest, err := applied.secretsDigest(testSecretsSalt)
		if err != nil {
			t.Fatalf("secretsDigest returned an error: %v", err)
		}
		state := ZeroTrustCasbIntegrationModel{
			Anthropic: &ZeroTrustCasbIntegrationAnthropicModel{
				WorkspaceAPIKey: &ZeroTrustCasbIntegrationAnthropicWorkspaceAPIKeyModel{
					APIKey:   types.StringNull(),
					TenantID: types.StringValue(tenantID),
				},
			},
		}
		state.SecretsDigest = types.StringValue(digest)
		return state
	}

	applied := anthropic("wrkspc_1", "key")

	tests := map[string]struct {
		config ZeroTrustCasbIntegrationModel
		state  ZeroTrustCasbIntegrationModel
		want   bool
	}{
		"nothing changed": {
			config: anthropic("wrkspc_1", "key"),
			state:  stateAfterApply("wrkspc_1", applied),
			want:   false,
		},
		// The secret is absent from state, so only the digest can reveal this.
		"secret rotated": {
			config: anthropic("wrkspc_1", "rotated"),
			state:  stateAfterApply("wrkspc_1", applied),
			want:   true,
		},
		// The digest is unchanged here, so the public argument comparison is
		// the only thing that can catch it.
		"public argument changed": {
			config: anthropic("wrkspc_2", "key"),
			state:  stateAfterApply("wrkspc_1", applied),
			want:   true,
		},
		// A freshly imported resource has no digest, so the configured
		// secrets must be pushed on the first apply.
		"imported without a digest": {
			config: anthropic("wrkspc_1", "key"),
			state:  anthropic("wrkspc_1", ""),
			want:   true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			digest, err := test.config.secretsDigest(testSecretsSalt)
			if err != nil {
				t.Fatalf("secretsDigest returned an error: %v", err)
			}
			config, state := test.config, test.state
			if got := credentialsChanged(&config, &state, digest); got != test.want {
				t.Errorf("expected credentialsChanged=%t, got %t", test.want, got)
			}
		})
	}
}

// TestModifyPlanRecordsDigestWithoutReplacement is the behavioural change that
// makes credential rotation cheap: the digest is refreshed so Terraform sees a
// diff, but the integration is updated in place rather than recreated.
func TestModifyPlanRecordsDigestWithoutReplacement(t *testing.T) {
	ctx := context.Background()
	schema := ResourceSchema(ctx)
	r := &ZeroTrustCasbIntegrationResource{}

	rotated := testResourceModel(ctx, false)
	rotated.GoogleCloudPlatform.ServiceAccount.ServiceAccountKeyJSON = jsontypes.NewNormalizedValue(
		`{"type":"service_account","private_key":"rotated"}`,
	)
	rotatedDigest, err := rotated.secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}

	imported := testResourceModel(ctx, false)
	imported.SecretsDigest = types.StringNull()

	unchangedDigest := testResourceModel(ctx, false).SecretsDigest.ValueString()

	tests := map[string]struct {
		state          ZeroTrustCasbIntegrationModel
		config         ZeroTrustCasbIntegrationModel
		wantPlanDigest string
	}{
		"unchanged credentials": {
			state:          testResourceModel(ctx, false),
			config:         testResourceModel(ctx, false),
			wantPlanDigest: unchangedDigest,
		},
		"rotated credentials": {
			state:          testResourceModel(ctx, false),
			config:         rotated,
			wantPlanDigest: rotatedDigest,
		},
		"imported without a digest": {
			state:  imported,
			config: testResourceModel(ctx, false),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			state := tfsdk.State{Schema: schema}
			if diags := state.Set(ctx, &test.state); diags.HasError() {
				t.Fatalf("failed to prepare state: %v", diags)
			}
			planState := tfsdk.State{Schema: schema}
			if diags := planState.Set(ctx, &test.config); diags.HasError() {
				t.Fatalf("failed to prepare plan: %v", diags)
			}

			req := frameworkresource.ModifyPlanRequest{
				Config: tfsdk.Config{Schema: schema, Raw: planState.Raw},
				Plan:   tfsdk.Plan{Schema: schema, Raw: planState.Raw},
				State:  tfsdk.State{Schema: schema, Raw: state.Raw},
			}
			resp := &frameworkresource.ModifyPlanResponse{
				Plan: tfsdk.Plan{Schema: schema, Raw: planState.Raw},
			}
			r.ModifyPlan(ctx, req, resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("ModifyPlan returned errors: %v", resp.Diagnostics)
			}

			// Rotating a secret must never recreate the integration.
			if len(resp.RequiresReplace) > 0 {
				t.Errorf("expected no replacement, got RequiresReplace=%v", resp.RequiresReplace)
			}

			var planned ZeroTrustCasbIntegrationModel
			if diags := resp.Plan.Get(ctx, &planned); diags.HasError() {
				t.Fatalf("failed to read planned state: %v", diags)
			}
			if test.wantPlanDigest != "" {
				if planned.SecretsDigest.ValueString() != test.wantPlanDigest {
					t.Errorf("unexpected planned digest\n got: %s\nwant: %s",
						planned.SecretsDigest.ValueString(), test.wantPlanDigest)
				}
			} else if _, ok := secretsDigestSalt(planned.SecretsDigest.ValueString()); !ok {
				t.Errorf("expected a valid salted digest, got %q", planned.SecretsDigest.ValueString())
			}
		})
	}
}

// TestModifyPlanDefersUnknownSecrets covers a secret that comes from another
// resource: the digest cannot be computed at plan time, so it must be planned
// as unknown rather than as a stale value.
func TestModifyPlanDefersUnknownSecrets(t *testing.T) {
	ctx := context.Background()
	schema := ResourceSchema(ctx)
	r := &ZeroTrustCasbIntegrationResource{}

	config := testResourceModel(ctx, false)
	config.GoogleCloudPlatform.ServiceAccount.ServiceAccountKeyJSON = jsontypes.NewNormalizedUnknown()
	config.SecretsDigest = types.StringUnknown()

	planState := tfsdk.State{Schema: schema}
	if diags := planState.Set(ctx, &config); diags.HasError() {
		t.Fatalf("failed to prepare plan: %v", diags)
	}
	stateModel := testResourceModel(ctx, false)
	state := tfsdk.State{Schema: schema}
	if diags := state.Set(ctx, &stateModel); diags.HasError() {
		t.Fatalf("failed to prepare state: %v", diags)
	}

	resp := &frameworkresource.ModifyPlanResponse{
		Plan: tfsdk.Plan{Schema: schema, Raw: planState.Raw},
	}
	r.ModifyPlan(ctx, frameworkresource.ModifyPlanRequest{
		Config: tfsdk.Config{Schema: schema, Raw: planState.Raw},
		Plan:   tfsdk.Plan{Schema: schema, Raw: planState.Raw},
		State:  tfsdk.State{Schema: schema, Raw: state.Raw},
	}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ModifyPlan returned errors: %v", resp.Diagnostics)
	}

	var planned ZeroTrustCasbIntegrationModel
	if diags := resp.Plan.Get(ctx, &planned); diags.HasError() {
		t.Fatalf("failed to read planned state: %v", diags)
	}
	if !planned.SecretsDigest.IsUnknown() {
		t.Errorf("expected the digest to be planned as unknown, got %v", planned.SecretsDigest)
	}
}

// ── Collection helpers ──────────────────────────────────────────────────────

func TestEnabledUseCases(t *testing.T) {
	tests := map[string]struct {
		useCases []map[string]any
		want     []string
	}{
		"none": {
			useCases: []map[string]any{},
			want:     []string{},
		},
		"only enabled use cases are reported": {
			useCases: []map[string]any{
				{"id": "casb", "is_enabled": true},
				{"id": "ces", "is_enabled": false},
				{"id": "auto_remediation", "is_enabled": true},
			},
			want: []string{"casb", "auto_remediation"},
		},
		// Malformed entries must be skipped rather than panic or leak a
		// zero-valued identifier into state.
		"malformed entries are skipped": {
			useCases: []map[string]any{
				{"id": "casb", "is_enabled": true},
				{"id": 42, "is_enabled": true},
				{"is_enabled": true},
				{"id": "ces"},
			},
			want: []string{"casb"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := enabledUseCases(test.useCases); !reflect.DeepEqual(got, test.want) {
				t.Errorf("unexpected enabled use cases\n got: %v\nwant: %v", got, test.want)
			}
		})
	}
}

// resolveComputedDLPProfiles must never contradict a configured value: only an
// unknown, that is unset, attribute is filled from the API response.
func TestResolveComputedDLPProfiles(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		configured customfield.Set[types.String]
		fromAPI    []string
		want       customfield.Set[types.String]
	}{
		"unset is filled from the API": {
			configured: customfield.UnknownSet[types.String](ctx),
			fromAPI:    []string{"profile-a", "profile-b"},
			want:       stringsToSet(ctx, []string{"profile-a", "profile-b"}),
		},
		"configured value is preserved": {
			configured: stringsToSet(ctx, []string{"profile-a"}),
			fromAPI:    []string{"profile-b"},
			want:       stringsToSet(ctx, []string{"profile-a"}),
		},
		"explicitly empty is preserved": {
			configured: stringsToSet(ctx, []string{}),
			fromAPI:    []string{"profile-b"},
			want:       stringsToSet(ctx, []string{}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			data := ZeroTrustCasbIntegrationModel{DLPProfiles: test.configured}
			resolveComputedDLPProfiles(ctx, &data, test.fromAPI)
			if !data.DLPProfiles.Equal(test.want) {
				t.Errorf("unexpected dlp_profiles\n got: %v\nwant: %v", data.DLPProfiles, test.want)
			}
		})
	}
}

// ── Import ──────────────────────────────────────────────────────────────────

func TestValidateImportableIntegration(t *testing.T) {
	tests := map[string]struct {
		result  zero_trust.CasbIntegrationGetResponse
		wantErr bool
	}{
		"valid google cloud platform": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "GOOGLE_CLOUD_PLATFORM"},
				AuthMethod:  map[string]string{"id": authMethodGoogleCloudPlatformServiceAccount},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
		},
		"valid anthropic": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "ANTHROPIC"},
				AuthMethod:  map[string]string{"id": authMethodAnthropicWorkspaceAPIKey},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
		},
		"valid aws": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "AWS"},
				AuthMethod:  map[string]string{"id": authMethodAWSIAMRole},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
		},
		"valid box": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "BOX"},
				AuthMethod:  map[string]string{"id": authMethodBoxServerAuthentication},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
		},
		"valid openai": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "OPENAI"},
				AuthMethod:  map[string]string{"id": authMethodOpenAIComplianceAPIKey},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
		},
		"unsupported application": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "SLACK"},
				AuthMethod:  map[string]string{"id": "oauth2_standard"},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
			wantErr: true,
		},
		"auth method from another vendor": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "GOOGLE_CLOUD_PLATFORM"},
				AuthMethod:  map[string]string{"id": authMethodAWSIAMRole},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
			wantErr: true,
		},
		"oauth auth method on a supported vendor": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "BOX"},
				AuthMethod:  map[string]string{"id": "oauth2_standard"},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
			wantErr: true,
		},
		"alpha only workload identity federation": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "GOOGLE_WORKSPACE"},
				AuthMethod:  map[string]string{"id": "google_workspace_workload_identity_federation"},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": true}},
			},
			wantErr: true,
		},
		"casb disabled": {
			result: zero_trust.CasbIntegrationGetResponse{
				Application: map[string]string{"id": "GOOGLE_CLOUD_PLATFORM"},
				AuthMethod:  map[string]string{"id": authMethodGoogleCloudPlatformServiceAccount},
				UseCases:    []map[string]any{{"id": "casb", "is_enabled": false}},
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateImportableIntegration(&test.result)
			if test.wantErr && err == nil {
				t.Fatal("expected validation to fail")
			}
			if !test.wantErr && err != nil {
				t.Fatalf("expected validation to succeed: %v", err)
			}
		})
	}
}

func TestImportStateImportsIntegration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			t.Errorf("unexpected request method: %s", req.Method)
		}
		if req.URL.Path != "/accounts/account-id/one/integrations/integration-id" {
			t.Errorf("unexpected request path: %s", req.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"result": map[string]any{
				"id":           "integration-id",
				"name":         "gcp integration",
				"is_paused":    true,
				"application":  map[string]string{"id": "GOOGLE_CLOUD_PLATFORM"},
				"auth_method":  map[string]string{"id": authMethodGoogleCloudPlatformServiceAccount},
				"dlp_profiles": []string{"profile-a"},
				"use_cases": []map[string]any{
					{"id": "casb", "is_enabled": true},
					{"id": "ces", "is_enabled": false},
				},
			},
		}); err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	ctx := context.Background()
	r := testResource(server.URL)
	resp := frameworkresource.ImportStateResponse{
		State: tfsdk.State{Schema: ResourceSchema(ctx)},
	}
	r.ImportState(ctx, frameworkresource.ImportStateRequest{ID: "account-id/integration-id"}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("ImportState returned errors: %v", resp.Diagnostics)
	}

	var got ZeroTrustCasbIntegrationModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read imported state: %v", diags)
	}
	if got.AccountID.ValueString() != "account-id" || got.ID.ValueString() != "integration-id" {
		t.Errorf("unexpected imported identifiers: account_id=%q id=%q", got.AccountID.ValueString(), got.ID.ValueString())
	}
	if got.Name.ValueString() != "gcp integration" || !got.Paused.ValueBool() {
		t.Errorf("unexpected imported integration: name=%q paused=%t", got.Name.ValueString(), got.Paused.ValueBool())
	}
	if got.GoogleCloudPlatform != nil {
		t.Error("expected imported credentials to be null because the API does not return them")
	}
	// A null digest is what makes the first apply after an import push the
	// configured secrets.
	if !got.SecretsDigest.IsNull() {
		t.Errorf("expected the imported digest to be null, got %v", got.SecretsDigest)
	}
	if want := stringsToSet(ctx, []string{"profile-a"}); !got.DLPProfiles.Equal(want) {
		t.Errorf("unexpected imported dlp_profiles\n got: %v\nwant: %v", got.DLPProfiles, want)
	}
	// Only the enabled use cases are imported.
	if want := stringsToSet(ctx, []string{"casb"}); !got.UseCases.Equal(want) {
		t.Errorf("unexpected imported use_cases\n got: %v\nwant: %v", got.UseCases, want)
	}
	// Permissions are never returned, so they cannot be imported.
	if !got.Permissions.IsNull() {
		t.Errorf("expected imported permissions to be null, got %v", got.Permissions)
	}
}

// ── CRUD ────────────────────────────────────────────────────────────────────

func TestDeleteIgnoresNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			t.Errorf("unexpected request method: %s", req.Method)
		}
		if req.URL.Path != "/accounts/account-id/one/integrations/integration-id" {
			t.Errorf("unexpected request path: %s", req.URL.Path)
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	ctx := context.Background()
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	model := testResourceModel(ctx, false)
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("failed to prepare state: %v", diags)
	}

	r := testResource(server.URL)
	resp := frameworkresource.DeleteResponse{State: state}
	r.Delete(ctx, frameworkresource.DeleteRequest{State: state}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Delete returned errors for a missing integration: %v", resp.Diagnostics)
	}
}

// TestCreatePausesIntegration also asserts the core guarantee: after a create,
// no confidential argument appears anywhere in the persisted state.
func TestCreatePausesIntegration(t *testing.T) {
	var requestCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestNumber := requestCount.Add(1)
		w.Header().Set("Content-Type", "application/json")

		switch requestNumber {
		case 1:
			if req.Method != http.MethodPost || req.URL.Path != "/accounts/account-id/one/integrations" {
				t.Errorf("unexpected create request: %s %s", req.Method, req.URL.Path)
			}
			var body map[string]any
			if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
				t.Errorf("failed to decode create request: %v", err)
			}
			if body["application"] != "GOOGLE_CLOUD_PLATFORM" || body["auth_method"] != authMethodGoogleCloudPlatformServiceAccount {
				t.Errorf("unexpected create request body: %#v", body)
			}
			if err := json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"result": map[string]any{
					"id":           "integration-id",
					"name":         "gcp integration",
					"is_paused":    false,
					"dlp_profiles": []string{"assigned-by-the-api"},
					"use_cases":    []map[string]any{{"id": "casb", "is_enabled": true}},
				},
			}); err != nil {
				t.Errorf("failed to write create response: %v", err)
			}
		case 2:
			if req.Method != http.MethodPost || req.URL.Path != "/accounts/account-id/one/integrations/integration-id/pause" {
				t.Errorf("unexpected pause request: %s %s", req.Method, req.URL.Path)
			}
			if err := json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"result": map[string]any{
					"id":        "integration-id",
					"name":      "gcp integration",
					"is_paused": true,
				},
			}); err != nil {
				t.Errorf("failed to write pause response: %v", err)
			}
		default:
			t.Errorf("unexpected extra request: %s %s", req.Method, req.URL.Path)
		}
	}))
	defer server.Close()

	ctx := context.Background()
	schema := ResourceSchema(ctx)
	model := testResourceModel(ctx, true)
	model.ID = types.StringUnknown()
	// dlp_profiles is left unset, so the API assigns it.
	model.DLPProfiles = customfield.UnknownSet[types.String](ctx)

	// Configuration carries the plaintext credentials.
	configState := tfsdk.State{Schema: schema}
	if diags := configState.Set(ctx, &model); diags.HasError() {
		t.Fatalf("failed to prepare config: %v", diags)
	}
	config := tfsdk.Config{Schema: schema, Raw: configState.Raw}

	// The plan has write-only arguments nulled, exactly as Terraform delivers
	// them once the framework has stripped them.
	planModel := model
	planModel.GoogleCloudPlatform = &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
		ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{
			ServiceAccountKeyJSON: jsontypes.NewNormalizedNull(),
		},
	}
	plan := tfsdk.Plan{Schema: schema}
	if diags := plan.Set(ctx, &planModel); diags.HasError() {
		t.Fatalf("failed to prepare plan: %v", diags)
	}

	r := testResource(server.URL)
	resp := frameworkresource.CreateResponse{State: tfsdk.State{Schema: schema}}
	r.Create(ctx, frameworkresource.CreateRequest{Plan: plan, Config: config}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Create returned errors: %v", resp.Diagnostics)
	}
	if requestCount.Load() != 2 {
		t.Fatalf("expected create and pause requests, got %d", requestCount.Load())
	}

	var got ZeroTrustCasbIntegrationModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read created state: %v", diags)
	}
	if got.ID.ValueString() != "integration-id" || !got.Paused.ValueBool() {
		t.Errorf("unexpected created state: id=%q paused=%t", got.ID.ValueString(), got.Paused.ValueBool())
	}
	// The core guarantee: no confidential material in state, only the digest.
	if !got.GoogleCloudPlatform.ServiceAccount.ServiceAccountKeyJSON.IsNull() {
		t.Error("expected the service account key to be absent from state")
	}
	if got.SecretsDigest.ValueString() != model.SecretsDigest.ValueString() {
		t.Errorf("unexpected stored digest\n got: %s\nwant: %s",
			got.SecretsDigest.ValueString(), model.SecretsDigest.ValueString())
	}
	if want := stringsToSet(ctx, []string{"assigned-by-the-api"}); !got.DLPProfiles.Equal(want) {
		t.Errorf("unexpected dlp_profiles\n got: %v\nwant: %v", got.DLPProfiles, want)
	}
}

func TestUpdatePauseOnly(t *testing.T) {
	var requestCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestCount.Add(1)
		if req.Method != http.MethodPost || req.URL.Path != "/accounts/account-id/one/integrations/integration-id/pause" {
			t.Errorf("unexpected update request: %s %s", req.Method, req.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"result": map[string]any{
				"id":        "integration-id",
				"name":      "gcp integration",
				"is_paused": true,
			},
		}); err != nil {
			t.Errorf("failed to write pause response: %v", err)
		}
	}))
	defer server.Close()

	ctx := context.Background()
	stateModel := testResourceModel(ctx, false)
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	if diags := state.Set(ctx, &stateModel); diags.HasError() {
		t.Fatalf("failed to prepare state: %v", diags)
	}
	planModel := testResourceModel(ctx, true)
	plan := tfsdk.Plan{Schema: ResourceSchema(ctx)}
	if diags := plan.Set(ctx, &planModel); diags.HasError() {
		t.Fatalf("failed to prepare plan: %v", diags)
	}

	r := testResource(server.URL)
	resp := frameworkresource.UpdateResponse{State: state}
	config := tfsdk.Config{Schema: ResourceSchema(ctx), Raw: plan.Raw}
	r.Update(ctx, frameworkresource.UpdateRequest{Plan: plan, State: state, Config: config}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Update returned errors: %v", resp.Diagnostics)
	}
	if requestCount.Load() != 1 {
		t.Fatalf("expected only a pause request, got %d requests", requestCount.Load())
	}

	var got ZeroTrustCasbIntegrationModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read updated state: %v", diags)
	}
	if !got.Paused.ValueBool() {
		t.Error("expected updated state to be paused")
	}
}

// TestUpdateRotatesCredentialsInPlace is the behaviour the digest exists to
// enable: a rotated secret is pushed through the update endpoint rather than
// forcing the integration to be recreated.
func TestUpdateRotatesCredentialsInPlace(t *testing.T) {
	var requestCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestNumber := requestCount.Add(1)
		w.Header().Set("Content-Type", "application/json")

		switch requestNumber {
		case 1:
			if req.Method != http.MethodPatch || req.URL.Path != "/accounts/account-id/one/integrations/integration-id" {
				t.Errorf("unexpected update request: %s %s", req.Method, req.URL.Path)
			}
			var body map[string]any
			if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
				t.Errorf("failed to decode update request: %v", err)
			}
			credentials, ok := body["credentials"].(map[string]any)
			if !ok {
				t.Fatalf("unexpected credentials: %#v", body["credentials"])
			}
			serviceAccount, ok := credentials["service_account_credentials"].(map[string]any)
			if !ok || serviceAccount["private_key"] != "new-private-key" {
				t.Errorf("unexpected service account credentials: %#v", credentials)
			}
			if body["name"] != "renamed integration" {
				t.Errorf("unexpected integration name: %#v", body["name"])
			}
			if err := json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"result": map[string]any{
					"id":        "integration-id",
					"name":      "renamed integration",
					"is_paused": true,
				},
			}); err != nil {
				t.Errorf("failed to write update response: %v", err)
			}
		case 2:
			if req.Method != http.MethodPost || req.URL.Path != "/accounts/account-id/one/integrations/integration-id/resume" {
				t.Errorf("unexpected resume request: %s %s", req.Method, req.URL.Path)
			}
			if err := json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"result": map[string]any{
					"id":        "integration-id",
					"name":      "renamed integration",
					"is_paused": false,
				},
			}); err != nil {
				t.Errorf("failed to write resume response: %v", err)
			}
		default:
			t.Errorf("unexpected extra request: %s %s", req.Method, req.URL.Path)
		}
	}))
	defer server.Close()

	ctx := context.Background()
	stateModel := testResourceModel(ctx, true)
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	if diags := state.Set(ctx, &stateModel); diags.HasError() {
		t.Fatalf("failed to prepare state: %v", diags)
	}
	planModel := testResourceModel(ctx, false)
	planModel.Name = types.StringValue("renamed integration")
	planModel.GoogleCloudPlatform.ServiceAccount.ServiceAccountKeyJSON = jsontypes.NewNormalizedValue(
		`{"type":"service_account","private_key":"new-private-key"}`,
	)
	plan := tfsdk.Plan{Schema: ResourceSchema(ctx)}
	if diags := plan.Set(ctx, &planModel); diags.HasError() {
		t.Fatalf("failed to prepare plan: %v", diags)
	}

	r := testResource(server.URL)
	resp := frameworkresource.UpdateResponse{State: state}
	config := tfsdk.Config{Schema: ResourceSchema(ctx), Raw: plan.Raw}
	r.Update(ctx, frameworkresource.UpdateRequest{Plan: plan, State: state, Config: config}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Update returned errors: %v", resp.Diagnostics)
	}
	if requestCount.Load() != 2 {
		t.Fatalf("expected update and resume requests, got %d", requestCount.Load())
	}

	var got ZeroTrustCasbIntegrationModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read updated state: %v", diags)
	}
	if got.Name.ValueString() != "renamed integration" || got.Paused.ValueBool() {
		t.Errorf("unexpected updated state: name=%q paused=%t", got.Name.ValueString(), got.Paused.ValueBool())
	}
	wantDigest, err := planModel.secretsDigest(testSecretsSalt)
	if err != nil {
		t.Fatalf("secretsDigest returned an error: %v", err)
	}
	if got.SecretsDigest.ValueString() != wantDigest {
		t.Error("expected updated state to record the digest of the replacement key")
	}
}

// TestUpdateSendsChangedCollections asserts that changing only a collection
// still produces an update request carrying that collection.
func TestUpdateSendsChangedCollections(t *testing.T) {
	var body map[string]any
	var requestCount atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestCount.Add(1)
		if req.Method != http.MethodPatch || req.URL.Path != "/accounts/account-id/one/integrations/integration-id" {
			t.Errorf("unexpected update request: %s %s", req.Method, req.URL.Path)
		}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Errorf("failed to decode update request: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"result": map[string]any{
				"id":           "integration-id",
				"name":         "gcp integration",
				"is_paused":    false,
				"dlp_profiles": []string{"profile-b"},
				"use_cases": []map[string]any{
					{"id": "casb", "is_enabled": true},
					{"id": "ces", "is_enabled": true},
				},
			},
		}); err != nil {
			t.Errorf("failed to write update response: %v", err)
		}
	}))
	defer server.Close()

	ctx := context.Background()
	stateModel := testResourceModel(ctx, false)
	stateModel.DLPProfiles = stringsToSet(ctx, []string{"profile-a"})
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	if diags := state.Set(ctx, &stateModel); diags.HasError() {
		t.Fatalf("failed to prepare state: %v", diags)
	}

	planModel := testResourceModel(ctx, false)
	planModel.DLPProfiles = stringsToSet(ctx, []string{"profile-b"})
	planModel.Permissions = stringsToSet(ctx, []string{"read"})
	planModel.UseCases = stringsToSet(ctx, []string{"casb", "ces"})
	plan := tfsdk.Plan{Schema: ResourceSchema(ctx)}
	if diags := plan.Set(ctx, &planModel); diags.HasError() {
		t.Fatalf("failed to prepare plan: %v", diags)
	}

	r := testResource(server.URL)
	resp := frameworkresource.UpdateResponse{State: state}
	config := tfsdk.Config{Schema: ResourceSchema(ctx), Raw: plan.Raw}
	r.Update(ctx, frameworkresource.UpdateRequest{Plan: plan, State: state, Config: config}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Update returned errors: %v", resp.Diagnostics)
	}
	if requestCount.Load() != 1 {
		t.Fatalf("expected a single update request, got %d", requestCount.Load())
	}

	if got, want := body["dlp_profiles"], []any{"profile-b"}; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected dlp_profiles in the request\n got: %#v\nwant: %#v", got, want)
	}
	if got, want := body["permissions"], []any{"read"}; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected permissions in the request\n got: %#v\nwant: %#v", got, want)
	}
	if got, want := body["use_cases"], []any{"casb", "ces"}; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected use_cases in the request\n got: %#v\nwant: %#v", got, want)
	}
	// The credentials are unchanged, so they must not be resent.
	if _, resent := body["credentials"]; resent {
		t.Error("expected unchanged credentials to be omitted from the update request")
	}

	var got ZeroTrustCasbIntegrationModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read updated state: %v", diags)
	}
	if want := stringsToSet(ctx, []string{"profile-b"}); !got.DLPProfiles.Equal(want) {
		t.Errorf("unexpected dlp_profiles\n got: %v\nwant: %v", got.DLPProfiles, want)
	}
}

func TestReadReconcilesCollections(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/one/integrations/integration-id" {
			t.Errorf("unexpected read request: %s %s", req.Method, req.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"result": map[string]any{
				"id":           "integration-id",
				"name":         "renamed out of band",
				"is_paused":    true,
				"dlp_profiles": []string{"profile-b", "profile-c"},
				"use_cases": []map[string]any{
					{"id": "casb", "is_enabled": true},
					{"id": "ces", "is_enabled": false},
				},
			},
		}); err != nil {
			t.Errorf("failed to write read response: %v", err)
		}
	}))
	defer server.Close()

	ctx := context.Background()
	model := testResourceModel(ctx, false)
	model.DLPProfiles = stringsToSet(ctx, []string{"profile-a"})
	model.Permissions = stringsToSet(ctx, []string{"read"})
	model.UseCases = stringsToSet(ctx, []string{"casb", "ces"})
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("failed to prepare state: %v", diags)
	}

	r := testResource(server.URL)
	resp := frameworkresource.ReadResponse{State: state}
	r.Read(ctx, frameworkresource.ReadRequest{State: state}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read returned errors: %v", resp.Diagnostics)
	}

	var got ZeroTrustCasbIntegrationModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read refreshed state: %v", diags)
	}
	if got.Name.ValueString() != "renamed out of band" || !got.Paused.ValueBool() {
		t.Errorf("unexpected refreshed state: name=%q paused=%t", got.Name.ValueString(), got.Paused.ValueBool())
	}
	if want := stringsToSet(ctx, []string{"profile-b", "profile-c"}); !got.DLPProfiles.Equal(want) {
		t.Errorf("unexpected refreshed dlp_profiles\n got: %v\nwant: %v", got.DLPProfiles, want)
	}
	// A use case that has been disabled out of band must drop out of state.
	if want := stringsToSet(ctx, []string{"casb"}); !got.UseCases.Equal(want) {
		t.Errorf("unexpected refreshed use_cases\n got: %v\nwant: %v", got.UseCases, want)
	}
	// Permissions are not returned, so refresh must leave them untouched
	// rather than blank them out and provoke a spurious diff.
	if want := stringsToSet(ctx, []string{"read"}); !got.Permissions.Equal(want) {
		t.Errorf("unexpected refreshed permissions\n got: %v\nwant: %v", got.Permissions, want)
	}
	// Refresh cannot learn the secrets, so the digest must survive it.
	if got.SecretsDigest.ValueString() != model.SecretsDigest.ValueString() {
		t.Error("expected refresh to preserve the secrets digest")
	}
}

func TestReadRemovesMissingIntegration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/one/integrations/integration-id" {
			t.Errorf("unexpected read request: %s %s", req.Method, req.URL.Path)
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	ctx := context.Background()
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	model := testResourceModel(ctx, false)
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("failed to prepare state: %v", diags)
	}

	r := testResource(server.URL)
	resp := frameworkresource.ReadResponse{State: state}
	r.Read(ctx, frameworkresource.ReadRequest{State: state}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Read returned errors for a missing integration: %v", resp.Diagnostics)
	}
	if !resp.State.Raw.IsNull() {
		t.Error("expected missing integration to be removed from state")
	}
}

// ── Helpers ─────────────────────────────────────────────────────────────────

func TestSetToStringsRoundTrip(t *testing.T) {
	ctx := context.Background()

	// A null or unknown set carries no value, so it must convert to nil and
	// therefore be omitted from the request rather than sent as an empty list.
	for name, set := range map[string]customfield.Set[types.String]{
		"null":    customfield.NullSet[types.String](ctx),
		"unknown": customfield.UnknownSet[types.String](ctx),
	} {
		t.Run(name, func(t *testing.T) {
			got, diags := setToStrings(ctx, set)
			if diags.HasError() {
				t.Fatalf("setToStrings returned errors: %v", diags)
			}
			if got != nil {
				t.Errorf("expected nil, got %#v", got)
			}
		})
	}

	t.Run("values", func(t *testing.T) {
		got, diags := setToStrings(ctx, stringsToSet(ctx, []string{"a", "b"}))
		if diags.HasError() {
			t.Fatalf("setToStrings returned errors: %v", diags)
		}
		sort.Strings(got)
		if want := []string{"a", "b"}; !reflect.DeepEqual(got, want) {
			t.Errorf("unexpected values\n got: %v\nwant: %v", got, want)
		}
	})

	// An explicitly empty set is a real value and must survive as one.
	t.Run("empty", func(t *testing.T) {
		got, diags := setToStrings(ctx, stringsToSet(ctx, []string{}))
		if diags.HasError() {
			t.Fatalf("setToStrings returned errors: %v", diags)
		}
		if got == nil || len(got) != 0 {
			t.Errorf("expected an empty non-nil slice, got %#v", got)
		}
	})
}

func testResourceModel(ctx context.Context, paused bool) ZeroTrustCasbIntegrationModel {
	model := ZeroTrustCasbIntegrationModel{
		ID:          types.StringValue("integration-id"),
		AccountID:   types.StringValue("account-id"),
		Name:        types.StringValue("gcp integration"),
		Paused:      types.BoolValue(paused),
		DLPProfiles: customfield.NullSet[types.String](ctx),
		Permissions: customfield.NullSet[types.String](ctx),
		UseCases:    stringsToSet(ctx, []string{"casb"}),
		GoogleCloudPlatform: &ZeroTrustCasbIntegrationGoogleCloudPlatformModel{
			ServiceAccount: &ZeroTrustCasbIntegrationGoogleCloudPlatformServiceAccountModel{
				ServiceAccountKeyJSON: jsontypes.NewNormalizedValue(`{"type":"service_account"}`),
			},
		},
	}
	digest, err := model.secretsDigest(testSecretsSalt)
	if err != nil {
		panic(err)
	}
	model.SecretsDigest = types.StringValue(digest)
	return model
}

func testResource(baseURL string) ZeroTrustCasbIntegrationResource {
	return ZeroTrustCasbIntegrationResource{
		client: cloudflare.NewClient(
			option.WithBaseURL(baseURL),
			option.WithAPIToken("test-token"),
		),
	}
}
