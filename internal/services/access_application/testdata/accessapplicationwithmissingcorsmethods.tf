
    resource "cloudflare_access_application" "%[1]s" {
      name    = "%[1]s-updated"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
      type    = "self_hosted"

    cors_headers = {
  allow_all_origins = true
}
  }
  