data "cloudflare_zones" "example_zones" {
  account = {
    id = "id"
    name = "name"
  }
  direction = "asc"
  name = "name"
  order = "name"
  status = "initializing"
}
