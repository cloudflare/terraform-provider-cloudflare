package main

import (
	"testing"
	"strings"
)

func TestPageRuleMinifyRemoval(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "remove minify block from actions",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    minify = {
      html = "on"
      css  = "on"
      js   = "on"
    }
    ssl = "flexible"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    
    ssl = "flexible"
  }
}`,
		},
		{
			name: "remove minify with different formatting",
			input: `resource "cloudflare_page_rule" "example" {
  actions = {
    minify = { html = "on", css = "on", js = "on" }
    cache_level = "aggressive"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  actions = {
    
    cache_level = "aggressive"
  }
}`,
		},
		{
			name: "handle resource without minify",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    ssl = "flexible"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    ssl = "flexible"
  }
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeMinifyFromActions(tt.input)
			// Check that minify was removed
			if strings.Contains(result, "minify") {
				t.Errorf("removeMinifyFromActions() still contains minify:\n%s", result)
			}
			// Check that other content is preserved
			if !strings.Contains(result, "cache_level") && strings.Contains(tt.input, "cache_level") {
				t.Errorf("removeMinifyFromActions() removed cache_level:\n%s", result)
			}
			if !strings.Contains(result, "ssl") && strings.Contains(tt.input, "ssl") {
				t.Errorf("removeMinifyFromActions() removed ssl:\n%s", result)
			}
		})
	}
}

func TestConsolidateCacheTTLByStatus(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "consolidate multiple cache_ttl_by_status entries",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    cache_ttl_by_status = {
      codes = "200"
      ttl   = 3600
    }
    cache_ttl_by_status = {
      codes = "301"
      ttl   = 1800
    }
    cache_ttl_by_status = {
      codes = "404"
      ttl   = 300
    }
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    cache_ttl_by_status = { "200" = 3600, "301" = 1800, "404" = 300 }
    
    
    
  }
}`,
		},
		{
			name: "single cache_ttl_by_status",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    cache_ttl_by_status = {
      codes = "200"
      ttl   = 3600
    }
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    cache_ttl_by_status = { "200" = 3600 }
    
  }
}`,
		},
		{
			name: "no cache_ttl_by_status",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    ssl = "flexible"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_level = "aggressive"
    ssl = "flexible"
  }
}`,
		},
		{
			name: "cache_ttl_by_status without cache_level",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    cache_ttl_by_status = {
      codes = "200"
      ttl   = 3600
    }
    cache_ttl_by_status = {
      codes = "404"
      ttl   = 300
    }
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    
    
  }
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := consolidateCacheTTLByStatus(tt.input)
			
			// Check if the consolidated map format is present when expected
			if strings.Contains(tt.input, "cache_ttl_by_status") {
				// Check for consolidated map format
				if strings.Contains(tt.expected, `"200" = 3600`) && !strings.Contains(result, `"200" = 3600`) {
					t.Errorf("consolidateCacheTTLByStatus() didn't consolidate properly:\n%s", result)
				}
				if strings.Contains(tt.expected, `"301" = 1800`) && !strings.Contains(result, `"301" = 1800`) {
					t.Errorf("consolidateCacheTTLByStatus() didn't include 301 code:\n%s", result)
				}
				if strings.Contains(tt.expected, `"404"`) && !strings.Contains(result, `"404"`) {
					t.Errorf("consolidateCacheTTLByStatus() didn't include 404 code:\n%s", result)
				}
				
				// Check that old format is removed
				if strings.Contains(result, `codes =`) {
					t.Errorf("consolidateCacheTTLByStatus() still contains old format:\n%s", result)
				}
			} else {
				// No cache_ttl_by_status, should remain unchanged (except whitespace)
				if strings.Contains(result, "cache_ttl_by_status") {
					t.Errorf("consolidateCacheTTLByStatus() added cache_ttl_by_status when not present:\n%s", result)
				}
			}
		})
	}
}

func TestPageRuleCompleteTransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "complete transformation with minify and cache_ttl_by_status",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  priority = 1
  actions = {
    cache_level = "aggressive"
    minify = {
      html = "on"
      css  = "on"
      js   = "on"
    }
    cache_ttl_by_status = {
      codes = "200"
      ttl   = 3600
    }
    cache_ttl_by_status = {
      codes = "301"
      ttl   = 1800
    }
    ssl = "flexible"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  priority = 1
  actions = {
    cache_level = "aggressive"
    cache_ttl_by_status = { "200" = 3600, "301" = 1800 }
    
    
    
    ssl = "flexible"
  }
}`,
		},
		{
			name: "multiple page rules",
			input: `resource "cloudflare_page_rule" "rule1" {
  zone_id = "abc123"
  target  = "example.com/api/*"
  actions = {
    cache_level = "bypass"
    minify = {
      html = "off"
    }
  }
}

