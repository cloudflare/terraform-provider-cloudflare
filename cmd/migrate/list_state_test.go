package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCFListStateV4_UnmarshalWithNestedItems(t *testing.T) {
	jsonData := `{
  "version": 4,
  "terraform_version": "1.9.8",
  "serial": 1,
  "lineage": "af188f41-9b09-e5a9-4d9c-3a619f44dcf3",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "terraform_managed_resource",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "description": "This is a note",
            "id": "ecdd727a04e946b9b9301b7664dc4b55",
            "item": [
              {
                "comment": "one",
                "value": [
                  {
                    "asn": 123,
                    "hostname": [],
                    "ip": null,
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "two",
                "value": [
                  {
                    "asn": 456,
                    "hostname": [],
                    "ip": null,
                    "redirect": []
                  }
                ]
              }
            ],
            "kind": "asn",
            "name": "asn_list"
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}`

	// Test unmarshaling
	state, err := ParseCFListStateV4([]byte(jsonData))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the data
	if state.Version != 4 {
		t.Errorf("Expected version 4, got %d", state.Version)
	}

	if state.TerraformVersion != "1.9.8" {
		t.Errorf("Expected terraform_version 1.9.8, got %s", state.TerraformVersion)
	}

	if len(state.Resources) != 1 {
		t.Fatalf("Expected 1 resource, got %d", len(state.Resources))
	}

	resource := state.Resources[0]
	if resource.Type != "cloudflare_list" {
		t.Errorf("Expected resource type cloudflare_list, got %s", resource.Type)
	}

	if resource.Name != "terraform_managed_resource" {
		t.Errorf("Expected resource name terraform_managed_resource, got %s", resource.Name)
	}

	if len(resource.Instances) != 1 {
		t.Fatalf("Expected 1 instance, got %d", len(resource.Instances))
	}

	attrs := resource.Instances[0].Attributes
	if attrs.AccountID != "f037e56e89293a057740de681ac9abbe" {
		t.Errorf("Expected account_id f037e56e89293a057740de681ac9abbe, got %s", attrs.AccountID)
	}

	if attrs.Kind != "asn" {
		t.Errorf("Expected kind asn, got %s", attrs.Kind)
	}

	if attrs.Name != "asn_list" {
		t.Errorf("Expected name asn_list, got %s", attrs.Name)
	}

	// Check the nested items
	if len(attrs.Item) != 2 {
		t.Fatalf("Expected 2 items, got %d", len(attrs.Item))
	}

	// Check first item
	if attrs.Item[0].Comment != "one" {
		t.Errorf("Expected first item comment 'one', got %s", attrs.Item[0].Comment)
	}

	if len(attrs.Item[0].Value) != 1 {
		t.Fatalf("Expected 1 value in first item, got %d", len(attrs.Item[0].Value))
	}

	// Check ASN value - handle both int and float64 (JSON unmarshaling)
	asn1 := attrs.Item[0].Value[0].ASN
	switch v := asn1.(type) {
	case float64:
		if v != 123 {
			t.Errorf("Expected first ASN to be 123, got %v", v)
		}
	case int:
		if v != 123 {
			t.Errorf("Expected first ASN to be 123, got %v", v)
		}
	default:
		t.Errorf("Expected ASN to be numeric, got type %T", asn1)
	}

	// Check second item
	if attrs.Item[1].Comment != "two" {
		t.Errorf("Expected second item comment 'two', got %s", attrs.Item[1].Comment)
	}

	asn2 := attrs.Item[1].Value[0].ASN
	switch v := asn2.(type) {
	case float64:
		if v != 456 {
			t.Errorf("Expected second ASN to be 456, got %v", v)
		}
	case int:
		if v != 456 {
			t.Errorf("Expected second ASN to be 456, got %v", v)
		}
	default:
		t.Errorf("Expected ASN to be numeric, got type %T", asn2)
	}
}

