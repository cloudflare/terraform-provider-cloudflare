resource "cloudflare_device_posture_integration" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Device posture integration"
  type       = "workspace_one"
  interval   = "24h"
  config {
    api_url       = "https://example.com/api"
    auth_url      = "https://example.com/connect/token"
    client_id     = "client-id"
    client_secret = "client-secret"
  }
}
