resource "cloudflare_user_group" "example_user_group" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "My New User Group"
  policies = [{
    access = "allow"
    permission_groups = [{
      id = "c8fed203ed3043cba015a93ad1616f1f"
    }, {
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resource_groups = [{
      id = "6d7f2f5f5b1d4a0e9081fdc98d432fd1"
    }]
  }]
}
