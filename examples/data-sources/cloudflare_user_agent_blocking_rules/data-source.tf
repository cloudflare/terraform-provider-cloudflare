data "cloudflare_user_agent_blocking_rules" "example_user_agent_blocking_rules" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  description = "abusive"
  description_search = "abusive"
  ua_search = "Safari"
}
