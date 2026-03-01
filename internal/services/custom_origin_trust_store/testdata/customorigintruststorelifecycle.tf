# Known drift issue with certificate field.
# Workaround:
# Certificate must be in this format or if hardcoded (ends with \n).

resource "cloudflare_custom_origin_trust_store" "%[1]s" {
  zone_id     = "%[2]s"
  certificate = <<EOT
%[3]sEOT
}
 