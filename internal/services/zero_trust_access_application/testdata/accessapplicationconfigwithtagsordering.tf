resource "cloudflare_zero_trust_access_tag" "tag_ccc_%[1]s" {
  account_id = "%[3]s"
  name       = "ccc"
}

resource "cloudflare_zero_trust_access_tag" "tag_aaa_%[1]s" {
  account_id = "%[3]s"
  name       = "aaa"
}

resource "cloudflare_zero_trust_access_tag" "tag_bbb_%[1]s" {
  account_id = "%[3]s"
  name       = "bbb"
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[3]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[2]s"
  type             = "self_hosted"
  session_duration = "24h"
  tags             = [
    cloudflare_zero_trust_access_tag.tag_ccc_%[1]s.name,
    cloudflare_zero_trust_access_tag.tag_aaa_%[1]s.name,
    cloudflare_zero_trust_access_tag.tag_bbb_%[1]s.name
  ]
}