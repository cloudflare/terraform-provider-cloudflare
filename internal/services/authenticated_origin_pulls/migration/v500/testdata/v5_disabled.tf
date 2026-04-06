resource "cloudflare_authenticated_origin_pulls_hostname_certificate" "%s" {
  zone_id     = "%s"
  certificate = <<-EOT
%s
  EOT
  private_key = <<-EOT
%s
  EOT
}

resource "cloudflare_authenticated_origin_pulls" "%s" {
  zone_id = "%s"
  config = [{
    hostname = "%s"
    cert_id  = cloudflare_authenticated_origin_pulls_hostname_certificate.%s.id
    enabled  = false
  }]
}
