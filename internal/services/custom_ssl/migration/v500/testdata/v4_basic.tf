resource "cloudflare_custom_ssl" "%[1]s" {
  zone_id = "%[2]s"
  custom_ssl_options {
    certificate   = <<EOT
%[3]s
EOT
    private_key   = <<EOT
%[4]s
EOT
    bundle_method = "force"
    type          = "legacy_custom"
  }
}
