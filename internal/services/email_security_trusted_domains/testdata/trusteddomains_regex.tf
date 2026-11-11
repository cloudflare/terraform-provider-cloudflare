resource "cloudflare_email_security_trusted_domains" "%[1]s" {
  account_id    = "%[2]s"
  pattern       = ".*\\.%[1]s\\.example\\.com"
  is_recent     = false
  is_regex      = true
  is_similarity = true
  comments      = "%[1]s"
}
