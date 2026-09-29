resource "cloudflare_zone_tracing_rules" "example_zone_tracing_rules" {
  zone_id = "zone_id"
  rules = [{
    action = "set_trace_settings"
    action_parameters = {
      sampling_ratio = 0
    }
    description = "description"
    enabled = true
    expression = "x"
  }]
}
