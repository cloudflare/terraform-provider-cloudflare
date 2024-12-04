data "cloudflare_zero_trust_tunnel_cloudflareds" "example_zero_trust_tunnel_cloudflareds" {
  account_id = "699d98642c564d2e855e9661899b7252"
  exclude_prefix = "vpc1-"
  existed_at = "2019-10-12T07:20:50.52Z"
  include_prefix = "vpc1-"
  is_deleted = true
  name = "blog"
  status = "inactive"
  uuid = "f70ff985-a4ef-4643-bbbc-4a0ed4fc8415"
  was_active_at = "2009-11-10T23:00:00Z"
  was_inactive_at = "2009-11-10T23:00:00Z"
}
