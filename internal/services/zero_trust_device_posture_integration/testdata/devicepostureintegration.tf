
resource "cloudflare_zero_trust_device_posture_integration" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "workspace_one"
	interval                  = "24h"
	config = {
  api_url       =  "%[5]s"
		auth_url      =  "%[6]s"
		client_id     =  "%[3]s"
		client_secret =  "%[4]s"
}
}
