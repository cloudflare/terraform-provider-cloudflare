# For account level applications

resource "cloudflare_access_application" "example" {
  account_id = "%[2]s"
  name = "example"
  domain = "example.com"
}

# You can use either `name` or `domain` to identify the application

data "cloudflare_access_application" "example" {
  account_id = "%[2]s"
  name = "example"
}

data "cloudflare_access_application" "example" {
  account_id = "%[2]s"
  domain = "example.com"
}

# For zone level applications

resource "cloudflare_access_application" "example" {
  zone_id = "%[3]s"
  name = "example"
  domain = "example.com"
}

# You can use either `name` or `domain` to identify the application

data "cloudflare_access_application" "example" {
  zone_id = "%[3]s"
  name = "example"
}

data "cloudflare_access_application" "example" {
  zone_id = "%[3]s"
  domain = "example.com"
}