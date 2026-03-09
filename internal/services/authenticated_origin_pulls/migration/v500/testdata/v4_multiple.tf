resource "cloudflare_authenticated_origin_pulls_certificate" "%s" {
  zone_id     = "%s"
  certificate = <<-EOT
%s
  EOT
  private_key = <<-EOT
%s
  EOT
  type        = "per-hostname"
}

resource "cloudflare_authenticated_origin_pulls_certificate" "%s" {
  zone_id     = "%s"
  certificate = <<-EOT
%s
  EOT
  private_key = <<-EOT
%s
  EOT
  type        = "per-hostname"
}

resource "cloudflare_authenticated_origin_pulls" "%s" {
  zone_id                                = "%s"
  hostname                               = "%s"
  authenticated_origin_pulls_certificate = cloudflare_authenticated_origin_pulls_certificate.%s.id
  enabled                                = true
}

resource "cloudflare_authenticated_origin_pulls" "%s" {
  zone_id                                = "%s"
  hostname                               = "%s"
  authenticated_origin_pulls_certificate = cloudflare_authenticated_origin_pulls_certificate.%s.id
  enabled                                = false
}
