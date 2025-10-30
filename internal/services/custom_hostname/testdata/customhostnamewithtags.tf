resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl = {
    method = "txt"
    type = "dv"
  }
  tags = {
    environment = "production"
    team        = "web-platform"
    managed_by  = "terraform"
  }

  lifecycle {
    ignore_changes = [
      created_at,
      ownership_verification,
      ownership_verification_http,
      ssl.certificate_authority,
      ssl.wildcard,
      status,
      verification_errors
    ]
  }
}
