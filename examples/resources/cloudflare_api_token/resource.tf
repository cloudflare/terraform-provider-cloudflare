resource "cloudflare_api_token" "example_api_token" {
  name = "readonly token"
  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "c8fed203ed3043cba015a93ad1616f1f"
      meta = {
        key = "key"
        value = "value"
      }
    }, {
      id = "82e64a83756745bbbb1c9c2701bf816b"
      meta = {
        key = "key"
        value = "value"
      }
    }]
    resources = {
      foo = {
        foo = "string"
      }
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
