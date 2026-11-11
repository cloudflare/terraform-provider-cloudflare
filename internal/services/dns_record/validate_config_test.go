package dns_record_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// nullObject returns an object of the given type with every attribute null,
// except those set in values.
func nullObject(objectType tftypes.Object, values map[string]tftypes.Value) tftypes.Value {
	attrs := make(map[string]tftypes.Value, len(objectType.AttributeTypes))
	for name, attrType := range objectType.AttributeTypes {
		attrs[name] = tftypes.NewValue(attrType, nil)
	}
	for name, value := range values {
		attrs[name] = value
	}
	return tftypes.NewValue(objectType, attrs)
}

func TestDNSRecordValidateConfig_ContentRequired(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	schema := dns_record.ResourceSchema(ctx)
	configType := schema.Type().TerraformType(ctx).(tftypes.Object)
	dataType := configType.AttributeTypes["data"].(tftypes.Object)

	str := func(v interface{}) tftypes.Value { return tftypes.NewValue(tftypes.String, v) }
	caaData := nullObject(dataType, map[string]tftypes.Value{
		"tag":   str("issue"),
		"value": str("letsencrypt.org"),
	})

	for _, tt := range []struct {
		name      string
		values    map[string]tftypes.Value
		wantError bool
	}{
		{"a_without_content", map[string]tftypes.Value{"type": str("A")}, true},
		{"a_with_unknown_content", map[string]tftypes.Value{"type": str("A"), "content": str(tftypes.UnknownValue)}, false},
		{"caa_with_data", map[string]tftypes.Value{"type": str("CAA"), "data": caaData}, false},
		{"lowercase_a_without_content", map[string]tftypes.Value{"type": str("a")}, true},
		{"mx_with_content", map[string]tftypes.Value{"type": str("MX"), "content": str("mail.example.com")}, false},
		{"txt_without_content", map[string]tftypes.Value{"type": str("TXT")}, true},
		{"unknown_type", map[string]tftypes.Value{"type": str(tftypes.UnknownValue)}, false},
		{"a_with_unknown_data", map[string]tftypes.Value{"type": str("A"), "data": tftypes.NewValue(dataType, tftypes.UnknownValue)}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			values := map[string]tftypes.Value{
				"zone_id": str("zone-id"),
				"name":    str("terraform.example.com"),
				"ttl":     tftypes.NewValue(tftypes.Number, 1),
			}
			for name, value := range tt.values {
				values[name] = value
			}
			config := tfsdk.Config{Schema: schema, Raw: nullObject(configType, values)}

			response := &resource.ValidateConfigResponse{}
			resourceUnderTest, ok := dns_record.NewResource().(resource.ResourceWithValidateConfig)
			if !ok {
				t.Fatal("DNS record resource does not implement ValidateConfig")
			}
			resourceUnderTest.ValidateConfig(ctx, resource.ValidateConfigRequest{Config: config}, response)

			if response.Diagnostics.HasError() != tt.wantError {
				t.Fatalf("diagnostics = %v, want error = %v", response.Diagnostics, tt.wantError)
			}
			if tt.wantError {
				for _, d := range response.Diagnostics.Errors() {
					withPath, ok := d.(interface{ Path() path.Path })
					if !ok || !withPath.Path().Equal(path.Root("content")) {
						t.Errorf("error %q is not on content", d.Summary())
					}
				}
			}
		})
	}
}
