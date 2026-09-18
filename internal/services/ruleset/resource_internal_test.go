package ruleset

import (
	"fmt"
	"testing"

	"github.com/tidwall/gjson"
)

func TestTransformQueryStringJSONLegacyWildcardStrings(t *testing.T) {
	t.Parallel()

	transformed := transformQueryStringJSON([]byte(`{
		"result": {
			"rules": [{
				"action_parameters": {
					"cache_key": {
						"custom_key": {
							"query_string": {"include": "*", "exclude": "*"}
						}
					}
				}
			}]
		}
	}`))

	for _, field := range []string{"include", "exclude"} {
		path := "result.rules.0.action_parameters.cache_key.custom_key.query_string." + field
		value := gjson.GetBytes(transformed, path)
		if !value.IsObject() || !value.Get("all").Bool() {
			t.Errorf("expected %s legacy wildcard to be normalized to {all: true}, got %s", field, value.Raw)
		}
	}
}

func TestTransformQueryStringJSONPreservesObjectAndListSemantics(t *testing.T) {
	t.Parallel()

	input := []byte(`{
		"result": {
			"rules": [
				{
					"action_parameters": {
						"cache_key": {
							"custom_key": {
								"query_string": {"include": {"all": true}, "exclude": {"list": ["skip"]}}
							}
						}
					}
				},
				{
					"action_parameters": {
						"cache_key": {
							"custom_key": {
								"query_string": {"include": ["*"], "exclude": ["*"]}
							}
						}
					}
				},
				{
					"action_parameters": {
						"cache_key": {
							"custom_key": {
								"query_string": {"include": ["keep"], "exclude": ["drop"]}
							}
						}
					}
				}
			]
		}
	}`)

	transformed := transformQueryStringJSON(input)
	queryStringPath := "result.rules.%d.action_parameters.cache_key.custom_key.query_string."

	include := gjson.GetBytes(transformed, "result.rules.0.action_parameters.cache_key.custom_key.query_string.include")
	if !include.IsObject() || !include.Get("all").Bool() {
		t.Errorf("object include = %s, want object with all=true", include.Raw)
	}
	exclude := gjson.GetBytes(transformed, "result.rules.0.action_parameters.cache_key.custom_key.query_string.exclude")
	if !exclude.IsObject() || !exclude.Get("list").IsArray() || exclude.Get("list.#").Int() != 1 || exclude.Get("list.0").String() != "skip" {
		t.Errorf("object exclude = %s, want object with list=[skip]", exclude.Raw)
	}
	if !gjson.GetBytes(transformed, fmt.Sprintf(queryStringPath, 1)+"include.all").Bool() {
		t.Error("expected include array wildcard to retain existing all=true normalization")
	}
	if got := gjson.GetBytes(transformed, fmt.Sprintf(queryStringPath, 1)+"exclude.list.0").String(); got != "*" {
		t.Errorf("exclude array wildcard = %q, want list value", got)
	}
	if gjson.GetBytes(transformed, fmt.Sprintf(queryStringPath, 1)+"exclude.all").Exists() {
		t.Error("exclude array wildcard must not be normalized to all=true")
	}
	if got := gjson.GetBytes(transformed, fmt.Sprintf(queryStringPath, 2)+"include.list.0").String(); got != "keep" {
		t.Errorf("ordinary include array = %q, want list value", got)
	}
	if got := gjson.GetBytes(transformed, fmt.Sprintf(queryStringPath, 2)+"exclude.list.0").String(); got != "drop" {
		t.Errorf("ordinary exclude array = %q, want list value", got)
	}
}
