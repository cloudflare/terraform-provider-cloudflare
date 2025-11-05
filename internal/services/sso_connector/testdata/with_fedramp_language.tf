resource "cloudflare_sso_connector" "%[1]s" {
  account_id           = "%[2]s"
  email_domain         = "%[1]s.example.com"
  begin_verification   = false
  enabled              = false
  use_fedramp_language = true
}
