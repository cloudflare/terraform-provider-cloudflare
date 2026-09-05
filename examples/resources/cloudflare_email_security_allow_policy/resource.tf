resource "cloudflare_email_security_allow_policy" "example_email_security_allow_policy" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  is_acceptable_sender = false
  is_exempt_recipient = false
  is_regex = false
  is_trusted_sender = true
  pattern = "test@example.com"
  pattern_type = "EMAIL"
  verify_sender = true
  comments = "Trust all messages send from test@example.com"
  is_recipient = false
  is_sender = true
  is_spoof = false
}
