resource "cloudflare_custom_pages" "%[1]s" {
  zone_id    = "%[2]s"
  identifier = "500_errors"
  state      = "default"
  url        = ""
}

data "cloudflare_custom_pages" "%[1]s" {
  zone_id    = cloudflare_custom_pages.%[1]s.zone_id
  identifier = cloudflare_custom_pages.%[1]s.identifier
}