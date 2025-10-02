resource "cloudflare_zero_trust_access_application" "complete_app" {
  account_id                   = "f037e56e89293a057740de681ac9abbe"
  name                         = "Complete Application"
  domain                       = "complete.example.com"
  type                         = "self_hosted"
  session_duration             = "12h"
  domain_type                  = "public"
  auto_redirect_to_identity    = true
  enable_binding_cookie        = false
  http_only_cookie_attribute   = true
  same_site_cookie_attribute   = "lax"
  custom_deny_message          = "Access denied to this application"
  custom_deny_url              = "https://example.com/denied"
  skip_interstitial            = true
  logo_url                     = "https://example.com/logo.png"
  app_launcher_visible         = true
  service_auth_401_redirect    = false
  custom_non_identity_deny_url = "https://example.com/non-identity-denied"
  allowed_idps                 = toset(["idp-1", "idp-2", "idp-3"])
  
  # Block-style destinations (v4 format)
  destinations {
    uri = "https://app1.example.com"
  }
  
  destinations {
    uri = "tcp://db.example.com:5432"
  }
  
  destinations {
    uri = "ssh://server.example.com:22"
  }

  # String list policies (v4 format)
  policies = [
    cloudflare_zero_trust_access_policy.allow_policy.id,
    cloudflare_zero_trust_access_policy.deny_policy.id,
    "literal-policy-id-123"
  ]

  cors_headers {
    allowed_methods   = ["GET", "POST", "OPTIONS"]
    allowed_origins   = ["https://example.com"]
    allow_credentials = true
    max_age          = 86400
  }

  saas_app {
    sp_entity_id         = "https://example.com/saml"
    consumer_service_url = "https://example.com/saml/acs"
    name_id_format       = "email"
    default_relay_state  = "https://example.com/dashboard"
    
    custom_attribute {
      name   = "department"
      source {
        name = "user.department"
      }
    }
    
    custom_attribute {
      name   = "role"
      source {
        name = "user.role"
      }
    }
  }

  tags = ["production", "critical", "external"]
}

resource "cloudflare_zero_trust_access_application" "app_launcher_type" {
  account_id                   = "f037e56e89293a057740de681ac9abbe"
  name                         = "App Launcher Application"
  type                         = "app_launcher"
  skip_app_launcher_login_page = true  # This should be preserved for app_launcher type
  app_launcher_visible         = true
  logo_url                     = "https://example.com/launcher-logo.png"
}

resource "cloudflare_zero_trust_access_application" "warp_type" {
  account_id   = "f037e56e89293a057740de681ac9abbe"
  name         = "WARP Application"
  type         = "warp"
  domain_type  = "private"  # Should be removed in v5
  allowed_idps = toset([])
  
  skip_app_launcher_login_page = false  # Should be removed for non-app_launcher types
  
  destinations {
    uri = "10.0.0.0/8"
  }
}

# Referenced policies for testing
resource "cloudflare_zero_trust_access_policy" "allow_policy" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Allow Policy"
  decision   = "allow"
  
  include {
    everyone = true
  }
}

resource "cloudflare_zero_trust_access_policy" "deny_policy" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Deny Policy"
  decision   = "deny"
  
  include {
    everyone = true
  }
}