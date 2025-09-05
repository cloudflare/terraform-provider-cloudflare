package main

import (
	"encoding/json"
	"testing"
)

func TestTransformStateJSON(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms_load_balancer_pool_with_empty_arrays",
			Input: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"identity_schema_version": 0,
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"load_shedding": [],
							"origin_steering": [],
							"origins": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"origins": []
						}
					}]
				}]
			}`,
		},
		{
			Name: "keeps_non_empty_arrays",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"load_shedding": [{"default_percent": 10}],
							"origin_steering": [{"policy": "random"}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"load_shedding": {"default_percent": 10},
							"origin_steering": {"policy": "random"}
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_non_load_balancer_pool_resources",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zone",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"some_array": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_zone",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"some_array": []
						}
					}]
				}]
			}`,
		},
		{
			Name: "removes_empty_header_arrays_in_origins",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"origins": [
								{
									"name": "origin-1",
									"address": "192.0.2.1",
									"header": []
								},
								{
									"name": "origin-2",
									"address": "192.0.2.2",
									"header": []
								}
							]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"name": "test-pool",
							"origins": [
								{
									"name": "origin-1",
									"address": "192.0.2.1"
								},
								{
									"name": "origin-2",
									"address": "192.0.2.2"
								}
							]
						}
					}]
				}]
			}`,
		},
		{
			Name: "transforms_header_formats_in_state",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"origins": [{
								"name": "origin-1",
								"address": "192.0.2.1",
								"header": [{
									"header": "Host",
									"values": ["example.com"]
								}]
							}]
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer_pool",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"origins": [{
								"name": "origin-1",
								"address": "192.0.2.1",
								"header": {
									"host": ["example.com"]
								}
							}]
						}
					}]
				}]
			}`,
		},
		{
			Name: "handles_load_balancer_transformations",
			Input: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"fallback_pool_id": "pool-123",
							"default_pool_ids": ["pool-1", "pool-2"],
							"adaptive_routing": [],
							"country_pools": [],
							"rules": []
						}
					}]
				}]
			}`,
			Expected: `{
				"version": 4,
				"resources": [{
					"type": "cloudflare_load_balancer",
					"name": "test",
					"instances": [{
						"attributes": {
							"id": "test-id",
							"fallback_pool": "pool-123",
							"default_pools": ["pool-1", "pool-2"],
							"country_pools": {},
							"rules": []
						}
					}]
				}]
			}`,
		},
		{
			Name: "complete_transformation_with_multiple_resources",
			Input: `{
				"version": 4,
				"resources": [
					{
						"type": "cloudflare_load_balancer_pool",
						"name": "pool1",
						"instances": [{
							"identity_schema_version": 0,
							"attributes": {
								"id": "pool-id",
								"name": "test-pool",
								"load_shedding": [{"default_percent": 10}],
								"origin_steering": [],
								"origins": [{
									"name": "origin-1",
									"address": "192.0.2.1",
									"header": [{
										"header": "Host",
										"values": ["example.com"]
									}]
								}]
							}
						}]
					},
					{
						"type": "cloudflare_load_balancer",
						"name": "lb1",
						"instances": [{
							"attributes": {
								"id": "lb-id",
								"name": "test-lb",
								"fallback_pool_id": "pool-123",
								"default_pool_ids": ["pool-1", "pool-2"],
								"adaptive_routing": [],
								"country_pools": [],
								"pop_pools": [],
								"rules": []
							}
						}]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"resources": [
					{
						"type": "cloudflare_load_balancer_pool",
						"name": "pool1",
						"instances": [{
							"attributes": {
								"id": "pool-id",
								"name": "test-pool",
								"load_shedding": {"default_percent": 10},
								"origins": [{
									"name": "origin-1",
									"address": "192.0.2.1",
									"header": {
										"host": ["example.com"]
									}
								}]
							}
						}]
					},
					{
						"type": "cloudflare_load_balancer",
						"name": "lb1",
						"instances": [{
							"attributes": {
								"id": "lb-id",
								"name": "test-lb",
								"fallback_pool": "pool-123",
								"default_pools": ["pool-1", "pool-2"],
								"country_pools": {},
								"pop_pools": {},
								"rules": []
							}
						}]
					}
				]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}

func TestTransformManagedTransformsState(t *testing.T) {
	// Test state with managed_transforms that's missing response headers
	testState := `{
		"version": 4,
		"resources": [{
			"type": "cloudflare_managed_transforms",
			"instances": [{
				"attributes": {
					"id": "test-zone-id",
					"zone_id": "test-zone-id",
					"managed_request_headers": [
						{"id": "add_true_client_ip_headers", "enabled": true}
					]
				}
			}]
		}]
	}`

	// Transform the state
	result, err := transformStateJSON([]byte(testState))
	if err != nil {
		t.Fatalf("Failed to transform state: %v", err)
	}

	// Parse the result
	var resultMap map[string]interface{}
	if err := json.Unmarshal(result, &resultMap); err != nil {
		t.Fatalf("Failed to parse result: %v", err)
	}

	// Check if managed_response_headers was added
	resources := resultMap["resources"].([]interface{})
	resource := resources[0].(map[string]interface{})
	instances := resource["instances"].([]interface{})
	instance := instances[0].(map[string]interface{})
	attributes := instance["attributes"].(map[string]interface{})

	if _, exists := attributes["managed_response_headers"]; !exists {
		t.Error("managed_response_headers was not added to the state")
	}

	// Verify it's an empty array
	responseHeaders := attributes["managed_response_headers"].([]interface{})
	if len(responseHeaders) != 0 {
		t.Errorf("Expected empty managed_response_headers, got %v", responseHeaders)
	}
}

func TestCustomPagesStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms type to identifier for account-level custom pages",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"type": "500_errors",
									"state": "customized",
									"url": "https://example.com/500.html",
									"id": "custom-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "500_errors",
									"state": "customized",
									"url": "https://example.com/500.html",
									"id": "custom-page-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "transforms type to identifier for zone-level custom pages",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"zone_id": "abcdef123456",
									"type": "ip_block",
									"state": "default",
									"url": "",
									"id": "custom-page-zone-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"zone_id": "abcdef123456",
									"identifier": "ip_block",
									"state": "default",
									"url": "",
									"id": "custom-page-zone-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "handles multiple custom pages in one state file",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "account_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"type": "basic_challenge",
									"state": "customized",
									"url": "https://example.com/challenge.html",
									"id": "account-page-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "zone_page",
						"instances": [
							{
								"attributes": {
									"zone_id": "fedcba987654321",
									"type": "managed_challenge",
									"state": "default",
									"url": "",
									"id": "zone-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "account_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "basic_challenge",
									"state": "customized",
									"url": "https://example.com/challenge.html",
									"id": "account-page-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "zone_page",
						"instances": [
							{
								"attributes": {
									"zone_id": "fedcba987654321",
									"identifier": "managed_challenge",
									"state": "default",
									"url": "",
									"id": "zone-page-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "leaves already migrated custom pages unchanged",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "already_migrated",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "waf_challenge",
									"state": "customized",
									"url": "https://example.com/waf.html",
									"id": "migrated-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "already_migrated",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "waf_challenge",
									"state": "customized",
									"url": "https://example.com/waf.html",
									"id": "migrated-page-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "handles mixed resource types with custom pages",
			Input: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zone",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"account": {
										"id": "123456789abcdef0"
									},
									"name": "example.com",
									"id": "zone-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "error_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"type": "1000_errors",
									"state": "customized",
									"url": "https://example.com/1000.html",
									"id": "error-page-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_zone",
						"name": "test_zone",
						"instances": [
							{
								"attributes": {
									"account": {
										"id": "123456789abcdef0"
									},
									"name": "example.com",
									"id": "zone-id"
								}
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_custom_pages",
						"name": "error_page",
						"instances": [
							{
								"attributes": {
									"account_id": "123456789abcdef0",
									"identifier": "1000_errors",
									"state": "customized",
									"url": "https://example.com/1000.html",
									"id": "error-page-id"
								}
							}
						]
					}
				]
			}`,
		},
	}

	RunFullStateTransformationTests(t, tests)
}