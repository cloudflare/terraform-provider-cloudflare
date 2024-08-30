
resource "cloudflare_keyless_certificate" "%[1]s" {
  zone_id       = "%[2]s"
  bundle_method = "force"
  name          = "%[1]s"
  host          = "%[3]s"
  port          = 24008
  certificate   = <<EOT
%[4]s
  EOT
}