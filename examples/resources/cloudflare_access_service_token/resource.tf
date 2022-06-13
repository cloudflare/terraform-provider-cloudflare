resource "cloudflare_access_service_token" "my_app" {
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
  name       = "CI/CD app"
}

# Generate a service token that will renew if terraform is ran within 30 days of expiration
resource "cloudflare_access_service_token" "my_app" {
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
  name       = "CI/CD app renewed"

  min_days_for_renewal = 30

  # This flag is important to set if min_days_for_renewal is defined otherwise
  # there will be a brief period where the service relying on that token
  # will not have access due to the resource being deleted
  lifecycle {
    create_before_destroy = true
  }
}
