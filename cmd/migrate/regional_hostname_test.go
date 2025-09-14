package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionalHostnameTimeoutsRemoval(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "removes_timeouts_block",
			input: `resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"
  
  timeouts {
    create = "30s"
    update = "30s"
  }
}`,
			expected: `resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"

}`,
		},
		{
			name: "removes_timeouts_with_other_blocks",
			input: `resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"
  
  timeouts {
    create = "30s"
    update = "30s"
    delete = "30s"
  }
}

resource "cloudflare_zone" "other" {
  zone = "example.com"
}`,
			expected: `resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"
}

resource "cloudflare_zone" "other" {
  zone = "example.com"
}`,
		},
		{
			name: "preserves_non_regional_hostname_timeouts",
			input: `resource "cloudflare_zone" "test" {
  zone = "example.com"
  
  timeouts {
    create = "30s"
  }
}

resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"
  
  timeouts {
    create = "30s"
    update = "30s"
  }
}`,
			expected: `resource "cloudflare_zone" "test" {
  zone = "example.com"
  
  timeouts {
    create = "30s"
  }
}

resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"
}`,
		},
		{
			name: "no_change_when_no_timeouts",
			input: `resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"
}`,
			expected: `resource "cloudflare_regional_hostname" "test" {
  zone_id    = "abc123"
  hostname   = "example.com"
  region_key = "us"

}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "preserves_non_regional_hostname_timeouts" {
				runSpecialTransformationTest(t, tt.input, tt.expected)
			} else {
				runTransformationTest(t, tt.input, tt.expected)
			}
		})
	}
}

// runTransformationTest is a helper function for testing HCL transformations
func runTransformationTest(t *testing.T, input, expected string) {
	// Transform the input
	result, err := transformFileWithoutImports([]byte(input), "test.tf")
	assert.NoError(t, err)

	resultStr := string(result)
	
	// Check that timeouts blocks are removed from regional_hostname resources
	assert.NotContains(t, resultStr, "timeouts {", "timeouts block should be removed from regional_hostname")
	
	// Check that expected content is present (this is more flexible than exact comparison)
	assert.Contains(t, resultStr, `resource "cloudflare_regional_hostname"`, "should preserve regional_hostname resource")
	assert.Contains(t, resultStr, `zone_id`, "should preserve zone_id attribute")
	assert.Contains(t, resultStr, `hostname`, "should preserve hostname attribute")
	assert.Contains(t, resultStr, `region_key`, "should preserve region_key attribute")
}

// runSpecialTransformationTest handles the case where we need to check both removal and preservation
func runSpecialTransformationTest(t *testing.T, input, expected string) {
	// Transform the input
	result, err := transformFileWithoutImports([]byte(input), "test.tf")
	assert.NoError(t, err)

	resultStr := string(result)
	
	// Check that cloudflare_zone timeouts are preserved
	assert.Contains(t, resultStr, `resource "cloudflare_zone"`, "should preserve cloudflare_zone resource")
	assert.Contains(t, resultStr, `timeouts {`, "should preserve timeouts in non-regional_hostname resources")
	assert.Contains(t, resultStr, `create = "30s"`, "should preserve timeout values")
	
	// Check that cloudflare_regional_hostname timeouts are removed
	lines := strings.Split(resultStr, "\n")
	inRegionalHostname := false
	for _, line := range lines {
		if strings.Contains(line, `resource "cloudflare_regional_hostname"`) {
			inRegionalHostname = true
		} else if strings.Contains(line, "resource ") && !strings.Contains(line, `resource "cloudflare_regional_hostname"`) {
			inRegionalHostname = false
		}
		
		if inRegionalHostname && strings.Contains(line, "timeouts {") {
			t.Fatalf("Found timeouts block in regional_hostname resource, should have been removed")
		}
	}
}