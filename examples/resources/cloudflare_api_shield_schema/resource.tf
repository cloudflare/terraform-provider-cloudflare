resource "cloudflare_api_shield_schema" "petstore_schema" {
  zone_id            = "0da42c8d2132a9ddaf714f9e7c920711"
  name               = "myschema"
  kind               = "openapi_v3" # optional
  validation_enabled = true         # optional, default false
  source             = file("./schemas/petstore.json")
}
