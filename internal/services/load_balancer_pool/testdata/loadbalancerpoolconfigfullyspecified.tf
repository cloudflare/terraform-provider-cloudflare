resource "cloudflare_load_balancer_pool" "%[1]s" {
  name       = "my-tf-pool-basic-%[1]s"
  account_id = "%[3]s"
  origins = [
    {
      name    = "example-1"
      address = "192.0.2.1"
      enabled = false
      weight  = 1.0
      header = {
        host = ["test1.%[2]s"]
      }
    },
    {
      name    = "example-2"
      address = "192.0.2.2"
      weight  = 0.5
      header = {
        host = ["test2.%[2]s"]
      }
    },
  ]

  load_shedding = {
    default_percent = 55
    default_policy  = "random"
    session_percent = 12
    session_policy  = "hash"
  }

  latitude  = 12.3
  longitude = 55

  origin_steering = {
    policy = "random"
  }

  check_regions   = ["WEU"]
  description     = "tfacc-fully-specified"
  enabled         = false
  minimum_origins = 2
  // monitor = abcd TODO: monitor resource
  notification_email = "someone@example.com"
}
