package main

import (
	"testing"
)

// State transformation tests
func TestLoadBalancerStateTransformation(t *testing.T) {
	tests := []StateTestCase{
		{
			Name: "renames_fallback_pool_id_to_fallback_pool",
			Input: `{
				"id": "test-id",
				"name": "test-lb",
				"fallback_pool_id": "pool-123"
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-lb",
				"fallback_pool": "pool-123"
			}`,
		},
		{
			Name: "renames_default_pool_ids_to_default_pools",
			Input: `{
				"id": "test-id",
				"name": "test-lb",
				"default_pool_ids": ["pool-1", "pool-2"]
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-lb",
				"default_pools": ["pool-1", "pool-2"]
			}`,
		},
		{
			Name: "removes_empty_single_object_attribute_arrays",
			Input: `{
				"id": "test-id",
				"name": "test-lb",
				"adaptive_routing": [],
				"location_strategy": [],
				"random_steering": [],
				"session_affinity_attributes": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-lb"
			}`,
		},
		{
			Name: "converts_empty_map_attribute_arrays_to_empty_maps",
			Input: `{
				"id": "test-id",
				"name": "test-lb",
				"country_pools": [],
				"pop_pools": [],
				"region_pools": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-lb",
				"country_pools": {},
				"pop_pools": {},
				"region_pools": {}
			}`,
		},
		{
			Name: "keeps_rules_as_array",
			Input: `{
				"id": "test-id",
				"name": "test-lb",
				"rules": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-lb",
				"rules": []
			}`,
		},
		{
			Name: "handles_all_transformations_together",
			Input: `{
				"id": "test-id",
				"name": "test-lb",
				"fallback_pool_id": "pool-123",
				"default_pool_ids": ["pool-1", "pool-2"],
				"adaptive_routing": [],
				"location_strategy": [],
				"random_steering": [],
				"session_affinity_attributes": [],
				"country_pools": [],
				"pop_pools": [],
				"region_pools": [],
				"rules": []
			}`,
			Expected: `{
				"id": "test-id",
				"name": "test-lb",
				"fallback_pool": "pool-123",
				"default_pools": ["pool-1", "pool-2"],
				"country_pools": {},
				"pop_pools": {},
				"region_pools": {},
				"rules": []
			}`,
		},
	}

	RunStateTransformationTests(t, tests, transformLoadBalancerState)
}

