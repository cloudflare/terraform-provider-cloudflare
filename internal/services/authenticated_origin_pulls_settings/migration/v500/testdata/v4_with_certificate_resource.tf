resource "cloudflare_authenticated_origin_pulls_certificate" "%[1]s_cert" {
  zone_id     = "%[2]s"
  certificate = <<-EOT
%[3]s
  EOT
  private_key = <<-EOT
%[4]s
  EOT
  type        = "per-zone"
}

resource "cloudflare_authenticated_origin_pulls" "%[1]s" {
  zone_id                                = "%[2]s"
  authenticated_origin_pulls_certificate = cloudflare_authenticated_origin_pulls_certificate.%[1]s_cert.id
  enabled                                = true
}
