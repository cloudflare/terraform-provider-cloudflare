resource "cloudflare_load_balancer_pool" "%s" {
  account_id = "%s"
  name       = "my-tf-pool-objattrs-%s"

  origins = [
    {
      name    = "example-1"
      address = "192.0.2.1"
      enabled = true
      weight  = 1.0
    },
    {
      name    = "example-2"
      address = "192.0.2.2"
      enabled = true
      weight  = 0.5
    }
  ]

  # The two attributes that broke the v5.0–v5.18 → v5.19 upgrade (#7098).
  # In early v5 these were stored as JSON objects at schema_version=0; the
  # slot-0 upgrader expected the v4 (SDKv2) array shape and rejected them
  # pre-handler. The fix detects shape from req.RawState directly.
  load_shedding = {
    default_percent = 55
    default_policy  = "random"
    session_percent = 12
    session_policy  = "hash"
  }

  origin_steering = {
    policy = "random"
  }

  description     = "tfacc-object-attrs"
  enabled         = false
  minimum_origins = 2
}
