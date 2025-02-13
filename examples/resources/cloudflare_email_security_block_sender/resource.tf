resource "cloudflare_email_security_block_sender" "example_email_security_block_sender" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  is_regex = true
  pattern = "x"
  pattern_type = "EMAIL"
  comments = "comments"
}
