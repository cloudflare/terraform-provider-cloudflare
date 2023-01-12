resource "cloudflare_teams_location" "example" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "office"
  client_default = true

  networks {
    network = "203.0.113.1/32"
  }

  networks {
    network = "203.0.113.2/32"
  }
}
