resource "cloudflare_sso_connector" "example_sso_connector" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  email_domain = "example.com"
  begin_verification = true
  use_fedramp_language = false
}
