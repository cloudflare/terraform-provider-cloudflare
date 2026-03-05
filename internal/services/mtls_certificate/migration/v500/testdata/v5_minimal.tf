resource "cloudflare_mtls_certificate" "%s" {
  account_id   = "%s"
  ca           = false
  certificates = <<EOT
%s
EOT
  private_key  = <<EOT
%s
EOT
}
