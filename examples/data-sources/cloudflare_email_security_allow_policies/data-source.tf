data "cloudflare_email_security_allow_policies" "example_email_security_allow_policies" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  is_acceptable_sender = true
  is_exempt_recipient = true
  is_trusted_sender = true
  order = "pattern"
  pattern = "pattern"
  pattern_type = "EMAIL"
  search = "search"
  verify_sender = true
}
