
resource "cloudflare_load_balancer_pool" "%[1]s" {
  account_id = "%[2]s"
  name       = "my-tf-pool-order-%[1]s"

  # Origins deliberately listed in DESCENDING name order. The API canonically
  # returns them sorted ascending by name, so without the response-reordering
  # fix this config produces a perpetual diff (LB-5712 / GitHub #7179).
  origins = [
    {
      name    = "clusterB"
      address = "192.0.2.2"
      weight  = 0
      enabled = true
    },
    {
      name    = "clusterA"
      address = "192.0.2.1"
      weight  = 1
      enabled = true
    },
  ]
}
