# Given you have the following zones in Cloudflare.
#
#  - example.com
#  - example.net
#  - not-example.com
#
# Look for a single zone that you know exists using an exact match.
# API request will be for zones?name=example.com. Will not match not-example.com
# or example.net.
data "cloudflare_zones" "example" {
  filter {
    name = "example.com"
  }
}

# Look for all zones which include "example".
# API request will be for zones?name=contains:example. Will return all three
# zones.
data "cloudflare_zones" "example" {
  filter {
    name        = "example"
    lookup_type = "contains"
  }
}

# Look for all zones which include "example" but start with "not-".
# API request will be for zones?name=contains:example. Will perform client side
# filtering using the provided regex and will only match the single zone,
# not-example.com.
data "cloudflare_zones" "example" {
  filter {
    name        = "example"
    lookup_type = "contains"
    match       = "^not-"
  }
}

# Look for all active zones in an account.
data "cloudflare_zones" "example" {
  filter {
    account_id = "f037e56e89293a057740de681ac9abbe"
    status     = "active"
  }
}
