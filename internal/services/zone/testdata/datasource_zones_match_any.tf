# Zones with match any - match zones with either status
data "cloudflare_zones" "%[1]s" {
  match  = "any"
  status = "%[2]s"  # Would match zones with this status
  # In a real scenario you'd have multiple filter conditions
  # but we're keeping it simple for testing
}