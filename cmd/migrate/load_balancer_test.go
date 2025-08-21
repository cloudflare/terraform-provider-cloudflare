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
