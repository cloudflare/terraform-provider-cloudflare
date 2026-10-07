package zero_trust_organization_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	apiFieldName   = "mfa_piv_key_requirements"
	staleFieldName = "mfa_ssh_piv_key_requirements"
)

func TestPivKeyRequirementsSchemaName(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	t.Run("resource", func(t *testing.T) {
		schema := zero_trust_organization.ResourceSchema(ctx)
		if _, ok := schema.Attributes[apiFieldName]; !ok {
			t.Fatalf("resource schema is missing %q", apiFieldName)
		}
		if _, ok := schema.Attributes[staleFieldName]; ok {
			t.Fatalf("resource schema still accepts %q", staleFieldName)
		}
	})
	t.Run("data source", func(t *testing.T) {
		schema := zero_trust_organization.DataSourceSchema(ctx)
		if _, ok := schema.Attributes[apiFieldName]; !ok {
			t.Fatalf("data source schema is missing %q", apiFieldName)
		}
		if _, ok := schema.Attributes[staleFieldName]; ok {
			t.Fatalf("data source schema still exposes %q", staleFieldName)
		}
	})
}

func pivRequirements() *zero_trust_organization.ZeroTrustOrganizationMfaPivKeyRequirementsModel {
	keyTypes := []types.String{types.StringValue("ecdsa")}
	keySizes := []types.Int64{types.Int64Value(256)}
	return &zero_trust_organization.ZeroTrustOrganizationMfaPivKeyRequirementsModel{
		PinPolicy:         types.StringValue("once"),
		RequireFipsDevice: types.BoolValue(true),
		SSHKeySize:        &keySizes,
		SSHKeyType:        &keyTypes,
		TouchPolicy:       types.StringValue("always"),
	}
}

func TestUnmarshalDataSourcePopulatesPivKeyRequirements(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	const apiResponse = `{
	  "result": {
	    "mfa_piv_key_requirements": {
	      "pin_policy": "once",
	      "require_fips_device": false,
	      "ssh_key_size": [256],
	      "ssh_key_type": ["ecdsa"],
	      "touch_policy": "always"
	    }
	  }
	}`

	var env zero_trust_organization.ZeroTrustOrganizationResultDataSourceEnvelope
	if err := apijson.Unmarshal([]byte(apiResponse), &env); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	got, diags := env.Result.MfaPivKeyRequirements.Value(ctx)
	if diags.HasError() {
		t.Fatalf("decode PIV requirements: %v", diags)
	}
	if got == nil {
		t.Fatalf("%s was not decoded from the API response", apiFieldName)
	}
	if got.PinPolicy.ValueString() != "once" || got.TouchPolicy.ValueString() != "always" {
		t.Errorf("unexpected PIV policies: pin=%q touch=%q", got.PinPolicy.ValueString(), got.TouchPolicy.ValueString())
	}
	if !got.RequireFipsDevice.Equal(types.BoolValue(false)) {
		t.Errorf("require_fips_device = %v, want false", got.RequireFipsDevice)
	}
	var keyTypes []types.String
	if diags := got.SSHKeyType.ElementsAs(ctx, &keyTypes, false); diags.HasError() {
		t.Fatalf("decode ssh_key_type: %v", diags)
	}
	if len(keyTypes) != 1 || keyTypes[0].ValueString() != "ecdsa" {
		t.Errorf("ssh_key_type = %v, want [ecdsa]", keyTypes)
	}
	var keySizes []types.Int64
	if diags := got.SSHKeySize.ElementsAs(ctx, &keySizes, false); diags.HasError() {
		t.Fatalf("decode ssh_key_size: %v", diags)
	}
	if len(keySizes) != 1 || keySizes[0].ValueInt64() != 256 {
		t.Errorf("ssh_key_size = %v, want [256]", keySizes)
	}
}

func assertPivKey(t *testing.T, raw []byte) {
	t.Helper()
	var body map[string]json.RawMessage
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("body is not valid JSON: %v\nbody: %s", err, raw)
	}
	if _, ok := body[staleFieldName]; ok {
		t.Errorf("body contains stale field %q, which the API ignores", staleFieldName)
	}
	if _, ok := body[apiFieldName]; !ok {
		t.Errorf("body is missing %q; the endpoint is a full-replace PUT, so omitting it clears the live value", apiFieldName)
	}
}

// Create path.
func TestMarshalJSONSendsPivKeyRequirements(t *testing.T) {
	t.Parallel()

	model := zero_trust_organization.ZeroTrustOrganizationModel{
		Name:                  types.StringValue("Example"),
		AuthDomain:            types.StringValue("example.cloudflareaccess.com"),
		MfaPivKeyRequirements: pivRequirements(),
	}

	raw, err := model.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON: %v", err)
	}
	assertPivKey(t, raw)
}

// Update path: changing an unrelated field must still send the PIV block.
func TestMarshalJSONForUpdateSendsPivKeyRequirements(t *testing.T) {
	t.Parallel()

	state := zero_trust_organization.ZeroTrustOrganizationModel{
		Name:                  types.StringValue("Example"),
		AuthDomain:            types.StringValue("example.cloudflareaccess.com"),
		MfaPivKeyRequirements: pivRequirements(),
	}
	plan := state
	plan.Name = types.StringValue("Example Renamed")

	raw, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate: %v", err)
	}
	assertPivKey(t, raw)
}

// Read path: a wrong tag leaves the attribute null and every plan shows a diff.
func TestUnmarshalPopulatesPivKeyRequirements(t *testing.T) {
	t.Parallel()

	const apiResponse = `{
	  "result": {
	    "name": "Example",
	    "auth_domain": "example.cloudflareaccess.com",
	    "mfa_piv_key_requirements": {
	      "pin_policy": "once",
	      "require_fips_device": true,
	      "ssh_key_size": [256],
	      "ssh_key_type": ["ecdsa"],
	      "touch_policy": "always"
	    }
	  }
	}`

	var env zero_trust_organization.ZeroTrustOrganizationResultEnvelope
	if err := apijson.Unmarshal([]byte(apiResponse), &env); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	got := env.Result.MfaPivKeyRequirements
	if got == nil {
		t.Fatalf("%s was not decoded from the API response", apiFieldName)
	}
	if got.PinPolicy.ValueString() != "once" {
		t.Errorf("pin_policy = %q, want %q", got.PinPolicy.ValueString(), "once")
	}
	if got.TouchPolicy.ValueString() != "always" {
		t.Errorf("touch_policy = %q, want %q", got.TouchPolicy.ValueString(), "always")
	}
	if !got.RequireFipsDevice.ValueBool() {
		t.Error("require_fips_device = false, want true")
	}
	if got.SSHKeyType == nil || len(*got.SSHKeyType) != 1 || (*got.SSHKeyType)[0].ValueString() != "ecdsa" {
		t.Errorf("ssh_key_type = %v, want [ecdsa]", got.SSHKeyType)
	}
	if got.SSHKeySize == nil || len(*got.SSHKeySize) != 1 || (*got.SSHKeySize)[0].ValueInt64() != 256 {
		t.Errorf("ssh_key_size = %v, want [256]", got.SSHKeySize)
	}
}
