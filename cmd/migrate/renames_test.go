package main

import (
	"testing"
)

func TestCustomPagesTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "simple attributes",
			Config: `
resource "cloudflare_custom_pages" "example" {
  account_id = "987654321fedcba0123456789abcdef0"
  type       = "basic_challenge"
  state      = "active"
  url        = "https://example.com/404"
}`,
			Expected: []string{`
resource "cloudflare_custom_pages" "example" {
  account_id = "987654321fedcba0123456789abcdef0"
  identifier = "basic_challenge"
  state      = "active"
  url        = "https://example.com/404"
}`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}
