package main

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDynamicListTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "dynamic block with for_each should wrap with toset and use each.value",
			Config: `resource "cloudflare_list" "salesforce_ips" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "salesforce_ips"
  kind       = "ip"

  dynamic "item" {
    for_each = local.salesforce_ips
    content {
      value {
        ip = item.value
      }
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "salesforce_ips" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "salesforce_ips"
  kind       = "ip"
}`,
				`resource "cloudflare_list_item" "salesforce_ips_items" {
  for_each   = toset(local.salesforce_ips)
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.salesforce_ips.id
  ip         = each.value
}`,
			},
		},
		{
			Name: "dynamic block already with toset should not double wrap",
			Config: `resource "cloudflare_list" "my_ips" {
  account_id = "abc123"
  name       = "my_ips"
  kind       = "ip"

  dynamic "item" {
    for_each = toset(var.ip_addresses)
    content {
      value {
        ip = item.value
      }
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "my_ips" {
  account_id = "abc123"
  name       = "my_ips"
  kind       = "ip"
}`,
				`resource "cloudflare_list_item" "my_ips_items" {
  for_each   = toset(var.ip_addresses)
  account_id = "abc123"
  list_id    = cloudflare_list.my_ips.id
  ip         = each.value
}`,
			},
		},
		{
			Name: "dynamic block with item.key should use each.key",
			Config: `resource "cloudflare_list" "asn_list" {
  account_id = "abc123"
  name       = "asn_list"
  kind       = "asn"

  dynamic "item" {
    for_each = var.asn_map
    content {
      comment = item.key
      value {
        asn = item.value
      }
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "asn_list" {
  kind       = "asn"
  account_id = "abc123"
  name       = "asn_list"
}`,
				`resource "cloudflare_list_item" "asn_list_items" {
  for_each   = toset(var.asn_map)
  account_id = "abc123"
  list_id    = cloudflare_list.asn_list.id
  comment    = each.key
  asn        = each.value
}`,
			},
		},
		{
			Name: "mixed static and dynamic items",
			Config: `resource "cloudflare_list" "mixed_list" {
  account_id = "abc123"
  name       = "mixed_list"
  kind       = "ip"

  item {
    value {
      ip = "192.168.1.1"
    }
  }

  dynamic "item" {
    for_each = var.dynamic_ips
    content {
      value {
        ip = item.value
      }
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "mixed_list" {
  name       = "mixed_list"
  kind       = "ip"
  account_id = "abc123"
}`,
				`resource "cloudflare_list_item" "mixed_list_item_0" {
  account_id = "abc123"
  list_id    = cloudflare_list.mixed_list.id
  ip         = "192.168.1.1"
}`,
				`resource "cloudflare_list_item" "mixed_list_items" {
  for_each   = toset(var.dynamic_ips)
  account_id = "abc123"
  list_id    = cloudflare_list.mixed_list.id
  ip         = each.value
}`,
			},
		},
	}

	RunHCLTransformationTests(t, tests, transformFile)
}

func RunHCLTransformationTests(t *testing.T, tests []TestCase, transformFunc func([]byte, string) ([]byte, error)) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Transform the input
			result, err := transformFunc([]byte(tt.Config), "test.tf")
			require.NoError(t, err)

			// Parse the result to get all blocks
			resultFile, diags := hclwrite.ParseConfig(result, "result.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors())

			// For each expected output, verify it exists in the result
			for _, expected := range tt.Expected {
				// Parse expected as HCL
				expectedFile, diags := hclwrite.ParseConfig([]byte(expected), "expected.tf", hcl.InitialPos)
				require.False(t, diags.HasErrors())

				// Get all blocks from expected
				for _, expectedBlock := range expectedFile.Body().Blocks() {
					// Find matching block in result
					found := false
					for _, resultBlock := range resultFile.Body().Blocks() {
						if blocksMatch(expectedBlock, resultBlock) {
							found = true
							break
						}
					}
					assert.True(t, found, "Expected block not found: %s %v", expectedBlock.Type(), expectedBlock.Labels())
				}
			}
		})
	}
}

// blocksMatch compares two HCL blocks semantically
func blocksMatch(a, b *hclwrite.Block) bool {
	// Check type and labels
	if a.Type() != b.Type() {
		return false
	}

	aLabels := a.Labels()
	bLabels := b.Labels()
	if len(aLabels) != len(bLabels) {
		return false
	}
	for i, label := range aLabels {
		if label != bLabels[i] {
			return false
		}
	}

	// Compare attributes
	aAttrs := getAttributeMap(a.Body())
	bAttrs := getAttributeMap(b.Body())

	// Check that all expected attributes exist and match
	for name, aAttr := range aAttrs {
		bAttr, exists := bAttrs[name]
		if !exists {
			return false
		}
		if !attributesMatch(aAttr, bAttr) {
			return false
		}
	}

	// Check that there are no extra attributes in b
	for name := range bAttrs {
		if _, exists := aAttrs[name]; !exists {
			return false
		}
	}

	return true
}

// getAttributeMap returns a map of attribute names to expressions
func getAttributeMap(body *hclwrite.Body) map[string]*hclwrite.Attribute {
	attrs := make(map[string]*hclwrite.Attribute)
	for name, attr := range body.Attributes() {
		attrs[name] = attr
	}
	return attrs
}

// attributesMatch compares two attributes
func attributesMatch(a, b *hclwrite.Attribute) bool {
	// Convert expressions to strings and normalize
	aStr := normalizeExpression(string(a.Expr().BuildTokens(nil).Bytes()))
	bStr := normalizeExpression(string(b.Expr().BuildTokens(nil).Bytes()))
	return aStr == bStr
}

// normalizeExpression normalizes an HCL expression for comparison
func normalizeExpression(s string) string {
	// Remove extra whitespace
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\t", " ")

	// Normalize spaces around operators and punctuation
	replacements := []struct {
		old, new string
	}{
		{" = ", "="},
		{" .", "."},
		{". ", "."},
		{" [", "["},
		{"[ ", "["},
		{" ]", "]"},
		{"] ", "]"},
		{" {", "{"},
		{"{ ", "{"},
		{" }", "}"},
		{"} ", "}"},
		{" ,", ","},
		{", ", ","},
	}

	for _, r := range replacements {
		s = strings.ReplaceAll(s, r.old, r.new)
	}

	return s
}
