resource "cloudflare_zero_trust_access_group" "example_zero_trust_access_group" {
  include = [{
    certificate = {

    }
  }]
  name = "Allow devs"
  zone_id = "zone_id"
  exclude = [{
    certificate = {

    }
  }]
  is_default = true
  require = [{
    certificate = {

    }
  }]
}
