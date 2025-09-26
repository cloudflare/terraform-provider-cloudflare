resource "cloudflare_zero_trust_access_service_token" "%[1]s" {
  %[3]s_id                           = "%[4]s"
  name                               = "%[2]s"
  client_secret_version              = %[5]s
  previous_client_secret_expires_at  = "%[6]s"
}