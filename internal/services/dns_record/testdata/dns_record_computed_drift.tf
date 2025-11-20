# Test configuration to verify computed fields don't cause drift
# This tests the issue from https://github.com/cloudflare/terraform-provider-cloudflare/issues/5517

resource "cloudflare_dns_record" "%[1]s_minimal" {
  zone_id = "%[2]s"
  name    = "tf-acctest-minimal.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.20"
  ttl     = 3600
  proxied = false

  # Minimal configuration - let provider handle defaults
}

resource "cloudflare_dns_record" "%[1]s_cname_settings" {
  zone_id = "%[2]s"
  name    = "tf-acctest-cname.%[1]s.%[3]s"
  type    = "CNAME"
  content = "target.%[3]s"
  ttl     = 60
  proxied = false

  # Explicitly set settings that were showing drift
  settings = {
    flatten_cname = false
  }

  # Empty tags list as reported in issue
  tags = []
}

resource "cloudflare_dns_record" "%[1]s_with_comment" {
  zone_id = "%[2]s"
  name    = "tf-acctest-comment.%[1]s.%[3]s"
  type    = "A"
  content = "192.168.0.21"
  ttl     = 3600
  proxied = false
  comment = "Test comment for drift"
  tags    = []
}
