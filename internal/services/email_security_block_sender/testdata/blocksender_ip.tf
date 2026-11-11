// 192.0.2.0/24 is the RFC 5737 documentation range: globally reachable as far as
// the API's validator is concerned (it rejects only loopback, broadcast,
// link-local, unspecified and private addresses), but never routed in practice.
resource "cloudflare_email_security_block_sender" "%[1]s" {
  account_id   = "%[2]s"
  pattern      = "192.0.2.0/24"
  pattern_type = "IP"
  is_regex     = false
  comments     = "%[1]s"
}
