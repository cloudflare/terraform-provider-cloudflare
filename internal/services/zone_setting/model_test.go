package zone_setting

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func objectValue(t *testing.T, attrs map[string]attr.Value) customfield.NormalizedDynamicValue {
	t.Helper()
	attrTypes := make(map[string]attr.Type, len(attrs))
	for k, v := range attrs {
		attrTypes[k] = v.Type(t.Context())
	}
	obj, diags := types.ObjectValue(attrTypes, attrs)
	if diags.HasError() {
		t.Fatalf("building object value: %v", diags)
	}
	return customfield.RawNormalizedDynamicValueFrom(obj)
}

func settingModel(value customfield.NormalizedDynamicValue) ZoneSettingModel {
	return ZoneSettingModel{
		ID:        types.StringValue("aegis"),
		SettingID: types.StringValue("aegis"),
		ZoneID:    types.StringValue("0da42c8d2132a9ddaf714f9e7c920711"),
		Value:     value,
		Enabled:   types.BoolUnknown(),
	}
}

// The API omits pool_id from the aegis setting once it is disabled, so the
// refreshed state no longer matches a config that still sets pool_id. The
// update must still send enabled=false; sending only the changed pool_id key
// enables the setting.
func TestMarshalJSONForUpdate_ObjectValueSendsUnchangedKeys(t *testing.T) {
	const poolID = "0123456789abcdef0123456789abcdef"

	plan := settingModel(objectValue(t, map[string]attr.Value{
		"enabled": types.BoolValue(false),
		"pool_id": types.StringValue(poolID),
	}))
	state := settingModel(objectValue(t, map[string]attr.Value{
		"enabled": types.BoolValue(false),
		"pool_id": types.StringNull(),
	}))

	got, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate: %v", err)
	}

	want := `{"value":{"enabled":false,"pool_id":"` + poolID + `"}}`
	if string(got) != want {
		t.Errorf("MarshalJSONForUpdate() = %s, want %s", got, want)
	}
}

func TestMarshalJSONForUpdate_ObjectValueDisable(t *testing.T) {
	const poolID = "0123456789abcdef0123456789abcdef"

	plan := settingModel(objectValue(t, map[string]attr.Value{
		"enabled": types.BoolValue(false),
		"pool_id": types.StringValue(poolID),
	}))
	state := settingModel(objectValue(t, map[string]attr.Value{
		"enabled": types.BoolValue(true),
		"pool_id": types.StringValue(poolID),
	}))

	got, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate: %v", err)
	}

	want := `{"value":{"enabled":false,"pool_id":"` + poolID + `"}}`
	if string(got) != want {
		t.Errorf("MarshalJSONForUpdate() = %s, want %s", got, want)
	}
}

func TestMarshalJSONForUpdate_ScalarValue(t *testing.T) {
	plan := settingModel(customfield.RawNormalizedDynamicValueFrom(types.StringValue("on")))
	state := settingModel(customfield.RawNormalizedDynamicValueFrom(types.StringValue("off")))

	got, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate: %v", err)
	}

	if want := `{"value":"on"}`; string(got) != want {
		t.Errorf("MarshalJSONForUpdate() = %s, want %s", got, want)
	}
}
