resource "cloudflare_custom_origin_trust_store" "%[1]s" {
  zone_id     = "%[2]s"
  certificate = <<EOT
%[3]sEOT
}
 