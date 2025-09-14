package main

import (
	"testing"
)

func TestCloudflareListTransformation(t *testing.T) {
	tests := []TestCase{
		{
			Name: "IP list migration",
			Config: `
resource "cloudflare_list" "ip_list" {
  account_id = "abc123"
  name = "ip_list"
  kind = "ip"
  description = "List of IP addresses"
  
  item {
    comment = "First IP"
    value {
      ip = "1.1.1.1"
    }
  }
  
  item {
    comment = "Second IP"
    value {
      ip = "1.1.1.2"
    }
  }
  
  item {
    value {
      ip = "1.1.1.3"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "ip_list" {
  account_id  = "abc123"
  name        = "ip_list"
  kind        = "ip"
  description = "List of IP addresses"
  items = [{
    ip      = "1.1.1.1"
    comment = "First IP"
    }, {
    ip      = "1.1.1.2"
    comment = "Second IP"
    }, {
    ip = "1.1.1.3"
  }]
}`},
		},
		{
			Name: "ASN list migration",
			Config: `
resource "cloudflare_list" "asn_list" {
  account_id = "abc123"
  name = "asn_list"
  kind = "asn"
  
  item {
    comment = "Google ASN"
    value {
      asn = 15169
    }
  }
  
  item {
    value {
      asn = 13335
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "asn_list" {
  account_id = "abc123"
  name       = "asn_list"
  kind       = "asn"
  items = [{
    asn     = 15169
    comment = "Google ASN"
    }, {
    asn = 13335
  }]
}`},
		},
		{
			Name: "Hostname list migration",
			Config: `
resource "cloudflare_list" "hostname_list" {
  account_id = "abc123"
  name = "hostname_list"
  kind = "hostname"
  
  item {
    comment = "Example hostname"
    value {
      hostname {
        url_hostname = "example.com"
      }
    }
  }
  
  item {
    value {
      hostname {
        url_hostname = "test.example.com"
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "hostname_list" {
  account_id = "abc123"
  name       = "hostname_list"
  kind       = "hostname"
  items = [{
    hostname = {
      url_hostname = "example.com"
    }
    comment = "Example hostname"
    }, {
    hostname = {
      url_hostname = "test.example.com"
    }
  }]
}`},
		},
		{
			Name: "Redirect list migration with boolean conversions",
			Config: `
resource "cloudflare_list" "redirect_list" {
  account_id = "abc123"
  name = "redirect_list"
  kind = "redirect"
  
  item {
    comment = "Main redirect"
    value {
      redirect {
        source_url = "example.com/old"
        target_url = "example.com/new"
        include_subdomains = "enabled"
        subpath_matching = "disabled"
        preserve_query_string = "enabled"
        preserve_path_suffix = "disabled"
        status_code = 301
      }
    }
  }
  
  item {
    value {
      redirect {
        source_url = "test.com"
        target_url = "newtest.com"
        include_subdomains = "disabled"
      }
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "redirect_list" {
  account_id = "abc123"
  name       = "redirect_list"
  kind       = "redirect"
  items = [{
    redirect = {
      source_url            = "example.com/old"
      target_url            = "example.com/new"
      include_subdomains    = true
      subpath_matching      = false
      preserve_query_string = true
      preserve_path_suffix  = false
      status_code           = 301
    }
    comment = "Main redirect"
    }, {
    redirect = {
      source_url         = "test.com"
      target_url         = "newtest.com"
      include_subdomains = false
    }
  }]
}`},
		},
		{
			Name: "Empty list (no items)",
			Config: `
resource "cloudflare_list" "empty_list" {
  account_id = "abc123"
  name = "empty_list"
  kind = "ip"
  description = "Empty list"
}`,
			Expected: []string{`
resource "cloudflare_list" "empty_list" {
  account_id  = "abc123"
  name        = "empty_list"
  kind        = "ip"
  description = "Empty list"
}`},
		},
		{
			Name: "Mixed list with comments and without",
			Config: `
resource "cloudflare_list" "mixed_list" {
  account_id = "abc123"
  name = "mixed_list"
  kind = "ip"
  
  item {
    comment = "With comment"
    value {
      ip = "10.0.0.1"
    }
  }
  
  item {
    value {
      ip = "10.0.0.2"
    }
  }
  
  item {
    comment = "Another comment"
    value {
      ip = "10.0.0.3"
    }
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "mixed_list" {
  account_id = "abc123"
  name       = "mixed_list"
  kind       = "ip"
  items = [{
    ip      = "10.0.0.1"
    comment = "With comment"
    }, {
    ip = "10.0.0.2"
    }, {
    ip      = "10.0.0.3"
    comment = "Another comment"
  }]
}`},
		},
	}
	
	RunTransformationTests(t, tests, transformFileWithoutImports)
}