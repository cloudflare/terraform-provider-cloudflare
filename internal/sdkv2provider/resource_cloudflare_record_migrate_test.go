package sdkv2provider

import (
	"context"
	"reflect"
	"testing"
)

func testCloudflareRecordStateDataV1() map[string]interface{} {
	return map[string]interface{}{
		"data": map[string]interface{}{},
	}
}

func testCloudflareRecordStateDataV2() map[string]interface{} {
	v0 := testCloudflareRecordStateDataV1()
	return map[string]interface{}{
		"data": []interface{}{v0["data"]},
	}
}

func TestCloudflareRecordStateUpgradeV0(t *testing.T) {
	expected := testCloudflareRecordStateDataV2()
	actual, err := resourceCloudflareRecordStateUpgradeV2(context.TODO(), testCloudflareRecordStateDataV1(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
