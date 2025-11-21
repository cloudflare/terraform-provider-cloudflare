package main

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
)

func TestListItemMerge(t *testing.T) {
	tests := []TestCase{
		{
			Name: "merge static cloudflare_list_item resources into parent list",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
  description = "Example IP list"
}

resource "cloudflare_list_item" "item1" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example.id
  ip         = "192.0.2.1"
  comment    = "First IP"
}

resource "cloudflare_list_item" "item2" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example.id
  ip         = "192.0.2.2"
  comment    = "Second IP"
}`,
			Expected: []string{`
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
  description = "Example IP list"
  items = [{
    comment = "First IP"
    ip      = "192.0.2.1"
    }, {
    comment = "Second IP"
    ip      = "192.0.2.2"
  }]
}`},
		},
		{
			Name: "merge cloudflare_list_item with for_each into parent list",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
}

resource "cloudflare_list_item" "dynamic_items" {
  for_each   = var.ip_addresses
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example.id
  ip         = each.value.ip
  comment    = each.value.comment
}`,
			Expected: []string{`
resource "cloudflare_list" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example_list"
  kind       = "ip"
  items = [
    for k, v in var.ip_addresses : {
      comment = v.comment,
      ip      = v.ip
    }
  ]
}`},
		},
		{
			Name: "merge cloudflare_list_item with count into parent list",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
}

resource "cloudflare_list_item" "counted_items" {
  count      = length(var.ip_list)
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example.id
  ip         = var.ip_list[count.index]
}`,
			Expected: []string{`
resource "cloudflare_list" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example_list"
  kind       = "ip"
  items = [
    for i in range(length(var.ip_list)) : {
      ip = var.ip_list[i]
    }
  ]
}`},
		},
		{
			Name: "merge ASN list_item resources",
			Config: `
resource "cloudflare_list" "asn_list" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "asn_list"
  kind        = "asn"
}

resource "cloudflare_list_item" "asn1" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.asn_list.id
  asn        = 12345
  comment    = "Example ASN"
}`,
			Expected: []string{`
resource "cloudflare_list" "asn_list" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "asn_list"
  kind       = "asn"
  items = [{
    comment = "Example ASN"
    asn     = 12345
  }]
}`},
		},
		{
			Name: "merge hostname list_item resources",
			Config: `
resource "cloudflare_list" "hostname_list" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "hostname_list"
  kind        = "hostname"
}

resource "cloudflare_list_item" "hostname1" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.hostname_list.id
  hostname = {
    url_hostname = "example.com"
  }
  comment = "Example hostname"
}`,
			Expected: []string{`
resource "cloudflare_list" "hostname_list" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "hostname_list"
  kind       = "hostname"
  items = [{
    comment = "Example hostname"
    hostname = {
      url_hostname = "example.com"
    }
  }]
}`},
		},
		{
			Name: "merge redirect list_item resources",
			Config: `
resource "cloudflare_list" "redirect_list" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "redirect_list"
  kind        = "redirect"
}

resource "cloudflare_list_item" "redirect1" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.redirect_list.id
  redirect = {
    source_url            = "example.com/old"
    target_url            = "example.com/new"
    include_subdomains    = true
    preserve_query_string = false
    status_code           = 301
  }
}`,
			Expected: []string{`
resource "cloudflare_list" "redirect_list" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "redirect_list"
  kind       = "redirect"
  items = [{
    redirect = {
      source_url            = "example.com/old"
      target_url            = "example.com/new"
      include_subdomains    = true
      preserve_query_string = false
      status_code           = 301
    }
  }]
}`},
		},
		{
			Name: "list without any list_item resources remains unchanged",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
  description = "Example IP list"
}

resource "cloudflare_list_item" "other_item" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.other.id
  ip         = "192.0.2.1"
}`,
			Expected: []string{`
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
  description = "Example IP list"
}`,
				`resource "cloudflare_list_item" "other_item"`, // Item with non-matching list_id remains
			},
		},
		{
			Name: "handle bracket notation list references",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
}

resource "cloudflare_list_item" "item1" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list["example"].id
  ip         = "192.0.2.1"
}`,
			Expected: []string{`
resource "cloudflare_list" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example_list"
  kind       = "ip"
  items = [{
    ip = "192.0.2.1"
  }]
}`},
		},
		{
			Name: "preserve complex expressions in list_item fields",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
}

resource "cloudflare_list_item" "item1" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example.id
  ip         = aws_instance.example.public_ip
  comment    = "IP for ${aws_instance.example.name}"
}`,
			Expected: []string{`
resource "cloudflare_list" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example_list"
  kind       = "ip"
  items = [{
    comment = "IP for ${aws_instance.example.name}"
    ip      = aws_instance.example.public_ip
  }]
}`},
		},
	}

	RunTransformationTests(t, tests, transformFileDefault)
}

