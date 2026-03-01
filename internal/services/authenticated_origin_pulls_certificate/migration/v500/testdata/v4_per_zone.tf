resource "cloudflare_authenticated_origin_pulls_certificate" "%s" {
  zone_id     = "%s"
  certificate = <<-EOT
%sEOT
  private_key = <<-EOT
%sEOT
  type        = "per-zone"
}
