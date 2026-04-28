resource "cloudflare_zero_trust_organization" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  auth_domain = "%[1]s-%[3]s"

  mfa_config = {
    allowed_authenticators = ["totp", "security_key", "ssh_piv_key"]
    session_duration       = "24h"
  }
}

resource "cloudflare_zero_trust_tunnel_cloudflared_virtual_network" "%[1]s" {
  account_id         = "%[2]s"
  name               = "%[1]s"
  comment            = "%[1]s"
  is_default_network = "false"
}

resource "cloudflare_zero_trust_access_infrastructure_target" "%[1]s" {
  account_id = "%[2]s"
  hostname   = "%[1]s"
  ip = {
    ipv4 = {
      ip_addr            = "127.0.0.1"
      virtual_network_id = cloudflare_zero_trust_tunnel_cloudflared_virtual_network.%[1]s.id
    }
  }
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "infrastructure"
  target_criteria = [
    {
      port     = 22
      protocol = "SSH"
      target_attributes = {
        "hostname" = ["%[1]s"]
      }
    }
  ]
  policies = [
    {
      name     = "%[1]s-policy-1"
      decision = "allow"
      include = [
        {
          email = { email = "example@cloudflare.com" }
        }
      ]
      connection_rules = {
        ssh = {
          usernames = ["root"]
        }
      }
      mfa_config = {
        allowed_authenticators = ["ssh_piv_key"]
        session_duration       = "12h"
        mfa_disabled           = false
      }
    }
  ]

  depends_on = [cloudflare_zero_trust_organization.%[1]s]
}
