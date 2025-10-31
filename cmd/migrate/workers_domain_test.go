package main

import (
	"testing"
	
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestWorkersDomainResourceRename(t *testing.T) {
	tests := []TestCase{
		{
			Name: "rename cloudflare_worker_domain to cloudflare_workers_custom_domain",
			Config: `resource "cloudflare_worker_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "subdomain.example.com"
  service    = "my-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`,
			Expected: []string{`resource "cloudflare_workers_custom_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "subdomain.example.com"
  service    = "my-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`},
		},
		{
			Name: "already renamed resource should not change",
			Config: `resource "cloudflare_workers_custom_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "app.example.com"
  service    = "worker-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`,
			Expected: []string{`resource "cloudflare_workers_custom_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "app.example.com"
  service    = "worker-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`},
		},
		{
			Name: "multiple worker domain resources",
			Config: `resource "cloudflare_worker_domain" "primary" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "primary.example.com"
  service    = "primary-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}

resource "cloudflare_worker_domain" "secondary" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "secondary.example.com"
  service    = "secondary-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`,
			Expected: []string{`resource "cloudflare_workers_custom_domain" "primary" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "primary.example.com"
  service    = "primary-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}

resource "cloudflare_workers_custom_domain" "secondary" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "secondary.example.com"
  service    = "secondary-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`},
		},
		{
			Name: "worker domain with environment attribute",
			Config: `resource "cloudflare_worker_domain" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  hostname    = "subdomain.example.com"
  service     = "my-service"
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  environment = "production"
}`,
			Expected: []string{`resource "cloudflare_workers_custom_domain" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  hostname    = "subdomain.example.com"
  service     = "my-service"
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  environment = "production"
}`},
		},
		{
			Name: "mixed old and new resource types",
			Config: `resource "cloudflare_worker_domain" "old_style" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "old.example.com"
  service    = "old-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}

resource "cloudflare_workers_custom_domain" "new_style" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "new.example.com"
  service    = "new-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`,
			Expected: []string{`resource "cloudflare_workers_custom_domain" "old_style" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "old.example.com"
  service    = "old-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}

resource "cloudflare_workers_custom_domain" "new_style" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "new.example.com"
  service    = "new-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestWorkersDomainStateTransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		path     string
	}{
		{
			name:     "basic state transformation",
			input:    `{"version": 4, "terraform_version": "1.0.0", "resources": []}`,
			expected: `{"version": 4, "terraform_version": "1.0.0", "resources": []}`,
			path:     "resources",
		},
		{
			name:     "empty state",
			input:    `{}`,
			expected: `{}`,
			path:     "resources",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformWorkersDomainStateJSON(tt.input, tt.path)
			if result != tt.expected {
				t.Errorf("transformWorkersDomainStateJSON() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsWorkersDomainResource(t *testing.T) {
	tests := []struct {
		name     string
		config   string
		expected bool
	}{
		{
			name: "cloudflare_worker_domain resource",
			config: `resource "cloudflare_worker_domain" "example" {
  hostname = "subdomain.example.com"
}`,
			expected: true,
		},
		{
			name: "cloudflare_workers_custom_domain resource",
			config: `resource "cloudflare_workers_custom_domain" "example" {
  hostname = "subdomain.example.com"
}`,
			expected: true,
		},
		{
			name: "non-worker domain resource",
			config: `resource "cloudflare_record" "example" {
  name = "example"
}`,
			expected: false,
		},
		{
			name: "data source should not match",
			config: `data "cloudflare_worker_domain" "example" {
  hostname = "subdomain.example.com"
}`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.config), "test.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse HCL: %s", diags.Error())
			}

			blocks := file.Body().Blocks()
			if len(blocks) != 1 {
				t.Fatalf("Expected 1 block, got %d", len(blocks))
			}

			result := isWorkersDomainResource(blocks[0])
			if result != tt.expected {
				t.Errorf("isWorkersDomainResource() = %v, want %v", result, tt.expected)
			}
		})
	}
}