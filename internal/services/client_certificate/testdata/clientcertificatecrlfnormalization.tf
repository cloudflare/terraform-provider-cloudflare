resource "cloudflare_client_certificate" "%[1]s" {
  zone_id = "%[2]s"
  csr     = <<EOT
%[3]s
EOT
  validity_days = %[4]d
}
