resource "cloudflare_monetization_rules" "example_monetization_rules" {
  zone_id = "9f1839b6152d298aca64c4e906b6d074"
  rules = [{
    address = "0x1234567890abcdef1234567890abcdef12345678==="
    expression = "(http.request.uri.path eq \"/premium\" and http.request.method in {\"GET\" \"POST\"})"
    price = "250000"
    scheme = "exact"
    id = "023e105f4ecef8ad9ca31a8372d0c353"
    description = "Premium API endpoint"
    enabled = true
  }]
}
