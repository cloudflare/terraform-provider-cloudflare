resource "cloudflare_zero_trust_dlp_sensitivity_group" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-%[1]s-group"
  description = "Parent group for data_class acceptance test"
}

resource "cloudflare_zero_trust_dlp_sensitivity_level" "%[1]s" {
  account_id           = "%[2]s"
  sensitivity_group_id = cloudflare_zero_trust_dlp_sensitivity_group.%[1]s.id
  name                 = "tf-acc-%[1]s-level"
  description          = "Sensitivity level for data_class acceptance test"
}

resource "cloudflare_zero_trust_dlp_data_tag_category" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-%[1]s-category"
  description = "Parent category for data_class acceptance test"
}

resource "cloudflare_zero_trust_dlp_data_tag" "%[1]s" {
  account_id  = "%[2]s"
  category_id = cloudflare_zero_trust_dlp_data_tag_category.%[1]s.id
  name        = "tf-acc-%[1]s-tag"
  description = "Data tag for data_class acceptance test"
}

resource "cloudflare_zero_trust_dlp_data_class" "%[1]s" {
  account_id  = "%[2]s"
  name        = "tf-acc-%[1]s-renamed"
  description = "Updated description"
  # Expression references a predefined DLP entry by UUID. Predefined
  # entries have stable UUIDs across accounts. This one is "AI Prompt
  # Content: Customer data" — different from basic.tf to exercise the
  # update path.
  expression  = "dlp_match(dlp.entries[\"5bce03be-0b03-434c-9f56-7b42512e0ff6\"])"
  data_tags   = [cloudflare_zero_trust_dlp_data_tag.%[1]s.id]
  sensitivity_levels = [
    {
      group_id = cloudflare_zero_trust_dlp_sensitivity_group.%[1]s.id
      level_id = cloudflare_zero_trust_dlp_sensitivity_level.%[1]s.id
    },
  ]
}
