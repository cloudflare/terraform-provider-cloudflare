# v5 minimal per-zone certificate - only required fields
resource "cloudflare_authenticated_origin_pulls_certificate" "%s" {
  zone_id     = "%s"
  certificate = <<-EOT
%sEOT
  private_key = <<-EOT
%sEOT
}
