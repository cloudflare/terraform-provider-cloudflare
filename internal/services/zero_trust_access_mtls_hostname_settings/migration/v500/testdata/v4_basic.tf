resource "cloudflare_zero_trust_access_mtls_hostname_settings" "%[1]s" {
  account_id = "%[2]s"
  settings {
    hostname                     = "%[3]s"
    client_certificate_forwarding = true
    china_network                 = false
  }
}
