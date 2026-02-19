resource "cloudflare_authenticated_origin_pulls_certificate" "%[1]s" {
  zone_id     = "%[2]s"
  certificate = <<EOT
%[3]s
EOT
  private_key = <<EOT
%[4]s
EOT
}
