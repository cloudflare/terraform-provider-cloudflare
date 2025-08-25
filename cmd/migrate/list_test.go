package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// assertConfigsEquivalent checks if the transformed config contains all expected resources
func assertConfigsEquivalent(t *testing.T, actual string, expected []string) {
	// Parse the actual output to get structured data
	actualFile, diags := hclwrite.ParseConfig([]byte(actual), "actual.tf", hcl.InitialPos)
	if diags.HasErrors() {
		t.Fatalf("Failed to parse actual output: %v", diags)
	}

	// Build a map of actual resources for easier comparison
	actualResources := extractResources(actualFile.Body())

	// For each expected resource, verify it exists with the right attributes
	for i, exp := range expected {
		expFile, diags := hclwrite.ParseConfig([]byte(exp), "expected.tf", hcl.InitialPos)
		if diags.HasErrors() {
			t.Fatalf("Failed to parse expected output %d: %v", i, diags)
		}

		// Extract expected resources
		expResources := extractResources(expFile.Body())

		// Check each expected resource
		for expKey, expRes := range expResources {
			actRes, found := actualResources[expKey]
			if !found {
				t.Errorf("Expected resource %s not found in output.\n\n=== Expected ===\n%s\n\n=== Actual Resources ===\n%s",
					expKey, exp, formatResources(actualResources))
				continue
			}

			// Compare attributes (order doesn't matter)
			if !compareResourceAttributes(expRes, actRes) {
				t.Errorf("Resource %s has different attributes.\n\n=== Expected ===\n%s\n\n=== Actual ===\n%s",
					expKey, formatResource(expKey, expRes), formatResource(expKey, actRes))
			}
		}
	}
}

// ResourceInfo holds information about a resource
type ResourceInfo struct {
	Type       string
	Name       string
	Attributes map[string]string
}

// extractResources extracts all resources from an HCL body
func extractResources(body *hclwrite.Body) map[string]*ResourceInfo {
	resources := make(map[string]*ResourceInfo)

	for _, block := range body.Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 2 {
			resType := block.Labels()[0]
			resName := block.Labels()[1]
			key := resType + "." + resName

			attrs := make(map[string]string)
			for name, attr := range block.Body().Attributes() {
				// Convert attribute to string representation
				tokens := attr.Expr().BuildTokens(nil)
				attrs[name] = tokensToString(tokens)
			}

			resources[key] = &ResourceInfo{
				Type:       resType,
				Name:       resName,
				Attributes: attrs,
			}
		}
	}

	return resources
}

// tokensToString converts HCL tokens to a normalized string
func tokensToString(tokens hclwrite.Tokens) string {
	var result strings.Builder
	for _, token := range tokens {
		result.Write(token.Bytes)
	}
	return strings.TrimSpace(result.String())
}

/*// transformFile is the main transformation function for tests
func transformFile(content []byte, filename string) ([]byte, error) {
	file, diags := hclwrite.ParseConfig(content, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, diags
	}

	rootBody := file.Body()
	var newBlocks []*hclwrite.Block
	var blocksToRemove []*hclwrite.Block

	for _, block := range rootBody.Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 2 {
			if IsCloudflareListResource(block) {
				hasItems := false
				for _, itemBlock := range block.Body().Blocks() {
					if itemBlock.Type() == "item" || (itemBlock.Type() == "dynamic" && len(itemBlock.Labels()) > 0 && itemBlock.Labels()[0] == "item") {
						hasItems = true
						break
					}
				}

				if hasItems {
					blocksToRemove = append(blocksToRemove, block)
					transformedBlocks := TransformCloudflareListBlock(block)
					newBlocks = append(newBlocks, transformedBlocks...)
				}
			}
		}
	}

	// Remove old blocks and add new ones
	for _, block := range blocksToRemove {
		rootBody.RemoveBlock(block)
	}
	for _, block := range newBlocks {
		rootBody.AppendBlock(block)
	}

	return hclwrite.Format(file.Bytes()), nil
}

// transformStateJSON transforms state JSON for tests
func transformStateJSON(content []byte) ([]byte, error) {
	// Parse JSON to map for transformation
	var stateMap map[string]interface{}
	if err := json.Unmarshal(content, &stateMap); err != nil {
		return nil, err
	}

	// Transform cloudflare_list resources by removing item and items fields
	if resources, ok := stateMap["resources"].([]interface{}); ok {
		for _, resource := range resources {
			if resMap, ok := resource.(map[string]interface{}); ok {
				if resType, ok := resMap["type"].(string); ok && resType == "cloudflare_list" {
					if instances, ok := resMap["instances"].([]interface{}); ok {
						for _, instance := range instances {
							if instMap, ok := instance.(map[string]interface{}); ok {
								if attrs, ok := instMap["attributes"].(map[string]interface{}); ok {
									// Remove item and items fields
									delete(attrs, "item")
									delete(attrs, "items")
								}
							}
						}
					}
				}
			}
		}
	}

	// Marshal back to JSON
	result, err := json.Marshal(stateMap)
	if err != nil {
		return nil, err
	}

	return result, nil
}*/

