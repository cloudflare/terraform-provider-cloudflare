package main

import (
	"testing"
)

// Config transformation tests
func TestTieredCacheTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "tiered_cache with cache_type=smart",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id    = "test-zone-id"
  cache_type = "smart"
}`,
			Expected: []string{`
resource "cloudflare_tiered_cache" "example" {
  zone_id = "test-zone-id"
  value   = "on"
}`},
		},
		{
			Name: "tiered_cache with cache_type=generic creates argo_tiered_caching",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id    = "test-zone-id"
  cache_type = "generic"
}`,
			Expected: []string{
				`resource "cloudflare_argo_tiered_caching" "example"`,
				`moved {
  from = cloudflare_tiered_cache.example
  to   = cloudflare_argo_tiered_caching.example
}`,
			},
		},
		{
			Name: "tiered_cache with cache_type=off",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id    = "test-zone-id"
  cache_type = "off"
}`,
			Expected: []string{`
resource "cloudflare_tiered_cache" "example" {
  zone_id = "test-zone-id"
  value   = "off"
}`},
		},
		{
			Name: "tiered_cache with variable cache_type",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id    = "test-zone-id"
  cache_type = var.cache_type_value
}`,
			Expected: []string{`
resource "cloudflare_tiered_cache" "example" {
  zone_id = "test-zone-id"
  value   = var.cache_type_value
}`},
		},
		{
			Name: "tiered_cache without cache_type",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id = "test-zone-id"
  value   = "on"
}`,
			Expected: []string{`
resource "cloudflare_tiered_cache" "example" {
  zone_id = "test-zone-id"
  value   = "on"
}`},
		},
		{
			Name: "multiple resources including tiered_cache",
			Config: `
resource "cloudflare_zone" "example" {
  zone = "example.com"
}

resource "cloudflare_tiered_cache" "example" {
  zone_id    = cloudflare_zone.example.id
  cache_type = "smart"
}

resource "cloudflare_record" "example" {
  zone_id = cloudflare_zone.example.id
  name    = "test"
  value   = "192.0.2.1"
  type    = "A"
}`,
			Expected: []string{
				`resource "cloudflare_zone" "example"`,
				`resource "cloudflare_tiered_cache" "example" {
  zone_id = cloudflare_zone.example.id
  value   = "on"
}`,
				`resource "cloudflare_record" "example"`,
			},
		},
		{
			Name: "tiered_cache with cache_type=generic should create argo_tiered_caching resource and moved block",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id    = "test-zone-id"
  cache_type = "generic"
}`,
			Expected: []string{
				`resource "cloudflare_argo_tiered_caching" "example" {
  zone_id = "test-zone-id"
  value   = "on"
}`,
				`moved {
  from = cloudflare_tiered_cache.example
  to   = cloudflare_argo_tiered_caching.example
}`,
			},
		},
		{
			Name: "multiple tiered_cache resources with mixed types",
			Config: `
resource "cloudflare_tiered_cache" "generic_one" {
  zone_id    = "zone1"
  cache_type = "generic"
}

resource "cloudflare_tiered_cache" "smart_one" {
  zone_id    = "zone2"
  cache_type = "smart"
}

resource "cloudflare_tiered_cache" "generic_two" {
  zone_id    = "zone3"
  cache_type = "generic"
}`,
			Expected: []string{
				`resource "cloudflare_argo_tiered_caching" "generic_one"`,
				`zone_id = "zone1"`,
				`moved {
  from = cloudflare_tiered_cache.generic_one
  to   = cloudflare_argo_tiered_caching.generic_one
}`,
				`resource "cloudflare_tiered_cache" "smart_one" {
  zone_id = "zone2"
  value   = "on"
}`,
				`resource "cloudflare_argo_tiered_caching" "generic_two"`,
				`zone_id = "zone3"`,
				`moved {
  from = cloudflare_tiered_cache.generic_two
  to   = cloudflare_argo_tiered_caching.generic_two
}`,
			},
		},
		{
			Name: "tiered_cache with dynamic cache_type should not generate moved block",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id    = "test-zone-id"
  cache_type = var.cache_type
}`,
			Expected: []string{
				`resource "cloudflare_tiered_cache" "example" {
  zone_id = "test-zone-id"
  value   = var.cache_type
}`,
			},
		},
		{
			Name: "tiered_cache with cache_type=generic and other attributes",
			Config: `
