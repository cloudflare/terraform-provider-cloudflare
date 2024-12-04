data "cloudflare_firewall_rules" "example_firewall_rules" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "372e67954025e0ba6aaa6d586b9e0b60"
  action = "block"
  description = "mir"
  paused = false
}
