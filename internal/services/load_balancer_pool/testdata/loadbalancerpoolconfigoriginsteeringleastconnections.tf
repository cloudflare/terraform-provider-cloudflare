
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-basic-%[1]s"
  latitude = 12.3
  longitude = 55
  origins =[ {
    name = "example-1"
    address = "192.0.2.1"
    enabled = true
  }]
  origin_steering = {
  policy = "least_connections"
}
}