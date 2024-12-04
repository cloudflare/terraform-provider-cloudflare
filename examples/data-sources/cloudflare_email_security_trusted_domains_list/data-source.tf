data "cloudflare_email_security_trusted_domains_list" "example_email_security_trusted_domains_list" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  is_recent = true
  is_similarity = true
  order = "pattern"
  search = "search"
}
