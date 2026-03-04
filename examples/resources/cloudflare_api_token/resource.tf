resource "cloudflare_api_token" "example_api_token" {
  name       = "workers read-only token"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "1a71c399035b4950a1bd1466bbe4f420"
    }, {
      id = "8b47d2786a534c08a1f94ee8f9f599ef"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.b67e14daa5f8dceeb91fe5449ba496eb" = "*"
    })
  }]

  condition = {
    request_ip = {
      in     = ["123.123.123.0/24", "2606:4700::/32"]
      not_in = ["123.123.123.0/28", "2606:4700:4700::/48"]
    }
  }

  expires_on = "2027-10-01T00:00:00Z"
  not_before = "2025-10-01T00:00:00Z"
}