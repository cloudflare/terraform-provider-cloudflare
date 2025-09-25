package main

import (
	"strings"
	"testing"
)

func TestModuleStateMigration(t *testing.T) {
	// Test state transformation for module = true
	originalState := `{
  "resources": [{
    "instances": [{
      "attributes": {
        "account_id": "f037e56e89293a057740de681ac9abbe",
        "script_name": "my-worker",
        "content": "export default { fetch(request) { return new Response('Hello World'); } };",
        "module": true
      }
    }]
  }]
}`

	result := transformModuleInState(originalState, "resources.0.instances.0")

	// Verify main_module was created
	if !strings.Contains(result, `"main_module":"worker.js"`) {
		t.Error("Expected main_module attribute with worker.js value for module=true")
	}

	// Verify original module attribute was removed
	if strings.Contains(result, `"module":true`) {
		t.Error("Expected original module attribute to be removed from state")
	}
}

func TestModuleStateMigrationFalse(t *testing.T) {
	// Test state transformation for module = false
	originalState := `{
  "resources": [{
    "instances": [{
      "attributes": {
        "account_id": "f037e56e89293a057740de681ac9abbe",
        "script_name": "my-worker",
        "content": "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });",
        "module": false
      }
    }]
  }]
}`

	result := transformModuleInState(originalState, "resources.0.instances.0")

	// Verify body_part was created
	if !strings.Contains(result, `"body_part":"worker.js"`) {
		t.Error("Expected body_part attribute with worker.js value for module=false")
	}

	// Verify original module attribute was removed
	if strings.Contains(result, `"module":false`) {
		t.Error("Expected original module attribute to be removed from state")
	}
}

func TestModuleStateMigrationNoModule(t *testing.T) {
	// Test state transformation when no module attribute exists
	originalState := `{
  "resources": [{
    "instances": [{
      "attributes": {
        "account_id": "f037e56e89293a057740de681ac9abbe",
        "script_name": "my-worker",
        "content": "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });"
      }
    }]
  }]
}`

	result := transformModuleInState(originalState, "resources.0.instances.0")

	// Should be unchanged
	if strings.Contains(result, `"main_module"`) || strings.Contains(result, `"body_part"`) {
		t.Error("Expected no module-related attributes to be added when module attribute doesn't exist")
	}

	// Should contain original content  
	if !strings.Contains(result, `"script_name": "my-worker"`) {
		t.Error("Expected original attributes to be preserved")
	}
}