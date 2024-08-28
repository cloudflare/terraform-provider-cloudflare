resource "cloudflare_zero_trust_tunnel_cloudflared" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "my-tunnel"
  secret     = "AQIDBAUGBwgBAgMEBQYHCAECAwQFBgcIAQIDBAUGBwg="
}
