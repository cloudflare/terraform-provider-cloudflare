resource "cloudflare_ai_gateway" "example_ai_gateway" {
  account_id = "3ebbcb006d4d46d7bb6a8c7f14676cb0"
  id = "my-gateway"
  cache_invalidate_on_update = true
  cache_ttl = 0
  collect_logs = true
  rate_limiting_interval = 0
  rate_limiting_limit = 0
  authentication = true
  log_management = 10000
  log_management_strategy = "STOP_INSERTING"
  logpush = true
  logpush_public_key = "xxxxxxxxxxxxxxxx"
  rate_limiting_technique = "fixed"
  retry_backoff = "constant"
  retry_delay = 0
  retry_max_attempts = 1
  workers_ai_billing_mode = "postpaid"
  zdr = true
}
