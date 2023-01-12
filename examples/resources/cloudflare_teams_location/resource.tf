resource "cloudflare_teams_location" "example" {
  account_id     = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name           = "office"
  client_default = true

  networks {
    network = "203.0.113.1/32"
  }

  networks {
    network = "203.0.113.2/32"
  }
}