func TestCFListStateV4_ExtractItems(t *testing.T) {
	state := &CFListStateV4{
		Version: 4,
		Resources: []ResourceV4{
			{
				Type: "cloudflare_list",
				Name: "test_list",
				Instances: []InstanceV4{
					{
						Attributes: CloudflareListAttributesV4{
							Kind: "ip",
							Item: []ListItemV4{
								{
									Comment: "First IP",
									Value: []ListItemValueV4{
										{IP: "192.0.2.0"},
									},
								},
								{
									Comment: "Second IP",
									Value: []ListItemValueV4{
										{IP: "192.0.2.1"},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	items := state.ExtractItems("test_list")
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items[0].Comment != "First IP" {
		t.Errorf("Expected first comment 'First IP', got %s", items[0].Comment)
	}
}

func TestCFListStateV4_RemoveItems(t *testing.T) {
	state := &CFListStateV4{
		Resources: []ResourceV4{
			{
				Type: "cloudflare_list",
				Name: "test_list",
				Instances: []InstanceV4{
					{
						Attributes: CloudflareListAttributesV4{
							Item: []ListItemV4{
								{Comment: "test"},
							},
						},
					},
				},
			},
		},
	}

	// Verify items exist before removal
	if state.Resources[0].Instances[0].Attributes.Item == nil {
		t.Error("Expected items to exist before removal")
	}

	state.RemoveItems()

	// Verify items are removed
	if state.Resources[0].Instances[0].Attributes.Item != nil {
		t.Error("Expected items to be nil after removal")
	}
}

func TestCFListStateV4_GetItemValue(t *testing.T) {
	tests := []struct {
		name     string
		item     ListItemV4
		kind     string
		expected interface{}
	}{
		{
			name: "ASN item",
			item: ListItemV4{
				Value: []ListItemValueV4{{ASN: 123}},
			},
			kind:     "asn",
			expected: 123,
		},
		{
			name: "IP item",
			item: ListItemV4{
				Value: []ListItemValueV4{{IP: "192.0.2.0"}},
			},
			kind:     "ip",
			expected: "192.0.2.0",
		},
		{
			name: "Hostname item",
			item: ListItemV4{
				Value: []ListItemValueV4{{
					Hostname: []HostnameValueV4{{URLHostname: "example.com"}},
				}},
			},
			kind:     "hostname",
			expected: HostnameValueV4{URLHostname: "example.com"},
		},
		{
			name: "Redirect item",
			item: ListItemV4{
				Value: []ListItemValueV4{{
					Redirect: []RedirectValueV4{{
						SourceURL: "old.com",
						TargetURL: "new.com",
					}},
				}},
			},
			kind: "redirect",
			expected: RedirectValueV4{
				SourceURL: "old.com",
				TargetURL: "new.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.item.GetItemValue(tt.kind)

			// For complex types, compare as JSON
			expectedJSON, _ := json.Marshal(tt.expected)
			resultJSON, _ := json.Marshal(result)

			if string(expectedJSON) != string(resultJSON) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestCFListStateV4_RoundTrip(t *testing.T) {
	// Create a state with various list types
	state := &CFListStateV4{
		Version:          4,
		TerraformVersion: "1.9.8",
		Serial:           1,
		Lineage:          "test-lineage",
		Outputs:          make(map[string]interface{}),
		Resources: []ResourceV4{
			{
				Mode:     "managed",
				Type:     "cloudflare_list",
				Name:     "test_list",
				Provider: `provider["registry.terraform.io/cloudflare/cloudflare"]`,
				Instances: []InstanceV4{
					{
						SchemaVersion: 0,
						Attributes: CloudflareListAttributesV4{
							AccountID:   "abc123",
							Description: "Test list",
							ID:          "list123",
							Kind:        "asn",
							Name:        "my_list",
							Item: []ListItemV4{
								{
									Comment: "First ASN",
									Value: []ListItemValueV4{
										{
											ASN:      123,
											Hostname: []HostnameValueV4{},
											IP:       nil,
											Redirect: []RedirectValueV4{},
										},
									},
								},
							},
						},
						SensitiveAttributes: []interface{}{},
					},
				},
			},
		},
		CheckResults: nil,
	}

	// Marshal to JSON
	jsonBytes, err := state.ToJSON()
	if err != nil {
		t.Fatalf("Failed to marshal to JSON: %v", err)
	}

	// Unmarshal back
	var unmarshaled CFListStateV4
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify key fields
	if unmarshaled.Version != state.Version {
		t.Errorf("Version mismatch after round-trip")
	}

	if len(unmarshaled.Resources) != len(state.Resources) {
		t.Errorf("Resource count mismatch after round-trip")
	}

	if len(unmarshaled.Resources[0].Instances[0].Attributes.Item) != 1 {
		t.Errorf("Item count mismatch after round-trip")
	}
}

func TestCFListStateV4_ConvertToV5State(t *testing.T) {
	v4State := &CFListStateV4{
		Version:          4,
		TerraformVersion: "1.9.8",
		Serial:           1,
		Lineage:          "test-lineage",
		Outputs:          make(map[string]interface{}),
		Resources: []ResourceV4{
			{
				Mode:     "managed",
				Type:     "cloudflare_list",
				Name:     "test_list",
				Provider: `provider["registry.terraform.io/cloudflare/cloudflare"]`,
				Instances: []InstanceV4{
					{
						SchemaVersion: 0,
						Attributes: CloudflareListAttributesV4{
							AccountID:   "abc123",
							Description: "Test list",
							ID:          "list123",
							Kind:        "ip",
							Name:        "my_list",
							Item: []ListItemV4{
								{
									Comment: "First IP",
									Value: []ListItemValueV4{
										{IP: "192.0.2.0"},
									},
								},
								{
									Comment: "Second IP",
									Value: []ListItemValueV4{
										{IP: "192.0.2.1"},
									},
								},
							},
						},
						SensitiveAttributes: []interface{}{},
					},
				},
			},
		},
	}

	// Convert to v5
	v5State := v4State.ConvertToV5State()

	// Verify basic structure
	if v5State.Version != v4State.Version {
		t.Errorf("Version should be preserved")
	}

	if len(v5State.Resources) != 1 {
		t.Fatalf("Expected 1 resource in v5 state, got %d", len(v5State.Resources))
	}

	// Get the list attributes from v5
	listAttrs, err := v5State.Resources[0].Instances[0].GetListAttributes()
	if err != nil {
		t.Fatalf("Failed to get list attributes: %v", err)
	}

	// Verify items were not copied
	if listAttrs.Item != nil {
		t.Error("Expected items to be nil in v5 state")
	}

	// Verify other fields were preserved
	if listAttrs.AccountID != "abc123" {
		t.Errorf("AccountID not preserved, got %s", listAttrs.AccountID)
	}

	if listAttrs.Kind != "ip" {
		t.Errorf("Kind not preserved, got %s", listAttrs.Kind)
	}
}

func TestCFListStateV5_Unmarshal(t *testing.T) {
	jsonData := `{
  "version": 4,
  "terraform_version": "1.9.8",
  "serial": 5,
  "lineage": "6ccc4e8b-d09c-3f57-fc59-6eb44bf4079c",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "test-list",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "created_on": "2025-08-18T18:36:43Z",
            "description": "List for demo purposes",
            "id": "649d24ad1dc84fe4ac21ff543db5cab1",
            "kind": "redirect",
            "modified_on": "2025-08-18T18:36:43Z",
            "name": "demo_list2",
            "num_items": 0,
            "num_referencing_filters": 0
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}`

	// Test unmarshaling
	state, err := ParseCFListStateV5([]byte(jsonData))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the data
	if state.Version != 4 {
		t.Errorf("Expected version 4, got %d", state.Version)
	}

	if state.TerraformVersion != "1.9.8" {
		t.Errorf("Expected terraform_version 1.9.8, got %s", state.TerraformVersion)
	}

	if len(state.Resources) != 1 {
		t.Fatalf("Expected 1 resource, got %d", len(state.Resources))
	}

	resource := state.Resources[0]
	if resource.Type != "cloudflare_list" {
		t.Errorf("Expected resource type cloudflare_list, got %s", resource.Type)
	}

	if resource.Name != "test-list" {
		t.Errorf("Expected resource name test-list, got %s", resource.Name)
	}

	if len(resource.Instances) != 1 {
		t.Fatalf("Expected 1 instance, got %d", len(resource.Instances))
	}

	attrs, err := resource.Instances[0].GetListAttributes()
	if err != nil {
		t.Fatalf("Failed to get list attributes: %v", err)
	}

	if attrs.AccountID != "f037e56e89293a057740de681ac9abbe" {
		t.Errorf("Expected account_id f037e56e89293a057740de681ac9abbe, got %s", attrs.AccountID)
	}

	if attrs.Kind != "redirect" {
		t.Errorf("Expected kind redirect, got %s", attrs.Kind)
	}

	if attrs.Name != "demo_list2" {
		t.Errorf("Expected name demo_list2, got %s", attrs.Name)
	}
}

func TestCFListStateV5_Marshal(t *testing.T) {
	// Create a state structure
	state := &CFListStateV5{
		Version:          4,
		TerraformVersion: "1.9.8",
		Serial:           5,
		Lineage:          "6ccc4e8b-d09c-3f57-fc59-6eb44bf4079c",
		Outputs:          make(map[string]interface{}),
		Resources: []ResourceV5{
			{
				Mode:     "managed",
				Type:     "cloudflare_list",
				Name:     "test-list",
				Provider: `provider["registry.terraform.io/cloudflare/cloudflare"]`,
				Instances: []InstanceV5{
					{
						SchemaVersion:       0,
						SensitiveAttributes: []interface{}{},
					},
				},
			},
		},
		CheckResults: nil,
	}

	// Set the list attributes
	listAttrs := &CloudflareListAttributes{
		AccountID:             "f037e56e89293a057740de681ac9abbe",
		CreatedOn:             "2025-08-18T18:36:43Z",
		Description:           "List for demo purposes",
		ID:                    "649d24ad1dc84fe4ac21ff543db5cab1",
		Kind:                  "redirect",
		ModifiedOn:            "2025-08-18T18:36:43Z",
		Name:                  "demo_list2",
		NumItems:              0,
		NumReferencingFilters: 0,
	}
	if err := state.Resources[0].Instances[0].SetListAttributes(listAttrs); err != nil {
		t.Fatalf("Failed to set list attributes: %v", err)
	}

	// Marshal to JSON
	jsonBytes, err := state.ToJSON()
	if err != nil {
		t.Fatalf("Failed to marshal to JSON: %v", err)
	}

	// Unmarshal back to verify round-trip
	var unmarshaled CFListStateV5
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify key fields
	if unmarshaled.Version != state.Version {
		t.Errorf("Version mismatch after round-trip")
	}

	if len(unmarshaled.Resources) != len(state.Resources) {
		t.Errorf("Resource count mismatch after round-trip")
	}
}

func TestCFListStateV5_WithItems(t *testing.T) {
	jsonWithItems := `{
  "version": 4,
  "terraform_version": "1.9.8",
  "serial": 5,
  "lineage": "test-lineage",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "ip-list",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "created_on": "2025-08-18T18:36:43Z",
            "description": "IP list with items",
            "id": "list123",
            "kind": "ip",
            "modified_on": "2025-08-18T18:36:43Z",
            "name": "my_ip_list",
            "num_items": 2,
            "num_referencing_filters": 0,
            "item": [
              {
                "value": {
                  "ip": "192.0.2.0"
                },
                "comment": "First IP"
              },
              {
                "value": {
                  "ip": "192.0.2.1"
                },
                "comment": "Second IP"
              }
            ]
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}`

	state, err := ParseCFListStateV5([]byte(jsonWithItems))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON with items: %v", err)
	}

	resource := state.Resources[0]
	attrs, err := resource.Instances[0].GetListAttributes()
	if err != nil {
		t.Fatalf("Failed to get list attributes: %v", err)
	}

	if len(attrs.Item) != 2 {
		t.Errorf("Expected 2 items, got %d", len(attrs.Item))
	}

	if attrs.Item[0].Value.IP != "192.0.2.0" {
		t.Errorf("Expected first IP to be 192.0.2.0, got %s", attrs.Item[0].Value.IP)
	}

	if attrs.Item[0].Comment != "First IP" {
		t.Errorf("Expected first comment to be 'First IP', got %s", attrs.Item[0].Comment)
	}

	// Test RemoveListItems method
	if err := state.RemoveListItems(); err != nil {
		t.Fatalf("Failed to remove list items: %v", err)
	}

	// Get attributes again to check if items were removed
	attrsAfter, _ := state.Resources[0].Instances[0].GetListAttributes()
	if attrsAfter.Item != nil {
		t.Error("Expected Item field to be nil after RemoveListItems")
	}
}

func TestCFListStateV5_HelperMethods(t *testing.T) {
	state := &CFListStateV5{
		Version: 4,
		Resources: []ResourceV5{
			{
				Type: "cloudflare_list",
				Name: "list1",
			},
			{
				Type: "cloudflare_list",
				Name: "list2",
			},
			{
				Type: "cloudflare_zone",
				Name: "zone1",
			},
		},
	}

	// Test FindResource
	found := state.FindResource("cloudflare_list", "list1")
	if len(found) == 0 {
		t.Error("Expected to find cloudflare_list.list1")
	}
	if len(found) > 0 && found[0].Name != "list1" {
		t.Errorf("Expected resource name list1, got %s", found[0].Name)
	}

	notFound := state.FindResource("cloudflare_list", "list3")
	if len(notFound) != 0 {
		t.Error("Expected not to find cloudflare_list.list3")
	}

	// Test GetCloudflareListResources
	lists := state.GetCloudflareListResources()
	if len(lists) != 2 {
		t.Errorf("Expected 2 cloudflare_list resources, got %d", len(lists))
	}
}

func TestCFListState_FindResourceMultipleMatches(t *testing.T) {
	// Test case where multiple resources have the same type and name
	// This can happen in different modules or with count/for_each

	t.Run("V4 Multiple Matches", func(t *testing.T) {
		stateV4 := &CFListStateV4{
			Resources: []ResourceV4{
				{
					Type:     "cloudflare_list",
					Name:     "duplicate_name",
					Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"].us",
					Instances: []InstanceV4{
						{
							Attributes: CloudflareListAttributesV4{
								AccountID: "account1",
								Name:      "list_us",
							},
						},
					},
				},
				{
					Type:     "cloudflare_list",
					Name:     "duplicate_name",
					Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"].eu",
					Instances: []InstanceV4{
						{
							Attributes: CloudflareListAttributesV4{
								AccountID: "account2",
								Name:      "list_eu",
							},
						},
					},
				},
				{
					Type: "cloudflare_list",
					Name: "unique_name",
				},
			},
		}

		// Test FindResource with duplicates
		duplicates := stateV4.FindResource("cloudflare_list", "duplicate_name")
		if len(duplicates) != 2 {
			t.Errorf("Expected to find 2 resources with name duplicate_name, got %d", len(duplicates))
		}

		// Verify both resources are returned
		if duplicates[0].Instances[0].Attributes.AccountID != "account1" {
			t.Error("Expected first duplicate to have account1")
		}
		if duplicates[1].Instances[0].Attributes.AccountID != "account2" {
			t.Error("Expected second duplicate to have account2")
		}

		// Test FindResource with unique name
		unique := stateV4.FindResource("cloudflare_list", "unique_name")
		if len(unique) != 1 {
			t.Errorf("Expected to find 1 resource with name unique_name, got %d", len(unique))
		}

		// Test GetCloudflareListResources
		allLists := stateV4.GetCloudflareListResources()
		if len(allLists) != 3 {
			t.Errorf("Expected 3 cloudflare_list resources total, got %d", len(allLists))
		}
	})

	t.Run("V5 Multiple Matches", func(t *testing.T) {
		stateV5 := &CFListStateV5{
			Resources: []ResourceV5{
				{
					Type:     "cloudflare_list_item",
					Name:     "item_duplicate",
					Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"].us",
				},
				{
					Type:     "cloudflare_list_item",
					Name:     "item_duplicate",
					Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"].eu",
				},
				{
					Type: "cloudflare_list_item",
					Name: "item_unique",
				},
			},
		}

		// Test FindResource with duplicates
		duplicates := stateV5.FindResource("cloudflare_list_item", "item_duplicate")
		if len(duplicates) != 2 {
			t.Errorf("Expected to find 2 resources with name item_duplicate, got %d", len(duplicates))
		}

		// Test FindResource with unique name
		unique := stateV5.FindResource("cloudflare_list_item", "item_unique")
		if len(unique) != 1 {
			t.Errorf("Expected to find 1 resource with name item_unique, got %d", len(unique))
		}

		// Test empty result
		notFound := stateV5.FindResource("cloudflare_list_item", "nonexistent")
		if len(notFound) != 0 {
			t.Errorf("Expected to find 0 resources with name nonexistent, got %d", len(notFound))
		}
	})
}

func TestCFListStateV5_WithListItem(t *testing.T) {
	jsonWithListItem := `{
  "version": 4,
  "terraform_version": "1.9.8",
  "serial": 2,
  "lineage": "73a6b6e3-fa06-75df-02c0-e815dead6dd9",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "test-list",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "created_on": "2025-08-18T18:50:27Z",
            "description": "List for demo purposes",
            "id": "49b9df90e1c249e8a46e3f8bd24bbcf8",
            "kind": "redirect",
            "modified_on": "2025-08-18T18:50:27Z",
            "name": "demo_list2",
            "num_items": 0,
            "num_referencing_filters": 0
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list_item",
      "name": "item1",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "asn": null,
            "comment": null,
            "created_on": "2025-08-18T18:50:27Z",
            "hostname": null,
            "id": "34570aa7549a4b07ae6d2af5777afddb",
            "ip": null,
            "list_id": "49b9df90e1c249e8a46e3f8bd24bbcf8",
            "modified_on": "2025-08-18T18:50:27Z",
            "operation_id": null,
            "redirect": {
              "include_subdomains": false,
              "preserve_path_suffix": false,
              "preserve_query_string": false,
              "source_url": "example-two-url.com/",
              "status_code": 301,
              "subpath_matching": false,
              "target_url": "https://www.example.com"
            }
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}`

	state, err := ParseCFListStateV5([]byte(jsonWithListItem))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON with list_item: %v", err)
	}

	// Check we have both resources
	if len(state.Resources) != 2 {
		t.Fatalf("Expected 2 resources, got %d", len(state.Resources))
	}

	// Check the list resource
	listResources := state.FindResource("cloudflare_list", "test-list")
	if len(listResources) == 0 {
		t.Fatal("Failed to find cloudflare_list resource")
	}
	listResource := listResources[0]

	listAttrs, err := listResource.Instances[0].GetListAttributes()
	if err != nil {
		t.Fatalf("Failed to get list attributes: %v", err)
	}

	if listAttrs.Kind != "redirect" {
		t.Errorf("Expected list kind to be redirect, got %s", listAttrs.Kind)
	}

	// Check the list_item resource
	itemResources := state.FindResource("cloudflare_list_item", "item1")
	if len(itemResources) == 0 {
		t.Fatal("Failed to find cloudflare_list_item resource")
	}
	itemResource := itemResources[0]

	itemAttrs, err := itemResource.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatalf("Failed to get list item attributes: %v", err)
	}

	if itemAttrs.ListID != "49b9df90e1c249e8a46e3f8bd24bbcf8" {
		t.Errorf("Expected list_id to match list ID, got %s", itemAttrs.ListID)
	}

	if itemAttrs.Redirect == nil {
		t.Fatal("Expected redirect to be non-nil")
	}

	if itemAttrs.Redirect.SourceURL != "example-two-url.com/" {
		t.Errorf("Expected source_url to be example-two-url.com/, got %s", itemAttrs.Redirect.SourceURL)
	}

	// Check boolean values are correctly set
	if itemAttrs.Redirect.IncludeSubdomains != false {
		t.Errorf("Expected include_subdomains to be false, got %v", itemAttrs.Redirect.IncludeSubdomains)
	}

	// Check that ASN is nil for a redirect item
	if itemAttrs.ASN != nil {
		t.Errorf("Expected ASN to be nil for redirect item, got %v", itemAttrs.ASN)
	}

	// Test GetCloudflareListItemResources
	listItems := state.GetCloudflareListItemResources()
	if len(listItems) != 1 {
		t.Errorf("Expected 1 cloudflare_list_item resource, got %d", len(listItems))
	}
}

func TestCFListStateV4_RedirectBooleanToStringMigration(t *testing.T) {
	// Test that redirect boolean properties in v4 are converted to strings in v5
	v4State := &CFListStateV4{
		Resources: []ResourceV4{
			{
				Type:     "cloudflare_list",
				Name:     "redirect_list",
				Provider: "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
				Instances: []InstanceV4{
					{
						Attributes: CloudflareListAttributesV4{
							AccountID: "test_account",
							ID:        "redirect_list_id",
							Kind:      "redirect",
							Name:      "test_redirect_list",
							Item: []ListItemV4{
								{
									Comment: "Test redirect with all true",
									Value: []ListItemValueV4{{
										Redirect: []RedirectValueV4{{
											SourceURL:           "source1.com",
											TargetURL:           "target1.com",
											IncludeSubdomains:   "enabled",
											SubpathMatching:     "enabled",
											PreserveQueryString: "enabled",
											PreservePathSuffix:  "enabled",
											StatusCode:          301,
										}},
									}},
								},
								{
									Comment: "Test redirect with all false",
									Value: []ListItemValueV4{{
										Redirect: []RedirectValueV4{{
											SourceURL:           "source2.com",
											TargetURL:           "target2.com",
											IncludeSubdomains:   "disabled",
											SubpathMatching:     "disabled",
											PreserveQueryString: "disabled",
											PreservePathSuffix:  "disabled",
											StatusCode:          302,
										}},
									}},
								},
							},
						},
					},
				},
			},
		},
	}

	// Convert to v5
	v5State := v4State.ConvertToV5State()

	// Manually migrate items (normally done by MigrateCloudflareListToV5)
	// For this test, we'll directly create the items
	for i, item := range v4State.Resources[0].Instances[0].Attributes.Item {
		itemResource := ResourceV5{
			Mode:     "managed",
			Type:     "cloudflare_list_item",
			Name:     fmt.Sprintf("redirect_list_item_%d", i),
			Provider: v4State.Resources[0].Provider,
			Instances: []InstanceV5{{
				SchemaVersion:       0,
				SensitiveAttributes: []interface{}{},
			}},
		}

		itemAttrs := CloudflareListItemAttributes{
			AccountID: "test_account",
			ListID:    "redirect_list_id",
			ID:        fmt.Sprintf("item_%d", i),
			Comment:   &item.Comment,
		}

		if len(item.Value) > 0 && len(item.Value[0].Redirect) > 0 {
			redirect := item.Value[0].Redirect[0]
			v5Redirect := &RedirectValue{
				SourceURL:  redirect.SourceURL,
				TargetURL:  redirect.TargetURL,
				StatusCode: redirect.StatusCode,
			}

			// Convert string to boolean for v5
			v5Redirect.IncludeSubdomains = (redirect.IncludeSubdomains == "enabled")
			v5Redirect.SubpathMatching = (redirect.SubpathMatching == "enabled")
			v5Redirect.PreserveQueryString = (redirect.PreserveQueryString == "enabled")
			v5Redirect.PreservePathSuffix = (redirect.PreservePathSuffix == "enabled")

			itemAttrs.Redirect = v5Redirect
		}

		attrBytes, _ := json.Marshal(itemAttrs)
		itemResource.Instances[0].Attributes = attrBytes
		v5State.Resources = append(v5State.Resources, itemResource)
	}

	// Verify the first item (all true -> all "enabled")
	item1 := v5State.Resources[1] // First list item
	attrs1, err := item1.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatal(err)
	}

	if attrs1.Redirect.IncludeSubdomains != true {
		t.Errorf("Expected include_subdomains to be true, got %v", attrs1.Redirect.IncludeSubdomains)
	}
	if attrs1.Redirect.SubpathMatching != true {
		t.Errorf("Expected subpath_matching to be true, got %v", attrs1.Redirect.SubpathMatching)
	}
	if attrs1.Redirect.PreserveQueryString != true {
		t.Errorf("Expected preserve_query_string to be true, got %v", attrs1.Redirect.PreserveQueryString)
	}
	if attrs1.Redirect.PreservePathSuffix != true {
		t.Errorf("Expected preserve_path_suffix to be true, got %v", attrs1.Redirect.PreservePathSuffix)
	}

	// Verify the second item (all false -> all "disabled")
	item2 := v5State.Resources[2] // Second list item
	attrs2, err := item2.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatal(err)
	}

	if attrs2.Redirect.IncludeSubdomains != false {
		t.Errorf("Expected include_subdomains to be false, got %v", attrs2.Redirect.IncludeSubdomains)
	}
	if attrs2.Redirect.SubpathMatching != false {
		t.Errorf("Expected subpath_matching to be false, got %v", attrs2.Redirect.SubpathMatching)
	}
	if attrs2.Redirect.PreserveQueryString != false {
		t.Errorf("Expected preserve_query_string to be false, got %v", attrs2.Redirect.PreserveQueryString)
	}
	if attrs2.Redirect.PreservePathSuffix != false {
		t.Errorf("Expected preserve_path_suffix to be false, got %v", attrs2.Redirect.PreservePathSuffix)
	}
}

func TestCFListStateV5_ASNHandling(t *testing.T) {
	// Test JSON with ASN list item
	jsonWithASN := `{
  "version": 4,
  "terraform_version": "1.9.8",
  "serial": 1,
  "lineage": "test-lineage",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "asn_list",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "id": "list123",
            "kind": "asn",
            "name": "my_asn_list"
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list_item",
      "name": "asn_item1",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "asn": 64512,
            "comment": "Test ASN",
            "created_on": "2025-08-18T18:50:27Z",
            "hostname": null,
            "id": "item123",
            "ip": null,
            "list_id": "list123",
            "modified_on": "2025-08-18T18:50:27Z",
            "operation_id": null,
            "redirect": null
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list_item",
      "name": "asn_item2",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "asn": null,
            "comment": null,
            "created_on": "2025-08-18T18:50:27Z",
            "hostname": null,
            "id": "item124",
            "ip": "192.0.2.1",
            "list_id": "list124",
            "modified_on": "2025-08-18T18:50:27Z",
            "operation_id": null,
            "redirect": null
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}`

	state, err := ParseCFListStateV5([]byte(jsonWithASN))
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Check we have 3 resources
	if len(state.Resources) != 3 {
		t.Fatalf("Expected 3 resources, got %d", len(state.Resources))
	}

	// Test ASN item with value
	asnItem1Resources := state.FindResource("cloudflare_list_item", "asn_item1")
	if len(asnItem1Resources) == 0 {
		t.Fatal("Failed to find asn_item1")
	}
	asnItem1 := asnItem1Resources[0]

	itemAttrs1, err := asnItem1.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatalf("Failed to get item attributes: %v", err)
	}

	// Check ASN is properly parsed as *int64
	if itemAttrs1.ASN == nil {
		t.Fatal("Expected ASN to be non-nil")
	}

	if *itemAttrs1.ASN != 64512 {
		t.Errorf("Expected ASN to be 64512, got %d", *itemAttrs1.ASN)
	}

	// Check comment
	if itemAttrs1.Comment == nil || *itemAttrs1.Comment != "Test ASN" {
		t.Error("Expected comment to be 'Test ASN'")
	}

	// Check other fields are null
	if itemAttrs1.IP != nil {
		t.Error("Expected IP to be nil for ASN item")
	}

	if itemAttrs1.Hostname != nil {
		t.Error("Expected Hostname to be nil for ASN item")
	}

	if itemAttrs1.Redirect != nil {
		t.Error("Expected Redirect to be nil for ASN item")
	}

	// Test item with null ASN (IP item)
	asnItem2Resources := state.FindResource("cloudflare_list_item", "asn_item2")
	if len(asnItem2Resources) == 0 {
		t.Fatal("Failed to find asn_item2")
	}
	asnItem2 := asnItem2Resources[0]

	itemAttrs2, err := asnItem2.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatalf("Failed to get item attributes: %v", err)
	}

	// Check ASN is nil for IP item
	if itemAttrs2.ASN != nil {
		t.Errorf("Expected ASN to be nil for IP item, got %v", itemAttrs2.ASN)
	}

	// Check IP is set
	if itemAttrs2.IP == nil {
		t.Fatal("Expected IP to be non-nil")
	}

	if *itemAttrs2.IP != "192.0.2.1" {
		t.Errorf("Expected IP to be 192.0.2.1, got %s", *itemAttrs2.IP)
	}
}

func TestCFListStateV5_ASNMarshalUnmarshal(t *testing.T) {
	// Create attributes with ASN
	attrs := CloudflareListItemAttributes{
		AccountID:  "abc123",
		ID:         "item123",
		ListID:     "list123",
		CreatedOn:  "2025-08-18T18:50:27Z",
		ModifiedOn: "2025-08-18T18:50:27Z",
	}

	// Set ASN as *int64
	asn := int64(12345)
	attrs.ASN = &asn

	// Set comment
	comment := "Test comment"
	attrs.Comment = &comment

	// Marshal to JSON
	jsonBytes, err := json.Marshal(attrs)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal back
	var unmarshaled CloudflareListItemAttributes
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify ASN
	if unmarshaled.ASN == nil {
		t.Fatal("Expected ASN to be non-nil after unmarshal")
	}

	if *unmarshaled.ASN != 12345 {
		t.Errorf("Expected ASN to be 12345, got %d", *unmarshaled.ASN)
	}

	// Test with nil ASN
	attrs2 := CloudflareListItemAttributes{
		AccountID: "abc123",
		ID:        "item124",
		ListID:    "list123",
		ASN:       nil, // explicitly nil
	}

	jsonBytes2, err := json.Marshal(attrs2)
	if err != nil {
		t.Fatalf("Failed to marshal with nil ASN: %v", err)
	}

	var unmarshaled2 CloudflareListItemAttributes
	if err := json.Unmarshal(jsonBytes2, &unmarshaled2); err != nil {
		t.Fatalf("Failed to unmarshal with nil ASN: %v", err)
	}

	if unmarshaled2.ASN != nil {
		t.Errorf("Expected ASN to be nil, got %v", unmarshaled2.ASN)
	}
}

func TestCFListStateV5_LargeASN(t *testing.T) {
	// Test with a large ASN that requires int64
	largeASN := int64(4294967295) // Max uint32, needs int64

	attrs := CloudflareListItemAttributes{
		AccountID: "abc123",
		ID:        "item123",
		ListID:    "list123",
		ASN:       &largeASN,
	}

	// Marshal and unmarshal
	jsonBytes, err := json.Marshal(attrs)
	if err != nil {
		t.Fatalf("Failed to marshal large ASN: %v", err)
	}

	var unmarshaled CloudflareListItemAttributes
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal large ASN: %v", err)
	}

	if unmarshaled.ASN == nil || *unmarshaled.ASN != largeASN {
		t.Errorf("Large ASN not preserved: expected %d, got %v", largeASN, unmarshaled.ASN)
	}
}

func TestCFListStateV5_HostnameHandling(t *testing.T) {
	// Test JSON with hostname list item
	jsonWithHostname := `{
  "version": 4,
  "terraform_version": "1.9.8",
  "serial": 1,
  "lineage": "test-lineage",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "hostname_list",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "id": "list123",
            "kind": "hostname",
            "name": "my_hostname_list"
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list_item",
      "name": "hostname_item1",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "asn": null,
            "comment": "Wildcard domain",
            "created_on": "2025-08-18T18:50:27Z",
            "hostname": {
              "url_hostname": "*.example.com",
              "exclude_exact_hostname": true
            },
            "id": "item123",
            "ip": null,
            "list_id": "list123",
            "modified_on": "2025-08-18T18:50:27Z",
            "operation_id": null,
            "redirect": null
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list_item",
      "name": "hostname_item2",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "asn": null,
            "comment": "Specific subdomain",
            "created_on": "2025-08-18T18:50:27Z",
            "hostname": {
              "url_hostname": "api.example.com",
              "exclude_exact_hostname": false
            },
            "id": "item124",
            "ip": null,
            "list_id": "list123",
            "modified_on": "2025-08-18T18:50:27Z",
            "operation_id": null,
            "redirect": null
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list_item",
      "name": "hostname_item3",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "asn": null,
            "comment": null,
            "created_on": "2025-08-18T18:50:27Z",
            "hostname": null,
            "id": "item125",
            "ip": "192.0.2.1",
            "list_id": "list124",
            "modified_on": "2025-08-18T18:50:27Z",
            "operation_id": null,
            "redirect": null
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}`

	state, err := ParseCFListStateV5([]byte(jsonWithHostname))
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	// Check we have 4 resources
	if len(state.Resources) != 4 {
		t.Fatalf("Expected 4 resources, got %d", len(state.Resources))
	}

	// Test hostname item with wildcard
	hostnameItem1Resources := state.FindResource("cloudflare_list_item", "hostname_item1")
	if len(hostnameItem1Resources) == 0 {
		t.Fatal("Failed to find hostname_item1")
	}
	hostnameItem1 := hostnameItem1Resources[0]

	itemAttrs1, err := hostnameItem1.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatalf("Failed to get item attributes: %v", err)
	}

	// Check hostname is properly parsed
	if itemAttrs1.Hostname == nil {
		t.Fatal("Expected Hostname to be non-nil")
	}

	if itemAttrs1.Hostname.URLHostname != "*.example.com" {
		t.Errorf("Expected URLHostname to be *.example.com, got %s", itemAttrs1.Hostname.URLHostname)
	}

	if !itemAttrs1.Hostname.ExcludeExactHostname {
		t.Error("Expected ExcludeExactHostname to be true for wildcard")
	}

	// Check comment
	if itemAttrs1.Comment == nil || *itemAttrs1.Comment != "Wildcard domain" {
		t.Error("Expected comment to be 'Wildcard domain'")
	}

	// Check other fields are null
	if itemAttrs1.ASN != nil {
		t.Error("Expected ASN to be nil for hostname item")
	}

	if itemAttrs1.IP != nil {
		t.Error("Expected IP to be nil for hostname item")
	}

	if itemAttrs1.Redirect != nil {
		t.Error("Expected Redirect to be nil for hostname item")
	}

	// Test hostname item with specific subdomain
	hostnameItem2Resources := state.FindResource("cloudflare_list_item", "hostname_item2")
	if len(hostnameItem2Resources) == 0 {
		t.Fatal("Failed to find hostname_item2")
	}
	hostnameItem2 := hostnameItem2Resources[0]

	itemAttrs2, err := hostnameItem2.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatalf("Failed to get item attributes: %v", err)
	}

	if itemAttrs2.Hostname == nil {
		t.Fatal("Expected Hostname to be non-nil")
	}

	if itemAttrs2.Hostname.URLHostname != "api.example.com" {
		t.Errorf("Expected URLHostname to be api.example.com, got %s", itemAttrs2.Hostname.URLHostname)
	}

	if itemAttrs2.Hostname.ExcludeExactHostname {
		t.Error("Expected ExcludeExactHostname to be false for specific subdomain")
	}

	// Test item with null hostname (IP item)
	hostnameItem3Resources := state.FindResource("cloudflare_list_item", "hostname_item3")
	if len(hostnameItem3Resources) == 0 {
		t.Fatal("Failed to find hostname_item3")
	}
	hostnameItem3 := hostnameItem3Resources[0]

	itemAttrs3, err := hostnameItem3.Instances[0].GetListItemAttributes()
	if err != nil {
		t.Fatalf("Failed to get item attributes: %v", err)
	}

	// Check hostname is nil for IP item
	if itemAttrs3.Hostname != nil {
		t.Errorf("Expected Hostname to be nil for IP item, got %v", itemAttrs3.Hostname)
	}

	// Check IP is set
	if itemAttrs3.IP == nil {
		t.Fatal("Expected IP to be non-nil")
	}

	if *itemAttrs3.IP != "192.0.2.1" {
		t.Errorf("Expected IP to be 192.0.2.1, got %s", *itemAttrs3.IP)
	}
}

func TestCFListStateV5_HostnameMarshalUnmarshal(t *testing.T) {
	// Create attributes with hostname
	attrs := CloudflareListItemAttributes{
		AccountID:  "abc123",
		ID:         "item123",
		ListID:     "list123",
		CreatedOn:  "2025-08-18T18:50:27Z",
		ModifiedOn: "2025-08-18T18:50:27Z",
	}

	// Set hostname
	attrs.Hostname = &HostnameValue{
		URLHostname:          "*.test.com",
		ExcludeExactHostname: true,
	}

	// Set comment
	comment := "Test hostname"
	attrs.Comment = &comment

	// Marshal to JSON
	jsonBytes, err := json.Marshal(attrs)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	// Unmarshal back
	var unmarshaled CloudflareListItemAttributes
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify hostname
	if unmarshaled.Hostname == nil {
		t.Fatal("Expected Hostname to be non-nil after unmarshal")
	}

	if unmarshaled.Hostname.URLHostname != "*.test.com" {
		t.Errorf("Expected URLHostname to be *.test.com, got %s", unmarshaled.Hostname.URLHostname)
	}

	if !unmarshaled.Hostname.ExcludeExactHostname {
		t.Error("Expected ExcludeExactHostname to be true")
	}

	// Test with nil hostname
	attrs2 := CloudflareListItemAttributes{
		AccountID: "abc123",
		ID:        "item124",
		ListID:    "list123",
		Hostname:  nil, // explicitly nil
	}

	jsonBytes2, err := json.Marshal(attrs2)
	if err != nil {
		t.Fatalf("Failed to marshal with nil hostname: %v", err)
	}

	var unmarshaled2 CloudflareListItemAttributes
	if err := json.Unmarshal(jsonBytes2, &unmarshaled2); err != nil {
		t.Fatalf("Failed to unmarshal with nil hostname: %v", err)
	}

	if unmarshaled2.Hostname != nil {
		t.Errorf("Expected Hostname to be nil, got %v", unmarshaled2.Hostname)
	}
}

func TestCFListStateV5_WildcardHostname(t *testing.T) {
	tests := []struct {
		name                 string
		urlHostname          string
		excludeExactHostname bool
		description          string
	}{
		{
			name:                 "wildcard with subdomain only",
			urlHostname:          "*.example.com",
			excludeExactHostname: true,
			description:          "Blocks only subdomains (default for wildcards)",
		},
		{
			name:                 "wildcard with root and subdomains",
			urlHostname:          "*.example.com",
			excludeExactHostname: false,
			description:          "Blocks both root domain and subdomains",
		},
		{
			name:                 "specific subdomain",
			urlHostname:          "api.example.com",
			excludeExactHostname: false,
			description:          "Blocks only the specific subdomain",
		},
		{
			name:                 "nested wildcard",
			urlHostname:          "*.dev.example.com",
			excludeExactHostname: true,
			description:          "Blocks only subdomains of dev.example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrs := CloudflareListItemAttributes{
				AccountID: "abc123",
				ID:        "item123",
				ListID:    "list123",
				Hostname: &HostnameValue{
					URLHostname:          tt.urlHostname,
					ExcludeExactHostname: tt.excludeExactHostname,
				},
			}

			// Marshal and unmarshal
			jsonBytes, err := json.Marshal(attrs)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			var unmarshaled CloudflareListItemAttributes
			if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			if unmarshaled.Hostname == nil {
				t.Fatal("Expected Hostname to be non-nil")
			}

			if unmarshaled.Hostname.URLHostname != tt.urlHostname {
				t.Errorf("URLHostname not preserved: expected %s, got %s",
					tt.urlHostname, unmarshaled.Hostname.URLHostname)
			}

			if unmarshaled.Hostname.ExcludeExactHostname != tt.excludeExactHostname {
				t.Errorf("ExcludeExactHostname not preserved: expected %v, got %v",
					tt.excludeExactHostname, unmarshaled.Hostname.ExcludeExactHostname)
			}
		})
	}
}

var v4JSON = `{
  "version": 4,
  "terraform_version": "1.9.8",
  "serial": 73,
  "lineage": "68193e3b-1112-702e-a6c1-35f3b8860142",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "cloudfalre_ohttp_gateways",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "description": "IPs Cloudflare's OHTTP gateways that connect to api.fooomind.com",
            "id": "fa7bbf4bee694b6196d0d4fb6b0a0461",
            "item": [
              {
                "comment": "Cloudflare OHTTP gateway IPv6 subnet",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "2606:54c3::/32",
                    "redirect": []
                  }
                ]
              }
            ],
            "kind": "ip",
            "name": "cloudflare_ohttp_gateways"
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "enterprise_asns",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "description": "ASNs in this list will have relaxed bot checking rules for enterprise traffic to fooomind.com",
            "id": "911aebee39324833b7d58b7819ef8218",
            "item": [
              {
                "comment": "Bloomberg",
                "value": [
                  {
                    "asn": 67532,
                    "hostname": [],
                    "ip": null,
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "Sony Network Communications Inc.",
                "value": [
                  {
                    "asn": 9999,
                    "hostname": [],
                    "ip": null,
                    "redirect": []
                  }
                ]
              }
            ],
            "kind": "asn",
            "name": "enterprise_asns"
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "known_closedmind_ips",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "description": "The set of IPs that internal closedmind requests may come from. Used for exemption from firewall and bot detection rules. Do not use for general access control!",
            "id": "5f391766a5b44111b012f1c24255fd14",
            "item": [
              {
                "comment": "comment1",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.0",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "comment2",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.2",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "comment3",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.3",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "comment4",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.4",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "comment5",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.5",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "comment6",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.6",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "comment7",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.7",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": "comment8",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "1.1.1.8",
                    "redirect": []
                  }
                ]
              }
            ],
            "kind": "ip",
            "name": "known_closedmind_ips"
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "persona_webhook_ips",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "description": "Static IPs used by Persona to send webhook notifications",
            "id": "fd83c2d5bbc648e6a940f7e21dbf5cc9",
            "item": [
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.1",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.10",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.11",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.12",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.2",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.3",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.4",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.5",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.6",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.7",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.8",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "127.0.0.9",
                    "redirect": []
                  }
                ]
              }
            ],
            "kind": "ip",
            "name": "persona_webhook_ips"
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "redirect_test_foo",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "description": "ASNs in this list will be foo",
            "id": "bb212459cd4a4cd59ca104b28162dbbb",
            "item": [
              {
                "comment": "one",
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": null,
                    "redirect": [
                      {
                        "include_subdomains": "enabled",
                        "preserve_path_suffix": "disabled",
                        "preserve_query_string": "enabled",
                        "source_url": "example.com/foo",
                        "status_code": 301,
                        "subpath_matching": "enabled",
                        "target_url": "https://foo.example.com"
                      }
                    ]
                  }
                ]
              }
            ],
            "kind": "redirect",
            "name": "redirect_test"
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "salesforce_ips",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "f037e56e89293a057740de681ac9abbe",
            "description": "Static IPs used by Salesforce to send webhook notifications",
            "id": "ec7d88353b054ef08a27f66fb6aac2ca",
            "item": [
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "10.0.0.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "10.1.0.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "10.10.0.0/22",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "10.100.0.0/16",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "10.2.0.0/23",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "10.20.0.0/21",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "172.16.0.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "172.16.1.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "172.16.2.0/23",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "172.17.0.0/22",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "172.18.0.0/21",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "172.20.0.0/14",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "172.31.0.0/20",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.0.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.1.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.10.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.100.0/24",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.2.0/23",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.20.0/22",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.200.0/23",
                    "redirect": []
                  }
                ]
              },
              {
                "comment": null,
                "value": [
                  {
                    "asn": null,
                    "hostname": [],
                    "ip": "192.168.50.0/24",
                    "redirect": []
                  }
                ]
              }
            ],
            "kind": "ip",
            "name": "salesforce_ips"
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}`
