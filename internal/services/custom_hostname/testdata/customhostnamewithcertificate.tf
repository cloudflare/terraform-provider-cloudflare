
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl = {
    method             = "http"
    type               = "dv"
    bundle_method      = "force"
    custom_certificate = <<EOT
%[4]s
	EOT
    custom_key = <<EOT
%[5]s
	EOT
  }
}
