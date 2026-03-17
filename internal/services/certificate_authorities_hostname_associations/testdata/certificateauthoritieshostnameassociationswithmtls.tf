resource "cloudflare_certificate_authorities_hostname_associations" "%[1]s" {
  zone_id             = "%[2]s"
  hostnames           = ["%[3]s"]
  mtls_certificate_id = "%[4]s"
}
