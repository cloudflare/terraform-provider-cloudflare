resource "cloudflare_zero_trust_access_service_token" "example_zero_trust_access_service_token" {
  name = "CI/CD token"
  zone_id = "zone_id"
  client_secret_version = 0
  duration = "60m"
  previous_client_secret_expires_at = "2014-01-01T05:20:00.12345Z"
}
