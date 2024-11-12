resource "cloudflare_notification_policy_webhooks" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Webhooks destination"
  url        = "https://example.com"
  secret     = "my-secret"
}
