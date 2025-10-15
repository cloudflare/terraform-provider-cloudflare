resource "cloudflare_content_scanning" "%[2]s" {
  zone_id = "%[1]s"
  value   = "invalid"
}
