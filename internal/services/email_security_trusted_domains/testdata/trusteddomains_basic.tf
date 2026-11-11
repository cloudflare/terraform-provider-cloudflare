resource "cloudflare_email_security_trusted_domains" "%[1]s" {
  account_id    = "%[2]s"
  pattern       = "%[1]s.example.com"
  is_recent     = true
  is_similarity = false
  is_regex      = false
  comments      = "%[1]s"
}
