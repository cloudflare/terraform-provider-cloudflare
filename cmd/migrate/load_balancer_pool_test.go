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
		{
			Name: "dynamic_origins_block_simple",
			Config: `
locals {
  origin_list = ["192.0.2.1", "192.0.2.2"]
}

resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"
  
  dynamic "origins" {
    for_each = local.origin_list
    content {
      name    = "origin-${origins.key}"
      address = origins.value
      enabled = true
    }
  }
}`,
			Expected: []string{`
locals {
  origin_list = ["192.0.2.1", "192.0.2.2"]
}

resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"

  origins = [for key, value in local.origin_list : {
    address = value
    enabled = true
    name    = "origin-${key}"
  }]
}`},
		},
		{
			Name: "dynamic_origins_with_header_block",
			Config: `
locals {
  origin_configs = [
    {
      name    = "origin1"
      address = "192.0.2.1"
      host    = "example1.com"
    },
    {
      name    = "origin2"
      address = "192.0.2.2"
      host    = "example2.com"
    }
  ]
}

resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"
  
  dynamic "origins" {
    for_each = local.origin_configs
    content {
      name    = origins.value.name
      address = origins.value.address
      enabled = true
      
      header {
        header = "Host"
        values = [origins.value.host]
      }
    }
  }
}`,
			Expected: []string{`
locals {
  origin_configs = [
    {
      name    = "origin1"
      address = "192.0.2.1"
      host    = "example1.com"
    },
    {
      name    = "origin2"
      address = "192.0.2.2"
      host    = "example2.com"
    }
  ]
}

resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"

  origins = [for key, value in local.origin_configs : {
    address = value.address
    enabled = true
    header = {
      header = "Host"
      values = [value.host]
    }
    name = value.name
  }]
}`},
		},
		{
			Name: "dynamic_origins_with_custom_iterator",
			Config: `
locals {
  origins_data = ["192.0.2.1", "192.0.2.2"]
}

resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"
  
  dynamic "origins" {
    for_each = local.origins_data
    iterator = origin
    content {
      name    = "origin-${origin.key}"
      address = origin.value
      enabled = true
    }
  }
}`,
			Expected: []string{`
locals {
  origins_data = ["192.0.2.1", "192.0.2.2"]
}

resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"

  origins = [for key, value in local.origins_data : {
    address = value
    enabled = true
    name    = "origin-${key}"
  }]
}`},
		},
		{
			Name: "static_origins_unchanged",
			Config: `
resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"
  origins = [{
    name    = "origin1"
    address = "192.0.2.1"
    enabled = true
  }]
}`,
			Expected: []string{`
resource "cloudflare_load_balancer_pool" "test" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test-pool"
  origins = [{
    name    = "origin1"
    address = "192.0.2.1"
    enabled = true
  }]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
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

func TestLoadBalancerPoolMapIteration(t *testing.T) {
	tests := []TestCase{
		{
			Name: "dynamic block with map iteration using key reference",
			Config: `
locals {
  origins_map = {
    "backup1" = { address = "1.1.1.1" }
    "backup2" = { address = "2.2.2.2" }
  }
}

resource "cloudflare_load_balancer_pool" "test" {
  account_id = "test"
  name = "test-pool"
  
  dynamic "origins" {
    for_each = local.origins_map
    content {
      address = origins.value.address
      name = "prefix-${origins.key}"
      header {
        header = "Host"
        values = [origins.value.address]
      }
    }
  }
}`,
			Expected: []string{`resource "cloudflare_load_balancer_pool" "test" {
  account_id = "test"
  name       = "test-pool"

  origins = [for key, value in local.origins_map : {
    address = value.address
    header = {
      header = "Host"
      values = [value.address]
    }
    name = "prefix-${key}"
  }]
}`},
		},
		{
			Name: "dynamic block with both key and value references",
			Config: `
resource "cloudflare_load_balancer_pool" "test" {
  account_id = "test"
  name = "test-pool"
  
  dynamic "origins" {
    for_each = local.transceiver_map
    content {
      address = origins.value.host
      name = origins.key
      weight = origins.value.weight
      header {
        header = "Host"
        values = [origins.value.host]
      }
    }
  }
}`,
			Expected: []string{`resource "cloudflare_load_balancer_pool" "test" {
  account_id = "test"
  name       = "test-pool"

  origins = [for key, value in local.transceiver_map : {
    address = value.host
    header = {
      header = "Host"
      values = [value.host]
    }
    name   = key
    weight = value.weight
  }]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}
