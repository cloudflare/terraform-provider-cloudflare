# v4 config using cloudflare_zero_trust_access_policy with connection_rules
# This reproduces the research-team error:
#   Error: AttributeName("connection_rules"): invalid JSON, expected "{", got "["
resource "cloudflare_zero_trust_access_policy" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "non_identity"

  include {
    any_valid_service_token = true
  }

  # v4 stores connection_rules as a list block (array in state)
  # v5 stores connection_rules as a single nested object
  connection_rules {
    ssh {
      usernames = ["root", "admin"]
    }
  }
}
