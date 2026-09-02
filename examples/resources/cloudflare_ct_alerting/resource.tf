resource "cloudflare_ct_alerting" "example_ct_alerting" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  enabled = true
  emails = ["security@example.com", "admin@example.com"]
}
