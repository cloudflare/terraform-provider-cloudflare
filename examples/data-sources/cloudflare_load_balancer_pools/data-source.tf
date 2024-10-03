data "cloudflare_load_balancer_pools" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  filter {
    name = "example-lb-pool"
  }
}
