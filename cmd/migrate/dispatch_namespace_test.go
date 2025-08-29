package main

import (
	"strings"
	"testing"
)

func TestDispatchNamespaceStateMigration(t *testing.T) {
	// Test state transformation for dispatch_namespace  
	originalState := `{
  "resources": [{
    "instances": [{
      "attributes": {
        "account_id": "f037e56e89293a057740de681ac9abbe",
        "script_name": "my-worker",
        "content": "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });",
        "dispatch_namespace": "my-dispatch-namespace"
      }
    }]
  }]
}`

	result := transformDispatchNamespaceInState(originalState, "resources.0.instances.0")

	// Verify the dispatch_namespace binding was created
	if !strings.Contains(result, `"type":"dispatch_namespace"`) {
		t.Error("Expected dispatch_namespace binding type in migrated state")
	}

	if !strings.Contains(result, `"namespace_id":"my-dispatch-namespace"`) {
		t.Error("Expected namespace_id with correct value in migrated state")
	}

	// Verify the original dispatch_namespace attribute was removed
	if strings.Contains(result, `"dispatch_namespace":"my-dispatch-namespace"`) {
		t.Error("Expected original dispatch_namespace attribute to be removed from state")
	}
}

func TestDispatchNamespaceStateMigrationWithExistingBindings(t *testing.T) {
	// Test state transformation with existing bindings
	originalState := `{
  "resources": [{
    "instances": [{
      "attributes": {
        "account_id": "f037e56e89293a057740de681ac9abbe",
        "script_name": "my-worker",
        "content": "addEventListener('fetch', event => { event.respondWith(new Response('Hello World')); });",
        "dispatch_namespace": "my-dispatch-namespace",
        "bindings": [
          {
            "type": "plain_text",
            "name": "MY_VAR",
            "text": "my-value"
          }
        ]
      }
    }]
  }]
}`

	result := transformDispatchNamespaceInState(originalState, "resources.0.instances.0")

	// Should have both bindings
	bindingCount := strings.Count(result, `"type":`)
	if bindingCount != 2 {
		t.Errorf("Expected 2 bindings, got %d", bindingCount)
	}

	// Should have the original plain_text binding
	if !strings.Contains(result, `"type":"plain_text"`) {
		t.Error("Expected original plain_text binding to be preserved")
	}

	// Should have the new dispatch_namespace binding
	if !strings.Contains(result, `"type":"dispatch_namespace"`) {
		t.Error("Expected dispatch_namespace binding to be added")
	}
}