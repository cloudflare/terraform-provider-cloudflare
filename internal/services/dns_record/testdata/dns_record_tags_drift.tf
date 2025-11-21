resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  name    = "tf-acctest-tags-drift.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.10"
  ttl     = 3600
  proxied = false

  # Explicitly set tags to empty list to test drift behavior
  tags = []
}

resource "cloudflare_dns_record" "%[1]s_with_tags" {
  zone_id = "%[2]s"
  name    = "tf-acctest-tags-with.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.11"
  ttl     = 3600
  proxied = false

  # Set explicit tags
  tags = ["test:tag1", "env:test"]
}

resource "cloudflare_dns_record" "%[1]s_with_settings" {
  zone_id = "%[2]s"
  name    = "tf-acctest-settings.%[1]s.%[3]s"
  type    = "CNAME"
  content = "example.%[3]s"
  ttl     = 3600
  proxied = false

  # Explicitly set settings to test drift
  settings = {
    flatten_cname = false
  }

  tags = []
}

resource "cloudflare_dns_record" "%[1]s_no_tags" {
  zone_id = "%[2]s"
  name    = "tf-acctest-no-tags.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.12"
  ttl     = 3600
  proxied = false

  # Don't specify tags at all to test default behavior
}