func TestListItemMergeWithCount(t *testing.T) {
	tests := []TestCase{
		{
			Name: "merge cloudflare_list_item with count into parent list",
			Config: `
resource "cloudflare_list" "example" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  kind        = "ip"
  description = "Example IP list"
}

resource "cloudflare_list_item" "items" {
  count      = 3
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example.id
  ip         = "192.0.2.${count.index + 1}"
  comment    = "IP number ${count.index + 1}"
}`,
			Expected: []string{`items = [
    for i in
    range(3)
    : {
      comment = "IP number ${i + 1}"
      ip      = "192.0.2.${i + 1}"
    }
  ]`},
		},
		{
			Name: "merge hostname list items with count",
			Config: `
resource "cloudflare_list" "hostnames" {
  account_id = "test"
  name       = "hostnames"
  kind       = "hostname"
}

resource "cloudflare_list_item" "hosts" {
  count      = length(var.hostnames)
  account_id = "test"
  list_id    = cloudflare_list.hostnames.id
  hostname = {
    url_hostname = var.hostnames[count.index]
  }
}`,
			Expected: []string{`items = [
    for i in
    range(length(var.hostnames))
    : { hostname = { url_hostname = var.hostnames[i] } }
  ]`},
		},
		{
			Name: "merge redirect list items with count",
			Config: `
resource "cloudflare_list" "redirects" {
  account_id = "test"
  name       = "redirects"
  kind       = "redirect"
}

resource "cloudflare_list_item" "redirect_items" {
  count      = length(var.redirects)
  account_id = "test"
  list_id    = cloudflare_list.redirects.id
  redirect = {
    source_url = var.redirects[count.index].source
    target_url = var.redirects[count.index].target
    status_code = 301
    subpath_matching = true
  }
}`,
			Expected: []string{`items = [
    for i in
    range(length(var.redirects))
    : {
      redirect = {
        source_url       = source
        target_url       = target
        status_code      = 301
        subpath_matching = true
      }
    }
  ]`},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}

func TestListItemMergeComplexCases(t *testing.T) {
	tests := []TestCase{
		{
			Name: "list with multiple item resources - asn type",
			Config: `
resource "cloudflare_list" "asn_list" {
  account_id  = "test"
  name        = "asn_list"
  kind        = "asn"
}

resource "cloudflare_list_item" "asn1" {
  account_id = "test"
  list_id    = cloudflare_list.asn_list.id
  asn        = 12345
  comment    = "Example ASN"
}

resource "cloudflare_list_item" "asn2" {
  account_id = "test"
  list_id    = cloudflare_list.asn_list.id
  asn        = 67890
}`,
			Expected: []string{`items = [{
    comment = "Example ASN"
    asn     = 12345
    },
  { asn = 67890 }]`},
		},
		{
			Name: "handle list items with no matching parent list",
			Config: `
resource "cloudflare_list_item" "orphan" {
  account_id = "test"
  list_id    = "external-list-id"
  ip         = "192.0.2.1"
  comment    = "Orphaned item"
}

resource "cloudflare_list" "unrelated" {
  account_id = "test"
  name       = "different_list"
  kind       = "ip"
}`,
			// Orphaned list_item should remain, unrelated list should remain unchanged
			Expected: []string{
				`resource "cloudflare_list_item" "orphan"`,
				`resource "cloudflare_list" "unrelated"`,
			},
		},
		{
			Name: "merge items with different field types",
			Config: `
resource "cloudflare_list" "mixed" {
  account_id = "test"
  name       = "mixed"
  kind       = "ip"
}

resource "cloudflare_list_item" "cidr" {
  account_id = "test"
  list_id    = cloudflare_list.mixed.id
  ip         = "10.0.0.0/8"
  comment    = "CIDR range"
}

resource "cloudflare_list_item" "single_ip" {
  account_id = "test"  
  list_id    = cloudflare_list.mixed.id
  ip         = "192.168.1.1"
  comment    = "Single IP"
}`,
			Expected: []string{`items = [{
    comment = "CIDR range"
    ip      = "10.0.0.0/8"
    }, {
    comment = "Single IP"
    ip      = "192.168.1.1"
  }]`},
		},
	}
	
	RunTransformationTests(t, tests, transformFileDefault)
}
func TestAddMigrationWarning(t *testing.T) {
	tests := []struct {
		name            string
		message         string
		expectedComment string
	}{
		{
			name:            "add simple warning",
			message:         "This resource needs manual review",
			expectedComment: "# MIGRATION WARNING: This resource needs manual review",
		},
		{
			name:            "add warning with special characters",
			message:         "Complex patterns like [0-9]+ are not supported",
			expectedComment: "# MIGRATION WARNING: Complex patterns like [0-9]+ are not supported",
		},
		{
			name:            "add multi-word warning",
			message:         "Unable to automatically merge cloudflare_list_item resources",
			expectedComment: "# MIGRATION WARNING: Unable to automatically merge cloudflare_list_item resources",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := hclwrite.NewEmptyFile()
			body := file.Body()
			
			// Add a dummy attribute to ensure body is not empty
			body.SetAttributeValue("test", cty.StringVal("value"))
			
			// Add migration warning
			addMigrationWarning(body, tt.message)
			
			// Check that the comment was added
			result := string(file.Bytes())
			assert.Contains(t, result, tt.expectedComment)
			assert.Contains(t, result, "test = \"value\"") // Original content preserved
		})
	}
}
