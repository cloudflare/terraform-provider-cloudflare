package main

import (
	"testing"
)

// Config transformation tests
func TestLoadBalancerPoolTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "header_as_attribute_inside_origins_list",
			Config: `
resource "cloudflare_load_balancer_pool" "test" {
  origins = [{
    name = "test"
    address = "192.0.2.1"
    header = {
      header = "Host"
      values = ["example.com"]
    }
  }]
}`,
			Expected: []string{`
resource "cloudflare_load_balancer_pool" "test" {
  origins = [{
    name    = "test"
    address = "192.0.2.1"
    header = {
      header = "Host"
      values = ["example.com"]
    }
  }]
}`},
		},
		{
			Name: "no_header_block",
			Config: `
resource "cloudflare_load_balancer_pool" "test" {
  origins = [{
    name = "test"
    address = "192.0.2.1"
  }]
}`,
			Expected: []string{`
resource "cloudflare_load_balancer_pool" "test" {
  origins = [{
    name    = "test"
    address = "192.0.2.1"
  }]
}`},
		},
		{
			Name: "multiple_origins",
			Config: `
resource "cloudflare_load_balancer_pool" "test" {
  origins = [{
    name = "test1"
    address = "192.0.2.1"
  }, {
    name = "test2"
    address = "192.0.2.2"
  }]
}`,
			Expected: []string{`
resource "cloudflare_load_balancer_pool" "test" {
  origins = [{
    name    = "test1"
    address = "192.0.2.1"
  }, {
    name    = "test2"
    address = "192.0.2.2"
  }]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

// State transformation tests
func TestLoadBalancerPoolStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "removes_empty_load_shedding_array",
			Input: `{
				"id": "test-id",
				"name": "test-pool",
				"load_shedding": [],
				"origins": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": []
			}`,
		},
		{
			Name: "converts_load_shedding_array_to_object",
			Input: `{
				"id": "test-id",
				"name": "test-pool",
				"load_shedding": [{
					"default_percent": 10,
					"default_policy": "random"
				}],
				"origins": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-pool",
				"load_shedding": {
					"default_percent": 10,
					"default_policy": "random"
				},
				"origins": []
			}`,
		},
		{
			Name: "removes_empty_origin_steering_array",
			Input: `{
				"id": "test-id",
				"name": "test-pool",
				"origin_steering": [],
				"origins": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": []
			}`,
		},
		{
			Name: "converts_origin_steering_array_to_object",
			Input: `{
				"id": "test-id",
				"name": "test-pool",
				"origin_steering": [{
					"policy": "least_connections"
				}],
				"origins": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-pool",
				"origin_steering": {
					"policy": "least_connections"
				},
				"origins": []
			}`,
		},
		{
			Name: "removes_empty_header_arrays",
			Input: `{
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
			}`,
			Expected: `{
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
			}`,
		},
		{
			Name: "transforms_v4_header_format_to_v5",
			Input: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": [{
					"name": "origin-1",
					"address": "192.0.2.1",
					"header": [{
						"header": "Host",
						"values": ["example.com", "www.example.com"]
					}]
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": [{
					"name": "origin-1",
					"address": "192.0.2.1",
					"header": {
						"host": ["example.com", "www.example.com"]
					}
				}]
			}`,
		},
		{
			Name: "transforms_grit_intermediate_header_format_to_v5",
			Input: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": [{
					"name": "origin-1",
					"address": "192.0.2.1",
					"header": {
						"header": "Host",
						"values": ["example.com"]
					}
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": [{
					"name": "origin-1",
					"address": "192.0.2.1",
					"header": {
						"host": ["example.com"]
					}
				}]
			}`,
		},
		{
			Name: "leaves_v5_header_format_unchanged",
			Input: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": [{
					"name": "origin-1",
					"address": "192.0.2.1",
					"header": {
						"host": ["example.com"]
					}
				}]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-pool",
				"origins": [{
					"name": "origin-1",
					"address": "192.0.2.1",
					"header": {
						"host": ["example.com"]
					}
				}]
			}`,
		},
	}

	RunStateTransformationTests(t, tests, transformLoadBalancerPoolState)
}
