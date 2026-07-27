data "cloudflare_load_balancer" "example_load_balancer" {
  load_balancer_id = "699d98642c564d2e855e9661899b7252"
  account_id = "account_id"
  zone_id = "zone_id"
}
