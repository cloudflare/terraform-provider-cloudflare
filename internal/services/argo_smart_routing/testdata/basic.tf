resource "cloudflare_argo_smart_routing" "%[2]s" {
	 zone_id = "%[1]s"
  value   = "on"
}
