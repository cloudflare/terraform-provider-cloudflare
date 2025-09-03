data "cloudflare_schema_validation_schemas" "example_schema_validation_schemas" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  schema_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
  omit_source = true
}