resource "cloudflare_tiered_cache" "example" {
  zone_id    = cloudflare_zone.example.id
  cache_type = "generic"
  
  lifecycle {
    create_before_destroy = true
  }
}`,
			Expected: []string{
				`resource "cloudflare_argo_tiered_caching" "example"`,
				`zone_id = cloudflare_zone.example.id`,
				`create_before_destroy = true`,
				`moved {
  from = cloudflare_tiered_cache.example
  to   = cloudflare_argo_tiered_caching.example
}`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

// State transformation tests
func TestTieredCacheStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "transforms_tiered_cache_generic_to_argo_tiered_caching",
			Input: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"cache_type": "generic",
									"id": "test-id",
									"editable": true,
									"modified_on": "2024-01-01T00:00:00Z"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"resources": [
					{
						"type": "cloudflare_argo_tiered_caching",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"value": "on",
									"id": "test-id",
									"editable": true,
									"modified_on": "2024-01-01T00:00:00Z"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "transforms_tiered_cache_smart_value",
			Input: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"cache_type": "smart",
									"id": "test-id",
									"editable": true,
									"modified_on": "2024-01-01T00:00:00Z"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"value": "on",
									"id": "test-id",
									"editable": true,
									"modified_on": "2024-01-01T00:00:00Z"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "transforms_tiered_cache_off_value",
			Input: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"cache_type": "off",
									"id": "test-id",
									"editable": true,
									"modified_on": "2024-01-01T00:00:00Z"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"value": "off",
									"id": "test-id",
									"editable": true,
									"modified_on": "2024-01-01T00:00:00Z"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "handles_multiple_resources_with_mixed_types",
			Input: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "generic_test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone1",
									"cache_type": "generic",
									"id": "id1"
								}
							}
						]
					},
					{
						"type": "cloudflare_tiered_cache",
						"name": "smart_test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone2",
									"cache_type": "smart",
									"id": "id2"
								}
							}
						]
					},
					{
						"type": "cloudflare_load_balancer",
						"name": "lb_test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone3",
									"fallback_pool_id": "pool1",
									"id": "id3"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"version": 4,
				"terraform_version": "1.12.2",
				"resources": [
					{
						"type": "cloudflare_argo_tiered_caching",
						"name": "generic_test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone1",
									"value": "on",
									"id": "id1"
								}
							}
						]
					},
					{
						"type": "cloudflare_tiered_cache",
						"name": "smart_test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone2",
									"value": "on",
									"id": "id2"
								}
							}
						]
					},
					{
						"type": "cloudflare_load_balancer",
						"name": "lb_test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone3",
									"fallback_pool": "pool1",
									"id": "id3"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "handles_tiered_cache_without_cache_type",
			Input: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"id": "test-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "test-zone-id",
									"id": "test-id"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "handles_tiered_cache_with_multiple_instances",
			Input: `{
				"resources": [
					{
						"type": "cloudflare_tiered_cache",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone1",
									"cache_type": "generic",
									"id": "id1"
								}
							},
							{
								"attributes": {
									"zone_id": "zone2",
									"cache_type": "smart",
									"id": "id2"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"resources": [
					{
						"type": "cloudflare_argo_tiered_caching",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone_id": "zone1",
									"value": "on",
									"id": "id1"
								}
							},
							{
								"attributes": {
									"zone_id": "zone2",
									"value": "on",
									"id": "id2"
								}
							}
						]
					}
				]
			}`,
		},
		{
			Name: "preserves_non_tiered_cache_resources",
			Input: `{
				"resources": [
					{
						"type": "cloudflare_zone",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone": "example.com",
									"id": "test-id"
								}
							}
						]
					}
				]
			}`,
			Expected: `{
				"resources": [
					{
						"type": "cloudflare_zone",
						"name": "test",
						"instances": [
							{
								"attributes": {
									"zone": "example.com",
									"id": "test-id"
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
