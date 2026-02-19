resource "cloudflare_access_mutual_tls_certificate" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  certificate = <<EOT
%[4]sEOT
}
