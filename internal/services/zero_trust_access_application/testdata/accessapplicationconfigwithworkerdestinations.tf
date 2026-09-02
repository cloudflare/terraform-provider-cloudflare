resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  subdomain = {
    enabled          = true
    previews_enabled = true
  }
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id                = "%[2]s"
  name                      = "%[1]s"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = false

  # The Access `worker_id` is the Worker's immutable id, which the newer
  # `cloudflare_worker` resource exposes directly as its `id`.
  destinations = [
    {
      type      = "worker"
      worker_id = cloudflare_worker.%[1]s.id
    },
    {
      type      = "preview_worker"
      worker_id = cloudflare_worker.%[1]s.id
    }
  ]
}