resource "cloudflare_page_rule" "rule2" {
  zone_id = "abc123"
  target  = "example.com/static/*"
  actions = {
    cache_level = "aggressive"
    cache_ttl_by_status = {
      codes = "200"
      ttl   = 86400
    }
    cache_ttl_by_status = {
      codes = "404"
      ttl   = 60
    }
  }
}`,
			expected: `resource "cloudflare_page_rule" "rule1" {
  zone_id = "abc123"
  target  = "example.com/api/*"
  actions = {
    cache_level = "bypass"
    
  }
}

resource "cloudflare_page_rule" "rule2" {
  zone_id = "abc123"
  target  = "example.com/static/*"
  actions = {
    cache_level = "aggressive"
    cache_ttl_by_status = { "200" = 86400, "404" = 60 }
    
    
  }
}`,
		},
		{
			name: "page rule with no transformable content",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    ssl = "flexible"
    browser_check = "on"
    email_obfuscation = "on"
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    ssl = "flexible"
    browser_check = "on"
    email_obfuscation = "on"
  }
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformPageRuleConfig(tt.input)
			
			// Check that minify was removed
			if strings.Contains(tt.input, "minify") && strings.Contains(result, "minify") {
				t.Errorf("transformPageRuleConfig() didn't remove minify:\n%s", result)
			}
			
			// Check cache_ttl_by_status consolidation
			if strings.Contains(tt.input, "cache_ttl_by_status") && strings.Contains(tt.input, "codes") {
				// Should be consolidated to map format
				if strings.Contains(result, "codes =") {
					t.Errorf("transformPageRuleConfig() didn't consolidate cache_ttl_by_status:\n%s", result)
				}
				// Check for map format
				if strings.Contains(tt.expected, `"200" = `) && !strings.Contains(result, `"200" = `) {
					t.Errorf("transformPageRuleConfig() didn't create proper map format:\n%s", result)
				}
			}
			
			// Check that other attributes are preserved
			if strings.Contains(tt.input, "ssl") && !strings.Contains(result, "ssl") {
				t.Errorf("transformPageRuleConfig() removed ssl:\n%s", result)
			}
			if strings.Contains(tt.input, "browser_check") && !strings.Contains(result, "browser_check") {
				t.Errorf("transformPageRuleConfig() removed browser_check:\n%s", result)
			}
		})
	}
}

func TestPageRuleEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: "",
		},
		{
			name: "non-page-rule resource",
			input: `resource "cloudflare_record" "example" {
  zone_id = "abc123"
  name    = "example"
  value   = "192.0.2.1"
  type    = "A"
}`,
			expected: `resource "cloudflare_record" "example" {
  zone_id = "abc123"
  name    = "example"
  value   = "192.0.2.1"
  type    = "A"
}`,
		},
		{
			name: "page rule with complex nested structure",
			input: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    forwarding_url = {
      url = "https://www.example.com/$1"
      status_code = 301
    }
    minify = {
      html = "on"
      css  = "on"
      js   = "on"
    }
  }
}`,
			expected: `resource "cloudflare_page_rule" "example" {
  zone_id = "abc123"
  target  = "example.com/*"
  actions = {
    forwarding_url = {
      url = "https://www.example.com/$1"
      status_code = 301
    }
    
  }
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformPageRuleConfig(tt.input)
			
			// For empty input, expect empty output
			if tt.input == "" && result != "" {
				t.Errorf("transformPageRuleConfig() should return empty for empty input, got:\n%s", result)
			}
			
			// For non-page-rule resources, should remain unchanged (except whitespace)
			if !strings.Contains(tt.input, "cloudflare_page_rule") {
				// Should not modify non-page-rule resources
				if strings.Contains(tt.input, "cloudflare_record") && !strings.Contains(result, "cloudflare_record") {
					t.Errorf("transformPageRuleConfig() modified non-page-rule resource:\n%s", result)
				}
			}
			
			// Check minify removal for page rules
			if strings.Contains(tt.input, "cloudflare_page_rule") && strings.Contains(tt.input, "minify") {
				if strings.Contains(result, "minify") {
					t.Errorf("transformPageRuleConfig() didn't remove minify:\n%s", result)
				}
			}
			
			// Check that forwarding_url is preserved
			if strings.Contains(tt.input, "forwarding_url") && !strings.Contains(result, "forwarding_url") {
				t.Errorf("transformPageRuleConfig() removed forwarding_url:\n%s", result)
			}
		})
	}
}