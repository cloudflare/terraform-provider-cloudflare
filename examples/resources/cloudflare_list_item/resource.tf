resource "cloudflare_list" "example_ip_list" {
  account_id          = "f037e56e89293a057740de681ac9abbe"
  name                = "example_list"
  description         = "example IPs for a list"
  kind                = "ip"
}

# IP List Item
resource "cloudflare_list_item" example_ip_item" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = data.cloudflare_list.example_ip_list.id
  comment    = "List Item Comment"
  ip         = "192.0.2.0"
}


# Redirect List Item
resource "cloudflare_list_item" "test_two" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  list_id    = data.cloudflare_list.example_ip_list.id
  redirect {
    source_url       = "https://source.tld"
    target_url       = "https://target.tld"
    status_code      = 302
    subpath_matching = "enabled"
  }
}