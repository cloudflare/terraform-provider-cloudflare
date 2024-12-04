resource "cloudflare_byo_ip_prefix" "example_byo_ip_prefix" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  asn = 209242
  cidr = "192.0.2.0/24"
  loa_document_id = "d933b1530bc56c9953cf8ce166da8004"
}
