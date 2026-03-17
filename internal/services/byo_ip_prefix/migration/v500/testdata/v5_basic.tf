resource "cloudflare_byo_ip_prefix" "%[1]s" {
  account_id  = "%[2]s"
  asn         = %[3]d
  cidr        = "%[4]s"
  description = "Migration test prefix"
}
