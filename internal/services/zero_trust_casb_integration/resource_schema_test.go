package zero_trust_casb_integration_test

import (
	"context"
	"sort"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_casb_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// argSpec pins how a single credential argument is exposed. The writeOnly flag
// is the whole point of the table: it records, per argument, whether the value
// is confidential (and therefore must never reach state) or is ordinary
// configuration that Terraform can diff directly.
type argSpec struct {
	writeOnly bool
	required  bool
	jsonType  bool
}

// wantVendors is the expected vendor/auth method/argument matrix.
//
// It mirrors the vendor configuration that is live in production, not
// alpha-gated, and createable without an interactive OAuth redirect. Any
// divergence between this table and the schema is a bug in one of the two.
func wantVendors() map[string]map[string]map[string]argSpec {
	return map[string]map[string]map[string]argSpec{
		"anthropic": {
			"anthropic_admin_api_key": {
				"api_key":   {writeOnly: true, required: true},
				"tenant_id": {},
			},
			"anthropic_workspace_api_key": {
				"api_key":   {writeOnly: true, required: true},
				"tenant_id": {required: true},
			},
			"anthropic_compliance_api_key": {
				"compliance_api_key": {writeOnly: true, required: true},
				"tenant_id":          {},
			},
		},
		"aws": {
			// Nothing here is confidential: the external ID is a
			// customer-chosen value that is useless without Cloudflare's
			// AWS principal.
			"aws_iam_role": {
				"role_arn":    {required: true},
				"external_id": {required: true},
			},
		},
		"box": {
			// Box server authentication uses a Cloudflare-side application,
			// so the caller supplies no secret at all.
			"box_server_authentication": {
				"enterprise_id": {required: true},
			},
		},
		"google_cloud_platform": {
			"google_cloud_platform_service_account": {
				"service_account_key_json": {writeOnly: true, required: true, jsonType: true},
			},
		},
		"google_workspace": {
			"google_domain_wide_delegation_service_account": {
				"service_account_key_json": {writeOnly: true, required: true, jsonType: true},
				"administrator_email":      {required: true},
			},
		},
		"openai": {
			"chatgpt_standard_api_key": {
				"admin_api_key":   {writeOnly: true, required: true},
				"organization_id": {required: true},
				"project_api_key": {writeOnly: true},
				"project_id":      {},
			},
			"chatgpt_compliance_api_key": {
				"admin_api_key":      {writeOnly: true, required: true},
				"organization_id":    {required: true},
				"compliance_api_key": {writeOnly: true, required: true},
				"workspace_id":       {required: true},
				"project_api_key":    {writeOnly: true},
				"project_id":         {},
			},
		},
	}
}

func TestZeroTrustCasbIntegrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_casb_integration.ZeroTrustCasbIntegrationModel)(nil)
	schema := zero_trust_casb_integration.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

func TestZeroTrustCasbIntegrationResourceSchemaTopLevel(t *testing.T) {
	t.Parallel()
	schema := zero_trust_casb_integration.ResourceSchema(context.TODO())

	want := []string{
		"account_id",
		"dlp_profiles",
		"id",
		"name",
		"paused",
		"permissions",
		"secrets_digest",
		"use_cases",
	}
	for vendor := range wantVendors() {
		want = append(want, vendor)
	}
	sort.Strings(want)

	got := make([]string, 0, len(schema.Attributes))
	for name := range schema.Attributes {
		got = append(got, name)
	}
	sort.Strings(got)

	if len(got) != len(want) {
		t.Errorf("unexpected resource attributes\n got: %v\nwant: %v", got, want)
	}
	for _, name := range want {
		if _, ok := schema.Attributes[name]; !ok {
			t.Errorf("expected resource attribute %q", name)
		}
	}
}

// TestZeroTrustCasbIntegrationSecretsDigestIsProviderManaged guards the one
// attribute that stands in for every write-only argument. If a practitioner
// could set it, they could suppress a credential rotation.
func TestZeroTrustCasbIntegrationSecretsDigestIsProviderManaged(t *testing.T) {
	t.Parallel()
	schema := zero_trust_casb_integration.ResourceSchema(context.TODO())

	digest, ok := schema.Attributes["secrets_digest"].(resourceschema.StringAttribute)
	if !ok {
		t.Fatal("expected secrets_digest to be a string attribute")
	}
	if !digest.Computed {
		t.Error("expected secrets_digest to be computed")
	}
	if digest.Required || digest.Optional {
		t.Error("expected secrets_digest to be provider-managed, not practitioner-settable")
	}
	if digest.WriteOnly {
		t.Error("expected secrets_digest to be persisted: it is the only record of the secrets")
	}
}

