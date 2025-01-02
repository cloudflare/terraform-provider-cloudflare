resource "cloudflare_api_shield_operation" "%[1]s" {
	zone_id  = "%[2]s"
	method   = "%[3]s"
	host     = "%[4]s"
	endpoint = "%[5]s"
}
