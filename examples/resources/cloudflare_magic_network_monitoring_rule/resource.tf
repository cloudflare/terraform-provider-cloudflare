resource "cloudflare_magic_network_monitoring_rule" "example_magic_network_monitoring_rule" {
  account_id = "6f91088a406011ed95aed352566e8d4c"
  automatic_advertisement = true
  name = "my_rule_1"
  prefixes = ["203.0.113.1/32"]
  type = "zscore"
  bandwidth_threshold = 1000
  duration = "1m"
  packet_threshold = 10000
  prefix_match = "exact"
  zscore_sensitivity = "high"
  zscore_target = "bits"
}