// TestZeroTrustCasbIntegrationVendorSchema walks the whole vendor matrix and
// asserts, argument by argument, that confidential values are write-only and
// sensitive while everything else is plain configuration.
func TestZeroTrustCasbIntegrationVendorSchema(t *testing.T) {
	t.Parallel()
	schema := zero_trust_casb_integration.ResourceSchema(context.TODO())

	for vendor, authMethods := range wantVendors() {
		vendorAttr, ok := schema.Attributes[vendor].(resourceschema.SingleNestedAttribute)
		if !ok {
			t.Errorf("expected %q to be a single nested attribute", vendor)
			continue
		}
		// Every vendor block is optional so that exactly one can be chosen;
		// the choice is enforced by ConfigValidators rather than by Required.
		if !vendorAttr.Optional {
			t.Errorf("expected %q to be optional", vendor)
		}
		if len(vendorAttr.Attributes) != len(authMethods) {
			t.Errorf("expected %q to expose %d auth methods, got %d",
				vendor, len(authMethods), len(vendorAttr.Attributes))
		}

		for authMethod, args := range authMethods {
			methodAttr, ok := vendorAttr.Attributes[authMethod].(resourceschema.SingleNestedAttribute)
			if !ok {
				t.Errorf("expected %s.%s to be a single nested attribute", vendor, authMethod)
				continue
			}
			if !methodAttr.Optional {
				t.Errorf("expected %s.%s to be optional", vendor, authMethod)
			}
			// Neither the vendor nor the auth method can change in place:
			// the update endpoint carries no application or auth_method field.
			if len(methodAttr.PlanModifiers) == 0 {
				t.Errorf("expected %s.%s to force replacement when added or removed", vendor, authMethod)
			}
			if len(methodAttr.Attributes) != len(args) {
				t.Errorf("expected %s.%s to expose %d arguments, got %d",
					vendor, authMethod, len(args), len(methodAttr.Attributes))
			}

			for argName, want := range args {
				arg, ok := methodAttr.Attributes[argName].(resourceschema.StringAttribute)
				if !ok {
					t.Errorf("expected %s.%s.%s to be a string attribute", vendor, authMethod, argName)
					continue
				}
				if arg.WriteOnly != want.writeOnly {
					t.Errorf("expected %s.%s.%s WriteOnly=%t, got %t",
						vendor, authMethod, argName, want.writeOnly, arg.WriteOnly)
				}
				// Confidential and only confidential arguments are sensitive:
				// marking a plain argument sensitive would hide a legitimate
				// diff from the plan output.
				if arg.Sensitive != want.writeOnly {
					t.Errorf("expected %s.%s.%s Sensitive=%t, got %t",
						vendor, authMethod, argName, want.writeOnly, arg.Sensitive)
				}
				// A write-only attribute may never be computed, since the
				// framework has no value to compute it from.
				if arg.WriteOnly && arg.Computed {
					t.Errorf("expected %s.%s.%s not to be computed: write-only attributes cannot be",
						vendor, authMethod, argName)
				}
				if arg.Required != want.required {
					t.Errorf("expected %s.%s.%s Required=%t, got %t",
						vendor, authMethod, argName, want.required, arg.Required)
				}
				if arg.Optional != !want.required {
					t.Errorf("expected %s.%s.%s Optional=%t, got %t",
						vendor, authMethod, argName, !want.required, arg.Optional)
				}

				_, isJSON := arg.CustomType.(jsontypes.NormalizedType)
				if isJSON != want.jsonType {
					t.Errorf("expected %s.%s.%s normalized JSON type=%t, got %t",
						vendor, authMethod, argName, want.jsonType, isJSON)
				}
			}
		}
	}
}

// TestZeroTrustCasbIntegrationOmitsNonProductionVendors guards the deliberate
// exclusions so that a future schema edit cannot quietly reintroduce them.
func TestZeroTrustCasbIntegrationOmitsNonProductionVendors(t *testing.T) {
	t.Parallel()
	schema := zero_trust_casb_integration.ResourceSchema(context.TODO())

	// OAuth-only vendors cannot be configured without an interactive browser
	// redirect, which Terraform cannot perform. Google Chat, Okta and
	// 1Password are not live and have no value in the public application enum.
	for _, vendor := range []string{
		"bitbucket", "confluence", "dropbox", "github", "gitlab", "google_chat",
		"jira", "microsoft", "okta", "one_password", "salesforce", "servicenow",
		"slack", "zoom",
	} {
		if _, ok := schema.Attributes[vendor]; ok {
			t.Errorf("%s must not be exposed", vendor)
		}
	}

	workspace, ok := schema.Attributes["google_workspace"].(resourceschema.SingleNestedAttribute)
	if !ok {
		t.Fatal("expected google_workspace to be a single nested attribute")
	}
	if _, bad := workspace.Attributes["google_workspace_workload_identity_federation"]; bad {
		t.Error("workload identity federation must not be exposed: it is alpha-only")
	}
}

