resource "cloudflare_zero_trust_gateway_logging" "example_zero_trust_gateway_logging" {
  account_id = "699d98642c564d2e855e9661899b7252"
  redact_pii = true
  settings_by_rule_type = {
    dns = {
      log_all = false
      log_blocks = true
    }
    http = {
      log_all = false
      log_blocks = true
    }
    l4 = {
      log_all = false
      log_blocks = true
    }
  }
}
