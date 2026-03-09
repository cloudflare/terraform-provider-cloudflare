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
