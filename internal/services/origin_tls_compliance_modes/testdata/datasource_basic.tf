data "cloudflare_origin_tls_compliance_modes" "%[2]s" {
  zone_id = "%[1]s"
}
