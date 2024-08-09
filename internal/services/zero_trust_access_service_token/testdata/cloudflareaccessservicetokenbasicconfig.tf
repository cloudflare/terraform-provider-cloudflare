
resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  %[3]s_id = "%[4]s"
  name     = "%[2]s"
  # min_days_for_renewal = "0"
}
