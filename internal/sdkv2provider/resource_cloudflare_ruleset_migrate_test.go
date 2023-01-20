package sdkv2provider

import (
	"context"
	"reflect"
	"testing"
)

func testCloudflareRulesetStateDataV0() map[string]interface{} {
	return map[string]interface{}{
		"action":      "block",
		"description": "example http rate limit",
		"enabled":     true,
		"expression":  "(http.request.uri.path matches \"^/api/\")",
		"id":          "5a297db75a614c318edb406fe5eef502",
		"ratelimit": []map[string]interface{}{{
			"characteristics":       []string{"cf.colo.id", "ip.src"},
			"mitigation_timeout":    600,
			"period":                60,
			"requests_per_period":   100,
			"mitigation_expression": "",
		}},
	}
}

func testCloudflareRulesetStateDataV1() map[string]interface{} {
	return map[string]interface{}{
		"action":      "block",
		"description": "example http rate limit",
		"enabled":     true,
		"expression":  "(http.request.uri.path matches \"^/api/\")",
		"id":          "5a297db75a614c318edb406fe5eef502",
		"ratelimit": []map[string]interface{}{{
			"characteristics":     []string{"cf.colo.id", "ip.src"},
			"mitigation_timeout":  600,
			"period":              60,
			"requests_per_period": 100,
			"counting_expression": "",
		}},
	}
}

func TestCloudflareRulesetStateUpgradeV0(t *testing.T) {
	expected := testCloudflareRulesetStateDataV1()
	actual, err := resourceCloudflareRulesetStateUpgradeV0ToV1(context.Background(), testCloudflareRulesetStateDataV0(), nil)
	if err != nil {
		t.Fatalf("error migrating ruleset state V0 to V1: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
