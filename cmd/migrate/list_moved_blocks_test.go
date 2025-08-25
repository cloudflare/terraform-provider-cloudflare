package main

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateMovedBlocksFromMigration(t *testing.T) {
	tests := []struct {
		name          string
		beforeState   string   // v4 state with items in list
		afterState    string   // v5 state with separate list_item resources
		expectedMoved []string // Expected moved blocks
	}{
		{
			name: "IP list migration from static to for_each",
			beforeState: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"serial": 1,
				"lineage": "test",
				"outputs": {},
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_list",
						"name": "salesforce_ips",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": [
							{
								"schema_version": 0,
								"attributes": {
									"account_id": "abc123",
									"id": "list123",
									"name": "salesforce_ips",
									"kind": "ip",
									"item": [
										{
											"comment": "SF IP 1",
											"value": [{"ip": "192.168.1.1"}]
										},
										{
											"comment": "SF IP 2",
											"value": [{"ip": "192.168.1.2"}]
										}
									]
								},
								"sensitive_attributes": []
							}
						]
					}
				],
				"check_results": null
			}`,
			afterState: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"serial": 2,
				"lineage": "test",
				"outputs": {},
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_list",
						"name": "salesforce_ips",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": [
							{
								"schema_version": 0,
								"attributes": {
									"account_id": "abc123",
									"id": "list123",
									"name": "salesforce_ips",
									"kind": "ip"
								},
								"sensitive_attributes": []
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_list_item",
						"name": "salesforce_ips_items",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": [
							{
								"schema_version": 0,
								"attributes": {
									"account_id": "abc123",
									"list_id": "list123",
									"ip": "192.168.1.1",
									"comment": "SF IP 1"
								},
								"sensitive_attributes": []
							}
						]
					}
				],
				"check_results": null
			}`,
			expectedMoved: []string{
				`moved {
  from = cloudflare_list_item.salesforce_ips_item_0
  to   = cloudflare_list_item.salesforce_ips_items["192.168.1.1"]
}`,
				`moved {
  from = cloudflare_list_item.salesforce_ips_item_1
  to   = cloudflare_list_item.salesforce_ips_items["192.168.1.2"]
}`,
			},
		},
		{
			name: "ASN list migration",
			beforeState: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"serial": 1,
				"lineage": "test",
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
									"id": "list456",
									"name": "asn_list",
									"kind": "asn",
									"item": [
										{
											"value": [{"asn": 12345}]
										},
										{
											"value": [{"asn": 67890}]
										}
									]
								},
								"sensitive_attributes": []
							}
						]
					}
				],
				"check_results": null
			}`,
			afterState: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"serial": 2,
				"lineage": "test",
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
									"id": "list456",
									"name": "asn_list",
									"kind": "asn"
								},
								"sensitive_attributes": []
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_list_item",
						"name": "asn_list_items",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": [
							{
								"schema_version": 0,
								"attributes": {
									"account_id": "abc123",
									"list_id": "list456",
									"asn": 12345
								},
								"sensitive_attributes": []
							}
						]
					}
				],
				"check_results": null
			}`,
			expectedMoved: []string{
				`moved {
  from = cloudflare_list_item.asn_list_item_0
  to   = cloudflare_list_item.asn_list_items["12345"]
}`,
				`moved {
  from = cloudflare_list_item.asn_list_item_1
  to   = cloudflare_list_item.asn_list_items["67890"]
}`,
			},
		},
		{
			name: "hostname list migration",
			beforeState: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"serial": 1,
				"lineage": "test",
				"outputs": {},
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_list",
						"name": "blocked_hosts",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": [
							{
								"schema_version": 0,
								"attributes": {
									"account_id": "abc123",
									"id": "list789",
									"name": "blocked_hosts",
									"kind": "hostname",
									"item": [
										{
											"value": [{
												"hostname": [{
													"url_hostname": "bad.example.com"
												}]
											}]
										},
										{
											"value": [{
												"hostname": [{
													"url_hostname": "*.evil.com"
												}]
											}]
										}
									]
								},
								"sensitive_attributes": []
							}
						]
					}
				],
				"check_results": null
			}`,
			afterState: `{
				"version": 4,
				"terraform_version": "1.5.0",
				"serial": 2,
				"lineage": "test",
				"outputs": {},
				"resources": [
					{
						"mode": "managed",
						"type": "cloudflare_list",
						"name": "blocked_hosts",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": [
							{
								"schema_version": 0,
								"attributes": {
									"account_id": "abc123",
									"id": "list789",
									"name": "blocked_hosts",
									"kind": "hostname"
								},
								"sensitive_attributes": []
							}
						]
					},
					{
						"mode": "managed",
						"type": "cloudflare_list_item",
						"name": "blocked_hosts_items",
						"provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
						"instances": [
							{
								"schema_version": 0,
								"attributes": {
									"account_id": "abc123",
									"list_id": "list789",
									"hostname": {
										"url_hostname": "bad.example.com"
									}
								},
								"sensitive_attributes": []
							}
						]
					}
				],
				"check_results": null
			}`,
			expectedMoved: []string{
				`moved {
  from = cloudflare_list_item.blocked_hosts_item_0
  to   = cloudflare_list_item.blocked_hosts_items["bad.example.com"]
}`,
				`moved {
  from = cloudflare_list_item.blocked_hosts_item_1
  to   = cloudflare_list_item.blocked_hosts_items["*.evil.com"]
}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate moved blocks
			movedBlocks, err := GenerateMovedBlocksFromMigration(
				[]byte(tt.beforeState),
				[]byte(tt.afterState),
			)
			require.NoError(t, err)

			// Check that we got the expected number of moved blocks
			assert.Equal(t, len(tt.expectedMoved), len(movedBlocks))

			// Convert blocks to strings for comparison
			var actualMoved []string
			for _, block := range movedBlocks {
				// Create a file to format the block properly
				f := hclwrite.NewEmptyFile()
				f.Body().AppendBlock(block)
				blockStr := string(hclwrite.Format(f.Bytes()))
				actualMoved = append(actualMoved, strings.TrimSpace(blockStr))
			}

			// Compare each expected moved block
			for _, expected := range tt.expectedMoved {
				expectedNorm := normalizeHCL(expected)
				found := false
				for _, actual := range actualMoved {
					actualNorm := normalizeHCL(actual)
					if expectedNorm == actualNorm {
						found = true
						break
					}
				}
				assert.True(t, found, "Expected moved block not found:\n%s\n\nActual blocks:\n%s",
					expected, strings.Join(actualMoved, "\n---\n"))
			}
		})
	}
}

