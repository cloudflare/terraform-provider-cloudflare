
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id  = "%[1]s"
  hostname = "%[2]s.%[3]s"
  lifecycle {
    ignore_changes = [
      created_at,
      ownership_verification,
      ownership_verification_http,
      status,
      verification_errors,
    ]
  }
}
