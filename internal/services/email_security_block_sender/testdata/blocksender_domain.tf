resource "cloudflare_email_security_block_sender" "%[1]s" {
  account_id   = "%[2]s"
  pattern      = "%[1]s.example.com"
  pattern_type = "DOMAIN"
  is_regex     = false
}
