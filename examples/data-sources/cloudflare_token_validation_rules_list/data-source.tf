data "cloudflare_token_validation_rules_list" "example_token_validation_rules_list" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
  action = "log"
  enabled = true
  host = "www.example.com"
  hostname = "www.example.com"
  rule_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
  token_configuration = ["f174e90a-fafe-4643-bbbc-4a0ed4fc8415"]
}
