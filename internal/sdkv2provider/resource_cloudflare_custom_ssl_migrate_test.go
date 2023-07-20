package sdkv2provider

import (
	"context"
	"reflect"
	"testing"
)

func testCloudflareCustomSSLDataV0() map[string]interface{} {
	return map[string]interface{}{
		"custom_ssl_options": map[string]interface{}{},
	}
}

func testCloudflareCustomSSLDataV1() map[string]interface{} {
	v0 := testCloudflareCustomSSLDataV0()
	return map[string]interface{}{
		"custom_ssl_options": []interface{}{v0["custom_ssl_options"]},
	}
}

func TestCloudflareCustomSSLUpgradeV0(t *testing.T) {
	expected := testCloudflareCustomSSLDataV1()
	actual, err := resourceCloudflareCustomSSLStateUpgradeV1(context.TODO(), testCloudflareCustomSSLDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
