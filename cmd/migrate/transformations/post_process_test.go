package transformations

import (
	"strings"
	"testing"
)

func TestFixDoubleDollarInContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Fix variable interpolation",
			input:    `address = "$${var.origin_address}"`,
			expected: `address = "${var.origin_address}"`,
		},
		{
			name:     "Fix local reference",
			input:    `zone_id = "$${local.zone_id}"`,
			expected: `zone_id = "${local.zone_id}"`,
		},
		{
			name:     "Fix data source reference",
			input:    `value = "$${data.cloudflare_zones.example.zones[0].id}"`,
			expected: `value = "${data.cloudflare_zones.example.zones[0].id}"`,
		},
		{
			name:     "Fix module reference",
			input:    `endpoint = "$${module.api.endpoint}"`,
			expected: `endpoint = "${module.api.endpoint}"`,
		},
		{
			name:     "Fix resource reference",
			input:    `pool_id = "$${cloudflare_load_balancer_pool.main.id}"`,
			expected: `pool_id = "${cloudflare_load_balancer_pool.main.id}"`,
		},
		{
			name:     "Fix count reference",
			input:    `name = "pool-$${count.index}"`,
			expected: `name = "pool-${count.index}"`,
		},
		{
			name:     "Fix each reference",
			input:    `address = "$${each.value.address}"`,
			expected: `address = "${each.value.address}"`,
		},
		{
			name:     "Fix function call",
			input:    `value = "$${base64encode(var.api_key)}"`,
			expected: `value = "${base64encode(var.api_key)}"`,
		},
		{
			name:     "Fix complex interpolation",
			input:    `name = "$${var.prefix}-origin-$${count.index}"`,
			expected: `name = "${var.prefix}-origin-${count.index}"`,
		},
		{
			name:     "Preserve regex backreference",
			input:    `expression = "regex_replace(http.request.uri.path, \"^/backend-api/(.*)$\", \"/chat/backend/api/$${1}\")"`,
			expected: `expression = "regex_replace(http.request.uri.path, \"^/backend-api/(.*)$\", \"/chat/backend/api/$${1}\")"`,
		},
		{
			name: "Fix multiple interpolations",
			input: `resource "cloudflare_load_balancer_pool" "example" {
  name    = "$${var.pool_name}"
  origins {
    address = "$${var.origin_address}"
    name    = "$${var.origin_name}"
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "example" {
  name    = "${var.pool_name}"
  origins {
    address = "${var.origin_address}"
    name    = "${var.origin_name}"
  }
}`,
		},
		{
			name:     "Fix terraform workspace",
			input:    `env = "$${terraform.workspace}"`,
			expected: `env = "${terraform.workspace}"`,
		},
		{
			name:     "Fix path reference",
			input:    `file = "$${path.module}/config.json"`,
			expected: `file = "${path.module}/config.json"`,
		},
		{
			name:     "Fix conditional expression",
			input:    `value = "$${var.use_backup ? var.backup : var.primary}"`,
			expected: `value = "${var.use_backup ? var.backup : var.primary}"`,
		},
		{
			name:     "Fix for expression",
			input:    `pools = "$${[for p in var.pools : p.id]}"`,
			expected: `pools = "${[for p in var.pools : p.id]}"`,
		},
		{
			name:     "Already correct - no change",
			input:    `value = "${var.example}"`,
			expected: `value = "${var.example}"`,
		},
		{
			name:     "Literal dollar sign - no change",
			input:    `price = "$100"`,
			expected: `price = "$100"`,
		},
		{
			name: "Mix of correct and incorrect",
			input: `resource "example" "test" {
  correct = "${var.correct}"
  incorrect = "$${var.incorrect}"
  literal = "Price: $50"
  regex = "pattern/$${1}/replacement"
}`,
			expected: `resource "example" "test" {
  correct = "${var.correct}"
  incorrect = "${var.incorrect}"
  literal = "Price: $50"
  regex = "pattern/$${1}/replacement"
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fixDoubleDollarInContent(tt.input)
			if result != tt.expected {
				// Find the difference
				resultLines := strings.Split(result, "\n")
				expectedLines := strings.Split(tt.expected, "\n")

				t.Errorf("Fix double dollar mismatch")
				t.Logf("Input:\n%s", tt.input)
				t.Logf("Expected:\n%s", tt.expected)
				t.Logf("Got:\n%s", result)

				// Show line-by-line differences
				maxLines := len(expectedLines)
				if len(resultLines) > maxLines {
					maxLines = len(resultLines)
				}
				for i := 0; i < maxLines; i++ {
					var exp, got string
					if i < len(expectedLines) {
						exp = expectedLines[i]
					}
					if i < len(resultLines) {
						got = resultLines[i]
					}
					if exp != got {
						t.Logf("Line %d differs:", i+1)
						t.Logf("  Expected: %s", exp)
						t.Logf("  Got:      %s", got)
					}
				}
			}
		})
	}
}

func TestFixDoubleDollarTargetedResources(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Fix in cloudflare_load_balancer",
			input: `resource "cloudflare_load_balancer" "test" {
  name = "$${var.lb_name}"
  default_pool_ids = ["$${cloudflare_load_balancer_pool.main.id}"]
}`,
			expected: `resource "cloudflare_load_balancer" "test" {
  name = "${var.lb_name}"
  default_pool_ids = ["${cloudflare_load_balancer_pool.main.id}"]
}`,
		},
		{
			name: "Fix in cloudflare_load_balancer_pool",
			input: `resource "cloudflare_load_balancer_pool" "test" {
  name = "$${var.pool_name}"
  origins {
    address = "$${var.origin_address}"
  }
}`,
			expected: `resource "cloudflare_load_balancer_pool" "test" {
  name = "${var.pool_name}"
  origins {
    address = "${var.origin_address}"
  }
}`,
		},
		{
			name: "Fix in cloudflare_ruleset",
			input: `resource "cloudflare_ruleset" "test" {
  zone_id = "$${var.zone_id}"
  name    = "$${var.ruleset_name}"
  rules {
    expression = "$${var.expression}"
  }
}`,
			expected: `resource "cloudflare_ruleset" "test" {
  zone_id = "${var.zone_id}"
  name    = "${var.ruleset_name}"
  rules {
    expression = "${var.expression}"
  }
}`,
		},
		{
			name: "Don't fix in other resources",
			input: `resource "cloudflare_zone" "test" {
  zone = "$${var.zone_name}"
}

resource "cloudflare_record" "test" {
  name = "$${var.record_name}"
}`,
			expected: `resource "cloudflare_zone" "test" {
  zone = "$${var.zone_name}"
}

resource "cloudflare_record" "test" {
  name = "$${var.record_name}"
}`,
		},
		{
			name: "Mixed resources - only fix target ones",
			input: `resource "cloudflare_zone" "test" {
  zone = "$${var.zone_name}"
}

resource "cloudflare_load_balancer_pool" "test" {
  name = "$${var.pool_name}"
  origins {
    address = "$${var.origin_address}"
  }
}

resource "cloudflare_record" "test" {
  name = "$${var.record_name}"
}`,
			expected: `resource "cloudflare_zone" "test" {
  zone = "$${var.zone_name}"
}

resource "cloudflare_load_balancer_pool" "test" {
  name = "${var.pool_name}"
  origins {
    address = "${var.origin_address}"
  }
}

resource "cloudflare_record" "test" {
  name = "$${var.record_name}"
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fixVariableInterpolationInTargetResources(tt.input)
			if result != tt.expected {
				t.Errorf("Fix double dollar mismatch for targeted resources\nExpected:\n%s\n\nGot:\n%s", tt.expected, result)
			}
		})
	}
}

func TestFixDoubleDollarEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Heredoc with interpolation",
			input: `expression = <<-EOF
  not (
    any(http.request.headers["authorization"][*] contains "$${var.token_prefix}")
  )
EOF`,
			expected: `expression = <<-EOF
  not (
    any(http.request.headers["authorization"][*] contains "${var.token_prefix}")
  )
EOF`,
		},
		{
			name:     "JSON encode with interpolation",
			input:    `body = jsonencode({key = "$${var.value}"})`,
			expected: `body = jsonencode({key = "${var.value}"})`,
		},
		{
			name:     "Nested function calls",
			input:    `value = "$${join(",", split(":", var.input))}"`,
			expected: `value = "${join(",", split(":", var.input))}"`,
		},
		{
			name:     "Complex regex pattern",
			input:    `pattern = "$${replace(var.pattern, "/", "\\/")}"`,
			expected: `pattern = "${replace(var.pattern, "/", "\\/")}"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fixDoubleDollarInContent(tt.input)
			if result != tt.expected {
				t.Errorf("Fix double dollar mismatch\nExpected:\n%s\n\nGot:\n%s", tt.expected, result)
			}
		})
	}
}
