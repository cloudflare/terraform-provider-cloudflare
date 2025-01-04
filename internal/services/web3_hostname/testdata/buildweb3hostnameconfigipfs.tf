
resource "cloudflare_web3_hostname" "%[1]s" {
	zone_id = "%[2]s"
	name = "%[1]s.%[3]s"
	target = "ipfs"
	description = "test"
	dnslink = "/ipns/onboarding.ipfs.cloudflare.com"
}
