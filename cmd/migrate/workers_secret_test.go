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

	RunTransformationTests(t, tests, transformFileDefault)
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

func TestExtractAttributeString(t *testing.T) {
	tests := []TestCase{
		{
			Name: "workers_secret with invalid block structure",
			Config: `resource "cloudflare_workers_secret" {
  # Missing resource name label
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION WARNING: Invalid workers_secret block structure - please migrate manually`,
			},
		},
		{
			Name: "workers_secret with missing script_name attribute",
			Config: `resource "cloudflare_workers_secret" "secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION WARNING: Unable to extract script_name - please migrate manually`,
			},
		},
		{
			Name: "workers_secret with missing name attribute",
			Config: `resource "cloudflare_workers_secret" "secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION WARNING: Unable to extract secret name - please migrate manually`,
			},
		},
		{
			Name: "workers_secret with missing secret_text attribute",
			Config: `resource "cloudflare_workers_secret" "secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "MY_SECRET"
}`,
			Expected: []string{
				`# MIGRATION WARNING: Unable to extract secret_text - please migrate manually`,
			},
		},
		{
			Name: "workers_secret with missing account_id attribute",
			Config: `resource "cloudflare_workers_secret" "secret" {
  script_name = "my-worker"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION WARNING: Unable to extract account_id - please migrate manually`,
			},
		},
		{
			Name: "workers_secret with variable reference in script_name",
			Config: `resource "cloudflare_workers_secret" "secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = var.worker_name
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION WARNING: Unable to extract script_name - please migrate manually`,
			},
		},
		{
			Name: "workers_secret with local reference in secret_text",
			Config: `resource "cloudflare_workers_secret" "secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  name        = "MY_SECRET"
  secret_text = local.my_secret_value
}`,
			Expected: []string{
				`# MIGRATION WARNING: Unable to extract secret_text - please migrate manually`,
			},
		},
		{
			Name: "workers_secret with resource reference in script_name",
			Config: `resource "cloudflare_workers_script" "worker" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "my-worker"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_workers_secret" "secret" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = cloudflare_workers_script.worker.script_name
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`# MIGRATION WARNING: Unable to extract script_name - please migrate manually`,
			},
		},
		{
			Name: "workers_script with no matching secrets (edge case)",
			Config: `resource "cloudflare_workers_script" "worker" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "worker-without-secrets"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}

resource "cloudflare_workers_secret" "secret" {
  account_id  = "different-account"
  script_name = "different-worker"
  name        = "MY_SECRET"
  secret_text = "secret-value"
}`,
			Expected: []string{
				`resource "cloudflare_workers_script" "worker" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  script_name = "worker-without-secrets"
  content     = "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
}`,
			},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}

