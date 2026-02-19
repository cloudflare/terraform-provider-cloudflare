resource "cloudflare_api_shield" "%s" {
  zone_id = "%s"
  # No auth_id_characteristics specified (Optional in v4)
}
