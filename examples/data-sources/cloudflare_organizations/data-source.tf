data "cloudflare_organizations" "example_organizations" {
  id = ["a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8"]
  containing = {
    account = "account"
    organization = "organization"
    user = "user"
  }
  name = {
    contains = "contains"
    ends_with = "endsWith"
    starts_with = "startsWith"
  }
  page_size = 0
  page_token = "page_token"
  parent = {
    id = "a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8"
  }
}
