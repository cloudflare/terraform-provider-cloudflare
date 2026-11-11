resource "cloudflare_email_security_allow_policy" "%[1]s" {
  account_id           = "%[2]s"
  pattern              = "%[1]s_updated@example.com"
  pattern_type         = "EMAIL"
  is_regex             = false
  is_acceptable_sender = true
  is_exempt_recipient  = false
  is_trusted_sender    = false
  verify_sender        = false
  comments             = "%[1]s updated"
}
