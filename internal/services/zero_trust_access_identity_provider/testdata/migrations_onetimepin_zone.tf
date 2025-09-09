resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  zone_id = "%[1]s"
  name    = "%[2]s"
  type    = "onetimepin"
}