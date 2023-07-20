package sdkv2provider

import (
	"context"
	"reflect"
	"testing"
)

func testCloudflareAccessRuleStateDataV0() map[string]interface{} {
	return map[string]interface{}{
		"configuration": map[string]interface{}{},
	}
}

func testCloudflareAccessRuleStateDataV1() map[string]interface{} {
	v0 := testCloudflareAccessRuleStateDataV0()
	return map[string]interface{}{
		"configuration": []interface{}{v0["configuration"]},
	}
}

func TestCloudflareAccessRuleStateUpgradeV0(t *testing.T) {
	expected := testCloudflareAccessRuleStateDataV1()
	actual, err := resourceCloudflareAccessRuleStateUpgradeV1(context.TODO(), testCloudflareAccessRuleStateDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
