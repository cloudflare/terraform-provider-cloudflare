resource "cloudflare_zero_trust_access_policy" "example_zero_trust_access_policy" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  decision = "allow"
  include = [{
    group = {
      id = "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"
    }
  }]
  name = "Allow devs"
  approval_groups = [{
    approvals_needed = 1
    email_addresses = ["test1@cloudflare.com", "test2@cloudflare.com"]
    email_list_uuid = "email_list_uuid"
  }, {
    approvals_needed = 3
    email_addresses = ["test@cloudflare.com", "test2@cloudflare.com"]
    email_list_uuid = "597147a1-976b-4ef2-9af0-81d5d007fc34"
  }]
  approval_required = true
  connection_rules = {
    rdp = {
      allowed_clipboard_local_to_remote_formats = ["text"]
      allowed_clipboard_remote_to_local_formats = ["text"]
    }
  }
  exclude = [{
    group = {
      id = "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"
    }
  }]
  isolation_required = false
  mfa_config = {
    allowed_authenticators = ["totp", "biometrics", "security_key"]
    mfa_disabled = false
    session_duration = "24h"
  }
  purpose_justification_prompt = "Please enter a justification for entering this protected domain."
  purpose_justification_required = true
  require = [{
    group = {
      id = "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"
    }
  }]
  session_duration = "24h"
}