// normalizeHCL normalizes HCL for comparison
func normalizeHCL(s string) string {
	// Remove extra whitespace
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	var normalized []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			normalized = append(normalized, line)
		}
	}
	return strings.Join(normalized, " ")
}

func TestBuildIndexedResourceReference(t *testing.T) {
	tests := []struct {
		name         string
		resourceType string
		resourceName string
		indexKey     string
		expected     string
	}{
		{
			name:         "IP address key",
			resourceType: "cloudflare_list_item",
			resourceName: "my_ips_items",
			indexKey:     "192.168.1.1",
			expected:     `cloudflare_list_item.my_ips_items["192.168.1.1"]`,
		},
		{
			name:         "ASN key",
			resourceType: "cloudflare_list_item",
			resourceName: "asn_items",
			indexKey:     "12345",
			expected:     `cloudflare_list_item.asn_items["12345"]`,
		},
		{
			name:         "hostname key",
			resourceType: "cloudflare_list_item",
			resourceName: "blocked_hosts_items",
			indexKey:     "bad.example.com",
			expected:     `cloudflare_list_item.blocked_hosts_items["bad.example.com"]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokens := buildIndexedResourceReference(tt.resourceType, tt.resourceName, tt.indexKey)
			actual := string(tokens.Bytes())
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractListNameFromResourceName(t *testing.T) {
	tests := []struct {
		name         string
		resourceName string
		isStatic     bool
		expected     string
	}{
		{
			name:         "for_each item resource",
			resourceName: "salesforce_ips_items",
			isStatic:     false,
			expected:     "salesforce_ips",
		},
		{
			name:         "static item resource 0",
			resourceName: "salesforce_ips_item_0",
			isStatic:     true,
			expected:     "salesforce_ips",
		},
		{
			name:         "static item resource 10",
			resourceName: "my_list_item_10",
			isStatic:     true,
			expected:     "my_list",
		},
		{
			name:         "complex list name with underscores",
			resourceName: "my_complex_list_name_item_5",
			isStatic:     true,
			expected:     "my_complex_list_name",
		},
		{
			name:         "for_each with complex name",
			resourceName: "my_complex_list_items",
			isStatic:     false,
			expected:     "my_complex_list",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual string
			if tt.isStatic {
				actual = extractListNameFromStaticItem(tt.resourceName)
			} else {
				actual = extractListNameFromItemResource(tt.resourceName)
			}
			assert.Equal(t, tt.expected, actual)
		})
	}
}
