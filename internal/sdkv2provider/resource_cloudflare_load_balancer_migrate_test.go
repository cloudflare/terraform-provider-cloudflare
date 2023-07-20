package sdkv2provider

import (
	"context"
	"reflect"
	"testing"
)

func testCloudflareLoadBalancerDataV0() map[string]interface{} {
	return map[string]interface{}{
		"rules": []interface{}{map[string]interface{}{
			"fixed_response": map[string]interface{}{
				"message_body": "example",
			},
		}},
	}
}

func testCloudflareLoadBalancerDataV1() map[string]interface{} {
	v0 := testCloudflareLoadBalancerDataV0()
	return map[string]interface{}{
		"rules": []interface{}{map[string]interface{}{
			"fixed_response": []interface{}{
				v0["rules"].([]interface{})[0].(map[string]interface{})["fixed_response"],
			},
		}},
	}
}

func TestCloudflareLoadBalancerUpgradeV0(t *testing.T) {
	expected := testCloudflareLoadBalancerDataV1()
	actual, err := resourceCloudflareLoadBalancerStateUpgradeV1(context.TODO(), testCloudflareLoadBalancerDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
