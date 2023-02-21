resource "cloudflare_tunnel" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-tunnel"
  secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}
