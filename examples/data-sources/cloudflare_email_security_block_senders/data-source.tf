data "cloudflare_email_security_block_senders" "example_email_security_block_senders" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  order = "pattern"
  pattern_type = "EMAIL"
  search = "search"
}
