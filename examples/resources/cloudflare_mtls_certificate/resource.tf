resource "cloudflare_mtls_certificate" "example" {
  account_id   = "f037e56e89293a057740de681ac9abbe"
  name         = "example"
  certificates = "-----BEGIN CERTIFICATE-----\nMIIDmDCCAoCgAwIBAgIUKTOAZNj...i4JhqeoTewsxndhDDE\n-----END CERTIFICATE-----"
  private_key  = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQE...1IS3EnQRrz6WMYA=\n-----END PRIVATE KEY-----"
  ca           = true
}
