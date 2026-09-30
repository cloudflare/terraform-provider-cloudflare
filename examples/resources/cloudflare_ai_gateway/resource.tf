resource "cloudflare_ai_gateway" "example_ai_gateway" {
  account_id = "3ebbcb006d4d46d7bb6a8c7f14676cb0"
  id = "my-gateway"
  cache_invalidate_on_update = true
  cache_ttl = 0
  collect_logs = true
  rate_limiting_interval = 0
  rate_limiting_limit = 0
  authentication = true
  byok_only = true
  dlp = {
    action = "BLOCK"
    enabled = true
    profiles = ["string"]
  }
  guardrails = {
    prompt = {
      p1 = "FLAG"
      s1 = "FLAG"
      s10 = "FLAG"
      s11 = "FLAG"
      s12 = "FLAG"
      s13 = "FLAG"
      s2 = "FLAG"
      s3 = "FLAG"
      s4 = "FLAG"
      s5 = "FLAG"
      s6 = "FLAG"
      s7 = "FLAG"
      s8 = "FLAG"
      s9 = "FLAG"
    }
    response = {
      p1 = "FLAG"
      s1 = "FLAG"
      s10 = "FLAG"
      s11 = "FLAG"
      s12 = "FLAG"
      s13 = "FLAG"
      s2 = "FLAG"
      s3 = "FLAG"
      s4 = "FLAG"
      s5 = "FLAG"
      s6 = "FLAG"
      s7 = "FLAG"
      s8 = "FLAG"
      s9 = "FLAG"
    }
  }
  log_classification = true
  log_management = 10000
  log_management_strategy = "STOP_INSERTING"
  logpush = true
  logpush_public_key = "xxxxxxxxxxxxxxxx"
  otel = [{
    headers = {
      foo = "string"
    }
    url = "https://example.com"
    authorization = "authorization"
    content_type = "json"
  }]
  rate_limiting_technique = "fixed"
  retry_backoff = "constant"
  retry_delay = 0
  retry_max_attempts = 1
  spend_limits = {
    enabled = true
    rules = [{
      limit = 1
      limit_type = "cost"
      window = 1
      id = "x"
      enabled = true
      metadata = {
        foo = {
          mode = "partition"
        }
      }
      model = {
        mode = "filter"
        values = ["string"]
      }
      ai_gateway_provider = {
        mode = "filter"
        values = ["string"]
      }
      technique = "fixed"
    }]
  }
  store_id = "store_id"
  stripe = {
    authorization = "authorization"
    usage_events = [{
      payload = "payload"
    }]
  }
  workers_ai_billing_mode = "postpaid"
  zdr = true
}
