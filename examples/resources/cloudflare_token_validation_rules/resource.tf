resource "cloudflare_token_validation_rules" "example_token_validation_rules" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  action = "log"
  description = "Long description for Token Validation Rule"
  enabled = true
  expression = "is_jwt_valid(\"52973293-cb04-4a97-8f55-e7d2ad1107dd\") or is_jwt_valid(\"46eab8d1-6376-45e3-968f-2c649d77d423\")"
  selector = {
    exclude = [{
      operation_ids = ["f9c5615e-fe15-48ce-bec6-cfc1946f1bec", "56828eae-035a-4396-ba07-51c66d680a04"]
    }]
    include = [{
      host = ["v1.example.com", "v2.example.com"]
    }]
  }
  title = "Example Token Validation Rule"
}
