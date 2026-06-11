resource "cloudflare_zero_trust_dlp_sensitivity_group" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-%[1]s"
  description = "Acceptance test group for sensitivity_level_order data source"
}

# Levels are created serially via depends_on. See basic.tf for rationale.
resource "cloudflare_zero_trust_dlp_sensitivity_level" "%[1]s_public" {
  account_id           = "%[2]s"
  sensitivity_group_id = cloudflare_zero_trust_dlp_sensitivity_group.%[1]s.id
  name                 = "Public"
}

resource "cloudflare_zero_trust_dlp_sensitivity_level" "%[1]s_internal" {
  account_id           = "%[2]s"
  sensitivity_group_id = cloudflare_zero_trust_dlp_sensitivity_group.%[1]s.id
  name                 = "Internal"
  depends_on           = [cloudflare_zero_trust_dlp_sensitivity_level.%[1]s_public]
}

resource "cloudflare_zero_trust_dlp_sensitivity_level_order" "%[1]s" {
  account_id           = "%[2]s"
  sensitivity_group_id = cloudflare_zero_trust_dlp_sensitivity_group.%[1]s.id
  level_ids = [
    cloudflare_zero_trust_dlp_sensitivity_level.%[1]s_public.id,
    cloudflare_zero_trust_dlp_sensitivity_level.%[1]s_internal.id,
  ]
}

data "cloudflare_zero_trust_dlp_sensitivity_level_order" "%[1]s" {
  account_id           = "%[2]s"
  sensitivity_group_id = cloudflare_zero_trust_dlp_sensitivity_group.%[1]s.id
  depends_on           = [cloudflare_zero_trust_dlp_sensitivity_level_order.%[1]s]
}
