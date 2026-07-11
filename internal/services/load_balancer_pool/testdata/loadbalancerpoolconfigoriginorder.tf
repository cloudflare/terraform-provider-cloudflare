
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name = "my-tf-pool-reorder-%[1]s"
  origins = [{
    name = "origin-c"
    address = "192.0.2.3"
    enabled = true
  }, {
    name = "origin-a"
    address = "192.0.2.1"
    enabled = true
  }, {
    name = "origin-b"
    address = "192.0.2.2"
    enabled = true
  }]
}