// compareResourceAttributes compares two resources' attributes
func compareResourceAttributes(expected, actual *ResourceInfo) bool {
	// Check all expected attributes exist in actual
	for expKey, expVal := range expected.Attributes {
		actVal, found := actual.Attributes[expKey]
		if !found {
			return false
		}
		// Normalize and compare values
		if normalizeValue(expVal) != normalizeValue(actVal) {
			return false
		}
	}
	return true
}

// normalizeValue normalizes an attribute value for comparison
func normalizeValue(val string) string {
	// Remove extra whitespace
	val = strings.TrimSpace(val)
	// Normalize quotes
	val = strings.ReplaceAll(val, `"`, `"`)
	return val
}

// formatResource formats a resource for error output
func formatResource(key string, res *ResourceInfo) string {
	var b strings.Builder
	b.WriteString("resource \"")
	b.WriteString(res.Type)
	b.WriteString("\" \"")
	b.WriteString(res.Name)
	b.WriteString("\" {\n")
	for k, v := range res.Attributes {
		b.WriteString("  ")
		b.WriteString(k)
		b.WriteString(" = ")
		b.WriteString(v)
		b.WriteString("\n")
	}
	b.WriteString("}")
	return b.String()
}

// formatResources formats all resources for error output
func formatResources(resources map[string]*ResourceInfo) string {
	var b strings.Builder
	for key, res := range resources {
		b.WriteString(formatResource(key, res))
		b.WriteString("\n\n")
	}
	return b.String()
}

func TestCloudflareListTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "basic list without items",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "abc123"
  name        = "my_list"
  description = "Test list"
  kind        = "ip"
}`,
			Expected: []string{`
resource "cloudflare_list" "example" {
  account_id  = "abc123"
  name        = "my_list"
  description = "Test list"
  kind        = "ip"
}`},
		},
		{
			Name: "IP list with items",
			Config: `