func TestZeroTrustCasbIntegrationDLPProfilesSchema(t *testing.T) {
	t.Parallel()
	schema := zero_trust_casb_integration.ResourceSchema(context.TODO())

	profiles, ok := schema.Attributes["dlp_profiles"].(resourceschema.SetAttribute)
	if !ok {
		t.Fatal("expected dlp_profiles to be a set attribute")
	}
	if profiles.ElementType != types.StringType {
		t.Errorf("expected dlp_profiles elements to be strings, got %v", profiles.ElementType)
	}
	// The API assigns default profiles when none are given, so the attribute
	// has to be computed as well as optional.
	if !profiles.Optional || !profiles.Computed {
		t.Errorf("expected dlp_profiles to be optional and computed, got optional=%t computed=%t",
			profiles.Optional, profiles.Computed)
	}
	if len(profiles.PlanModifiers) == 0 {
		t.Error("expected dlp_profiles to keep its prior value rather than plan as unknown")
	}
}

// The API never returns permissions, so the provider cannot detect drift in
// them and must not claim to. Marking it Computed would let the provider
// invent a value it has no way to read back.
func TestZeroTrustCasbIntegrationPermissionsSchema(t *testing.T) {
	t.Parallel()
	schema := zero_trust_casb_integration.ResourceSchema(context.TODO())

	permissions, ok := schema.Attributes["permissions"].(resourceschema.SetAttribute)
	if !ok {
		t.Fatal("expected permissions to be a set attribute")
	}
	if permissions.ElementType != types.StringType {
		t.Errorf("expected permissions elements to be strings, got %v", permissions.ElementType)
	}
	if !permissions.Optional {
		t.Error("expected permissions to be optional")
	}
	if permissions.Computed {
		t.Error("expected permissions not to be computed: the API never returns it")
	}
}

func TestZeroTrustCasbIntegrationUseCasesSchema(t *testing.T) {
	t.Parallel()
	ctx := context.TODO()
	schema := zero_trust_casb_integration.ResourceSchema(ctx)

	useCases, ok := schema.Attributes["use_cases"].(resourceschema.SetAttribute)
	if !ok {
		t.Fatal("expected use_cases to be a set attribute")
	}
	if useCases.ElementType != types.StringType {
		t.Errorf("expected use_cases elements to be strings, got %v", useCases.ElementType)
	}
	if !useCases.Optional || !useCases.Computed {
		t.Errorf("expected use_cases to be optional and computed, got optional=%t computed=%t",
			useCases.Optional, useCases.Computed)
	}

	if useCases.Default == nil {
		t.Fatal("expected use_cases to default to the casb use case")
	}
	defaultResp := defaults.SetResponse{}
	useCases.Default.DefaultSet(ctx, defaults.SetRequest{Path: path.Root("use_cases")}, &defaultResp)
	wantDefault := types.SetValueMust(types.StringType, []attr.Value{types.StringValue("casb")})
	if !defaultResp.PlanValue.Equal(wantDefault) {
		t.Errorf("unexpected use_cases default\n got: %v\nwant: %v", defaultResp.PlanValue, wantDefault)
	}

	tests := map[string]struct {
		values  []string
		wantErr bool
	}{
		"casb":          {values: []string{"casb"}},
		"all use cases": {values: []string{"casb", "ces", "auto_remediation"}},
		"unknown use case": {
			values:  []string{"telepathy"},
			wantErr: true,
		},
		"one unknown among known": {
			values:  []string{"casb", "telepathy"},
			wantErr: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			elements := make([]attr.Value, 0, len(test.values))
			for _, value := range test.values {
				elements = append(elements, types.StringValue(value))
			}
			resp := validator.SetResponse{}
			req := validator.SetRequest{
				Path:        path.Root("use_cases"),
				ConfigValue: types.SetValueMust(types.StringType, elements),
			}
			for _, v := range useCases.Validators {
				v.ValidateSet(ctx, req, &resp)
			}
			if test.wantErr && !resp.Diagnostics.HasError() {
				t.Errorf("expected %v to be rejected", test.values)
			}
			if !test.wantErr && resp.Diagnostics.HasError() {
				t.Errorf("expected %v to be accepted: %v", test.values, resp.Diagnostics)
			}
		})
	}
}
