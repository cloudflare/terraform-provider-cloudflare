resource "cloudflare_email_auth_dmarc_report" "example_email_auth_dmarc_report" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  enabled = true
  skip_wizard = true
}
