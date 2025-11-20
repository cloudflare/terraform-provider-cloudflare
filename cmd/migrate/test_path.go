package main

import (
	"fmt"
	"testing"
	"github.com/tidwall/gjson"
)

func TestPathExtraction(t *testing.T) {
	input := `{
		"resources": [{
			"type": "cloudflare_ruleset",
			"instances": [{
				"schema_version": 1,
				"attributes": {
					"id": "abc123",
					"rules.#": "1",
					"rules.0.id": "rule1",
					"rules.0.action": "block"
				}
			}]
		}]
	}`
	
	// Test different path formats
	paths := []string{
		"resources.0.instances.0.attributes.rules\\.#",
		`resources.0.instances.0.attributes.rules\.#`,
		"resources.0.instances.0.attributes.rules.#",
		`resources.0.instances.0.attributes.rules.0.id`,
		"resources.0.instances.0.attributes.rules\\.0\\.id",
	}
	
	for _, path := range paths {
		val := gjson.Get(input, path)
		fmt.Printf("Path: %s -> Exists: %v, Value: %s\n", path, val.Exists(), val.String())
	}
}
