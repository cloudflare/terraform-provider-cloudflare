resource "cloudflare_account_token" "example_account_token" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "readonly token"
  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "c8fed203ed3043cba015a93ad1616f1f"
      meta = {
        category = "category"
        deprecated = "deprecated"
        description = "description"
        editable = "editable"
        eol_at = "2019-12-27T18:11:19.117Z"
        label = "load_balancer_admin"
        scopes = "com.cloudflare.api.account"
        visibility = "visibility"
      }
    }, {
      id = "82e64a83756745bbbb1c9c2701bf816b"
      meta = {
        category = "category"
        deprecated = "deprecated"
        description = "description"
        editable = "editable"
        eol_at = "2019-12-27T18:11:19.117Z"
        label = "fbm_user"
        scopes = "com.cloudflare.api.account"
        visibility = "visibility"
      }
    }]
    resources = {
      "com.cloudflare.api.account.zone.22b1de5f1c0e4b3ea97bb1e963b06a43" = "*"
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
