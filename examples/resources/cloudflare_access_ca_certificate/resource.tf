# account level
resource "cloudflare_access_ca_certificate" "example" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  application_id = "6cd6cea3-3ef2-4542-9aea-85a0bbcd5414"
}

# zone level
resource "cloudflare_access_ca_certificate" "another_example" {
  zone_id        = "0da42c8d2132a9ddaf714f9e7c920711"
  application_id = "fe2be0ff-7f13-4350-8c8e-a9b9795fe3c2"
}
