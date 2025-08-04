resource "cloudflare_zero_trust_access_application" "%[1]s" {
  name       = "%[1]s"
  account_id = "%[3]s"
  domain     = "%[1]s.%[2]s"
  type       = "self_hosted"
}

resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  name           = "%[1]s"
  account_id     = "%[3]s"
  decision       = "allow"
  include = [{
    email = {
      email = "test@example.com"
    }
  }]
  isolation_required             = "true"
  approval_required              = "false"
  purpose_justification_required = "false"
}

resource "cloudflare_zero_trust_gateway_settings" "%[1]s" {
  account_id = "%[3]s"
  settings = {
    activity_log = {
      enabled = true
    }
    protocol_detection = {
      enabled = true
    }
    tls_decrypt = {
      enabled = true
    }
    browser_isolation = {
      url_browser_isolation_enabled = true
      non_identity_enabled = false
    }
  }
}