resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  name             = "%[1]s"
  account_id       = "%[2]s"
  decision         = "bypass"
  session_duration = ""
  include = [{
    everyone = {}
  }]
}