resource "cloudflare_list" "ip_list" {
  account_id  = "abc123"
  name        = "ip_list"
  description = "IP list with items"
  kind        = "ip"
  
  item {
    value {
      ip = "192.0.2.0"
    }
    comment = "First IP"
  }
  
  item {
    value {
      ip = "192.0.2.1"
    }
    comment = "Second IP"
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "ip_list" {
  account_id  = "abc123"
  name        = "ip_list"
  description = "IP list with items"
  kind        = "ip"
}`,
				`resource "cloudflare_list_item" "ip_list_item_0" {
  account_id = "abc123"
  list_id    = cloudflare_list.ip_list.id
  comment    = "First IP"
  ip         = "192.0.2.0"
}`,
				`resource "cloudflare_list_item" "ip_list_item_1" {
  account_id = "abc123"
  list_id    = cloudflare_list.ip_list.id
  comment    = "Second IP"
  ip         = "192.0.2.1"
}`,
			},
		},
		{
			Name: "ASN list with items",
			Config: `
resource "cloudflare_list" "asn_list" {
  account_id  = var.account_id
  name        = "asn_list"
  description = "ASN list"
  kind        = "asn"
  
  item {
    value {
      asn = 64512
    }
    comment = "First ASN"
  }
  
  item {
    value {
      asn = 64513
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "asn_list" {
  account_id  = var.account_id
  name        = "asn_list"
  description = "ASN list"
  kind        = "asn"
}`,
				`resource "cloudflare_list_item" "asn_list_item_0" {
  account_id = var.account_id
  list_id    = cloudflare_list.asn_list.id
  comment    = "First ASN"
  asn        = 64512
}`,
				`resource "cloudflare_list_item" "asn_list_item_1" {
  account_id = var.account_id
  list_id    = cloudflare_list.asn_list.id
  asn        = 64513
}`,
			},
		},
		{
			Name: "hostname list with items",
			Config: `
resource "cloudflare_list" "hostname_list" {
  account_id  = "abc123"
  name        = "hostname_list"
  description = "Hostname list"
  kind        = "hostname"
  
  item {
    value {
      hostname {
        url_hostname = "*.example.com"
      }
    }
    comment = "Wildcard"
  }
  
  item {
    value {
      hostname {
        url_hostname = "test.example.com"
      }
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "hostname_list" {
  account_id  = "abc123"
  name        = "hostname_list"
  description = "Hostname list"
  kind        = "hostname"
}`,
				`resource "cloudflare_list_item" "hostname_list_item_0" {
  account_id = "abc123"
  list_id    = cloudflare_list.hostname_list.id
  comment    = "Wildcard"
  hostname = {
    url_hostname = "*.example.com"
  }
}`,
				`resource "cloudflare_list_item" "hostname_list_item_1" {
  account_id = "abc123"
  list_id    = cloudflare_list.hostname_list.id
  hostname = {
    url_hostname = "test.example.com"
  }
}`,
			},
		},
		{
			Name: "redirect list with items - booleans remain as booleans",
			Config: `
resource "cloudflare_list" "redirect_list" {
  account_id  = "abc123"
  name        = "redirect_list"
  description = "Redirect list"
  kind        = "redirect"
  
  item {
    value {
      redirect {
        source_url            = "example.com/old"
        target_url            = "https://example.com/new"
        include_subdomains    = true
        subpath_matching      = false
        preserve_query_string = true
        preserve_path_suffix  = false
        status_code          = 301
      }
    }
    comment = "Main redirect"
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "redirect_list" {
  account_id  = "abc123"
  name        = "redirect_list"
  description = "Redirect list"
  kind        = "redirect"
}`,
				`resource "cloudflare_list_item" "redirect_list_item_0" {
  account_id = "abc123"
  list_id    = cloudflare_list.redirect_list.id
  comment    = "Main redirect"
  redirect = {
    source_url            = "example.com/old"
    target_url            = "https://example.com/new"
    include_subdomains    = true
    subpath_matching      = false
    preserve_query_string = true
    preserve_path_suffix  = false
    status_code           = 301
  }
}`,
			},
		},
		{
			Name: "list with variable references",
			Config: `
resource "cloudflare_list" "var_list" {
  account_id  = var.account_id
  name        = var.list_name
  description = "List with variables"
  kind        = "ip"
  
  item {
    value {
      ip = var.ip_address
    }
    comment = var.comment
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "var_list" {
  account_id  = var.account_id
  name        = var.list_name
  description = "List with variables"
  kind        = "ip"
}`,
				`resource "cloudflare_list_item" "var_list_item_0" {
  account_id = var.account_id
  list_id    = cloudflare_list.var_list.id
  comment    = var.comment
  ip         = var.ip_address
}`,
			},
		},
		{
			Name: "list with multiple items of different complexity",
			Config: `
resource "cloudflare_list" "mixed_list" {
  account_id  = local.account_id
  name        = "mixed_${var.environment}"
  description = "Mixed complexity list"
  kind        = "ip"
  
  item {
    value {
      ip = "10.0.0.0/8"
    }
    comment = "Private network"
  }
  
  item {
    value {
      ip = cidrsubnet(var.base_network, 8, 1)
    }
    comment = "Dynamic subnet"
  }
  
  item {
    value {
      ip = "192.168.1.1"
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "mixed_list" {
  account_id  = local.account_id
  name        = "mixed_${var.environment}"
  description = "Mixed complexity list"
  kind        = "ip"
}`,
				`resource "cloudflare_list_item" "mixed_list_item_0" {
  account_id = local.account_id
  list_id    = cloudflare_list.mixed_list.id
  comment    = "Private network"
  ip         = "10.0.0.0/8"
}`,
				`resource "cloudflare_list_item" "mixed_list_item_1" {
  account_id = local.account_id
  list_id    = cloudflare_list.mixed_list.id
  ip         = "192.168.1.1"
}`,
				`resource "cloudflare_list_item" "mixed_list_item_2" {
  account_id = local.account_id
  list_id    = cloudflare_list.mixed_list.id
  comment    = "Dynamic subnet"
  ip         = cidrsubnet(var.base_network, 8, 1)
}`,
			},
		},
		{
			Name: "redirect list with variables in boolean fields",
			Config: `
resource "cloudflare_list" "redirect_vars" {
  account_id  = "abc123"
  name        = "redirect_vars"
  description = "Redirect with variable booleans"
  kind        = "redirect"
  
  item {
    value {
      redirect {
        source_url            = var.source
        target_url            = var.target
        include_subdomains    = var.include_subs ? "enabled" : "disabled"
        preserve_query_string = var.preserve_qs ? "enabled" : "disabled"
        status_code          = var.status_code
      }
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "redirect_vars" {
  account_id  = "abc123"
  name        = "redirect_vars"
  description = "Redirect with variable booleans"
  kind        = "redirect"
}`,
				`resource "cloudflare_list_item" "redirect_vars_item_0" {
  account_id = "abc123"
  list_id    = cloudflare_list.redirect_vars.id
  redirect = {
    source_url            = var.source
    target_url            = var.target
    include_subdomains    = var.include_subs ? "enabled" : "disabled"
    preserve_query_string = var.preserve_qs ? "enabled" : "disabled"
    status_code           = var.status_code
  }
}`,
			},
		},
		{
			Name: "ASN list with multiple items and comments",
			Config: `
resource "cloudflare_list" "cloud_providers" {
  account_id  = "xyz789"
  name        = "cloud_provider_asns"
  description = "Major cloud provider ASNs for rate limiting"
  kind        = "asn"

  item {
    value {
      asn = "398324"
    }
    comment = "Oracle Cloud"
  }
  item {
    value {
      asn = "63949"
    }
    comment = "Linode"
  }
  item {
    value {
      asn = "197540"
    }
    comment = "Netcup"
  }
  item {
    value {
      asn = "201814"
    }
    comment = "MEVSPACE"
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "cloud_providers" {
  account_id  = "xyz789"
  name        = "cloud_provider_asns"
  description = "Major cloud provider ASNs for rate limiting"
  kind        = "asn"
}`,
				`resource "cloudflare_list_item" "cloud_providers_item_0" {
  account_id = "xyz789"
  list_id    = cloudflare_list.cloud_providers.id
  comment    = "Linode"
  asn        = 63949
}`,
				`resource "cloudflare_list_item" "cloud_providers_item_1" {
  account_id = "xyz789"
  list_id    = cloudflare_list.cloud_providers.id
  comment    = "Netcup"
  asn        = 197540
}`,
				`resource "cloudflare_list_item" "cloud_providers_item_2" {
  account_id = "xyz789"
  list_id    = cloudflare_list.cloud_providers.id
  comment    = "MEVSPACE"
  asn        = 201814
}`,
				`resource "cloudflare_list_item" "cloud_providers_item_3" {
  account_id = "xyz789"
  list_id    = cloudflare_list.cloud_providers.id
  comment    = "Oracle Cloud"
  asn        = 398324
}`,
			},
		},
		{
			Name: "IP list with CIDR blocks and comments",
			Config: `
resource "cloudflare_list" "office_networks" {
  account_id  = local.account_id
  name        = "office_networks"
  description = "Corporate office network ranges"
  kind        = "ip"

  item {
    value {
      ip = "203.0.113.0/24"
    }
    comment = "London office"
  }
  item {
    value {
      ip = "198.51.100.0/24"
    }
    comment = "Tokyo office"
  }
  item {
    value {
      ip = "192.0.2.0/24"
    }
    comment = "New York office"
  }
  item {
    value {
      ip = "2001:db8:1234::/48"
    }
    comment = "IPv6 range for all offices"
  }
  item {
    value {
      ip = "172.16.0.0/12"
    }
    comment = "Internal VPN range"
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "office_networks" {
  account_id  = local.account_id
  name        = "office_networks"
  description = "Corporate office network ranges"
  kind        = "ip"
}`,
				`resource "cloudflare_list_item" "office_networks_item_0" {
  account_id = local.account_id
  list_id    = cloudflare_list.office_networks.id
  comment    = "Internal VPN range"
  ip         = "172.16.0.0/12"
}`,
				`resource "cloudflare_list_item" "office_networks_item_1" {
  account_id = local.account_id
  list_id    = cloudflare_list.office_networks.id
  comment    = "New York office"
  ip         = "192.0.2.0/24"
}`,
				`resource "cloudflare_list_item" "office_networks_item_2" {
  account_id = local.account_id
  list_id    = cloudflare_list.office_networks.id
  comment    = "Tokyo office"
  ip         = "198.51.100.0/24"
}`,
				`resource "cloudflare_list_item" "office_networks_item_3" {
  account_id = local.account_id
  list_id    = cloudflare_list.office_networks.id
  comment    = "IPv6 range for all offices"
  ip         = "2001:db8:1234::/48"
}`,
				`resource "cloudflare_list_item" "office_networks_item_4" {
  account_id = local.account_id
  list_id    = cloudflare_list.office_networks.id
  comment    = "London office"
  ip         = "203.0.113.0/24"
}`,
			},
		},
		{
			Name: "hostname list with wildcard and specific domains",
			Config: `
resource "cloudflare_list" "allowed_domains" {
  account_id  = "def456"
  name        = "allowed_partner_domains"
  description = "Partner domains allowed to access APIs"
  kind        = "hostname"
  
  item {
    value {
      hostname {
        url_hostname = "*.partner1.com"
      }
    }
    comment = "Partner 1 - all subdomains"
  }
  item {
    value {
      hostname {
        url_hostname = "api.partner2.net"
      }
    }
    comment = "Partner 2 - API only"
  }
  item {
    value {
      hostname {
        url_hostname = "secure.partner3.org"
      }
    }
    comment = "Partner 3 - secure endpoint"
  }
  item {
    value {
      hostname {
        url_hostname = "*.dev.partner4.io"
      }
    }
    comment = "Partner 4 - dev environments"
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "allowed_domains" {
  account_id  = "def456"
  name        = "allowed_partner_domains"
  description = "Partner domains allowed to access APIs"
  kind        = "hostname"
}`,
				`resource "cloudflare_list_item" "allowed_domains_item_0" {
  account_id = "def456"
  list_id    = cloudflare_list.allowed_domains.id
  comment    = "Partner 4 - dev environments"
  hostname = {
    url_hostname = "*.dev.partner4.io"
  }
}`,
				`resource "cloudflare_list_item" "allowed_domains_item_1" {
  account_id = "def456"
  list_id    = cloudflare_list.allowed_domains.id
  comment    = "Partner 1 - all subdomains"
  hostname = {
    url_hostname = "*.partner1.com"
  }
}`,
				`resource "cloudflare_list_item" "allowed_domains_item_2" {
  account_id = "def456"
  list_id    = cloudflare_list.allowed_domains.id
  comment    = "Partner 2 - API only"
  hostname = {
    url_hostname = "api.partner2.net"
  }
}`,
				`resource "cloudflare_list_item" "allowed_domains_item_3" {
  account_id = "def456"
  list_id    = cloudflare_list.allowed_domains.id
  comment    = "Partner 3 - secure endpoint"
  hostname = {
    url_hostname = "secure.partner3.org"
  }
}`,
			},
		},
		{
			Name: "redirect list with complex redirects",
			Config: `
resource "cloudflare_list" "url_redirects" {
  account_id  = "ghi789"
  name        = "marketing_redirects"
  description = "Marketing campaign URL redirects"
  kind        = "redirect"
  
  item {
    value {
      redirect {
        source_url            = "promo.example.com/summer2024"
        target_url            = "https://shop.example.com/sales/summer"
        include_subdomains    = "enabled"
        subpath_matching      = "enabled"
        status_code          = 302
        preserve_query_string = "enabled"
        preserve_path_suffix  = "disabled"
      }
    }
    comment = "Summer campaign redirect"
  }
  item {
    value {
      redirect {
        source_url            = "old.example.com/*"
        target_url            = "https://new.example.com/"
        include_subdomains    = "disabled"
        subpath_matching      = "enabled"
        status_code          = 301
        preserve_query_string = "disabled"
        preserve_path_suffix  = "enabled"
      }
    }
    comment = "Legacy domain redirect"
  }
  item {
    value {
      redirect {
        source_url            = "docs.example.com/v1/*"
        target_url            = "https://docs.example.com/latest/"
        include_subdomains    = "disabled"
        subpath_matching      = "enabled"
        status_code          = 308
        preserve_query_string = "enabled"
        preserve_path_suffix  = "enabled"
      }
    }
    comment = "Documentation version redirect"
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "url_redirects" {
  account_id  = "ghi789"
  name        = "marketing_redirects"
  description = "Marketing campaign URL redirects"
  kind        = "redirect"
}`,
				`resource "cloudflare_list_item" "url_redirects_item_0" {
  account_id = "ghi789"
  list_id    = cloudflare_list.url_redirects.id
  comment    = "Documentation version redirect"
  redirect = {
    source_url            = "docs.example.com/v1/*"
    target_url            = "https://docs.example.com/latest/"
    include_subdomains    = false
    subpath_matching      = true
    status_code           = 308
    preserve_query_string = true
    preserve_path_suffix  = true
  }
}`,
				`resource "cloudflare_list_item" "url_redirects_item_1" {
  account_id = "ghi789"
  list_id    = cloudflare_list.url_redirects.id
  comment    = "Legacy domain redirect"
  redirect = {
    source_url            = "old.example.com/*"
    target_url            = "https://new.example.com/"
    include_subdomains    = false
    subpath_matching      = true
    status_code           = 301
    preserve_query_string = false
    preserve_path_suffix  = true
  }
}`,
				`resource "cloudflare_list_item" "url_redirects_item_2" {
  account_id = "ghi789"
  list_id    = cloudflare_list.url_redirects.id
  comment    = "Summer campaign redirect"
  redirect = {
    source_url            = "promo.example.com/summer2024"
    target_url            = "https://shop.example.com/sales/summer"
    include_subdomains    = true
    subpath_matching      = true
    status_code           = 302
    preserve_query_string = true
    preserve_path_suffix  = false
  }
}`,
			},
		},
		{
			Name: "mixed IP list with single IPs and CIDR blocks",
			Config: `
resource "cloudflare_list" "monitoring_services" {
  account_id  = local.account_id
  name        = "monitoring_service_ips"
  description = "IPs for monitoring and alerting services"
  kind        = "ip"

  item {
    value {
      ip = "8.8.8.8"
    }
    comment = "Google DNS for health checks"
  }
  item {
    value {
      ip = "1.1.1.1"
    }
    comment = "Cloudflare DNS for health checks"
  }
  item {
    value {
      ip = "185.60.216.0/22"
    }
    comment = "StatusCake monitoring"
  }
  item {
    value {
      ip = "151.106.0.0/16"
    }
    comment = "Pingdom monitoring range"
  }
  item {
    value {
      ip = "2a03:b0c0:3:d0::1af1:1"
    }
    comment = "DigitalOcean monitoring IPv6"
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "monitoring_services" {
  account_id  = local.account_id
  name        = "monitoring_service_ips"
  description = "IPs for monitoring and alerting services"
  kind        = "ip"
}`,
				`resource "cloudflare_list_item" "monitoring_services_item_0" {
  account_id = local.account_id
  list_id    = cloudflare_list.monitoring_services.id
  comment    = "Cloudflare DNS for health checks"
  ip         = "1.1.1.1"
}`,
				`resource "cloudflare_list_item" "monitoring_services_item_1" {
  account_id = local.account_id
  list_id    = cloudflare_list.monitoring_services.id
  comment    = "Pingdom monitoring range"
  ip         = "151.106.0.0/16"
}`,
				`resource "cloudflare_list_item" "monitoring_services_item_2" {
  account_id = local.account_id
  list_id    = cloudflare_list.monitoring_services.id
  comment    = "StatusCake monitoring"
  ip         = "185.60.216.0/22"
}`,
				`resource "cloudflare_list_item" "monitoring_services_item_3" {
  account_id = local.account_id
  list_id    = cloudflare_list.monitoring_services.id
  comment    = "DigitalOcean monitoring IPv6"
  ip         = "2a03:b0c0:3:d0::1af1:1"
}`,
				`resource "cloudflare_list_item" "monitoring_services_item_4" {
  account_id = local.account_id
  list_id    = cloudflare_list.monitoring_services.id
  comment    = "Google DNS for health checks"
  ip         = "8.8.8.8"
}`,
			},
		},
		{
			Name: "ASN list without comments",
			Config: `
resource "cloudflare_list" "blocked_asns" {
  account_id  = var.account_id
  name        = "blocked_hosting_providers"
  description = "ASNs to block due to abuse"
  kind        = "asn"

  item {
    value {
      asn = "32934"
    }
  }
  item {
    value {
      asn = "19318"
    }
  }
  item {
    value {
      asn = "21859"
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "blocked_asns" {
  account_id  = var.account_id
  name        = "blocked_hosting_providers"
  description = "ASNs to block due to abuse"
  kind        = "asn"
}`,
				`resource "cloudflare_list_item" "blocked_asns_item_0" {
  account_id = var.account_id
  list_id    = cloudflare_list.blocked_asns.id
  asn        = 19318
}`,
				`resource "cloudflare_list_item" "blocked_asns_item_1" {
  account_id = var.account_id
  list_id    = cloudflare_list.blocked_asns.id
  asn        = 21859
}`,
				`resource "cloudflare_list_item" "blocked_asns_item_2" {
  account_id = var.account_id
  list_id    = cloudflare_list.blocked_asns.id
  asn        = 32934
}`,
			},
		},
		{
			Name: "IP list with dynamic blocks",
			Config: `
resource "cloudflare_list" "persona_webhook_ips" {
  account_id  = local.account_id
  name        = "persona_webhook_ips"
  description = "Static IPs used by Persona to send webhook notifications"
  kind        = "ip"

  dynamic "item" {
    for_each = local.persona_webhook_ips
    content {
      value {
        ip = item.value
      }
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "persona_webhook_ips" {
  account_id  = local.account_id
  name        = "persona_webhook_ips"
  description = "Static IPs used by Persona to send webhook notifications"
  kind        = "ip"
}`,
				`resource "cloudflare_list_item" "persona_webhook_ips_items" {
  for_each   = toset(local.persona_webhook_ips)
  account_id = local.account_id
  list_id    = cloudflare_list.persona_webhook_ips.id
  ip         = each.value
}`,
			},
		},
		{
			Name: "List with both static and dynamic items",
			Config: `
resource "cloudflare_list" "mixed_ips" {
  account_id  = var.account_id
  name        = "mixed_ip_list"
  description = "List with both static and dynamic items"
  kind        = "ip"

  item {
    value {
      ip = "192.0.2.1"
    }
    comment = "Static IP"
  }

  dynamic "item" {
    for_each = var.dynamic_ips
    content {
      value {
        ip = item.value
      }
      comment = "Dynamic IP"
    }
  }
}`,
			Expected: []string{
				`resource "cloudflare_list" "mixed_ips" {
  account_id  = var.account_id
  name        = "mixed_ip_list"
  description = "List with both static and dynamic items"
  kind        = "ip"
}`,
				`resource "cloudflare_list_item" "mixed_ips_item_0" {
  account_id = var.account_id
  list_id    = cloudflare_list.mixed_ips.id
  comment    = "Static IP"
  ip         = "192.0.2.1"
}`,
				`resource "cloudflare_list_item" "mixed_ips_items" {
  for_each   = toset(var.dynamic_ips)
  account_id = var.account_id
  list_id    = cloudflare_list.mixed_ips.id
  comment    = "Dynamic IP"
  ip         = each.value
}`,
			},
		},
	}

	// Run tests with custom assertion for list transformations
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Transform the config
			result, err := transformFile([]byte(tt.Config), "test.tf")
			if err != nil {
				t.Fatalf("transformation failed: %v", err)
			}

			// Use custom assertion for list transformations
			assertConfigsEquivalent(t, string(result), tt.Expected)
		})
	}
}

func TestCloudflareListStateTransformation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "basic list state with items",
			input: `{
  "version": 4,
  "terraform_version": "1.5.0",
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "example",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "name": "my_list",
            "description": "Test list",
            "kind": "ip",
            "id": "list123",
            "item": [
              {
                "value": [
                  {
                    "ip": "192.0.2.0",
                    "asn": null,
                    "hostname": [],
                    "redirect": []
                  }
                ],
                "comment": "First IP"
              },
              {
                "value": [
                  {
                    "ip": "192.0.2.1",
                    "asn": null,
                    "hostname": [],
                    "redirect": []
                  }
                ],
                "comment": "Second IP"
              }
            ]
          }
        }
      ]
    }
  ]
}`,
			expected: `{
  "version": 4,
  "terraform_version": "1.5.0",
  "resources": [
    {
      "mode": "managed",
      "type": "cloudflare_list",
      "name": "example",
      "provider": "provider[\"registry.terraform.io/cloudflare/cloudflare\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "account_id": "abc123",
            "name": "my_list",
            "description": "Test list",
            "kind": "ip",
            "id": "list123"
          }
        }
      ]
    }
  ]
}`,
		},
		{
			name: "list state with nested items",
			input: `{
  "resources": [
    {
      "type": "cloudflare_list",
      "instances": [
        {
          "attributes": {
            "account_id": "abc123",
            "name": "hostname_list",
            "kind": "hostname",
            "item": [
              {
                "value": [
                  {
                    "hostname": [
                      {
                        "url_hostname": "*.example.com"
                      }
                    ],
                    "ip": null,
                    "asn": null,
                    "redirect": []
                  }
                ]
              }
            ],
            "items": []
          }
        }
      ]
    }
  ]
}`,
			expected: `{
  "resources": [
    {
      "type": "cloudflare_list",
      "instances": [
        {
          "attributes": {
            "account_id": "abc123",
            "name": "hostname_list",
            "kind": "hostname"
          }
        }
      ]
    }
  ]
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := transformStateJSON([]byte(tt.input))
			if err != nil {
				t.Fatalf("transformation failed: %v", err)
			}

			// Check that item and items fields are removed
			resultStr := string(result)
			if contains(resultStr, `"item"`) {
				t.Errorf("Expected 'item' field to be removed from state")
			}
			if contains(resultStr, `"items"`) {
				t.Errorf("Expected 'items' field to be removed from state")
			}

			// Check that other fields are preserved
			if !contains(resultStr, `"account_id"`) {
				t.Errorf("Expected 'account_id' field to be preserved")
			}
			if !contains(resultStr, `"name"`) {
				t.Errorf("Expected 'name' field to be preserved")
			}
			if !contains(resultStr, `"kind"`) {
				t.Errorf("Expected 'kind' field to be preserved")
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestOAIMigration(t *testing.T) {
	hclBytes, err := os.ReadFile("/Users/vaishak/cf-repos/github/cf-terraforming/testdata/terraform/v4/cloudflare_list_asn/test.tf")
	if err != nil {
		t.Fatal(err)
	}
	transformed, err := transformFile(hclBytes, "test.tf")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("------------------------------------------------------")
	fmt.Println(string(transformed))
}
