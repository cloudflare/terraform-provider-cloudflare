resource "cloudflare_email_security_block_sender" "%[1]s" {
  account_id   = "%[2]s"
  pattern      = "%[1]s_updated@example.com"
  pattern_type = "EMAIL"
  is_regex     = false
  comments     = "%[1]s updated"
}
