resource "cloudflare_email_security_trusted_domains" "%[1]s" {
  account_id    = "%[2]s"
  pattern       = "%[1]s.example.net"
  is_recent     = false
  is_similarity = true
  is_regex      = false
}
