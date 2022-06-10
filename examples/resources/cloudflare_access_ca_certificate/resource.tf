# account level
resource "cloudflare_access_ca_certificate" "example" {
  account_id     = "1d5fdc9e88c8a8c4518b068cd94331fe"
  application_id = "6cd6cea3-3ef2-4542-9aea-85a0bbcd5414"
}

# zone level
resource "cloudflare_access_ca_certificate" "another_example" {
  zone_id        = "b6bc7eb6027c792a6bca3dc91fd2d7e0"
  application_id = "fe2be0ff-7f13-4350-8c8e-a9b9795fe3c2"
}
