resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_pipeline_sink" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "r2"
  format = {
    type = "json"
  }
  schema = {
    fields = []
  }
  config = {
    account_id = "%[2]s"
    bucket     = cloudflare_r2_bucket.%[1]s.name
    credentials = {
      access_key_id     = "%[3]s"
      secret_access_key = "%[4]s"
    }
  }
}
