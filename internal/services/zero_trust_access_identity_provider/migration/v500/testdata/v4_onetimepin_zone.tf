resource "cloudflare_access_identity_provider" "%[1]s" {
  zone_id = "%[2]s"
  name    = "%[1]s"
  type    = "onetimepin"
}
