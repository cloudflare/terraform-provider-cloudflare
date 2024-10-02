
resource "cloudflare_zero_trust_access_identity_provider" "%[1]s" {
  %[2]s_id = "%[3]s"
  name     = "%[1]s"
  type     = "onetimepin"
}