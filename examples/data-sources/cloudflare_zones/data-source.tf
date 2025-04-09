data "cloudflare_zones" "example_zones" {
  account = {
    id = "id"
    name = "name"
  }
  direction = "desc"
  name = "name"
  order = "status"
  status = "initializing"
}