// Configuration transformation tests for rules with region_pools
func TestLoadBalancerRulesTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transform region_pools with single region string to list",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  rules = [
    {
      condition = "true"
      disabled  = false
      name      = "geo-aware rule"
      overrides = {
        steering_policy = "geo"
        region_pools = [
          {
            pool_ids = ["pool1"]
            region   = "WNAM"
          },
          {
            pool_ids = ["pool2"]
            region   = "EEU"
          }
        ]
        default_pools = ["default_pool"]
        session_affinity = "none"
      }
      priority   = 0
      terminates = true
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_load_balancer" "test" {
  zone_id          = "test_zone"
  name             = "test-lb"
  default_pool_ids = ["pool1"]

  rules = [
    {
      condition  = "true"
      disabled   = false
      name       = "geo-aware rule"
      overrides  = {
        steering_policy = "geo"
        region_pools = [
          {
            pool_ids = ["pool1"]
            region   = ["WNAM"]
          },
          {
            pool_ids = ["pool2"]
            region   = ["EEU"]
          }
        ]
        default_pools    = ["default_pool"]
        session_affinity = "none"
      }
      priority   = 0
      terminates = true
    }
  ]
}`},
		},
		{
			Name: "handle multiple rules with region_pools",
			Config: `resource "cloudflare_load_balancer" "api-openai" {
  zone_id = cloudflare_zone.api-openai-com.id
  name = "api.openai.com"
  default_pool_ids = [var.api_primary_origins_lb_pool_id]
  
  rules = [
    {
      condition = "(${local.router_non_sticky_condition})"
      disabled  = false
      name      = "router non-sticky endpoints (geo-aware)"
      overrides = {
        steering_policy = "geo"
        region_pools = [
          {
            region   = "WNAM"
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["us-west"].id]
          },
          {
            region   = "ENAM"
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["us-east"].id]
          },
          {
            region   = "WEU"
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["eu"].id]
          },
          {
            region   = "EEU"
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["eu"].id]
          }
        ]
        default_pools    = [var.router_us_unified_origins_lb_pool_id]
        session_affinity = "none"
      }
      priority   = 0
      terminates = true
    },
    {
      condition = "(http.request.uri.path matches \"^/v1/models\")"
      disabled  = false
      name      = "apis"
      overrides = {
        default_pools   = [var.api_unified_origins_lb_pool_id]
        steering_policy = "random"
      }
      priority   = 10
      terminates = false
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_load_balancer" "api-openai" {
  zone_id          = cloudflare_zone.api-openai-com.id
  name             = "api.openai.com"
  default_pool_ids = [var.api_primary_origins_lb_pool_id]

  rules = [
    {
      condition  = "(${local.router_non_sticky_condition})"
      disabled   = false
      name       = "router non-sticky endpoints (geo-aware)"
      overrides  = {
        steering_policy = "geo"
        region_pools = [
          {
            region   = ["WNAM"]
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["us-west"].id]
          },
          {
            region   = ["ENAM"]
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["us-east"].id]
          },
          {
            region   = ["WEU"]
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["eu"].id]
          },
          {
            region   = ["EEU"]
            pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["eu"].id]
          }
        ]
        default_pools    = [var.router_us_unified_origins_lb_pool_id]
        session_affinity = "none"
      }
      priority   = 0
      terminates = true
    },
    {
      condition  = "(http.request.uri.path matches \"^/v1/models\")"
      disabled   = false
      name       = "apis"
      overrides  = {
        default_pools   = [var.api_unified_origins_lb_pool_id]
        steering_policy = "random"
      }
      priority   = 10
      terminates = false
    }
  ]
}`},
		},
		{
			Name: "handle rules without region_pools",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  rules = [
    {
      condition = "(http.request.uri.path matches \"^/api/\")"
      disabled  = false
      name      = "api rule"
      overrides = {
        default_pools   = ["api_pool"]
        steering_policy = "random"
      }
      priority   = 1
      terminates = false
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_load_balancer" "test" {
  zone_id          = "test_zone"
  name             = "test-lb"
  default_pool_ids = ["pool1"]

  rules = [
    {
      condition  = "(http.request.uri.path matches \"^/api/\")"
      disabled   = false
      name       = "api rule"
      overrides  = {
        default_pools   = ["api_pool"]
        steering_policy = "random"
      }
      priority   = 1
      terminates = false
    }
  ]
}`},
		},
		{
			Name: "transform region_pools as single object (not array)",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  rules = [
    {
      condition = "true"
      disabled  = false
      name      = "geo-aware rule"
      overrides = {
        steering_policy = "geo"
        region_pools = {
          pool_ids = ["pool1"]
          region   = "EEU"
        }
        default_pools = ["default_pool"]
        session_affinity = "none"
      }
      priority   = 0
      terminates = true
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_load_balancer" "test" {
  zone_id          = "test_zone"
  name             = "test-lb"
  default_pool_ids = ["pool1"]

  rules = [
    {
      condition  = "true"
      disabled   = false
      name       = "geo-aware rule"
      overrides  = {
        steering_policy = "geo"
        region_pools = {
          pool_ids = ["pool1"]
          region   = ["EEU"]
        }
        default_pools    = ["default_pool"]
        session_affinity = "none"
      }
      priority   = 0
      terminates = true
    }
  ]
}`},
		},
		{
			Name: "handle region_pools that already have region as list",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  rules = [
    {
      condition = "true"
      disabled  = false
      name      = "already migrated"
      overrides = {
        steering_policy = "geo"
        region_pools = [
          {
            pool_ids = ["pool1"]
            region   = ["WNAM", "ENAM"]
          }
        ]
        default_pools = ["default_pool"]
      }
      priority   = 0
      terminates = false
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_load_balancer" "test" {
  zone_id          = "test_zone"
  name             = "test-lb"
  default_pool_ids = ["pool1"]

  rules = [
    {
      condition  = "true"
      disabled   = false
      name       = "already migrated"
      overrides  = {
        steering_policy = "geo"
        region_pools = [
          {
            pool_ids = ["pool1"]
            region   = ["WNAM", "ENAM"]
          }
        ]
        default_pools = ["default_pool"]
      }
      priority   = 0
      terminates = false
    }
  ]
}`},
		},
		{
			Name: "handle empty rules list",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  rules = []
}`,
			Expected: []string{`resource "cloudflare_load_balancer" "test" {
  zone_id          = "test_zone"
  name             = "test-lb"
  default_pool_ids = ["pool1"]
  rules            = []
}`},
		},
		{
			Name: "handle rules without overrides",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  rules = [
    {
      condition = "true"
      disabled  = false
      name      = "simple rule"
      priority   = 0
      terminates = true
    }
  ]
}`,
			Expected: []string{`resource "cloudflare_load_balancer" "test" {
  zone_id          = "test_zone"
  name             = "test-lb"
  default_pool_ids = ["pool1"]

  rules = [
    {
      condition  = "true"
      disabled   = false
      name       = "simple rule"
      priority   = 0
      terminates = true
    }
  ]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileWithoutImports)
}

// Configuration transformation tests for pool blocks to maps
func TestLoadBalancerPoolBlockTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "transform region_pools blocks to map",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  region_pools {
    region = "WNAM"
    pool_ids = ["pool1", "pool2"]
  }
  
  region_pools {
    region = "EEU"
    pool_ids = ["pool3"]
  }
}`,
			Expected: []string{`region_pools = {
    "WNAM" = ["pool1", "pool2"]
    "EEU"  = ["pool3"]
  }`},
		},
		{
			Name: "transform country_pools blocks to map",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  country_pools {
    country = "US"
    pool_ids = ["pool1"]
  }
  
  country_pools {
    country = "GB"
    pool_ids = ["pool2"]
  }
}`,
			Expected: []string{`country_pools = {
    "US" = ["pool1"]
    "GB" = ["pool2"]
  }`},
		},
		{
			Name: "transform pop_pools blocks to map",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  pop_pools {
    pop = "LAX"
    pool_ids = ["pool1"]
  }
  
  pop_pools {
    pop = "ORD"
    pool_ids = ["pool2"]
  }
}`,
			Expected: []string{`pop_pools = {
    "LAX" = ["pool1"]
    "ORD" = ["pool2"]
  }`},
		},
		{
			Name: "transform country_pools from array to map (Grit pre-transformation)",
			Config: `resource "cloudflare_load_balancer" "test" {
  zone_id = "test_zone"
  name = "test-lb"
  default_pool_ids = ["pool1"]
  
  country_pools = [{
    country = "US"
    pool_ids = ["pool1"]
  }, {
    country = "GB"
    pool_ids = ["pool2"]
  }]
}`,
			Expected: []string{`country_pools = {
    US = ["pool1"]
    GB = ["pool2"]
  }`},
		},
	}

	RunTransformationTests(t, tests, transformFileWithoutImports)
}
