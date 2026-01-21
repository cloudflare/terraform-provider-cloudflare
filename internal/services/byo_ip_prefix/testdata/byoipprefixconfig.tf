resource "cloudflare_byo_ip_prefix" "%[6]s" {
  account_id      = "%[1]s"
  asn             = "%[2]d"
  cidr            = "%[3]s"
  loa_document_id = "%[4]s"
  description     = "%[5]s"
}
