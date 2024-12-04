resource "cloudflare_zero_trust_access_policy" "example_zero_trust_access_policy" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  decision = "allow"
  include = [{
    group = {
      id = "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"
    }
  }]
  name = "Allow devs"
  exclude = [{
    group = {
      id = "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"
    }
  }]
  require = [{
    group = {
      id = "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"
    }
  }]
}
