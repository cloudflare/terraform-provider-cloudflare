resource "cloudflare_byo_ip_prefix" "%[1]s" {
  account_id = "%[2]s"
  prefix_id  = "%[3]s"
}
