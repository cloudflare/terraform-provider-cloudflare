resource "cloudflare_custom_ssl" "%[2]s" {
  zone_id       = "%[1]s"
  certificate   = <<EOT
%[3]s
EOT
  private_key   = <<EOT
%[4]s
EOT
  bundle_method = "force"
  geo_restrictions = {
    label = "us"
  }
}
