resource "cloudflare_zero_trust_access_mtls_certificate" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  certificate = <<EOT
%[4]sEOT
}
