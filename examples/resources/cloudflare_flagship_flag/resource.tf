resource "cloudflare_flagship_flag" "example_flagship_flag" {
  account_id = "account_id"
  app_id = "app_id"
  default_variation = "x"
  enabled = true
  key = "x"
  rules = [{
    conditions = [{
      attribute = "x"
      operator = "equals"
      value = {

      }
    }]
    priority = 1
    serve_variation = "x"
    rollout = {
      percentage = 0
      attribute = "x"
    }
  }]
  variations = {
    foo = "string"
  }
  description = "description"
  type = "boolean"
}
