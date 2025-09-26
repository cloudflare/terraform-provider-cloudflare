resource "cloudflare_zero_trust_gateway_policy" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Local egress policy via WARP IPs"
  precedence  = 12403
  action      = "egress"
  filters     = ["egress"]
  traffic     = "net.dst.port in {443 80}"
  
  # Omit rule_settings.egress to use local egress via WARP IPs
}