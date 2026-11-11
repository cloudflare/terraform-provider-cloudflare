resource "cloudflare_email_security_allow_policy" "%[1]s" {
  account_id           = "%[2]s"
  pattern              = "%[1]s@example.com"
  pattern_type         = "EMAIL"
  is_regex             = false
  is_acceptable_sender = false
  is_exempt_recipient  = true
  is_trusted_sender    = false
  verify_sender        = true
}
