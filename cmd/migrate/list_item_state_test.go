package main

import (
	"strings"
	"testing"
)

func TestMigrateListItemsInState(t *testing.T) {
	input := `{
  "version": 4,
  "terraform_version": "1.5.0",
  "resources": [
    {
      "type": "cloudflare_list",
      "name": "example",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "attributes": {
            "account_id": "test-account",
            "id": "list-123",
            "kind": "ip",
            "name": "test-list"
          }
        }
      ]
    },
    {
      "type": "cloudflare_list_item",
      "name": "item1",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "attributes": {
            "account_id": "test-account",
            "list_id": "list-123",
            "ip": "192.0.2.1",
            "comment": "Test IP 1"
          }
        }
      ]
    },
    {
      "type": "cloudflare_list_item",
      "name": "item2",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "attributes": {
            "account_id": "test-account",
            "list_id": "list-123",
            "ip": "192.0.2.2",
            "comment": "Test IP 2"
          }
        }
      ]
    }
  ]
}`

	result := migrateListItemsInState(input)

	// Verify list_item resources are removed
	if strings.Contains(result, `"type":"cloudflare_list_item"`) {
		t.Error("Expected cloudflare_list_item resources to be removed from state")
	}

	// Verify items are merged into the list
	if !strings.Contains(result, `"items"`) {
		t.Error("Expected items array in cloudflare_list resource")
	}

	// Verify specific item data is present
	if !strings.Contains(result, `"ip":"192.0.2.1"`) {
		t.Error("Expected first IP item in merged state")
	}

	if !strings.Contains(result, `"ip":"192.0.2.2"`) {
		t.Error("Expected second IP item in merged state")
	}

	if !strings.Contains(result, `"comment":"Test IP 1"`) {
		t.Error("Expected first comment in merged state")
	}

	// Verify num_items is set correctly
	if !strings.Contains(result, `"num_items":2`) {
		t.Error("Expected num_items to be set to 2")
	}
}

func TestMigrateListItemsInStateWithHostname(t *testing.T) {
	input := `{
  "version": 4,
  "terraform_version": "1.5.0",
  "resources": [
    {
      "type": "cloudflare_list",
      "name": "example",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "attributes": {
            "account_id": "test-account",
            "id": "list-456",
            "kind": "hostname",
            "name": "test-hostname-list"
          }
        }
      ]
    },
    {
      "type": "cloudflare_list_item",
      "name": "hostname_item",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "attributes": {
            "account_id": "test-account",
            "list_id": "list-456",
            "hostname": [
              {
                "url_hostname": "example.com"
              }
            ],
            "comment": "Test hostname"
          }
        }
      ]
    }
  ]
}`

	result := migrateListItemsInState(input)

	// Verify list_item resources are removed
	if strings.Contains(result, `"type":"cloudflare_list_item"`) {
		t.Error("Expected cloudflare_list_item resources to be removed from state")
	}

	// Verify hostname is transformed from array to object
	if !strings.Contains(result, `"hostname":{"url_hostname":"example.com"}`) {
		t.Error("Expected hostname to be transformed from array to object in merged state")
	}
}

func TestMigrateListItemsInStateWithRedirect(t *testing.T) {
	input := `{
  "version": 4,
  "terraform_version": "1.5.0",
  "resources": [
    {
      "type": "cloudflare_list",
      "name": "example",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "attributes": {
            "account_id": "test-account",
            "id": "list-789",
            "kind": "redirect",
            "name": "test-redirect-list"
          }
        }
      ]
    },
    {
      "type": "cloudflare_list_item",
      "name": "redirect_item",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "attributes": {
            "account_id": "test-account",
            "list_id": "list-789",
            "redirect": [
              {
                "source_url": "example.com/old",
                "target_url": "example.com/new",
                "status_code": 301,
                "include_subdomains": "enabled",
                "subpath_matching": "disabled",
                "preserve_query_string": "enabled",
                "preserve_path_suffix": "disabled"
              }
            ]
          }
        }
      ]
    }
  ]
}`

	result := migrateListItemsInState(input)

	// Verify list_item resources are removed
	if strings.Contains(result, `"type":"cloudflare_list_item"`) {
		t.Error("Expected cloudflare_list_item resources to be removed from state")
	}

	// Verify redirect booleans are transformed
	if !strings.Contains(result, `"include_subdomains":true`) {
		t.Error("Expected include_subdomains to be transformed to boolean true")
	}

	if !strings.Contains(result, `"subpath_matching":false`) {
		t.Error("Expected subpath_matching to be transformed to boolean false")
	}

	if !strings.Contains(result, `"preserve_query_string":true`) {
		t.Error("Expected preserve_query_string to be transformed to boolean true")
	}

	if !strings.Contains(result, `"preserve_path_suffix":false`) {
		t.Error("Expected preserve_path_suffix to be transformed to boolean false")
	}
}