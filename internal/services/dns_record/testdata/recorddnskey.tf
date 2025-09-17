
	 resource "cloudflare_dns_record" "dnskey" {
 		zone_id = "%[1]s"
	   	name    = "%[2]s"
	   	type    = "DNSKEY"

	   	data = {
  algorithm  = 2
		 	flags      = 2371
		 	protocol   = 13
		 	public_key = "mdsswUyr3DPW132mOi8V9xESWE8jTo0dxCjjnopKl+GqJxpVXckHAeF+KkxLbxILfDLUT0rAK9iUzy1L53eKGQ=="
}
	 }
