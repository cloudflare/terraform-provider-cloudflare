package cloudflare

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func testCloudflareApiTokenStateDataV0() map[string]interface{} {
	return map[string]interface{}{
		"policy": schema.NewSet(
			resourceCloudflareApiTokenSchemaV0().Schema["policy"].Set,
			[]interface{}{
				map[string]interface{}{
					"resources":         map[string]interface{}{},
					"effect":            "",
					"permission_groups": []interface{}{"test"},
				},
			},
		),
	}
}

func testCloudflareApiTokenStateDataV1() map[string]interface{} {
	v0 := testCloudflareApiTokenStateDataV0()
	return map[string]interface{}{
		"policy": schema.NewSet(
			resourceCloudflareApiTokenSchema()["policy"].Set,
			[]interface{}{
				map[string]interface{}{
					"resources": map[string]interface{}{},
					"effect":    "",
					"permission_groups": schema.NewSet(
						schema.HashString,
						v0["policy"].(*schema.Set).List()[0].(map[string]interface{})["permission_groups"].([]interface{}),
					),
				},
			},
		),
	}
}

func TestCloudflareApiTokenStateUpgradeV0(t *testing.T) {
	expected := testCloudflareApiTokenStateDataV1()
	actual, err := resourceCloudflareApiTokenStateUpgradeV0(context.TODO(), testCloudflareApiTokenStateDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(
		expected["policy"].(*schema.Set).List()[0].(map[string]interface{})["permission_groups"].(*schema.Set).List(),
		actual["policy"].(*schema.Set).List()[0].(map[string]interface{})["permission_groups"].(*schema.Set).List()) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
