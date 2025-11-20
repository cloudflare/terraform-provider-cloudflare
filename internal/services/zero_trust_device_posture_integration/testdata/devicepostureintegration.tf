resource "cloudflare_zero_trust_device_posture_integration" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "crowdstrike_s2s"
	interval                  = "24h"
	config = {
		api_url       =  "%[5]s"
		client_id     =  "%[3]s"
		client_secret =  "%[4]s"
		customer_id   =  "%[6]s"
	}
}
