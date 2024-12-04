data "cloudflare_access_rules" "example_access_rules" {
  account_id = "account_id"
  zone_id = "zone_id"
  configuration = {
    target = "ip"
    value = "198.51.100.4"
  }
  direction = "asc"
  mode = "block"
  notes = "my note"
  order = "configuration.target"
}
