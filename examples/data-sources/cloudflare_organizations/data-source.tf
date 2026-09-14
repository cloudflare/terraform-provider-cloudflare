data "cloudflare_organizations" "example_organizations" {
  id = ["a7b9c3d2e8f4a1b5c6d0e9f2a3b7c4d8"]
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
    id = "a7b9c3d2e8f4a1b5c6d0e9f2a3b7c4d8"
  }
}
