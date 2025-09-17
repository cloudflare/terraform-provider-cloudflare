
resource "cloudflare_address_map" "%[1]s" {
  account_id  = "%[2]s"
  enabled = %t
  %[4]s
  %[5]s
  %[6]s
  %[7]s
}
