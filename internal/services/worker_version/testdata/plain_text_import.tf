resource "cloudflare_worker" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"
}

resource "cloudflare_worker_version" "%[1]s" {
  account_id = "%[2]s"
  worker_id = cloudflare_worker.%[1]s.id
  modules = [
    {
      name           = "index.js"
      content_base64 = "ZXhwb3J0IGRlZmF1bHQge2ZldGNoKCkge3JldHVybiBuZXcgUmVzcG9uc2UoKX19"
      content_type   = "application/javascript+module"
    }
  ]
  main_module = "index.js"
  bindings = [
    {
      type = "plain_text"
      name = "BLOCKED_CITIES"
      text = "kyiv,kharkiv"
    },
    {
      type = "plain_text"
      name = "COUNTRY_CODE"
      text = "UA"
    }
  ]
}
