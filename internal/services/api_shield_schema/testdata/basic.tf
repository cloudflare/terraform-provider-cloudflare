resource "cloudflare_api_shield_schema" "%[2]s" {
  zone_id = "%[1]s"
  file = jsonencode({
    "openapi" : "3.0.3",
    "info" : {
      "title" : "Example",
      "version" : "0.1.0"
    },
    "servers" : [
      {
        "url" : "api.example.com"
      }
    ],
    "paths" : {
      "/" : {}
    }
  })
  kind               = "openapi_v3"
  name               = "example_schema.json"
  validation_enabled = "true"
}
