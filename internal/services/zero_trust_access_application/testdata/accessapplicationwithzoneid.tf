
    resource "cloudflare_zero_trust_access_application" "%[1]s" {
      name    = "%[1]s"
      zone_id = "%[3]s"
      domain  = "%[1]s.%[2]s"
      type    = "self_hosted"
    }
  