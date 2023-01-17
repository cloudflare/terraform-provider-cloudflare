resource "cloudflare_load_balancer_pool" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "example-pool"
  origins {
    name    = "example-1"
    address = "192.0.2.1"
    enabled = false
    header {
      header = "Host"
      values = ["example-1"]
    }
  }
  origins {
    name    = "example-2"
    address = "192.0.2.2"
    header {
      header = "Host"
      values = ["example-2"]
    }
  }
  latitude           = 55
  longitude          = -12
  description        = "example load balancer pool"
  enabled            = false
  minimum_origins    = 1
  notification_email = "someone@example.com"
  load_shedding {
    default_percent = 55
    default_policy  = "random"
    session_percent = 12
    session_policy  = "hash"
  }
  origin_steering {
    policy = "random"
  }
}
