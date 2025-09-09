package main

import (
	"testing"
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

	RunTransformationTests(t, tests, transformFile)
}