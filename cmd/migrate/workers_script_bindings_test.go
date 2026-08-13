package main

import (
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/ast"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestBindingAttributeMappings(t *testing.T) {
	tests := []struct {
		name            string
		v4Config        string
		expectedV5      string
		expectedInV5    []string // strings that should be present in v5 config
		notExpectedInV5 []string // strings that should NOT be present in v5 config
	}{
		{
			name: "hyperdrive_config_binding mappings",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  hyperdrive_config_binding {
    binding = "MY_HYPERDRIVE"
    id = "hyperdrive-123"
  }
}`,
			expectedInV5: []string{
				`type = "hyperdrive"`,
				`name = "MY_HYPERDRIVE"`,
				`id   = "hyperdrive-123"`, // Note: HCL formatter adds extra spaces for alignment
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"hyperdrive_config_binding",
				`binding = "MY_HYPERDRIVE"`, // should be renamed to 'name'
			},
		},
		{
			name: "queue_binding mappings",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  queue_binding {
    binding = "MY_QUEUE"
    queue = "test-queue"
  }
}`,
			expectedInV5: []string{
				`type       = "queue"`, // Note: HCL formatter adds spaces for alignment
				`name       = "MY_QUEUE"`,
				`queue_name = "test-queue"`,
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"queue_binding",
				`binding = "MY_QUEUE"`, // should be renamed to 'name'
				`queue = "test-queue"`, // should be renamed to 'queue_name'
			},
		},
		{
			name: "d1_database_binding mappings",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  d1_database_binding {
    name = "MY_DB"
    database_id = "db-123"
  }
}`,
			expectedInV5: []string{
				`type = "d1"`,
				`name = "MY_DB"`,
				`id   = "db-123"`, // Note: HCL formatter adds spaces for alignment
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"d1_database_binding",
				`database_id = "db-123"`, // should be renamed to 'id'
			},
		},
		{
			name: "analytics_engine_binding (no attribute changes)",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  analytics_engine_binding {
    name = "MY_ANALYTICS"
    dataset = "analytics-dataset"
  }
}`,
			expectedInV5: []string{
				`type    = "analytics_engine"`,
				`name    = "MY_ANALYTICS"`,
				`dataset = "analytics-dataset"`,
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"analytics_engine_binding",
			},
		},
		{
			name: "kv_namespace_binding (no attribute changes)",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  kv_namespace_binding {
    name = "MY_KV"
    namespace_id = "kv-namespace-123"
  }
}`,
			expectedInV5: []string{
				`type         = "kv_namespace"`,
				`name         = "MY_KV"`,
				`namespace_id = "kv-namespace-123"`,
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"kv_namespace_binding",
			},
		},
		{
			name: "plain_text_binding (no attribute changes)",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  plain_text_binding {
    name = "MY_TEXT"
    text = "plain-text-value"
  }
}`,
			expectedInV5: []string{
				`type = "plain_text"`,
				`name = "MY_TEXT"`,
				`text = "plain-text-value"`,
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"plain_text_binding",
			},
		},
		{
			name: "r2_bucket_binding (no attribute changes)",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  r2_bucket_binding {
    name = "MY_BUCKET"
    bucket_name = "test-bucket"
  }
}`,
			expectedInV5: []string{
				`type        = "r2_bucket"`,
				`name        = "MY_BUCKET"`,
				`bucket_name = "test-bucket"`,
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"r2_bucket_binding",
			},
		},
		{
			name: "secret_text_binding (no attribute changes)",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  secret_text_binding {
    name = "MY_SECRET"
    text = "secret-value"
  }
}`,
			expectedInV5: []string{
				`type = "secret_text"`,
				`name = "MY_SECRET"`,
				`text = "secret-value"`,
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"secret_text_binding",
			},
		},
		{
			name: "service_binding (no attribute changes)",
			v4Config: `resource "cloudflare_workers_script" "test" {
  account_id = "test-account"
  name       = "test-script"
  content    = "export default { async fetch() { return new Response('ok'); } };"
  
  service_binding {
    name = "MY_SERVICE"
    service = "my-service"
  }
}`,
			expectedInV5: []string{
				`type    = "service"`,
				`name    = "MY_SERVICE"`,
				`service = "my-service"`,
				`bindings = [`,
			},
			notExpectedInV5: []string{
				"service_binding",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the v4 config
			file, diags := hclwrite.ParseConfig([]byte(tt.v4Config), "test.tf", hcl.InitialPos)
			if diags.HasErrors() {
				t.Fatalf("Failed to parse v4 config: %v", diags)
			}

			// Find the workers_script resource
			var workerBlock *hclwrite.Block
			for _, block := range file.Body().Blocks() {
				if isWorkersScriptResource(block) {
					workerBlock = block
					break
				}
			}
			if workerBlock == nil {
				t.Fatal("Could not find workers_script resource in config")
			}

			// Apply the transformation
			astDiags := ast.Diagnostics{}
			transformWorkersScriptBlock(workerBlock, astDiags)

			// Get the transformed config
			transformedConfig := string(hclwrite.Format(file.Bytes()))

			// Check expected strings are present
			for _, expected := range tt.expectedInV5 {
				if !strings.Contains(transformedConfig, expected) {
					t.Errorf("Expected string not found in transformed config: %q\n\nTransformed config:\n%s", expected, transformedConfig)
				}
			}

			// Check unwanted strings are NOT present
			for _, notExpected := range tt.notExpectedInV5 {
				if strings.Contains(transformedConfig, notExpected) {
					t.Errorf("Unwanted string found in transformed config: %q\n\nTransformed config:\n%s", notExpected, transformedConfig)
				}
			}
		})
	}
}
