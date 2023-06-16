resource "cloudflare_list" "example_ip_list" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_list"
  description = "example IPs for a list"
  kind        = "ip"
}

# IP List Item
resource "cloudflare_list_item" "example_ip_item" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example_ip_list.id
  comment    = "List Item Comment"
  ip         = "192.0.2.0"
}


# Redirect List Item
resource "cloudflare_list_item" "test_two" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example_ip_list.id
  redirect {
    source_url       = "https://source.tld"
    target_url       = "https://target.tld"
    status_code      = 302
    subpath_matching = "enabled"
  }
}

# ASN list
resource "cloudflare_list" "example_asn_list" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_asn_list"
  description = "example ASNs for a list"
  kind        = "asn"
}

# ASN List Item
resource "cloudflare_list_item" "example_asn_item" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example_asn_list.id
  comment    = "List Item Comment"
  asn         = 6789
}

# Hostname list
resource "cloudflare_list" "example_hostname_list" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "example_hostname_list"
  description = "example Hostnames for a list"
  kind        = "hostname"
}

# Hostname List Item
resource "cloudflare_list_item" "example_hostname_item" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = cloudflare_list.example_hostname_list.id
  comment    = "List Item Comment"
  asn         = "example.com"
}
