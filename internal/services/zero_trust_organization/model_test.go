package zero_trust_organization_test

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	apiFieldName = "mfa_ssh_piv_key_requirements"
)

func pivRequirements() *zero_trust_organization.ZeroTrustOrganizationMfaSSHPivKeyRequirementsModel {
	keyTypes := []types.String{types.StringValue("ecdsa")}
	keySizes := []types.Int64{types.Int64Value(256)}
	return &zero_trust_organization.ZeroTrustOrganizationMfaSSHPivKeyRequirementsModel{
		PinPolicy:         types.StringValue("once"),
		RequireFipsDevice: types.BoolValue(true),
		SSHKeySize:        &keySizes,
		SSHKeyType:        &keyTypes,
		TouchPolicy:       types.StringValue("always"),
	}
}

func assertPivKey(t *testing.T, raw []byte) {
	t.Helper()
	var body map[string]json.RawMessage
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("body is not valid JSON: %v\nbody: %s", err, raw)
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
		MfaSSHPivKeyRequirements: pivRequirements(),
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
		MfaSSHPivKeyRequirements: pivRequirements(),
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
	    "mfa_ssh_piv_key_requirements": {
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

	got := env.Result.MfaSSHPivKeyRequirements
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
