package main

import (
	"strings"
	"testing"
)

func TestWorkersSecretMigration(t *testing.T) {
	tests := []TestCase{
		{
			Name: "workers_secret migrated to workers_script secret_text binding",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_workers_secret" "my_secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION: cloudflare_workers_secret has been migrated to secret_text binding in cloudflare_workers_script`,
				`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "secret_text"
    name = "MY_SECRET"
    text = "secret-value"
  }]
}`,
			},
		},
		{
			Name: "workers_secret merged with existing bindings",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "plain_text"
    name = "MY_VAR"
    text = "my-value"
  }]
}

resource "cloudflare_workers_secret" "my_secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION: cloudflare_workers_secret has been migrated to secret_text binding in cloudflare_workers_script`,
				`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "plain_text"
    name = "MY_VAR"
    text = "my-value"
  }, {
    type = "secret_text"
    name = "MY_SECRET"
    text = "secret-value"
  }]
}`,
			},
		},
		{
			Name: "multiple workers_secret for same script",
			Config: `resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_workers_secret" "secret1" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "API_KEY"
  secret_text = "key-value"
}

resource "cloudflare_workers_secret" "secret2" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "DB_PASSWORD"
  secret_text = "password-value"
}`,
			Expected: []string{
				`# MIGRATION: cloudflare_workers_secret has been migrated to secret_text binding in cloudflare_workers_script`,
				`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "secret_text"
    name = "API_KEY"
    text = "key-value"
  }, {
    type = "secret_text"
    name = "DB_PASSWORD"
    text = "password-value"
  }]
}`,
			},
		},
		{
			Name: "workers_secret with v4 worker_script (singular) resource name",
			Config: `resource "cloudflare_worker_script" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-worker"
  content    = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_worker_secret" "my_secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION: cloudflare_workers_secret has been migrated to secret_text binding in cloudflare_workers_script`,
				`resource "cloudflare_workers_script" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
  bindings = [{
    type = "secret_text"
    name = "MY_SECRET"
    text = "secret-value"
  }]
}`,
			},
		},
		{
			Name: "orphaned workers_secret with no matching script (should add warning)",
			Config: `resource "cloudflare_workers_secret" "orphan" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "nonexistent-worker"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION: cloudflare_workers_secret has been migrated to secret_text binding in cloudflare_workers_script`,
				// The orphaned secret should be removed but may not have a corresponding script
			},
		},
	}

	RunTransformationTests(t, tests, transformFileWithoutImports)
}

func TestWorkersSecretStateMigration(t *testing.T) {
	// Test state transformation
	originalState := `{
  "resources": [
    {
      "type": "cloudflare_workers_script",
      "name": "example",
      "instances": [{
        "attributes": {
          "account_id": "f037e56e89293a057740de681ac9abbe",
          "script_name": "my-worker",
          "content": "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
        }
      }]
    },
    {
      "type": "cloudflare_workers_secret", 
      "name": "my_secret",
      "instances": [{
        "attributes": {
          "account_id": "f037e56e89293a057740de681ac9abbe",
          "script_name": "my-worker",
          "name": "MY_SECRET",
          "secret_text": "secret-value"
        }
      }]
    }
  ]
}`

	// Clear any previous state migration data
	workersSecretsForStateMigration = nil

	// Process workers_secret first to collect information
	result := transformWorkersSecretStateJSON(originalState, "resources.1.instances.0")

	// Then perform cross-resource migration
	result = migrateWorkersSecretsInState(result)

	// Verify the result contains the expected transformations
	if !contains(result, `"type":"secret_text"`) {
		t.Error("Expected secret_text binding type in migrated state")
	}

	if !contains(result, `"name":"MY_SECRET"`) {
		t.Error("Expected secret name in migrated state")
	}

	if !contains(result, `"text":"secret-value"`) {
		t.Error("Expected secret text in migrated state")
	}

	if contains(result, `"type":"cloudflare_workers_secret"`) {
		t.Error("Expected workers_secret resource to be removed from state")
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

