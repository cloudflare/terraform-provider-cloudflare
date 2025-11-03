package main

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
)

func TestWorkersCronTriggerTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "worker_cron_trigger singular resource rename",
			Config: `resource "cloudflare_worker_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "0 0 * * *"
}`,
			Expected: []string{`resource "cloudflare_workers_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "0 0 * * *"
}`},
		},
		{
			Name: "workers_cron_trigger plural stays same",
			Config: `resource "cloudflare_workers_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "*/5 * * * *"
}`,
			Expected: []string{`resource "cloudflare_workers_cron_trigger" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  cron        = "*/5 * * * *"
}`},
		},
		{
			Name: "multiple cron triggers with different patterns",
			Config: `resource "cloudflare_worker_cron_trigger" "daily" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "daily-worker"
  cron        = "0 0 * * *"
}

resource "cloudflare_worker_cron_trigger" "hourly" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "hourly-worker"
  cron        = "0 * * * *"
}`,
			Expected: []string{`resource "cloudflare_workers_cron_trigger" "daily"`, `resource "cloudflare_workers_cron_trigger" "hourly"`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestIsWorkersCronTriggerResource(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "cloudflare_workers_cron_trigger resource",
			input: `resource "cloudflare_workers_cron_trigger" "test" {
  account_id = "test"
  script_name = "test"
  cron = "* * * * *"
}`,
			expected: true,
		},
		{
			name: "cloudflare_worker_cron_trigger resource",
			input: `resource "cloudflare_worker_cron_trigger" "test" {
  account_id = "test"
  script_name = "test"
  cron = "* * * * *"
}`,
			expected: true,
		},
		{
			name: "non-cron-trigger resource",
			input: `resource "cloudflare_workers_script" "test" {
  account_id = "test"
  name = "test"
}`,
			expected: false,
		},
		{
			name: "data source not resource",
			input: `data "cloudflare_workers_cron_trigger" "test" {
  account_id = "test"
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

			result := isWorkersCronTriggerResource(blocks[0])
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTransformWorkersCronTriggerBlock(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "transforms singular to plural",
			input: `resource "cloudflare_worker_cron_trigger" "test" {
  account_id = "test"
  script_name = "worker"
  cron = "*/10 * * * *"
}`,
			expected: `resource "cloudflare_workers_cron_trigger" "test" {
  account_id  = "test"
  script_name = "worker"
  cron        = "*/10 * * * *"
}`,
		},
		{
			name: "keeps plural unchanged",
			input: `resource "cloudflare_workers_cron_trigger" "test" {
  account_id = "test"
  script_name = "worker"
  cron = "0 0 * * MON"
}`,
			expected: `resource "cloudflare_workers_cron_trigger" "test" {
  account_id  = "test"
  script_name = "worker"
  cron        = "0 0 * * MON"
}`,
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
				if isWorkersCronTriggerResource(block) {
					transformWorkersCronTriggerBlock(block, ds)
				}
			}

			output := string(hclwrite.Format(file.Bytes()))
			assert.Contains(t, output, tt.expected)
		})
	}
}

func TestTransformWorkersCronTriggerStateJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "returns json unchanged",
			input:    `{"version": 4, "terraform_version": "1.0.0"}`,
			expected: `{"version": 4, "terraform_version": "1.0.0"}`,
		},
		{
			name:     "empty json",
			input:    `{}`,
			expected: `{}`,
		},
		{
			name:     "complex state json",
			input:    `{"resources": [{"type": "cloudflare_worker_cron_trigger", "name": "test"}]}`,
			expected: `{"resources": [{"type": "cloudflare_worker_cron_trigger", "name": "test"}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transformWorkersCronTriggerStateJSON(tt.input, "test.tfstate")
			assert.Equal(t, tt.expected, result)
		})
	}
}