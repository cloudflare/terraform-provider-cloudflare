package main

import (
	"testing"
)

func TestCloudflareListDynamicBlockTransformations(t *testing.T) {
	tests := []TestCase{
		{
			Name: "simple dynamic block with for_each list",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  dynamic "item" {
    for_each = var.ip_list
    content {
      value {
        ip = item.value
      }
      comment = "IP ${item.value}"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  items = [
    for item in var.ip_list : {
      ip      = item
      comment = "IP ${item}"
    }
  ]
}`},
		},
		{
			Name: "dynamic block with custom iterator",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  dynamic "item" {
    for_each = var.ip_addresses
    iterator = ip
    content {
      value {
        ip = ip.value
      }
      comment = "Address: ${ip.value}"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  items = [
    for ip in var.ip_addresses : {
      ip      = ip
      comment = "Address: ${ip}"
    }
  ]
}`},
		},
		{
			Name: "dynamic block for hostname list",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "hostname_list"
  kind       = "hostname"

  dynamic "item" {
    for_each = var.hostnames
    content {
      value {
        hostname {
          url_hostname = item.value
        }
      }
      comment = "Hostname: ${item.value}"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "hostname_list"
  kind       = "hostname"

  items = [
    for item in var.hostnames : {
      hostname = {
        url_hostname = item
      }
      comment = "Hostname: ${item}"
    }
  ]
}`},
		},
		{
			Name: "mixed static and dynamic items",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "mixed_list"
  kind       = "ip"

  item {
    value {
      ip = "192.168.1.1"
    }
    comment = "Static IP"
  }

  dynamic "item" {
    for_each = var.dynamic_ips
    content {
      value {
        ip = item.value
      }
      comment = "Dynamic: ${item.value}"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "mixed_list"
  kind       = "ip"

  items = concat([
    {
      ip      = "192.168.1.1"
      comment = "Static IP"
    }
  ], [
    for item in var.dynamic_ips : {
      ip      = item
      comment = "Dynamic: ${item}"
    }
  ])
}`},
		},
		{
			Name: "dynamic block with ASN list",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "asn_list"
  kind       = "asn"

  dynamic "item" {
    for_each = var.asn_numbers
    content {
      value {
        asn = item.value
      }
      comment = "ASN ${item.value}"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "asn_list"
  kind       = "asn"

  items = [
    for item in var.asn_numbers : {
      asn     = item
      comment = "ASN ${item}"
    }
  ]
}`},
		},
		{
			Name: "dynamic block with redirect list and boolean conversions",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "redirect_list"
  kind       = "redirect"

  dynamic "item" {
    for_each = var.redirects
    content {
      value {
        redirect {
          source_url            = item.value.source
          target_url            = item.value.target
          include_subdomains    = "enabled"
          subpath_matching      = "disabled"
          preserve_query_string = item.value.preserve_query ? "enabled" : "disabled"
          status_code           = item.value.code
        }
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "redirect_list"
  kind       = "redirect"

  items = [
    for item in var.redirects : {
      redirect = {
        source_url            = item.source
        target_url            = item.target
        include_subdomains    = true
        subpath_matching      = false
        preserve_query_string = item.preserve_query ? true : false
        status_code           = item.code
      }
    }
  ]
}`},
		},
		{
			Name: "problematic toset() pattern with warning",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"
  for_each   = toset(var.lists)

  dynamic "item" {
    for_each = toset(var.ip_list)
    content {
      value {
        ip = item.value
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"
  for_each   = toset(var.lists)

  # MIGRATION WARNING: toset() in for_each makes keys and values identical. Consider using a map for distinct keys and values.
  items = [
    for item in toset(var.ip_list) : {
      ip = item
    }
  ]
}`},
		},
		{
			Name: "complex conditional in dynamic block",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  dynamic "item" {
    for_each = var.env == "prod" ? var.prod_ips : var.dev_ips
    content {
      value {
        ip = item.value
      }
      comment = var.env == "prod" ? "Production IP" : "Development IP"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  items = [
    for item in var.env == "prod" ? var.prod_ips : var.dev_ips : {
      ip      = item
      comment = var.env == "prod" ? "Production IP" : "Development IP"
    }
  ]
}`},
		},
		{
			Name: "multiple dynamic blocks",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "multi_list"
  kind       = "ip"

  dynamic "item" {
    for_each = var.internal_ips
    content {
      value {
        ip = item.value
      }
      comment = "Internal IP"
    }
  }

  dynamic "item" {
    for_each = var.external_ips
    content {
      value {
        ip = item.value
      }
      comment = "External IP"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "multi_list"
  kind       = "ip"

  items = concat([
    for item in var.internal_ips : {
      ip      = item
      comment = "Internal IP"
    }
  ], [
    for item in var.external_ips : {
      ip      = item
      comment = "External IP"
    }
  ])
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}

func TestCloudflareListEdgeCases(t *testing.T) {
	tests := []TestCase{
		{
			Name: "empty dynamic block",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  dynamic "item" {
    for_each = []
    content {
      value {
        ip = item.value
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  items = [
    for item in [] : {
      ip = item
    }
  ]
}`},
		},
		{
			Name: "nested for expressions",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  dynamic "item" {
    for_each = [for s in var.subnets : cidrhost(s, 0)]
    content {
      value {
        ip = item.value
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "test_list"
  kind       = "ip"

  items = [
    for item in [for s in var.subnets : cidrhost(s, 0)] : {
      ip = item
    }
  ]
}`},
		},
		{
			Name: "no item blocks",
			Config: `
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "empty_list"
  kind       = "ip"
  description = "An empty list"
}`,
			Expected: []string{`
resource "cloudflare_list" "test" {
  account_id = "abc123"
  name       = "empty_list"
  kind       = "ip"
  description = "An empty list"
}`},
		},
	}

	RunTransformationTests(t, tests, transformFile)
}
