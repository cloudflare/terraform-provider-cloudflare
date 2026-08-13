package main

import (
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

func TestWorkersRouteTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "workers_route script_name to script",
			Config: `resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
  script_name = "my-worker"
}`,
			Expected: []string{`resource "cloudflare_workers_route" "example" {
  zone_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern     = "example.com/*"
  script      = "my-worker"
}`},
		},
		{
			Name: "workers_route with no script_name",
			Config: `resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"  
  pattern = "example.com/*"
}`,
			Expected: []string{`resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
}`},
		},
		{
			Name: "worker_route (singular) with script_name",
			Config: `resource "cloudflare_worker_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
  script_name = "my-worker"
}`,
			Expected: []string{`resource "cloudflare_workers_route" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "example.com/*"
  script  = "my-worker"
}`},
		},
		{
			Name: "multiple workers routes with both singular and plural",
			Config: `resource "cloudflare_worker_route" "route1" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "api.example.com/*"
  script_name = "api-worker"
}

resource "cloudflare_workers_route" "route2" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  pattern = "www.example.com/*"
  script_name = "web-worker"
}`,
			Expected: []string{
				`resource "cloudflare_workers_route" "route1"`,
				`script  = "api-worker"`,
				`resource "cloudflare_workers_route" "route2"`,
				`script  = "web-worker"`,
			},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestIsWorkersRouteResource(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "cloudflare_workers_route resource",
			input: `resource "cloudflare_workers_route" "test" {
  zone_id = "test"
  pattern = "test.com/*"
  script_name = "test"
}`,
			expected: true,
		},
		{
			name: "cloudflare_worker_route resource (singular)",
			input: `resource "cloudflare_worker_route" "test" {
  zone_id = "test"
  pattern = "test.com/*"
  script_name = "test"
}`,
			expected: true,
		},
		{
			name: "non-route resource",
			input: `resource "cloudflare_workers_script" "test" {
  account_id = "test"
  name = "test"
}`,
			expected: false,
		},
		{
			name: "data source not resource",
			input: `data "cloudflare_workers_route" "test" {
  zone_id = "test"
}`,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse input: %v", diags.Error())
			}

			blocks := file.Body().Blocks()
			if len(blocks) != 1 {
				t.Fatalf("Expected 1 block, got %d", len(blocks))
			}

			result := isWorkersRouteResource(blocks[0])
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTransformWorkersRouteBlock(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "transforms singular to plural and renames script_name",
			input: `resource "cloudflare_worker_route" "test" {
  zone_id = "test"
  pattern = "test.com/*"
  script_name = "worker"
}`,
			expected: []string{
				`resource "cloudflare_workers_route" "test"`,
				`script  = "worker"`,
			},
		},
		{
			name: "keeps plural and renames script_name",
			input: `resource "cloudflare_workers_route" "test" {
  zone_id = "test"
  pattern = "test.com/*"
  script_name = "worker"
}`,
			expected: []string{
				`resource "cloudflare_workers_route" "test"`,
				`script  = "worker"`,
			},
		},
		{
			name: "handles missing script_name",
			input: `resource "cloudflare_workers_route" "test" {
  zone_id = "test"
  pattern = "test.com/*"
}`,
			expected: []string{
				`resource "cloudflare_workers_route" "test"`,
				`zone_id = "test"`,
				`pattern = "test.com/*"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, diags := hclwrite.ParseConfig([]byte(tt.input), "test.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse input: %v", diags.Error())
			}

			ds := ast.NewDiagnostics()
			for _, block := range file.Body().Blocks() {
				if isWorkersRouteResource(block) {
					transformWorkersRouteBlock(block, ds)
				}
			}

			output := string(hclwrite.Format(file.Bytes()))
			for _, exp := range tt.expected {
				assert.Contains(t, output, exp)
			}
		})
	}
}

func TestTransformWorkersRouteStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]interface{}
		path     string
		expected string
		check    string
	}{
		{
			name: "transforms script_name to script",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_worker_route",
						"name": "test",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"zone_id":     "test-zone",
									"pattern":     "test.com/*",
									"script_name": "my-worker",
								},
							},
						},
					},
				},
			},
			path:  "resources.0.instances.0",
			check: `"script":"my-worker"`,
		},
		{
			name: "handles missing script_name",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_worker_route",
						"name": "test",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{
									"zone_id": "test-zone",
									"pattern": "test.com/*",
								},
							},
						},
					},
				},
			},
			path:  "resources.0.instances.0",
			check: `"zone_id":"test-zone"`,
		},
		{
			name: "handles empty attributes",
			input: map[string]interface{}{
				"resources": []interface{}{
					map[string]interface{}{
						"type": "cloudflare_worker_route",
						"name": "test",
						"instances": []interface{}{
							map[string]interface{}{
								"attributes": map[string]interface{}{},
							},
						},
					},
				},
			},
			path:  "resources.0.instances.0",
			check: `"attributes":{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal input: %v", err)
			}

			result := transformWorkersRouteStateJSON(string(jsonBytes), tt.path)

			// Check that the result contains expected content
			assert.Contains(t, result, tt.check)

			// Ensure script_name is removed if it was present
			if tt.name == "transforms script_name to script" {
				assert.NotContains(t, result, `"script_name"`)
			}
		})
	}
}
