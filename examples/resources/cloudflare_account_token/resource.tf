resource "cloudflare_account_token" "example_account_token" {
  account_id = "eb78d65290b24279ba6f44721b3ea3c4"
  name = "readonly token"
  policies = [{
    id = "f267e341f3dd4697bd3b9f71dd96247f"
    effect = "allow"
    permission_groups = [{
      id = "c8fed203ed3043cba015a93ad1616f1f"
      meta = {
        key = "key"
        value = "value"
      }
      name = "Zone Read"
    }, {
      id = "82e64a83756745bbbb1c9c2701bf816b"
      meta = {
        key = "key"
        value = "value"
      }
      name = "Magic Network Monitoring"
    }]
    resources = {
      com_cloudflare_api_account_zone_22b1de5f1c0e4b3ea97bb1e963b06a43 = "*"
      com_cloudflare_api_account_zone_eb78d65290b24279ba6f44721b3ea3c4 = "*"
    }
  }]
  condition = {
    request_ip = {
      in = ["123.123.123.0/24", "2606:4700::/32"]
      not_in = ["123.123.123.100/24", "2606:4700:4700::/48"]
    }
  }
  expires_on = "2020-01-01T00:00:00Z"
  not_before = "2018-07-01T05:20:00Z"
}
