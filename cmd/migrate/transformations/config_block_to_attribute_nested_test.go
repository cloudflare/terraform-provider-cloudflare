package transformations

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadBalancerNestedAttributes(t *testing.T) {
	// Create test config
	tempDir, err := os.MkdirTemp("", "variable_interpolation_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "config.yaml")
	configContent := `
version: "1.0"
transformations:
 cloudflare_load_balancer:
  to_map:
   - adaptive_routing
   - location_strategy
   - random_steering
   - fixed_response
   - overrides
   - session_affinity_attributes
  to_list:
   - rules
   - country_pools
   - pop_pools
   - region_pools
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Test input with variable interpolation
	input := `resource "cloudflare_load_balancer" "api-openai" {
		 zone_id = cloudflare_zone.api-openai-com.id
		 name = "api.openai.com"
		 default_pool_ids = [var.api_primary_origins_lb_pool_id]

		 rules {
			condition = "(${local.router_non_sticky_condition})"
			disabled  = false
			name      = "router non-sticky endpoints (geo-aware)"
			overrides {
			  steering_policy = "geo"
			  region_pools {
				region   = "WNAM" # West North America
				pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["us-west"].id]
			  }
			  region_pools {
				region   = "ENAM" # East North America
				pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["us-east"].id]
			  }
			  region_pools {
				region   = "WEU" # Western Europe
				pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["eu"].id]
			  }
			  region_pools {
				region   = "EEU" # Eastern Europe
				pool_ids = [cloudflare_load_balancer_pool.router-geo-pools["eu"].id]
			  }
              country_pools {
				country  = "US"
				pool_ids = [cloudflare_load_balancer_pool.example.id]
			  }
			  pop_pools {
                 pop      = "LAX"
                 pool_ids = [cloudflare_load_balancer_pool.example.id]
              }
			  # Fallback to the default pool if the request doesn't match any of the above regions
			  default_pools    = [var.router_us_unified_origins_lb_pool_id]
			  session_affinity = "none"
			}
			priority   = 0
			terminates = true
		  }
		}`

	// Create transformer
	transformer, err := NewHCLTransformer(configPath)
	if err != nil {
		t.Fatal(err)
	}

	// Transform
	inputPath := filepath.Join(tempDir, "input.tf")
	outputPath := filepath.Join(tempDir, "output.tf")
	if err := os.WriteFile(inputPath, []byte(input), 0644); err != nil {
		t.Fatal(err)
	}

	if err := transformer.TransformFile(inputPath, outputPath); err != nil {
		t.Fatal(err)
	}

	// Check result
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatal(err)
	}
	result := string(output)

	// Check that rules has been converted to a list
	if !strings.Contains(result, "rules = [{") {
		t.Errorf("rules should be converted to a list attribute")
	}

	// Check that overrides has been converted to an attribute
	if !strings.Contains(result, "overrides = {") {
		t.Errorf("overrides should be converted to an attribute map")
	}

	// Check that all 4 region_pools are preserved as duplicate keys
	regionPoolsCount := strings.Count(result, "region_pools =")
	if regionPoolsCount != 4 {
		t.Errorf("Expected 4 'region_pools' but found %d occurrences", regionPoolsCount)
		t.Logf("Output:\n%s", result)
	}

	// Check for specific regions to ensure they're preserved
	if !strings.Contains(result, `"WNAM"`) {
		t.Errorf("Region WNAM was lost in transformation:\n%s", result)
	}
	if !strings.Contains(result, `"ENAM"`) {
		t.Errorf("Region ENAM was lost in transformation:\n%s", result)
	}
	if !strings.Contains(result, `"WEU"`) {
		t.Errorf("Region WEU was lost in transformation:\n%s", result)
	}
	if !strings.Contains(result, `"EEU"`) {
		t.Errorf("Region EEU was lost in transformation:\n%s", result)
	}

	// Check that pool references are preserved
	if !strings.Contains(result, `cloudflare_load_balancer_pool.router-geo-pools["us-west"].id`) {
		t.Errorf("Pool reference for us-west was lost:\n%s", result)
	}
	if !strings.Contains(result, `cloudflare_load_balancer_pool.router-geo-pools["us-east"].id`) {
		t.Errorf("Pool reference for us-east was lost:\n%s", result)
	}
	if !strings.Contains(result, `cloudflare_load_balancer_pool.router-geo-pools["eu"].id`) {
		t.Errorf("Pool reference for eu was lost:\n%s", result)
	}

	// Check that country_pools is preserved
	if !strings.Contains(result, "country_pools =") {
		t.Errorf("country_pools was lost in transformation:\n%s", result)
	}

	// Check that pop_pools is preserved
	if !strings.Contains(result, "pop_pools =") {
		t.Errorf("pop_pools was lost in transformation:\n%s", result)
	}

	// Check that other attributes in overrides are preserved
	if !strings.Contains(result, `steering_policy`) || !strings.Contains(result, `"geo"`) {
		t.Errorf("steering_policy was lost in overrides:\n%s", result)
	}
	if !strings.Contains(result, `session_affinity`) || !strings.Contains(result, `"none"`) {
		t.Errorf("session_affinity was lost in overrides:\n%s", result)
	}
}
