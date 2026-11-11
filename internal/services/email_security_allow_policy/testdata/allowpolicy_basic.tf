resource "cloudflare_email_security_allow_policy" "%[1]s" {
  account_id           = "%[2]s"
  pattern              = "%[1]s"
  pattern_type         = "DOMAIN"
  is_regex             = false
  is_acceptable_sender = false
  is_exempt_recipient  = false
  is_trusted_sender    = true
  verify_sender        = true
  comments             = "%[1]s"
}
