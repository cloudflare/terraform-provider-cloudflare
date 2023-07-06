# For account level applications
resource "cloudflare_access_application" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example"
  domain     = "example.com"
}

# You can use either `name` or `domain` to identify the application
data "cloudflare_access_application" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example"
}

data "cloudflare_access_application" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  domain     = "example.com"
}

# For zone level applications
resource "cloudflare_access_application" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "example"
  domain  = "example.com"
}

# You can use either `name` or `domain` to identify the application
data "cloudflare_access_application" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name    = "example"
}

data "cloudflare_access_application" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  domain  = "example.com"
}
