resource "cloudflare_teams_proxy_endpoint" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "office"
  ips        = ["192.0.2.0/24"]
}
