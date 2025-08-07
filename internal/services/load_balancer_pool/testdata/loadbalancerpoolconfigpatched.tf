resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-patched-%[1]s"
  description = "Patched load balancer pool"
  enabled = false
  latitude = 37.7749
  longitude = -122.4194
  minimum_origins = 2
  check_regions = ["WEU"]
  
  load_shedding {
    default_percent = 25
    default_policy = "random"
    session_percent = 10
    session_policy = "hash"
  }
  
  origin_steering {
    policy = "random"
  }
  
  origins = [
    {
      name = "patched-origin-1"
      address = "192.0.2.2"
      enabled = true
      weight = 0.5
    },
    {
      name = "patched-origin-2"
      address = "192.0.2.3"
      enabled = true
      weight = 0.5
    }
  ]
}
